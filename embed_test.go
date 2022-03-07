package golangembed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed golang-logo.jpeg
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("new_golang_logo.svg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
