package logAndErrors

import "fmt"

func Wrapper(err error, msg string) error {
	return fmt.Errorf("%s -> %s", msg, err.Error())
}
