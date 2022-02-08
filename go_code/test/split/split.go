package split

import "strings"

// Split 分割字符串
// Split("我爱你","爱")=>["我"，“你”]
func Split(s, sep string) (ret []string) {
	idx := strings.Index(s, sep)
	ret = make([]string, 0, strings.Count(s, sep)+1)
	for idx > -1 {
		ret = append(ret, s[:idx]) //	每一次都需要申请内存，故需要在函数开始是初始化
		s = s[idx+len(sep):]       // 中文在go语言中默认编码是三个字节 ，所以不能直接+1
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
