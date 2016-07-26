package tgz_builder

import (
	"os"
	"testing"
)

func TestNewTgz(t *testing.T) {
	err, tgz := New("./fixtures/bar.tar.gz")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tgz.Path)

	if tgz.Path != "./fixtures/bar.tar.gz" {
		t.Fatal("Path is not set correctly in the Tgz struct")
	}
}

func TestAddingAFileByContent(t *testing.T) {
	tarFile := "./fixtures/oneFileByContent.tar.gz"
	err, tgz := New(tarFile)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tgz.Path)

	err = tgz.AddFileByContent([]byte("sdfsdfsdfs\n"), "test.txt")
	if err != nil {
		t.Fatal(err)
	}

	tgz.Finish()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
}

func TestAddingTwoFilesByContent(t *testing.T) {
	tarFile := "./fixtures/twoFilesByContent.tar.gz"
	err, tgz := New(tarFile)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tgz.Path)

	err = tgz.AddFileByContent([]byte("sdfsdfsdfs\n"), "test.txt")
	if err != nil {
		t.Fatal(err)
	}
	err = tgz.AddFileByContent([]byte("sdfsdfsdfs\n"), "test2.txt")
	if err != nil {
		t.Fatal(err)
	}

	tgz.Finish()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
}

func TestAddingAFileByPath(t *testing.T) {
	tarFile := "./fixtures/oneFileByPath.tar.gz"
	err, tgz := New(tarFile)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tgz.Path)

	err = tgz.AddFileByPath("./fixtures/test.txt", "test.txt")
	if err != nil {
		t.Fatal(err)
	}

	tgz.Finish()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
}

func TestAddingTwoFilesByPath(t *testing.T) {
	tarFile := "./fixtures/twoFileByPath.tar.gz"
	err, tgz := New(tarFile)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tgz.Path)

	err = tgz.AddFileByPath("./fixtures/test.txt", "test.txt")
	if err != nil {
		t.Fatal(err)
	}
	err = tgz.AddFileByPath("./fixtures/test.txt", "test2.txt")
	if err != nil {
		t.Fatal(err)
	}

	tgz.Finish()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
}

func TestAddingMixedFiles(t *testing.T) {
	tarFile := "./fixtures/twoMixedFiles.tar.gz"
	err, tgz := New(tarFile)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tgz.Path)

	err = tgz.AddFileByContent([]byte("sdfsdfsdfs\n"), "test.txt")
	if err != nil {
		t.Fatal(err)
	}
	err = tgz.AddFileByPath("./fixtures/test.txt", "test2.txt")
	if err != nil {
		t.Fatal(err)
	}

	tgz.Finish()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
}

func TestAddingFilesInSubdirs(t *testing.T) {
	tarFile := "./fixtures/filesInSubdirs.tar.gz"
	err, tgz := New(tarFile)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tgz.Path)
	files := map[string]string{
		"files/test1.txt":       "sdfsdfsdfs\n",
		"files/test2.txt":       "sdfsdfsdfs\n",
		"files/test3.txt":       "sdfsdfsdfs\n",
		"files/test4.txt":       "sdfsdfsdfs\n",
		"other_files/test1.txt": "sdfsdfsdfs\n",
		"other_files/test2.txt": "sdfsdfsdfs\n",
		"other_files/test3.txt": "sdfsdfsdfs\n",
		"other_files/test4.txt": "sdfsdfsdfs\n",
	}

	for dest, content := range files {
		err = tgz.AddFileByContent([]byte(content), dest)
		if err != nil {
			t.Fatal(err)
		}
	}
	tgz.Finish()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
}
