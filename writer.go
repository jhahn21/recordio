package recordio

import (
	"encoding/binary"
	"io"
)

type compressor interface {
	compress(data []byte) ([]byte, error)
}

// Writer implements sequential writing of variable-length data.
type Writer struct {
	w io.Writer
	c compressor

	recHeader []byte
}

// NewWriter returns a new writer that writes records to w.
func NewWriter(w io.Writer, mode Mode) (*Writer, error) {
	var c compressor
	switch mode {
	case NoCompression:
		c = newNoCompressor()
	case Snappy:
		c = newSnappyCompressor()
	case Zlib:
		c = newZlibCompressor()
	default:
		return nil, ErrInvalidMode
	}

	recHeader := make([]byte, recHeaderSize)
	binary.LittleEndian.PutUint32(recHeader, uint32(mode))

	return &Writer{
		w:         w,
		c:         c,
		recHeader: recHeader,
	}, nil
}

// Write writes a record. It returns always equal to len(rec) on success.
func (w *Writer) Write(rec []byte) (int, error) {
	if len(rec) > maxRecordSize {
		return 0, ErrRecordTooLarge
	}

	binary.LittleEndian.PutUint32(w.recHeader[4:], uint32(len(rec)))

	p, err := w.c.compress(rec)
	if err != nil {
		return 0, err
	}

	binary.LittleEndian.PutUint32(w.recHeader[8:], uint32(len(p)))

	if _, err := w.w.Write(w.recHeader); err != nil {
		return 0, err
	}
	if _, err := w.w.Write(p); err != nil {
		return 0, err
	}

	return len(rec), nil
}
