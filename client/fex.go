package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 定义文件名格式的正则表达式
	pattern := `^(.+)_(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?\.zip$`

	// 编译正则表达式
	regex := regexp.MustCompile(pattern)

	// 传入的文件名
	fileName := "exampleFile_1.2.3-alpha.1+build123.zip"

	// 使用正则表达式来匹配文件名
	matches := regex.FindStringSubmatch(fileName)

	// 检查是否匹配成功
	if len(matches) > 1 {
		// 提取命名捕获组的内容
		fullName := matches[1]
		major := matches[2]
		minor := matches[3]
		patch := matches[4]
		prerelease := matches[5]
		buildmetadata := matches[6]

		fmt.Printf("文件名: %s\n", fullName)
		fmt.Printf("Major版本号: %s\n", major)
		fmt.Printf("Minor版本号: %s\n", minor)
		fmt.Printf("Patch版本号: %s\n", patch)
		fmt.Printf("预发布版本: %s\n", prerelease)
		fmt.Printf("构建元数据: %s\n", buildmetadata)
	} else {
		fmt.Println("文件名不符合指定格式")
	}
}
