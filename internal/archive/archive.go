package archive

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"io"
	"os"
)

type Zip struct {
	buff   *bytes.Buffer
	writer *zip.Writer
}

func New() *Zip {
	b := bytes.NewBuffer(make([]byte, 0, 10240))
	w := zip.NewWriter(b)
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})
	return &Zip{buff: b, writer: w}
}

func (v *Zip) File(filename string, call func(w io.Writer) error) error {
	f, err := v.writer.Create(filename)
	if err != nil {
		return err
	}
	return call(f)
}

func (v *Zip) Save(filepath string) error {
	if err := v.writer.Close(); err != nil {
		return err
	}
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.ReadFrom(v.buff)
	if err != nil {
		return err
	}
	return nil
}
