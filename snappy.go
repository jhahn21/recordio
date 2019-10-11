package recordio

import (
	"bytes"
	"io"

	"github.com/golang/snappy"
)

type snappyCompressor struct {
	w   *snappy.Writer
	buf bytes.Buffer
}

func newSnappyCompressor() compressor {
	return &snappyCompressor{w: snappy.NewBufferedWriter(nil)}
}

func (c *snappyCompressor) compress(p []byte) ([]byte, error) {
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

type snappyDecompressor struct {
	r   *snappy.Reader
	buf bytes.Buffer
}

func newSnappyDecompressor() decompressor {
	return &snappyDecompressor{r: snappy.NewReader(nil)}
}

func (d *snappyDecompressor) decompress(r io.Reader, usize, csize uint32) ([]byte, error) {
	d.buf.Reset()
	d.buf.Grow(int(usize))

	d.r.Reset(io.LimitReader(r, int64(csize)))
	if _, err := d.buf.ReadFrom(d.r); err != nil {
		return nil, err
	}

	return d.buf.Bytes(), nil
}
