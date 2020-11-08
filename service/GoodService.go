package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
)

type GoodService struct {
}

func NewGoodService() *GoodService {
	return &GoodService{}
}

func (gs *GoodService) GetFoods(shop_id int64) []model.Goods {
	goodDao := dao.NewGoodDao()
	goods, err := goodDao.QueryFoods(shop_id)
	if err != nil {
		return nil
	}
	return goods
}
