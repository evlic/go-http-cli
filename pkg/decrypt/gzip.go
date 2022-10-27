package decrypt

import (
	"compress/gzip"
	"fmt"
	"io"
)

func GzipDecode(src io.Reader) (string, error) {
	r, err := gzip.NewReader(src)
	if err != nil {
		fmt.Printf("err >> %#v\n", err)
		return "", err
	}
	defer r.Close()
	res, err := io.ReadAll(r)
	return string(res), err
}
