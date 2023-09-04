package zip

import (
	"archive/zip"
	"fmt"
	"github.com/Maicarons/bsdiffup/diff"
	"io"
	"os"
)

type FileDiffType int

const (
	Modified FileDiffType = iota
	Added
	Deleted
)

func Zipping(fileDiffs []diff.FileDiff) {
	zipFile, err := os.Create("update.zip")
	if err != nil {
		fmt.Println("Error creating zip file:", err)
		return
	}
	defer zipFile.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, difft := range fileDiffs {
		if difft.Type == diff.FileDiffTypeModified || difft.Type == diff.FileDiffTypeAdded {
			file, err := os.Open(difft.Path)
			if err != nil {
				fmt.Println("Error opening file:", err)
				continue
			}
			defer file.Close()

			// Create a new file header
			fileInfo, err := file.Stat()
			if err != nil {
				fmt.Println("Error getting file info:", err)
				continue
			}
			header := &zip.FileHeader{
				Name:   difft.Path,
				Method: zip.Deflate, // Use zip.Deflate for compression
			}
			header.SetModTime(fileInfo.ModTime())

			// Create a new file in the zip archive
			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				fmt.Println("Error creating zip file entry:", err)
				continue
			}

			// Copy the file data into the zip writer
			_, err = io.Copy(writer, file)
			if err != nil {
				fmt.Println("Error copying file to zip:", err)
				continue
			}
		}
	}

}
