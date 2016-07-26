package tgz_builder

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

type Tgz struct {
	Path      string
	tgzFile   *os.File
	tarWriter *tar.Writer
	gzWriter  *gzip.Writer
}

func CreateTgz(path string) (error, *Tgz) {
	tgz := Tgz{Path: path}
	var err error
	err, tgz.tgzFile = tgz.getTarFile()
	if err != nil {
		return err, nil
	}

	tgz.gzWriter = gzip.NewWriter(tgz.tgzFile)
	tgz.tarWriter = tar.NewWriter(tgz.gzWriter)
	return nil, &tgz
}

//although I understand the this ref it isn't how it is done in golang it would be something like tgz
func (this *Tgz) AddFile(src string, dest string) error {
	var (
		err error
		srcFile *os.File
	)

	//srcFile could be nil so close after we know there is no error
	srcFile, err = os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	//not idiomatic I would do the same as above
	stat, err := srcFile.Stat()
	if err != nil {
		return err
	}
	header := new(tar.Header)
	header.Name = dest
	header.Size = stat.Size()
	header.Mode = int64(stat.Mode())
	header.ModTime = stat.ModTime()
	// write the header to the tarball archive
	if err := this.tarWriter.WriteHeader(header); err != nil {
		return err
	}

	if _, err = io.Copy(this.tarWriter, srcFile); err != nil {
		return err
	}

	return nil
}

func (this *Tgz) Finish() {
	this.tarWriter.Close()
	this.gzWriter.Close()
	this.tgzFile.Close()
}

func (this *Tgz) getTarFile() (error, *os.File) {
	var (
		f   *os.File
		err error
	)

	if _, err = os.Stat(this.Path); os.IsNotExist(err) {
		f, err = os.Create(this.Path)
		if err != nil {
			return err, nil
		}
	} else {
		f, err = os.OpenFile(this.Path, os.O_RDWR, os.ModePerm)
		if err != nil {
			return err, nil
		}
	}

	return nil, f
}

//I don't think I would have it this way and would just add this to the Create method
//func (this *Tgz) init() (error) {
//	var err error
//	err, this.tgzFile = this.getTarFile()
//	if err != nil {
//		return err
//	}
//
//	this.gzWriter  = gzip.NewWriter(this.tgzFile)
//	this.tarWriter = tar.NewWriter(this.gzWriter)
//	return nil
//}
