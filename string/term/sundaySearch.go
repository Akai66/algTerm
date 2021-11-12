package term

/*
sunday字符串匹配算法

*/

func SundaySearch(str, needle string) int {
	sLen := len(str)
	nLen := len(needle)
	if sLen < 1 || nLen < 1 {
		return -1
	}
	//构造move字典，move["a"]表示的是needle字符串中，最右边的字母"a"的位置到尾部的距离+1
	//换算过来就是needle字符串长度-最右边字母"a"在串中的索引(从0开始)
	move := make(map[uint8]int)
	for i := 0; i < nLen; i++ {
		//后面相同的字符会覆盖前面的，最终存储的就是最靠右边的字符需要移动步数
		move[needle[i]] = nLen - i
	}
	index := 0 //str字符串中匹配needle字符串的起始索引，最初从0位置开始匹配
	for index <= sLen-nLen {
		j := 0
		for str[index+j] == needle[j] {
			j++
			if j >= nLen {
				return index
			}
		}

		if index == sLen-nLen {
			//说明index已经遍历到最后一个可能的位置，且没有匹配上needle，直接break，否则会进入死循环
			break
		}

		if index < sLen-nLen {
			//必须判断index是否< sLen-nLen,否则可能会越界
			//只有<sLen-nLen,继续向前移动,才有意义
			ch := str[index+nLen]
			//判断该字符在needle字符串中是否存在
			//如果不存在，则index直接向前跳nLen+1步
			//如果存在，则index往前跳move[ch]步
			moveLen := nLen + 1
			if x, ok := move[ch]; ok {
				moveLen = x
			}
			index += moveLen
		}
	}
	return -1
}
