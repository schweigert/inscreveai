package html

func If(logic bool, d1 Dom, d2 Dom) Dom {
	if logic {
		return d1
	}

	return d2
}
