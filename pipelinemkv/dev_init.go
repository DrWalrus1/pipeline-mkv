//go:build dev
// +build dev

package main

import (
	"fmt"
	dummymkv "pipelinemkv/dummyMkv"
)

func init() {
	fmt.Println("Hello from dummy loader!")
	commandHandler = &dummymkv.DummyMakeMkvHandler{}
}
