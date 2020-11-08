package dao

import (
	"CloudRestaurant/model"
	"CloudRestaurant/tool"
	"fmt"
)

type GoodDao struct {
	*tool.Orm
}

func NewGoodDao() *GoodDao {
	return &GoodDao{tool.DbEngine}
}

func (gd *GoodDao) QueryFoods(shop_id int64) ([]model.Goods, error) {
	var goods []model.Goods
	err := gd.Orm.Where(" shop_id = ? ", shop_id).Find(&goods)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return goods, nil
}
