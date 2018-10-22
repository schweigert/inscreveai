package html

func HtmlBytes(dom Dom) []byte {
	return []byte("<!DOCTYPE html>" + dom.HtmlSafe())
}
