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
	final := "\n"

	for _, el := range dom.children {
		final += el.HtmlSafe() + "\n"
	}

	if dom.attrs == "" {
		return "\n<" + dom.name + ">" + final + "</" + dom.name + ">"
	}
	return "\n<" + dom.name + " " + dom.attrs + ">" + final + "</" + dom.name + ">"
}
