package html

func If(logic bool, d1 Dom) Dom {
	if logic {
		return d1
	}

	return Tag("null", "")
}
