# [bsdiffup](https://github.com/Maicarons/bsdiffup)
bsdiffup - BSDiff Update Tool

[![GitHub license](https://img.shields.io/github/license/Maicarons/bsdiffup?style=flat-square)](https://github.com/Maicarons/bsdiffup/blob/master/LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Maicarons/bsdiffup?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/Maicarons/bsdiffup?style=flat-square)](https://github.com/Maicarons/bsdiffup)
[![GoDoc](https://godoc.org/github.com/Maicarons/bsdiffup?status.svg)](https://pkg.go.dev/github.com/Maicarons/bsdiffup)
[![Go Report Card](https://goreportcard.com/badge/github.com/Maicarons/bsdiffup?style=flat-square)](https://goreportcard.com/report/github.com/Maicarons/bsdiffup)
---

## Package

- [ ] Client - compare,unzip
- [ ] cmd - test 
- [ ] Server - compare,zip,upload
- [ ] Go package


## Usage
```go
	goDiff, err := diff.GoDiff("./testdir/dir1", "./testdir/dir2")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(goDiff)
```

## ZIP file Format

```text
{FileName}_{Semver}.zip
e.g. update_1.2.0.zip | exampleFile_1.2.3-alpha.1+build123.zip

Path
-.
-..
-diff.json
-VERSION
-other files
```

## diff.json Format
```json
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
```

## VERSION Format
```VERSION
1.2.3-alpha.1+build123
```