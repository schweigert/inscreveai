package html

import "regexp"

var textReg *regexp.Regexp

func init() {
	var err error
	textReg, err = regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		panic(err)
	}
}
