// Copyright 2023 Michael D Henderson. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package main implements ...
package main

import (
	"fmt"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	return fmt.Errorf("not implemented")
}
