# go-findfont

[![Go Report Card](https://goreportcard.com/badge/github.com/danielgatis/go-findfont?style=flat-square)](https://goreportcard.com/report/github.com/danielgatis/go-findfont)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/danielgatis/go-findfont/master/LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/danielgatis/go-findfont)

Find system fonts through the fontconfig library (a.k.a `fc-match`)

## Install

```bash
go get -u github.com/danielgatis/go-findfont
```

And then import the package in your code:

```go
import "github.com/danielgatis/go-findfont/findfont"
```

### Example

An example described below is one of the use cases.

```go
package main

import (
	"fmt"
	"os"

	"github.com/danielgatis/go-findfont/findfont"
)

func main() {
	fonts, err := findfont.Find("Emoji", findfont.FontRegular)

	if err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}

	for _, f := range fonts {
		fmt.Printf("Family: %v\nStyle : %v\nPath  : %v\n\n", f[0], f[1], f[2])
	}
}
```


```
❯ go run main.go
Family: Apple Color Emoji
Style : Regular,標準體,Ordinær,Normal,Normaali,Regolare,レギュラー,일반체,Regulier,Обычный,常规体,عادي
Path  : /System/Library/Fonts/Apple Color Emoji.ttc

Family: Twitter Color Emoji
Style : Regular
Path  : /Users/daniel/Library/Fonts/TwitterColorEmoji-SVGinOT.ttf

Family: .LastResort
Style : Regular
Path  : /System/Library/Fonts/LastResort.otf
```


## License

Copyright (c) 2020-present [Daniel Gatis](https://github.com/danielgatis)

Licensed under [MIT License](./LICENSE)
