package recordio

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Reader implements sequential reading of variable-length data.
type Reader struct {
	r io.Reader
	d decompressor

	mode      Mode
	recHeader []byte
}

type decompressor interface {
	decompress(r io.Reader, usize, csize uint32) ([]byte, error)
}

// NewReader returns a new reader that reads records from r.
func NewReader(r io.Reader) (*Reader, error) {
	return &Reader{
		r:         r,
		recHeader: make([]byte, recHeaderSize),
	}, nil
}

// Read reads a record. The returned slice is only valid until the next call.
func (r *Reader) Read() ([]byte, error) {
	n, err := io.ReadAtLeast(r.r, r.recHeader, recHeaderSize)
	if err != nil {
		return nil, err
	}
	if n != recHeaderSize {
		return nil, ErrInvalidHeader
	}

	mode := Mode(binary.LittleEndian.Uint32(r.recHeader[:4]))
	if mode != r.mode {
		if r.mode != 0 {
			return nil, fmt.Errorf("%w: %#08x", ErrInvalidHeader, mode)
		}

		switch mode {
		case NoCompression:
			r.d = newNoDecompressor()
		case Snappy:
			r.d = newSnappyDecompressor()
		case Zlib:
			r.d = newZlibDecompressor()
		default:
			return nil, fmt.Errorf("%w: %#08x", ErrInvalidMode, mode)
		}
		r.mode = mode
	}

	usize := binary.LittleEndian.Uint32(r.recHeader[4:])
	csize := binary.LittleEndian.Uint32(r.recHeader[8:])

	rec, err := r.d.decompress(r.r, usize, csize)
	if err != nil {
		return nil, err
	}
	if len(rec) != int(usize) {
		return nil, fmt.Errorf("%w: size %d not matched with header %d", ErrRecordCorrupted, len(rec), usize)
	}

	return rec, nil
}

// TODO(jungho): Add a function to skip corrupted records.
