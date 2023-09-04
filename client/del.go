package main

import (
	"encoding/json"
	"fmt"
	"github.com/Maicarons/bsdiffup/diff"
	"os"
)

func Del() {
	// 读取JSON文件内容
	fileContent := `
[
  {
    "path":"good\\apple.html",
    "type":"deleted",
    "md5":"6620d7be0e9e40f4bd3c0a5a86590ecf"
  },
  {
    "path":"good\\food.bin",
    "type":"modified",
    "md5":"d41d8cd98f00b204e9800998ecf8427e"
  },
  {
    "path":"img.bmp",
    "type":"modified",
    "md5":"9f7c76cbd9a189856db543a95654d5c5"
  },
  {
    "path":"main.py",
    "type":"added",
    "md5":"a53826cc458d6ba62653e9d66f23d4e9"
  }
]
`

	var fileInfos []diff.FileDiff
	err := json.Unmarshal([]byte(fileContent), &fileInfos)
	if err != nil {
		fmt.Println("解析JSON文件时发生错误:", err)
		return
	}

	// 遍历文件信息并删除"deleted"类型的文件
	for _, fileInfo := range fileInfos {
		if fileInfo.Type == "deleted" {
			err := os.Remove(fileInfo.Path)
			if err != nil {
				fmt.Printf("删除文件 %s 时发生错误: %v\n", fileInfo.Path, err)
			} else {
				fmt.Printf("成功删除文件 %s\n", fileInfo.Path)
			}
		}
	}
}
