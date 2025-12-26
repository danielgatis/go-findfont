# go-findfont

[![Go Report Card](https://goreportcard.com/badge/github.com/danielgatis/go-findfont?style=flat-square)](https://goreportcard.com/report/github.com/danielgatis/go-findfont)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/danielgatis/go-findfont/master/LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/danielgatis/go-findfont)

A pure Go library to find fonts on the system. Works on macOS, Linux, and Windows without any external dependencies.

## Install

```bash
go get -u github.com/danielgatis/go-findfont
```

And then import the package in your code:

```go
import "github.com/danielgatis/go-findfont/findfont"
```

## Usage

### Find a font by name

```go
package main

import (
	"fmt"
	"os"

	"github.com/danielgatis/go-findfont/findfont"
)

func main() {
	path, err := findfont.Find("Arial")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(path)
}
```

```
❯ go run main.go
/System/Library/Fonts/Supplemental/Arial.ttf
```

### List all fonts

```go
fonts := findfont.List()
for _, f := range fonts {
	fmt.Println(f)
}
```

### Find with custom suffixes

```go
path, err := findfont.FindWithSuffixes("MyFont", []string{".ttf", ".otf"})
```

## CLI

You can also use it as a command-line tool:

```bash
# Find a font
go-findfont Arial

# List all fonts
go-findfont --list
```

## Supported Platforms

| Platform | Font Directories |
|----------|------------------|
| macOS    | `~/Library/Fonts`, `/Library/Fonts`, `/System/Library/Fonts` |
| Linux    | `~/.fonts`, `~/.local/share/fonts`, `/usr/share/fonts`, `/usr/local/share/fonts` |
| Windows  | `%windir%\Fonts`, `%localappdata%\Microsoft\Windows\Fonts` |

## License

Copyright (c) 2020-present [Daniel Gatis](https://github.com/danielgatis)

Licensed under [MIT License](./LICENSE)
