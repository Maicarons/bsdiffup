# [bsdiffup](https://github.com/Maicarons/bsdiffup)
bsdiffup - BSDiff Update Tool

[![GitHub license](https://img.shields.io/github/license/Maicarons/bsdiffup?style=flat-square)](https://github.com/Maicarons/bsdiffup/blob/master/LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Maicarons/bsdiffup?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/Maicarons/bsdiffup?style=flat-square)](https://github.com/Maicarons/bsdiffup)
[![GoDoc](https://godoc.org/github.com/Maicarons/bsdiffup?status.svg)](https://pkg.go.dev/github.com/Maicarons/bsdiffup)
[![Go Report Card](https://goreportcard.com/badge/github.com/Maicarons/bsdiffup?style=flat-square)](https://goreportcard.com/report/github.com/Maicarons/bsdiffup)
---

## Usage
```go
	goDiff, err := diff.GoDiff("./testdir/dir1", "./testdir/dir2")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(goDiff)
```
