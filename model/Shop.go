package model

type Shop struct {
	Id            int64  `xorm:"pk autoincr" json:"id"`
	Name          string `xorm:"varchar(12)" json:"name"`
	PromotionInfo string `xorm:"varchar(30)" json:"promotion_info"`
	Address       string `xorm:"varchar(100)" json:"address"`
	Phone         string `xorm:"varchar(11)" json:"phone"`
	Status        int    `xorm:"tinyint" json:"status"`
	// 经纬度
	Longitude float64 `xorm:"" json:"longitude"`
	Latitude  float64 `xorm:"" json:"latitude"`
	ImagePath string  `xorm:"varchar(255)" json:"image_path"`
	IsNew     bool    `xorm:"bool" json:"is_new"`
	IsPremiun bool    `xorm:"bool" json:"is_premiun"`
	//店铺评分，评分总数，当前订单数量
	Rating         float32 `xorm:"float" json:"rating"`
	RatingCount    int64   `xorm:"int" json:"rating_count"`
	RecentOrderNum int64   `xorm:"int" json:"recent_order_num"`
	//配送起送价，配送费
	MinimumOrderAmount int32 `xorm:"int" json:"minimum_order_amount"`
	DeliveryFee        int32 `xorm:"int" json:"delivery_fee"`
	//营业时间
	OpeningHours string    `xorm:"varchar(20)" json:"opening_hours"`
	Supports     []Service `xorm:""`
}
