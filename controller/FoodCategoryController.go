package controller

import (
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

type FoodCategoryController struct {
}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {
	engine.GET("/api/food_category", fcc.foodCategory)
}

func (fcc *FoodCategoryController) foodCategory(ctx *gin.Context) {
	// service层获取食品种类
	foodCategoryService := &service.FoodCategoryService{}
	categories, err := foodCategoryService.Categories()
	if err != nil {
		tool.Failed(ctx, "食品种类获取失败")
		return
	}

	//格式转换，转换url的东西
	for _, category := range categories {
		if category.ImageUrl != "" {
			category.ImageUrl = tool.FileServerAddr() + "/" + category.ImageUrl
		}
	}
	tool.Success(ctx, categories)
}
