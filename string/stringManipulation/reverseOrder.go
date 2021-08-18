package stringManipulation

func Reverse(str string) (result string) {

	for i := len(str) - 1; i >= 0; i-- {
		result = result + string(str[i])
	} 

	return
}