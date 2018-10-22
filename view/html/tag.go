package html

type tag struct {
	children []Dom
	name     string
	attrs    string
}

func Tag(name, attrs string, children ...Dom) Dom {
	return &tag{children: children, name: name, attrs: attrs}
}

func (dom *tag) HtmlSafe() string {
	final := ""

	for _, el := range dom.children {
		final += el.HtmlSafe()
	}
	if dom.attrs == "" {
		return "<" + dom.name + ">" + final + "</" + dom.name + ">"
	}
	return "<" + dom.name + " " + dom.attrs + ">" + final + "</" + dom.name + ">"
}
