package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"fmt"
	"strconv"
)

type ShopService struct {
}

func (ss *ShopService) GetService(shopId int64) []model.Service {
	shopDao := dao.NewShopDao()
	return shopDao.QueryServiceByShopId(shopId)
}

func (ss *ShopService) ShopList(long, lat string, keyword string) []model.Shop {
	longitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		return nil
	}
	shopDao := dao.NewShopDao()
	fmt.Println(keyword)
	return shopDao.QueryShops(longitude, latitude, keyword)
}
