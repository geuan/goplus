package main

//type User struct {
//	UserId   int `name:"uid"`
//	UserName string
//}

/*
func main()  {
	//u := User{}
	//t := reflect.TypeOf(u)
	u := &User{
		UserId:10,
		UserName:"chuan",
	}
	t := reflect.ValueOf(u)
	if t.Kind() ==  reflect.Ptr{
		t=t.Elem()  // // 反射中使用 Elem()方法获取指针对应的值
	}

	//for i:=0;i<t.NumField();i++{
	//	if t.Field(i).Kind() == reflect.Int{
	//		t.Field(i).Set(reflect.ValueOf(10))
	//	}
	//	if t.Field(i).Kind() == reflect.String{
	//		t.Field(i).Set(reflect.ValueOf("list"))
	//	}
	//}


	values := []interface{}{101,"chuan"}
	for i:=0;i<t.NumField();i++ {
		if t.Field(i).Kind() == reflect.ValueOf(values[i]).Kind(){
			t.Field(i).Set(reflect.ValueOf(values[i]))
		}
	}

	fmt.Println(u)
}

*/

/*
func MapToStruct(m map[string]interface{},u interface{})  {
	v := reflect.ValueOf(u)
	if v.Kind() ==  reflect.Ptr{
		v = v.Elem()
		if v.Kind() != reflect.Struct{
			panic("must struct")
		}
		findFromMap := func(key string,nameTag string) interface{} {
			for k,v := range m {
				if k == key || k == nameTag {
					return v
				}
			}
			return nil
		}

		for i := 0;i<v.NumField();i++ {
			getValue := findFromMap(v.Type().Field(i).Name,v.Type().Field(i).Tag.Get("name"))
			if getValue != nil && reflect.ValueOf(getValue).Kind() == v.Field(i).Kind(){
				v.Field(i).Set(reflect.ValueOf(getValue))
			}
		}

	}else {
		panic("must ptr")
	}

}
func main()  {
	u := &User{}
	m := map[string]interface{}{
		"id":123,
		"uid":101,
		"UserName":"chuan",
		"age":19,
	}
	MapToStruct(m,u)
	fmt.Println(u)
}
*/

/*
// 限制协程执行的基本方法
func job(index int)  {
	time.Sleep(time.Millisecond*500)
	fmt.Printf("执行完毕，序号：…%d\n",index)
}

var pool chan struct{}
func main()  {
	MaxNum := 10
	pool = make(chan struct{},MaxNum)
	wg:=sync.WaitGroup{}
	for i:=0;i<100;i++ {
		pool <- struct{}{}  //到达最大长度阻塞
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			defer func() {
				<-pool
			}()
			job(index)
		}(i)
	}
	wg.Wait()

}

*/

/*
// 函数执行超时控制代码
// 1、业务过程放到协程
// 2、把业务结果塞入 channel

func job() chan string {
	ret := make(chan string)
	go func() {
		time.Sleep(time.Second * 5)
		ret <- "success"
	}()
	return ret
}

func run() (interface{}, error) {
	c := job()
	select {
	case r := <-c:
		return r, nil
	case <-time.After(time.Second * 3):
		return nil, fmt.Errorf("time out")
	}
}

func main() {
	fmt.Println(run())
}
*/



/*
func main() {
	var f func()   //对于接口而言，值是nil，但是类型不是nil
	var a *struct{}
	//fmt.Println(f, a)

	list := []interface{}{f, a}
	for _, item := range list {
		//if v, ok := item.(func()); ok && v == nil {
		//	fmt.Println("nil func")
		//	fmt.Println(v)
		//}
		//
		//if v, ok := item.(*struct{}); ok && v == nil {
		//	fmt.Println("nil struct")
		//	fmt.Println(v)
		if reflect.ValueOf(item).IsNil(){
			fmt.Println(nil)
		}
	}
}
*/

/*
// defer定义函数时的参数问题
func show(i *int)  {
	fmt.Println(*i)
}

func main()  {
	a := 1
	//defer  fmt.Println(a)  // defer后面语句的参数的值是在定义的时候被确定的，参数不是在执行的时候定义的，而是在定义的时候确定的
	defer show(&a)
	a ++
}
*/


/*
// defer里使用链式调用
type test struct {}

func NewTest() *test  {
	return &test{}
}

func (t *test) do(i int) *test {
	fmt.Println(i)
	return t
}

func main()  {
	t := NewTest()
	defer t.do(1).do(2)
	t.do(3)
}
*/

/*
// 循环执行 defer
func main()  {
	for i:=0;i<3;i++ {
		defer func() {
			fmt.Println(i)  // defer 会保存 i 的地址
		}()
	}
}
 // 3 3 3
*/

/*
// defer和panic哪个先执行，嵌套panic
func main()  {
	func(){
		defer func() {fmt.Println("打印前")}()
		defer func() {fmt.Println("打印中")}()
		defer func() {fmt.Println("打印后")}()
		panic("触发异常1")
	}()
	panic("触发异常2")
}
*/


/*
func main()  {
	n:=0
	locker := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i:=0;i<1000000;i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			defer locker.Unlock()
			locker.Lock()
			n++
		}()
	}
	wg.Wait()
	fmt.Println(n)
}
*/

