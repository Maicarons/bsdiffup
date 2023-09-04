package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Unzip 解压
func Unzip(Zip string, Dir string) error {
	zipFile := Zip
	destDir := Dir

	err := UnzipFile(zipFile, destDir)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	fmt.Println("Unzip completed!")
	return nil
}

func UnzipFile(zipFile, destDir string) error {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		err := extractFile(f, destDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func extractFile(f *zip.File, destDir string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	path := filepath.Join(destDir, f.Name)

	if f.FileInfo().IsDir() {
		os.MkdirAll(path, f.Mode())
	} else {
		os.MkdirAll(filepath.Dir(path), os.ModePerm)

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(f, rc)
		if err != nil {
			return err
		}
	}

	return nil
}
