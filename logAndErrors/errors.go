package logAndErrors

func IfErrorThen(err error, message string, fn func()) {
	if err == nil {
		return
	}
	Error(Wrapper(err, message).Error())
	fn()
}

func IfErrorThenFatal(err error, message string) {
	if err == nil {
		return
	}
	Fatal(Wrapper(err, message).Error())
}
