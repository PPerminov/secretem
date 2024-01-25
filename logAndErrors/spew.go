package logAndErrors

import "github.com/davecgh/go-spew/spew"

func Dump(i interface{}) {
	spew.Dump(i)
}
