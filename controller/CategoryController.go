package controller

import (
	"log"
	"strconv"

	"example.com/ginEssential/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"example.com/ginEssential/model"
	"example.com/ginEssential/response"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{DB: db}
}

func (cat CategoryController) Create(c *gin.Context) {
	var requestCategory model.Category
	c.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(c, nil, "数据验证错误，分类名称必填")
		return
	}

	cat.DB.Create(&requestCategory)

	response.Success(c, gin.H{"category": requestCategory}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	//绑定body中的参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory model.Category
	if c.DB.First(&updateCategory, categoryId).RecordNotFound() {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	//更新分类
	//map
	//struct
	//name value
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")

}

func (c CategoryController) Show(ctx *gin.Context) {

	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category
	if c.DB.First(&category, categoryId).RecordNotFound() {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	//更新分类
	//map
	//struct
	//name value
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.DB.Delete(model.Category{}, categoryId); err != nil {
		response.Fail(ctx, nil, "删除失败请重试")
		log.Printf(err.Error.Error())
		return
	}

	response.Success(ctx, nil, "删除成功")
}
