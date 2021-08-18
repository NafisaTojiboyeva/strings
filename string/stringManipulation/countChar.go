package stringManipulation

func CountChar(str string, char byte) (n int) {
	for _, v := range str {
		if byte(v) == char {
			n++
		} 
	}

	return
}