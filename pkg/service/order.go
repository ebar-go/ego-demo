package service

import (
	"ego-demo/pkg/dao"
	"ego-demo/pkg/entity"
	"github.com/ebar-go/ego/app"
	"github.com/ebar-go/ego/errors"
)

type orderService struct {

}

func Order() *orderService  {
	return new(orderService)
}

func (service *orderService) Get(id int) (*entity.OrderEntity, error){
	items, err := dao.Order(app.DB()).GetById(id)
	if err != nil {
		return nil, errors.Sprintf(1001, "获取数据失败:%v" ,err)
	}

	return items, nil
}
