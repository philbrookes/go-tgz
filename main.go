package tgz_builder

import(
	"os"
	"archive/tar"
	"compress/gzip"
	"io"
)

type Tgz struct {
	Path string
	tgzFile *os.File
	tarWriter *tar.Writer
	gzWriter *gzip.Writer
}

func CreateTgz(path string) (error, *Tgz) {
	tgz := Tgz{Path: path}
	if err := tgz.init(); err != nil {
		return err, nil
	}
	return nil, &tgz
}

func (this *Tgz) AddFile(src string, dest string) (error){
	var err error
	var srcFile *os.File


	srcFile, err = os.Open(src)
	defer srcFile.Close()

	if err != nil {
		return err
	}

	if stat, err := srcFile.Stat(); err == nil {
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
	}

	return nil
}

func (this *Tgz) Finish() {
	this.tarWriter.Close()
	this.gzWriter.Close()
	this.tgzFile.Close()
}

func (this *Tgz) getTarFile() (error, *os.File){
	var f *os.File
	var err error

	if _, err = os.Stat(this.Path); os.IsNotExist(err){
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

func (this *Tgz) init() (error) {
	var err error
	err, this.tgzFile = this.getTarFile()
	if err != nil {
		return err
	}

	this.gzWriter  = gzip.NewWriter(this.tgzFile)
	this.tarWriter = tar.NewWriter(this.gzWriter)
	return nil
}