/*
func main()  {
	var n int32 = 0
	wg := sync.WaitGroup{}
	for i:=0;i<1000000;i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			atomic.AddInt32(&n,1)
		}()
	}
	wg.Wait()
	fmt.Println(n)
}
*/


/*
// go的常见并发模式（1）基本模式
// 通过协程来执行并发任务。通过channel来进行通信
func main()  {
	c := make(chan int,3)
	for i:=0;i<3;i++ {
		go func(input int) {
			c <- input * 2
		}(i)
	}
	for i:=0;i<cap(c);i++ {
		fmt.Println(<-c)
	}
}
*/

/*
// go的常见并发模式（2）：生产者模式、多种写法
func Producer(out chan int)  {
	defer close(out)
	for i:=0;i<5;i++ {
		out <-i*2
		time.Sleep(time.Second*1)
	}
}

func Consumer(out chan  int)  (ret chan  struct{}) {
	ret = make(chan struct{})
	go func() {
		defer func() {
			ret <- struct{}{}
		}()
		for item := range out {
			fmt.Println(item)
		}
	}()
	return ret
}

func main()  {
	c := make(chan  int)
	go Producer(c)
	r := Consumer(c)
	<-r
}
*/


/*
// go的常见并发模式（3）：优胜劣汰模式
func job() int  {
	rand.Seed(time.Now().Unix())
	result := rand.Intn(5)
	time.Sleep(time.Second*time.Duration(result))  // 模拟延迟
	return result
}

// 应用场景在执行远程访问不可控的情况下，执行多协程，但是只返回响应最快的那一个结果
func main()  {
	c := make(chan int)
	for i:=0;i<=5;i++ {
		go func() {
			c <- job()
		}()
	}
	fmt.Println("最快用了",<-c,"秒")
}
*/


/*
// 协程为什么总是先输出倒数第一个
func main()  {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	for i:=0;i<5;i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%d ",i)
		}(i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("我要开始循环了")
	}()

	wg.Wait()
}
*/

/*
// 写一个带过期机制的kv获取map
var kv sync.Map
func Set(key string ,value interface{},expire time.Duration)   {
	kv.Store(key,value)
	time.AfterFunc(expire, func() {
		kv.Delete(key)
	})
}

func main()  {
	Set("id",101,time.Second*5)
	Set("name","zhangsan",time.Second*8)

	for {
		fmt.Println(kv.Load("id"))
		fmt.Println(kv.Load("name"))
		time.Sleep(time.Second*1)
	}
}

*/


/*
// 请用go实现一个简单的set
type Empty struct {}

type Set map[interface{}]Empty

func (s Set) Add(vs ...interface{})  Set {
	for _,v := range vs {
		s[v] = Empty{}
	}
	return s
}

func NewSet()  Set {
	return make(map[interface{}]Empty)
}

func main()  {
	 set := NewSet().Add(1,2,3,4,5,6,3,2)
	 fmt.Println(set)
}

*/


/*
func test() []int  {
	a := []int{1,2,3}
	a[1] = 4
	return a
}
*/



/*
// go语言中的单例模式
type WebConfig struct {
	Port int
}
var cc *WebConfig
//var mu sync.Mutex
//func GetConfig() *WebConfig  {
//	if cc == nil {
//		mu.Lock()
//		defer mu.Unlock()
//		cc =  &WebConfig{Port:80}
//	}
//	return cc
//}

var once sync.Once
func GetConfig() *WebConfig  {
	once.Do(func() {
		cc = &WebConfig{Port:8080}
	})
	return cc
}

func main()  {
	c1:= GetConfig()
	c2:= GetConfig()
	c3:=GetConfig()
	fmt.Println(c1==c2,c2==c3)
	fmt.Printf("%p,%p,%p",c1,c2,c3)
}
*/




/*
// go语言中的抽象工厂模式
type User interface {
	GetRole() string
}

type Member struct {}
func (m *Member) GetRole() string  {
	return  "会员用户"
}

type Admin struct {}
func (a *Admin) GetRole() string  {
	return  "后台管理用户"
}

type AbstractFactory interface {
	CreateUser() User
}
type MemberFactory struct {}

func (m *MemberFactory) CreateUser() User  {
	return &Member{}
}

func (a *Admin) CreateUser() User  {
	return &Admin{}
}

func main()  {
	var fact AbstractFactory=new(MemberFactory)
	fmt.Println(fact.CreateUser().GetRole())
}
*/



/*
// go语言中的装饰器模式
// 一般可以使用在校验参数
func CheckLogin(f http.HandlerFunc) http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Query().Get("token") == ""{
			_, _ = writer.Write([]byte("token error"))
		}else {
			f(writer,request)
		}
	}
}

func index(writer http.ResponseWriter,request *http.Request)  {
	_, _ = writer.Write([]byte("index"))
}

func main()  {
	http.HandleFunc("/",CheckLogin(index))
	_ = http.ListenAndServe(":8080", nil)
}
*/





/*
var p *sync.Pool

type User struct {
	Name 	string
}

func main()  {
	p = &sync.Pool{
		New: func() interface{} {
			log.Println("create user")
			return  &User{Name:"zhangsan"}
		},
	}
	u1 := p.Get().(*User)
	fmt.Println(u1)
	u1.Name = "lisi"
	p.Put(u1)
	runtime.GC()
	u2 := p.Get()
	fmt.Println(u2)
}
*/
























