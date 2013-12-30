package main

import (
	"os"
	"path/filepath"
	"time"
)

func RecursiveTouch(name string) {
	dirs := filepath.Dir(name)
	os.MkdirAll(dirs, 0755)

	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		os.Create(name)
	} else {
		os.Chtimes(name, time.Now(), time.Now())
	}
}
