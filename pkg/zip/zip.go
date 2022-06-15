package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Extract(source, destination string) ([]string, error) {
	source = filepath.Clean(source)
	destination = filepath.Clean(destination)

	z, err := zip.OpenReader(source)
	if err != nil {
		return nil, err
	}
	defer z.Close()

	err = os.MkdirAll(destination, os.ModePerm)
	if err != nil {
		return nil, err
	}

	var extractedFiles []string
	for _, f := range z.File {
		err := write(destination, f)
		if err != nil {
			return nil, err
		}

		extractedFiles = append(extractedFiles, f.Name)
	}

	return extractedFiles, nil
}

func write(destination string, f *zip.File) error {
	zz, err := f.Open()
	if err != nil {
		return err
	}
	defer zz.Close()

	path := filepath.Join(destination, f.Name)

	if f.FileInfo().IsDir() {
		err = os.MkdirAll(path, f.Mode())
		if err != nil {
			return err
		}
	} else {
		err = os.MkdirAll(filepath.Dir(path), f.Mode())
		if err != nil {
			return err
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(f, zz)
		if err != nil {
			return err
		}
	}

	return nil
}
