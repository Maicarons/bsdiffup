package diff

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

/*
FileDiff
Type:
deleted
modified
added
*/
type FileDiff struct {
	Path string `json:"path"`
	Type string `json:"type"`
	MD5  string `json:"md5"`
}

func GoDiff(Path1, Path2 string) (string, error) {
	folder1 := Path1
	folder2 := Path2

	diff, err := CompareFolders(folder1, folder2)
	if err != nil {
		log.Println(err)
		return "", err
	}

	jsonData, err := ToJSON(diff)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return jsonData, nil
}

func CompareFolders(folder1, folder2 string) ([]FileDiff, error) {
	diff := []FileDiff{}

	// 获取文件夹1中的所有文件
	files1, err := getAllFiles(folder1)
	if err != nil {
		return nil, err
	}

	// 获取文件夹2中的所有文件
	files2, err := getAllFiles(folder2)
	if err != nil {
		return nil, err
	}

	// 比较文件夹1中的文件
	for _, file1 := range files1 {
		relPath, err := filepath.Rel(folder1, file1)
		if err != nil {
			return nil, err
		}

		file2 := filepath.Join(folder2, relPath)

		// 检查文件是否存在于文件夹2中
		if containsFile(files2, file2) {
			// 比较文件内容
			if !compareFileContent(file1, file2) {
				md5sum, err := getFileMD5(file1)
				if err != nil {
					return nil, err
				}

				diff = append(diff, FileDiff{
					Path: relPath,
					Type: "modified",
					MD5:  md5sum,
				})
			}
		} else {
			md5sum, err := getFileMD5(file1)
			if err != nil {
				return nil, err
			}

			diff = append(diff, FileDiff{
				Path: relPath,
				Type: "deleted",
				MD5:  md5sum,
			})
		}
	}

	// 检查文件夹2中的新增文件
	for _, file2 := range files2 {
		relPath, err := filepath.Rel(folder2, file2)
		if err != nil {
			return nil, err
		}

		file1 := filepath.Join(folder1, relPath)

		// 检查文件是否存在于文件夹1中
		if !containsFile(files1, file1) {
			md5sum, err := getFileMD5(file2)
			if err != nil {
				return nil, err
			}

			diff = append(diff, FileDiff{
				Path: relPath,
				Type: "added",
				MD5:  md5sum,
			})
		}
	}

	return diff, nil
}

func getAllFiles(folder string) ([]string, error) {
	files := []string{}

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func containsFile(files []string, file string) bool {
	for _, f := range files {
		if f == file {
			return true
		}
	}
	return false
}

func compareFileContent(file1, file2 string) bool {
	content1, err := ioutil.ReadFile(file1)
	if err != nil {
		return false
	}

	content2, err := ioutil.ReadFile(file2)
	if err != nil {
		return false
	}

	return string(content1) == string(content2)
}

func getFileMD5(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	md5sum := hex.EncodeToString(hash.Sum(nil))
	return md5sum, nil
}

func ToJSON(diff []FileDiff) (string, error) {
	jsonData, err := json.Marshal(diff)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
