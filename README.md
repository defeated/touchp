# touchp

basically `mkdir -p` + `touch`

## install

### osx

brew install https://raw.githubusercontent.com/defeated/touchp/master/touchp.rb

## usage

```sh
touchp path/to/new/directory/file.ext
```

This will create any subdirectories that don't exist. if `file.ext` exists it
will update the last modified date, otherwise it will create `file.ext`

## developing

  1. `go get github.com/defeated/touchp`
  2. `go test -cover`
