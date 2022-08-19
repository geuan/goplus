package Object

import "log"

type ProdService struct {

}

func NewProdService() *ProdService {
	return &ProdService{}
}

func (p *ProdService)  Save(data interface{})  IServcie{
	log.Println("商品保存入库成功")
	return p
}

func (p *ProdService) List()  IServcie {
	log.Println("商品列表获取")
	return p
}