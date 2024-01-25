package logAndErrors

import "fmt"

func wrapper(err error, msg string) string {
	return fmt.Sprintf("%s -> %s", msg, err.Error())
}
