package utils

func reverseWords(s string) (ans string) {
	str := []byte(s)
	ouco := make([]byte, len(str))
	length := len(str)
	for i := length - 1; i >= 0; i-- {
		if str[i] == ' ' {
			str = str[:i]
		}
		if str[i] != ' ' {
			break
		}
	}
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			str = str[:i]
		}
		if str[i] != ' ' {
			break
		}
	}
	mark := len(str)
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' {
			ouco = append(ouco, str[i+1:mark]...)
			ouco = append(ouco, ' ')
			mark = i
		}
	}
	return string(ouco)
}
