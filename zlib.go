package recordio

import (
	"bytes"
	"compress/zlib"
	"io"
)

type zlibCompressor struct {
	w   *zlib.Writer
	buf bytes.Buffer
}

func newZlibCompressor() compressor {
	return &zlibCompressor{w: zlib.NewWriter(nil)}
}

func (c *zlibCompressor) compress(p []byte) ([]byte, error) {
	c.buf.Reset()

	c.w.Reset(&c.buf)
	if _, err := c.w.Write(p); err != nil {
		c.w.Close()
		return nil, err
	}
	if err := c.w.Close(); err != nil {
		return nil, err
	}

	return c.buf.Bytes(), nil
}

type zlibDecompressor struct {
	r        io.ReadCloser
	resetter zlib.Resetter
	buf      bytes.Buffer
}

// emptyCompressed is the zlib compressed output of an empty string.
var emptyCompressed = []byte{120, 156, 1, 0, 0, 255, 255, 0, 0, 0, 1}

func newZlibDecompressor() *zlibDecompressor {
	d := &zlibDecompressor{}
	// Note that zlib.NewReader requires valid compressed data.
	d.r, _ = zlib.NewReader(bytes.NewReader(emptyCompressed))
	// Cache zlib.Resetter to avoid a type cast assertion overhead although it's tiny.
	d.resetter = d.r.(zlib.Resetter)
	return d
}

func (d *zlibDecompressor) decompress(r io.Reader, usize, csize uint32) ([]byte, error) {
	d.buf.Reset()
	d.buf.Grow(int(usize))

	if err := d.resetter.Reset(io.LimitReader(r, int64(csize)), nil); err != nil {
		d.r.Close()
		return nil, err
	}
	if _, err := d.buf.ReadFrom(d.r); err != nil {
		d.r.Close()
		return nil, err
	}
	if err := d.r.Close(); err != nil {
		return nil, err
	}

	return d.buf.Bytes(), nil
}
