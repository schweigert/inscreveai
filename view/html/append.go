package html

func Append(doms ...Dom) Dom {
	return Tag("div", ``, doms...)
}
