package utilities

func ErrorChecker(e error) {
	if e != nil {
		panic(e)
	}
}
