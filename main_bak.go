package main

// .  引入 context的上下文环境，不需要在调用方法前面加任何调用者

// 技术 “两害取其轻”

/*
type Person struct {
	name string
	age int
}

func (p Person) test()  {
	fmt.Println(p)

}
func main() {
	s := String.From("abcde")
	//fmt.Println(s)
	//fmt.Println(s,s.Len())

	s.Each(func(item string) {  //理解调用的时候执行each方法，然后打印参数中函数的参数
		fmt.Println(item)
	})

	//测试值接收者具体是个什么东西
	p1:=Person{name:"a",age:24}
	p1.test()

}
*/

/*
func main()  {
	u:=Object.NewUser(
		Object.WithUserID(123  ),
		)
	fmt.Println(u)

	u1 := Object.NewUser1(123)
	fmt.Println(u1)
}
*/

/*
func change(u Map.User)  {
	u["id"] = "404"
}

func main()  {
	u1:=Map.NewUser()
	//u["id"] = "101"
	//u["name"] = "chuan"
	//change(u)

	u1.With("id","101").
		With("name","chuan").
		With("sex","male").
		With("age",18)
	//fmt.Println(u)
	//u.String()

	u2:=Map.NewUser()
	u2.With("id","107").
		With("name","chuan").
		With("sex","male").
		With("age",20)

	u3:=Map.NewUser()
	u3.With("id","102").
		With("name","chuan").
		With("sex","male").
		With("age",16)

	users := []Map.User{}
	users= append(users,u1,u2,u3)
	sort.Slice(users, func(i, j int) bool {
		age1 := users[i]["age"].(int)  // 断言的功能，专门针对 interface
		age2 := users[j]["age"].(int)
		return age1 > age2
	})
	fmt.Println(users)
}
*/


//实现了链式调用（链式调用即可以不断的调用对象的方法）
//func SaveModel(servcie IServcie)  {
//	servcie.Save()
//}


/*
// IService 代言 UserService 和 ProdService
func main() {
	//var service1 IServcie = NewUserService()
	//service1.Save()
	//
	//var service2 IServcie = NewProdService()
	//service2.Save()

	//SaveModel(NewProdService()).List()
	//SaveModel(NewUserService()).List()

	//user := NewUser(
	//	WithUserName("chuan"),
	//	WithUserSex(1),
	//	)
	//NewUserService().Save(user)

}
*/
