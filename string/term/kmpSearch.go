package term

/*
kmp字符串匹配算法
在原字符串str中寻找子字符串needle,如果存在则返回在str中的起始位置index，不存在则返回-1
*/

//kmp算法的思想是，利用统计相同前缀和后缀的字符串，然后在匹配失败时，直接跳到needle的相同前缀处，开始匹配，而不需要每次都从0开始匹配

func KmpSearch(str, needle string) int {
	sLen := len(str)
	nLen := len(needle)
	if sLen < 1 || nLen < 1 {
		return -1
	}

	i, j := 0, 0
	//构造next数组，next[j]存储的是needle字符串中，索引j之前的子字符串的最大相同前后缀的长度
	//例如:对于字符串"ababc",那么next[4]=2，在字符'c'之前的子串是"abab",最大相同前后缀为"ab",长度为2
	//思想就是如果找出相同的前后缀子串，那么下次匹配的时候，就直接跳到和后缀相同的前缀处开始匹配，而不是每次都从0开始匹配
	next := buildNext(needle)
	for i < sLen && j < nLen {
		if j == -1 || str[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	//循环结束后，j是nendle字符串的长度，i的索引是最后一个相同的字符再往后移动一位
	//所以用i-j刚好是str中最开始匹配needle的索引
	if j == nLen {
		return i - j
	}
	return -1
}

func buildNext(needle string) []int {
	nLen := len(needle)
	next := make([]int, nLen)
	//next[j]存储的是needle字符串中，索引j之前的子字符串的最大相同前后缀的长度
	//0位置之前没有字符串，默认为-1
	next[0] = -1
	k, j := -1, 0
	//这里是nLen-1，因为循环中会j++，防止越界
	for j < nLen-1 {
		//k为前缀索引，j为后缀索引
		if k == -1 || needle[k] == needle[j] {
			//由于next[j]存储的是索引j之前的子字符串的最大相同前后缀的长度
			//所以当needle[k] == needle[j]时，说明对于字符串needle,有索引0~k(前缀)和(j-k)~j(后缀)子字符串相等
			//所以索引j+1之前的子字符串的最大相同前后缀的长度为k+1
			k++
			j++
			next[j] = k
		} else {
			//如果上述没有找到needle[k] == needle[j],说明不存在索引0~k(前缀)和(j-k)~j(后缀)子字符串相等
			//那就需要寻找k'让索引0~k'(前缀)和(j-k')~j(后缀)子字符串相等
			//直接让k = next[k]是因为next[k]表示的是在索引k之前的子字符串的最大相同前后缀的长度
			//既然循环已经循环到k处了,只是此时needle[k] != needle[j],但是说明索引0~(k-1)(前缀)和(j-1-(k-1))~(j-1)(后缀)子字符串相等
			//如果找出0~(k-1)这个字符串内部相同前后缀长度，这个长度不就是next[k]吗？让k'=next[k]
			//因此索引0~(k'-1)(前缀)一定和(k-1-(k'-1))~(k-1)(后缀)相同,所以一定有0~(k'-1)(前缀)和(j-1-(k'-1))~(j-1)(后缀)相同
			//那么接下来就只需要继续判断needle[k'] 是否和 needle[j]相等了，如果相等，那么next[j+1] = k'+1
			//不相等，则继续递归k'' = next[k'],重复上述过程
			k = next[k]
		}
	}
	return next
}
