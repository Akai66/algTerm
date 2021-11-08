package term

/*
字符串转整型，不合法则返回0
1.字符串是否为空
2.正负符号处理
3.字符范围在'0'~'9'之间
4.整型溢出问题
5.返回有效性标识
*/

func Str2int(str string) (ret int, isValid bool) {
	length := len(str)
	if length < 1 || (length == 1 && (str == "+" || str == "-")) {
		return
	}
	isNegative := false
	if str[0] == '-' {
		isNegative = true
	}
	for i := 0; i < length; i++ {
		if i == 0 && (str[i] == '+' || str[i] == '-') {
			continue
		}
		if str[i] < '0' || str[i] > '9' {
			ret = 0
			return
		}
		//判断是否溢出,int32有符号整型的范围是2^-31 ~ 2^31-1，其中2^31=2147483648
		if ret > 214748364 || (ret == 214748364 && str[i] > '7') {
			ret = 0
			return
		}
		ret = ret*10 + int(str[i]-'0')
	}
	if isNegative {
		ret = -1 * ret
	}
	isValid = true
	return
}
