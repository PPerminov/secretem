package logAndErrors

func IfErrorThen(err error, message string, fn func()) {
	if err == nil {
		return
	}
	Error(wrapper(err, message))
	fn()
}

func IfErrorThenFatal(err error, message string) {
	if err == nil {
		return
	}
	Fatal(wrapper(err, message))
}
