package tgz_builder
import (
	"testing"
	"os"
)
func TestCreateTgz(t *testing.T){
	err, tgz := CreateTgz("./fixtures/bar.tar.gz")
	if err != nil {
		t.Fatal(err)
	}

	if(tgz.Path != "./fixtures/bar.tar.gz"){
		t.Fatal("Path is not set correctly in the Tgz struct");
	}
}

func TestAddingAFile(t *testing.T){
	tarFile := "./fixtures/oneFile.tar.gz"
	err, tgz := CreateTgz(tarFile)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tgz.Path)

	err = tgz.AddFile("./fixtures/test.txt", "test.txt")
	if err != nil{
		t.Fatal(err)
	}

	tgz.Finish()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() != 108 {
		t.Fatalf("tar file should be 108 bytes, but is %d bytes", tarStats.Size())
	}
}

func TestAddingTwoFiles(t *testing.T){
	tarFile := "./fixtures/twoFiles.tar.gz"
	err, tgz := CreateTgz(tarFile)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tgz.Path)

	err = tgz.AddFile("./fixtures/test.txt", "test.txt")
	if err != nil{
		t.Fatal(err)
	}
	err = tgz.AddFile("./fixtures/test.txt", "test2.txt")
	if err != nil{
		t.Fatal(err)
	}

	tgz.Finish()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() != 129 {
		t.Fatalf("tar file should be 129 bytes, but is %d bytes", tarStats.Size())
	}
}
