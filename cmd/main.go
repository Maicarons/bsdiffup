package main

import (
	"fmt"
	"github.com/Maicarons/bsdiffup/diff"
	"log"
)

func main() {
	goDiff, err := diff.GoDiff("./testdir/dir1", "./testdir/dir2")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(goDiff)
}
