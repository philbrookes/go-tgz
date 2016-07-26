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
err, tgz := tgz.New("./fixtures/bar.tar.gz")
if err != nil {
    t.Fatal(err)
}
```

Add a file by content:
```
//First argument is the content in a byte array, second argument is where in the tgz it should be stored
err = tgz.AddFileByContent([]byte("sdfsdfsdfs\n"), "test.txt")
if err != nil {
    t.Fatal(err)
}
```

Or add an existing file:
```
//First argument is the path to the local file, second argument is where in the tgz it should be stored
err = tgz.AddFileByPath("/path/to/file", "file.txt")
if err != nil {
    t.Fatal(err)
}
```

Commit the changes to the tar.gz file:
```
tgz.Finish()
```

## Contributing
Fork the repo and make a PR and verify sufficient test coverage with:
```
go test --coverprofile cover.out
go tool cover -func=cover.out
```
And include the output of this in the PR description.