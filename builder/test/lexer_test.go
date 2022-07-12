package test

import (
	"arcell/build"
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	l := build.NewLexer()
	l.Read("tell.conf")
	fmt.Println(l)
}
