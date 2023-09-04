package main

import (
	"flag"
	"github.com/Maicarons/bsdiffup/diff"
	"log"
	"os"
)

// main -d1 /path/to/first/folder -d2 /path/to/second/folder -p /path/to/output/file.txt

func main() {
	// 定义命令行参数
	var dir1, dir2, outputFile string
	var isZip bool
	flag.StringVar(&dir1, "d1", "", "第一个文件夹路径")
	flag.StringVar(&dir2, "d2", "", "第二个文件夹路径")
	flag.StringVar(&outputFile, "p", "", "输出文件位置")

	flag.BoolVar(&isZip, "zip", false, "是否拣择文件并压缩")

	flag.Parse()

	// 检查参数是否缺失
	if dir1 == "" || dir2 == "" || outputFile == "" {
		log.Println("请提供-d1、-d2和-p参数")
		return
	}

	// 输出用户提供的文件夹信息到指定文件
	outputContent, err := diff.GoDiff(dir1, dir2)
	if err != nil {
		log.Printf("比较文件时发生错误：%v\n", err)
		return
	}
	err = writeToFile(outputFile, outputContent)
	if err != nil {
		log.Printf("写入文件时发生错误：%v\n", err)
		return
	}

	log.Println("信息已写入", outputFile)
	if isZip {
		log.Printf("开始压缩")
	}

}

// 将内容写入文件
func writeToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
