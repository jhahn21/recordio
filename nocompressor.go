package recordio

import (
	"bytes"
	"io"
)

type noCompressor struct{}

func newNoCompressor() noCompressor {
	return noCompressor{}
}

func (c noCompressor) compress(p []byte) ([]byte, error) {
	return p, nil
}

type noDecompressor struct {
	buf bytes.Buffer
}

func newNoDecompressor() *noDecompressor {
	return &noDecompressor{}
}

func (d *noDecompressor) decompress(r io.Reader, usize, csize uint32) ([]byte, error) {
	d.buf.Reset()
	d.buf.Grow(int(usize))

	if _, err := d.buf.ReadFrom(io.LimitReader(r, int64(csize))); err != nil {
		return nil, err
	}

	return d.buf.Bytes(), nil
}
