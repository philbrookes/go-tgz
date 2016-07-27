package tgz

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"bytes"
)

func TestNewTgz(t *testing.T) {
	tgz, err := New("./fixtures/bar.tar.gz")
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
	tgz, err := New(tarFile)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tgz.Path)

	err = tgz.AddFileByContent([]byte("sdfsdfsdfs\n"), "test.txt")
	if err != nil {
		t.Fatal(err)
	}

	tgz.Close()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
	tar.Close()

	files, err := decompressAndListFiles(tgz.Path)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := files["test.txt"]; !ok {
		t.Fatal("Expected tgz to contain test.txt but it didnt")
	}

}

func TestAddingTwoFilesByContent(t *testing.T) {
	tarFile := "./fixtures/twoFilesByContent.tar.gz"
	tgz, err := New(tarFile)
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

	tgz.Close()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
	tar.Close()

	files, err := decompressAndListFiles(tgz.Path)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := files["test.txt"]; !ok {
		t.Fatal("Expected tgz to contain test.txt but it didnt")
	}
	if _, ok := files["test2.txt"]; !ok {
		t.Fatal("Expected tgz to contain test2.txt but it didnt")
	}
}

func TestAddingAFileByPath(t *testing.T) {
	tarFile := "./fixtures/oneFileByPath.tar.gz"
	tgz, err := New(tarFile)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tgz.Path)

	err = tgz.AddFileByPath("./fixtures/test.txt", "test.txt")
	if err != nil {
		t.Fatal(err)
	}

	tgz.Close()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
	tar.Close()

	files, err := decompressAndListFiles(tgz.Path)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := files["test.txt"]; !ok {
		t.Fatal("Expected tgz to contain test.txt but it didnt")
	}
}

func TestAddingAFileByBuffer(t *testing.T) {
	tarFile := "./fixtures/oneFileByBuffer.tar.gz"
	tgz, err := New(tarFile)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tgz.Path)

	b := &bytes.Buffer{}
	b.Write([]byte("test\n"))
	b.Write([]byte("test 2\n"))

	err = tgz.AddFileByBuffer(b, "test.txt")
	if err != nil {
		t.Fatal(err)
	}

	tgz.Close()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
	tar.Close()

	files, err := decompressAndListFiles(tgz.Path)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := files["test.txt"]; !ok {
		t.Fatal("Expected tgz to contain test.txt but it didnt")
	}
}


func TestAddingTwoFilesByBuffer(t *testing.T) {
	tarFile := "./fixtures/twoFilesByBuffer.tar.gz"
	tgz, err := New(tarFile)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tgz.Path)

	b := &bytes.Buffer{}
	b.Write([]byte("test\n"))
	b.Write([]byte("test 2\n"))

	err = tgz.AddFileByBuffer(b, "test.txt")
	err = tgz.AddFileByBuffer(b, "test2.txt")
	if err != nil {
		t.Fatal(err)
	}

	tgz.Close()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
	tar.Close()

	files, err := decompressAndListFiles(tgz.Path)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := files["test.txt"]; !ok {
		t.Fatal("Expected tgz to contain test.txt but it didnt")
	}
	if _, ok := files["test2.txt"]; !ok {
		t.Fatal("Expected tgz to contain test2.txt but it didnt")
	}
}

func TestAddingTwoFilesByPath(t *testing.T) {
	tarFile := "./fixtures/twoFileByPath.tar.gz"
	tgz, err := New(tarFile)

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

	tgz.Close()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
	tar.Close()

	files, err := decompressAndListFiles(tgz.Path)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := files["test.txt"]; !ok {
		t.Fatal("Expected tgz to contain test.txt but it didnt")
	}
	if _, ok := files["test2.txt"]; !ok {
		t.Fatal("Expected tgz to contain test.txt but it didnt")
	}
}

func TestAddingMixedFiles(t *testing.T) {
	tarFile := "./fixtures/twoMixedFiles.tar.gz"
	tgz, err := New(tarFile)

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

	tgz.Close()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
	tar.Close()

	files, err := decompressAndListFiles(tgz.Path)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := files["test.txt"]; !ok {
		t.Fatal("Expected tgz to contain test.txt but it didnt")
	}
	if _, ok := files["test2.txt"]; !ok {
		t.Fatal("Expected tgz to contain test.txt but it didnt")
	}
}

func TestAddingFilesInSubdirs(t *testing.T) {
	tarFile := "./fixtures/filesInSubdirs.tar.gz"
	tgz, err := New(tarFile)

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
	tgz.Close()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
	tar.Close()

	outFiles, err := decompressAndListFiles(tgz.Path)
	if err != nil {
		t.Fatal(err)
	}
	for dest, _ := range files {
		if _, ok := outFiles[dest]; !ok {
			t.Fatal("Expected tgz to contain " + dest + ", but it didnt")
		}
	}
}

func TestWritingToExistingTar(t *testing.T) {
	tarFile := "./fixtures/rewriteTo.tar.gz"
	tgz, err := New(tarFile)

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

	tgz.Close()

	tar, _ := os.Open(tgz.Path)
	tarStats, _ := tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}

	tgz, err = New(tarFile)
	if err != nil {
		t.Fatal(err)
	}

	err = tgz.AddFileByContent([]byte("sdfsdfsdfs\n"), "test.txt")
	if err != nil {
		t.Fatal(err)
	}
	err = tgz.AddFileByPath("./fixtures/test.txt", "test2.txt")
	if err != nil {
		t.Fatal(err)
	}

	tgz.Close()
	tar, _ = os.Open(tgz.Path)
	tarStats, _ = tar.Stat()
	if tarStats.Size() == 0 {
		t.Fatalf("tar file should be > 0 bytes, but is %d bytes", tarStats.Size())
	}
	if tarStats.Size() > 2048 {
		t.Fatalf("tar is much larger than expected, should be < 2048 but is %d byes", tarStats.Size())
	}
	tar.Close()

	files, err := decompressAndListFiles(tgz.Path)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := files["test.txt"]; !ok {
		t.Fatal("Expected tgz to contain test.txt but it didnt")
	}
	if _, ok := files["test2.txt"]; !ok {
		t.Fatal("Expected tgz to contain test.txt but it didnt")
	}
}

func decompressAndListFiles(pathToTgz string) (map[string]string, error) {
	os.Mkdir("./fixtures/uncompressed", 0755)
	defer os.RemoveAll("./fixtures/uncompressed")

	cmd := exec.Command("tar", "-xf", pathToTgz, "-C", "./fixtures/uncompressed")
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	ret := map[string]string{}

	err = filepath.Walk("./fixtures/uncompressed", func(path string, f os.FileInfo, err error) error {
		ret[strings.Replace(path, "fixtures/uncompressed/", "", 1)] = strings.Replace(path, "fixtures/uncompressed", "", 1)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return ret, nil
}
