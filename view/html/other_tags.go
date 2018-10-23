package html

func HtmlTag(attrs string, children ...Dom) Dom {
	return Tag("html", attrs, children...)
}

func HeadTag(attrs string, children ...Dom) Dom {
	return Tag("head", attrs, children...)
}

func BodyTag(attrs string, children ...Dom) Dom {
	return Tag("body", attrs, children...)
}

func TitleTag(attrs string, children ...Dom) Dom {
	return Tag("title", attrs, children...)
}

func PTag(attrs string, children ...Dom) Dom {
	return Tag("p", attrs, children...)
}

func ATag(attrs string, children ...Dom) Dom {
	return Tag("a", attrs, children...)
}

func DivTag(attrs string, children ...Dom) Dom {
	return Tag("div", attrs, children...)
}

func NavTag(attrs string, children ...Dom) Dom {
	return Tag("nav", attrs, children...)
}

func MainTag(attrs string, children ...Dom) Dom {
	return Tag("div", attrs, children...)
}

func ButtonTag(attrs string, children ...Dom) Dom {
	return Tag("button", attrs, children...)
}

func ScriptTag(attrs string, children ...Dom) Dom {
	return Tag("script", attrs, children...)
}

func FooterTag(attrs string, children ...Dom) Dom {
	return Tag("footer", attrs, children...)
}

func FormTag(attrs string, children ...Dom) Dom {
	return Tag("form", attrs, children...)
}

func LabelTag(attrs string, children ...Dom) Dom {
	return Tag("label", attrs, children...)
}

func SpanTag(attrs string, children ...Dom) Dom {
	return Tag("span", attrs, children...)
}

func InputTag(attrs string, children ...Dom) Dom {
	return Tag("input", attrs, children...)
}

func UlTag(attrs string, children ...Dom) Dom {
	return Tag("ul", attrs, children...)
}

func LiTag(attrs string, children ...Dom) Dom {
	return Tag("li", attrs, children...)
}

func ITag(attrs string, children ...Dom) Dom {
	return Tag("i", attrs, children...)
}
