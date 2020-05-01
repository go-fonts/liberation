# liberation

[![GitHub release](https://img.shields.io/github/release/go-fonts/liberation.svg)](https://github.com/go-fonts/liberation/releases)
[![GoDoc](https://godoc.org/github.com/go-fonts/liberation?status.svg)](https://godoc.org/github.com/go-fonts/liberation)
[![License](https://img.shields.io/badge/License-BSD--3-blue.svg)](https://github.com/go-fonts/liberation/raw/master/LICENSE)

`liberation` provides the [liberation](https://github.com/liberationfonts/liberation-fonts/) fonts as importable Go packages.

The fonts are released under the [SIL Open Font](https://github.com/go-fonts/liberation/raw/master/LICENSE-SIL) license.
The Go packages under the [BSD-3](https://github.com/go-fonts/liberation/raw/master/LICENSE) license.

## Example

```go
import (
	"fmt"
	"log"

	"github.com/go-fonts/liberation/liberationserifregular"
	"golang.org/x/image/font/sfnt"
)

func Example() {
	ttf, err := sfnt.Parse(liberationserifregular.TTF)
	if err != nil {
		log.Fatalf("could not parse Liberation Serif font: %+v", err)
	}

	fmt.Printf("num glyphs: %d\n", ttf.NumGlyphs())

	// Output:
	// num glyphs: 2589
}
```
