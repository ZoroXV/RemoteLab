package utils

import (
	"os"
	"path/filepath"
	"io"
	"mime/multipart"
)

func FileExists(directory string, filename string) bool {
	path := filepath.Join(directory, filename)

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GetFullPath(directory string, filename string) string {
    rootPath, _ := os.Getwd()

    path := filepath.Join(directory, filename)
    path = filepath.Join(rootPath, path)

    return path
}

func DirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func MakeDirectory(dirname string) {
	if !DirExists(dirname) {
		os.MkdirAll(dirname, os.ModePerm)
	}
}

func SaveFile(directory string, filename string, inputFile multipart.File) error {
	if filename == "" {
		return nil
	}

	MakeDirectory(directory)

	path := filepath.Join(directory, filename)
	outputFile, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}

	return nil
}
