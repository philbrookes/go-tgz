package tgz

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type Tgz struct {
	Path      string
	tgzFile   *os.File
	tarWriter *tar.Writer
	gzWriter  *gzip.Writer
	finished  bool
}

func New(path string) (*Tgz, error) {
	tgz := Tgz{Path: path}
	var err error
	tgz.tgzFile, err = tgz.getTarFile()
	if err != nil {
		return nil, err
	}

	tgz.gzWriter = gzip.NewWriter(tgz.tgzFile)
	tgz.tarWriter = tar.NewWriter(tgz.gzWriter)
	tgz.finished = false

	return &tgz, nil
}

func (tgz *Tgz) AddFileByBuffer(b *bytes.Buffer, dest string) error {
	return tgz.AddFileByContent(b.Bytes(), dest)
}

func (tgz *Tgz) AddFileByPath(srcFile string, dest string) error {
	if src, err := ioutil.ReadFile(srcFile); err == nil {
		return tgz.AddFileByContent(src, dest)
	} else {
		return err
	}
}

func (tgz *Tgz) AddFileByContent(src []byte, dest string) error {
	if tgz.finished == true {
		return errors.New("Gzip file has already been finished, cannot add more files")
	}
	var (
		err error
	)

	header := new(tar.Header)
	header.Name = dest
	header.Size = int64(len(src))
	header.Mode = int64(uint32(0775))
	header.ModTime = time.Now()

	if err := tgz.tarWriter.WriteHeader(header); err != nil {
		return err
	}

	if _, err = io.Copy(tgz.tarWriter, bytes.NewReader(src)); err != nil {
		return err
	}

	return nil
}

func (tgz *Tgz) Close() {
	tgz.finished = true
	tgz.tarWriter.Close()
	tgz.gzWriter.Close()
	tgz.tgzFile.Close()
}

func (tgz *Tgz) getTarFile() (*os.File, error) {
	var (
		f   *os.File
		err error
	)

	if _, err = os.Stat(tgz.Path); os.IsNotExist(err) {
		f, err = os.Create(tgz.Path)
		if err != nil {
			return nil, err
		}
	} else {
		f, err = os.OpenFile(tgz.Path, os.O_RDWR, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	return f, nil
}
