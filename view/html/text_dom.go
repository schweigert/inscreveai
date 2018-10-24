package html

import (
	"fmt"
	"log"
	"regexp"
)

type textDom struct {
	element interface{}
}

func Text(element interface{}) Dom {
	return &textDom{element: element}
}

func (dom *textDom) HtmlSafe() string {
	return fmt.Sprintf("%v", dom.element)
}

func SafeText(element string) Dom {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(element, " ")

	return Text(processedString)
}
