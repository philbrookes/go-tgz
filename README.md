This tool is intended to ease the generation of multi-file tar.gz files.


## Usage
Import the package:
```
import (
    "github.com/philbrookes/go-tgz"
)
```

Create a new tgz object:
```
archiveFile, _ := os.Open("path/to/file.tar.gz")
tar, _ := NewTgz(archiveFile)
writer := tarFile.GetWriterToFile(filepath.Join(projectPath, resource+"."+extension))

writer.Write([]byte("data"))

writer.Close()
tar.Close()
archiveFile.Close()
```

## Contributing
Fork the repo and make a PR and verify sufficient test coverage with:
```
go test --coverprofile cover.out
go tool cover -func=cover.out
```
And include the output of this in the PR description.