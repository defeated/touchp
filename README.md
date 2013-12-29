# touchp

basically `mkdir -p` + `touch`

## usage

```sh
touchp path/to/new/directory/file.ext
```

This will create any subdirectories that don't exist. if `file.ext` exists it
will update the last modified date, otherwise it will create `file.ext`

## developing

  1. `go get github.com/defeated/touchp`
  2. `go test -cover`