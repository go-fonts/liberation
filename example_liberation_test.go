// Copyright ©2020 The go-fonts Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package liberation_test

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
