package main

import (
	"path/filepath"
	"testing"
)

func TestPath(t *testing.T) {
	path := "../../script/testdata/file1.txt"
	filename := filepath.Base(path)

	t.Log(filename)
}

func TestScanDir(t *testing.T) {
	path := "../../script/testdata"
	scanDir(path)
}

func TestScanFile(t *testing.T) {
	path := `../../script/testdata/JavaScript 语言入门.md`
	scanFile(path)
}