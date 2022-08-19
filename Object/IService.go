package Object

type IServcie interface {
	Save(data interface{}) IServcie
	List() IServcie
}