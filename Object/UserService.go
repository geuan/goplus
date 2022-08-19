package Object

import "log"

type UserService struct {
	
}

func NewUserService() *UserService  {
	return &UserService{}
}

// 保存入库
// 断言； 针对 interface 进行 类型断言
func (u *UserService) Save(data interface{})  IServcie {
	if user,ok := data.(*User);ok{
		log.Printf("%v",user.Name)
		log.Println("用户保存入库成功")
	}else {
		log.Fatal("用户参数错误")
	}
	return u
}

func (u *UserService) List()  IServcie {
	log.Println("用户列表获取")
	return u
}

