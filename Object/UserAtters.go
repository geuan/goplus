package Object

type UserAttrFunc func(*User)  // 设置User属性的函数类型

type UserAttrFuncs []UserAttrFunc

func (this UserAttrFuncs) Apply(u *User)  {
	for _,f := range this {
		f(u)
	}
}


func WithUserID(id int) UserAttrFunc {
	return func(u *User) {
		u.Id = id
	}
}

func WithUserName(name string) UserAttrFunc {
	return func(u *User) {
		u.Name = name
	}
}

func WithUserSex(sex byte) UserAttrFunc {
	return func(u *User) {
		u.Sex = sex
	}
}
