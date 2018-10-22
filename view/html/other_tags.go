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
