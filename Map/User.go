package Map

import (
	"fmt"
	"sort"
)

type User map[string]interface{}

func NewUser() User  {
	return make(map[string]interface{})
}

func (u User) With(k string,v interface{}) User  {
	u[k] = v
	return u
}

func (u User) Fileds() []string  {
	var kSlice []string
	for k,_ := range u {
		kSlice = append(kSlice, k)
	}
	sort.Strings(kSlice)
	return kSlice
}

func (u User) String() string {   //为啥打印的时候会自动调用该方法？
	str := ""
	for index,k := range u.Fileds() {
		str+=fmt.Sprintf("%d、%v-->%v\n",index+1,k,u[k])
	}
	return str

}