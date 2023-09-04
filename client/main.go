package main

import (
	"fmt"
	"github.com/Maicarons/bsdiffup/zip"
	"github.com/Masterminds/semver/v3"
	"log"
)

func main() {
	zip.Unzip("update.zip", "bsdiffup")
	
	goupdate, err := updateIfHigher("1.1.9", "1.2.0")
	if err != nil {
		log.Fatal(err)
		return
	}
	if !goupdate {
		return
	}

	Del()
}

func updateIfHigher(currentVersionStr, inputVersionStr string) (bool, error) {
	// 解析当前版本号
	currentVersion, err := semver.NewVersion(currentVersionStr)
	if err != nil {
		return false, fmt.Errorf("无法解析当前版本号: %v", err)
	}

	// 解析用户输入的版本号
	inputVersion, err := semver.NewVersion(inputVersionStr)
	if err != nil {
		return false, fmt.Errorf("无法解析输入的版本号: %v", err)
	}

	// 判断是否需要更新
	if inputVersion.GreaterThan(currentVersion) {
		fmt.Printf("需要更新到新版本: %s\n", inputVersionStr)
		return true, nil
	}

	// 输入版本不高于当前版本，输出错误信息
	return false, fmt.Errorf("输入的版本不高于当前版本，无需更新")
}
