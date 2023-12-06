package controller

import (
	"AeonBackend/entity" // Update with the correct import path
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProductController(g *gin.Engine, db *gorm.DB) {
	router := g.Group("/products")
	{
		router.GET("/category/:cate", getListProductByCategory(db))
		router.GET("/", getListProductHavePromotion(db))
		router.GET("/store/:id", getListProductByStoreID(db))
		router.POST("/", createProduct(db))
	}

}
func createProduct(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entity.ProductCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"new_customer": data.ProductID,
		})
	}
}
func getListProductByCategory(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		targetCategory := c.Param("category")
		var products []entity.Product
		if err := db.Table("product").Where("Category = ?", targetCategory).Find(&products).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

func getListProductHavePromotion(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var promotedProducts []entity.Product
		if err := db.Table("product").Joins("JOIN promoteproduct ON product.ProductID = promoteproduct.ProductID").Find(&promotedProducts).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, promotedProducts)
	}
}

func getListProductByStoreID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var products []entity.Product
		if err := db.Table("product").Joins("JOIN at ON product.ProductID = at.ProductID").
			Where("at.StoreID = ?", id).Find(&products).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}
