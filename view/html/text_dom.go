package html

import "fmt"

type textDom struct {
	element interface{}
}

func Text(element interface{}) Dom {
	return &textDom{element: element}
}

func (dom *textDom) HtmlSafe() string {
	return fmt.Sprintf("%v", dom.element)
}
