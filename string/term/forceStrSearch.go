package term

/*
暴力字符串匹配算法

每次都从0开始匹配，时间复杂度O(m*n)
*/

func ForceStrSearch(str, needle string) int {
	sLen := len(str)
	nLen := len(needle)
	if sLen < 1 || nLen < 1 {
		return -1
	}
	i, j := 0, 0
	for i < sLen && j < nLen {
		if str[i] == needle[j] {
			//相等时，都往前走一步
			i++
			j++
		} else {
			//不相等时，将i相对i上一次的初始位置，往前走一步，j从0开始
			i = i - j + 1
			j = 0
		}
	}
	if j == nLen {
		return i - j
	}
	return -1
}
