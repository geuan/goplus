package Object

type User struct {
	Id int
	Sex  byte
	Name 	string
}

func setid(u *User)  {
	u.Id = 132
}
// 有选择性的对id进行赋值
func NewUser(fs ...UserAttrFunc) *User {
	//return new(User)
	//return &User{}
	u := new(User)
	UserAttrFuncs(fs).Apply(u)
	return u
}

func NewUser1(id int)  *User{
	u := new(User)
	u.Id = id
	return u
}

