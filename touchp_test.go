package main

// TODO: mock the filesystem instead? http://nf.wh3rd.net/10things/#8

import (
	"os"
	"testing"
)

const (
	dir = "test"
)

func init() {
	os.RemoveAll(dir)
	os.Mkdir(dir, 0755)
}

func TestRecursiveTouchNew(t *testing.T) {
	name := dir + "/new/file.ext"
	RecursiveTouch(name)

	_, err := os.Stat(name)
	if err != nil {
		t.Errorf("no such file or directory: '%v'", name)
	}
}
