package recordio_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/jhahn21/recordio"
)

func TestRecordIO(t *testing.T) {
	const numRecords = 100

	rec := func(n int) []byte {
		return bytes.Repeat([]byte{byte(n)}, n)
	}

	for m := recordio.NoCompression; m < recordio.EndOfMode; m++ {
		t.Run(m.String(), func(t *testing.T) {
			var buf bytes.Buffer
			w, err := recordio.NewWriter(&buf, m)
			if err != nil {
				t.Fatal(err)
			}

			for i := 0; i < numRecords; i++ {
				n, err := w.Write(rec(i))
				if err != nil {
					t.Fatal(err)
				}
				if n != i {
					t.Errorf("rec[%d]: wanted to write %d, but got %d", i, i, n)
				}
			}

			r, err := recordio.NewReader(&buf)
			if err != nil {
				t.Fatal(err)
			}

			i := 0
			for ; ; i++ {
				got, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Fatal(err)
				}

				if wanted := rec(i); bytes.Compare(got, wanted) != 0 {
					t.Errorf("rec[%d]: wanted to read %x, but got %x", i, wanted, got)
				}
			}

			if i != numRecords {
				t.Errorf("wanted to read %d records, but got %d", numRecords, i)
			}
		})
	}
}

// infiniteRecs emits an infinite stream of records.
type infiniteRecs struct {
	buf []byte
	off int
}

func newInfiniteRecs(b []byte) *infiniteRecs {
	return &infiniteRecs{buf: b}
}

func (r *infiniteRecs) Read(p []byte) (int, error) {
	n := copy(p, r.buf[r.off:])
	r.off = (r.off + n) % len(r.buf)
	return n, nil
}

func genRec(rand *rand.Rand, n int) []byte {
	rec := make([]byte, n)
	rand.Read(rec)
	return rec
}

func BenchmarkRead(b *testing.B) {
	rand := rand.New(rand.NewSource(1000))

	sizes := []int{100, 1000, 10000}

	for m := recordio.NoCompression; m < recordio.EndOfMode; m++ {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("%v_%d", m, size), func(b *testing.B) {
				var buf bytes.Buffer
				w, _ := recordio.NewWriter(&buf, m)
				w.Write(genRec(rand, size))

				r, _ := recordio.NewReader(newInfiniteRecs(buf.Bytes()))

				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					if _, err := r.Read(); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}

func BenchmarkWrite(b *testing.B) {
	rand := rand.New(rand.NewSource(1000))

	sizes := []int{100, 1000, 10000}

	for m := recordio.NoCompression; m < recordio.EndOfMode; m++ {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("%v_%d", m, size), func(b *testing.B) {
				rec := genRec(rand, size)

				w, _ := recordio.NewWriter(ioutil.Discard, m)

				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					if _, err := w.Write(rec); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}
