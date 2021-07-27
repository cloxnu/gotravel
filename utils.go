package main

import (
	"io"
	"os"
)

func CopyRes(dir string)  {
	entries, err := resFS.ReadDir("res/" + dir)
	if err != nil {
		return
	}

	err = os.MkdirAll("res/" + dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		_, err = CopyResFile("res/" + dir + "/" + entry.Name())
		if err != nil {
			panic(err)
		}
	}
}

func CopyResFile(filename string) (written int64, err error) {
	file, err := resFS.Open(filename)
	if err != nil {
		return
	}

	desFile, err := os.Create(filename)
	if err != nil {
		return
	}
	defer func(desFile *os.File) {
		err = desFile.Close()
		if err != nil {
			panic(err)
		}
	}(desFile)

	return io.Copy(desFile, file)
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
