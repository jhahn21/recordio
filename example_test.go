package recordio_test

import (
	"bytes"
	"fmt"
	"io"

	"github.com/jhahn21/recordio"
)

func ExampleNewWriter() {
	var b bytes.Buffer

	w, err := recordio.NewWriter(&b, recordio.Snappy)
	if err != nil {
		panic(err)
	}

	w.Write([]byte("first"))
	w.Write([]byte("second"))

	fmt.Println(b.Bytes())

	// Output:
	// [210 159 49 224 5 0 0 0 23 0 0 0 255 6 0 0 115 78 97 80 112 89 1 9 0 0 85 255 35 229 102 105 114 115 116 210 159 49 224 6 0 0 0 24 0 0 0 255 6 0 0 115 78 97 80 112 89 1 10 0 0 211 224 211 202 115 101 99 111 110 100]
}

func ExampleNewReader() {
	b := bytes.NewReader([]byte{210, 159, 49, 224, 5, 0, 0, 0, 23, 0, 0, 0, 255, 6, 0, 0, 115, 78, 97, 80, 112, 89, 1, 9, 0, 0, 85, 255, 35, 229, 102, 105, 114, 115, 116, 210, 159, 49, 224, 6, 0, 0, 0, 24, 0, 0, 0, 255, 6, 0, 0, 115, 78, 97, 80, 112, 89, 1, 10, 0, 0, 211, 224, 211, 202, 115, 101, 99, 111, 110, 100})

	r, err := recordio.NewReader(b)
	if err != nil {
		panic(err)
	}

	for {
		d, err := r.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(string(d))
	}

	// Output:
	// first
	// second
}
