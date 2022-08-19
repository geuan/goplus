package String

import "fmt"

type String string


// go语言中的类方法
func (s String) Len() int  {
	return len(s)
}

func (s String) Each(f func(item string))  {
	for i := 0;i<len(s);i++ {
		f(fmt.Sprintf("%c",s[i]))
	}
}

func From(s string) String  {
	return String(s)
}

func FromInt(n int)  String {
	//return  String(strconv.Itoa(n))
	return String(fmt.Sprintf("%d",n))
}