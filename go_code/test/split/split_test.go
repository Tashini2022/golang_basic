package split

import (
	"reflect"
	"testing"
)

// 测试

func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	/*	单个测试用例
		got := Split("我爱你", "爱") // 实际得到的值
		want := []string{"我", "你"} // 期望得到的值
		if !reflect.DeepEqual(got, want) { // 因为slice不能比较直接，借助反射包中的方法比较
			// 调用Errorf会自动判断测试失败
			t.Errorf("want:%v\tgot:%v\t\n", want, got) // 失败之后打印结果
		}*/

	// 多个测试用例
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"simple": {input: "我爱你", sep: "爱", want: []string{"我", "你"}},
		// 注意开头有空字符串
		"chinese": {input: "一花一世界，一曲一永恒", sep: "一", want: []string{"", "花", "世界，", "曲", "永恒"}},
		"sep1":    {input: "a:b:c:d", sep: ":", want: []string{"a", "b", "c", "d"}},
		// 注意尾部有空字符串
		"sep2": {input: "how do you do", sep: "o", want: []string{"h", "w d", " y", "u d", ""}},
	}
	for name, ts := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(ts.input, ts.sep)
			if !reflect.DeepEqual(got, ts.want) { // 因为slice不能比较直接，借助反射包中的方法比较
				// 调用Errorf会自动判断测试失败
				t.Errorf("name:%#v faile, want:%#v\tgot:%#v\t\n", name, ts.want, got) // 失败之后打印结果
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("一花一世界，一曲一永恒", "一")
	}
}
