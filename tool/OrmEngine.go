package tool

import (
	"CloudRestaurant/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

//数据库连接
func OrmEngine(cfg *Config) error {
	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" +
		database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine(database.Driver, conn)
	if err != nil {
		return err
	}
	engine.ShowSQL(database.ShowSql)

	// 同步数据库表
	err = engine.Sync2(new(model.SmsCode), new(model.Member),
		new(model.FoodCategory), new(model.Shop), new(model.Service),
		new(model.ShopService), new(model.Goods))
	if err != nil {
		return err
	}

	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	// 插入shop信息
	InitShopDate()
	// 插入商户服务表信息
	InitServiceDate()
	InitShopServiceDate()
	InitGoodsData()
	return nil
}

func InitGoodsData() {
	Goods := []model.Goods{
		model.Goods{Id: 1, Name: "大虾", Description: "王婆大虾", SellCount: 201, Price: 108, OldPrice: 98, ShopId: 1},
		model.Goods{Id: 2, Name: "扯面", Description: "好吃", SellCount: 652, Price: 3, OldPrice: 2, ShopId: 1},
		model.Goods{Id: 3, Name: "烤鸭", Description: "附赠卷饼", SellCount: 198, Price: 128, OldPrice: 118, ShopId: 2},
		model.Goods{Id: 4, Name: "荷塘小炒", Description: "不辣", SellCount: 51, Price: 45, OldPrice: 38, ShopId: 2},
	}
	session := DbEngine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, good := range Goods {
		_, err := session.Insert(&good)
		if err != nil {
			fmt.Println(err.Error())
			session.Rollback()
			return
		}
	}
	err = session.Commit()
	if err != nil {
		fmt.Println(err)
	}
}

func InitShopServiceDate() {
	shopServices := []model.ShopService{
		model.ShopService{ShopId: 1, ServiceId: 1},
		model.ShopService{ShopId: 1, ServiceId: 2},
		model.ShopService{ShopId: 1, ServiceId: 3},
		model.ShopService{ShopId: 2, ServiceId: 1},
		model.ShopService{ShopId: 2, ServiceId: 3},
	}
	session := DbEngine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, shopService := range shopServices {
		_, err := session.Insert(&shopService)
		if err != nil {
			fmt.Println(err.Error())
			session.Rollback()
			return
		}
	}
	err = session.Commit()
	if err != nil {
		fmt.Println(err.Error())
	}
}

//初始化商家服务信息表
func InitServiceDate() {
	services := []model.Service{
		model.Service{
			Id:          1,
			Name:        "外面保",
			Description: "已经加入外卖宝计划",
			IconName:    "保",
			IconColor:   "999999",
		},
		model.Service{
			Id:          2,
			Name:        "准时达",
			Description: "准时必达",
			IconName:    "准",
			IconColor:   "57A9FF",
		},
		model.Service{
			Id:          3,
			Name:        "开发票",
			Description: "商家支持开发票",
			IconName:    "票",
			IconColor:   "999999",
		},
	}
	session := DbEngine.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, service := range services {
		_, err := session.Insert(&service)
		if err != nil {
			fmt.Println(err.Error())
			session.Rollback()
			return
		}
	}
	err = session.Commit()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// 初始化shop测试数据
func InitShopDate() {
	shops := []model.Shop{
		model.Shop{
			Id:                 1,
			Name:               "王婆大虾",
			PromotionInfo:      "欢迎光临",
			Address:            "北京市昌平区回龙观大街111号",
			Phone:              "18611823849",
			Status:             1,
			Longitude:          116.34,
			Latitude:           40.08,
			ImagePath:          "",
			IsNew:              true,
			IsPremiun:          true,
			Rating:             4.9,
			RatingCount:        125,
			RecentOrderNum:     50,
			MinimumOrderAmount: 20,
			DeliveryFee:        5,
			OpeningHours:       "09:00/22:00",
		},
		model.Shop{
			Id:                 2,
			Name:               "大鸭梨（回龙观店）",
			PromotionInfo:      "欢迎光临",
			Address:            "北京市昌平区回龙观大街112号",
			Phone:              "13888281234",
			Status:             1,
			Longitude:          116.34,
			Latitude:           40.07,
			ImagePath:          "",
			IsNew:              true,
			IsPremiun:          true,
			Rating:             4.4,
			RatingCount:        472,
			RecentOrderNum:     89,
			MinimumOrderAmount: 50,
			DeliveryFee:        5,
			OpeningHours:       "09:00/22:00",
		},
	}

	//事务
	session := DbEngine.NewSession()
	defer session.Close()
	//开始，执行，错误回滚，结束
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, shop := range shops {
		_, err := session.Insert(&shop)
		if err != nil {
			fmt.Println(err.Error())
			session.Rollback()
			return
		}
	}
	err = session.Commit()
	if err != nil {
		fmt.Println(err.Error())
	}
}
