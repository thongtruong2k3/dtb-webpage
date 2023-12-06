package controller

import (
	"AeonBackend/entity"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getRandomCashierID(db *gorm.DB, storeID int) (int, error) {
	// Seed the random number generator to get different values on each run
	rand.Seed(time.Now().UnixNano())

	var cashier entity.Cashier
	if err := db.Table("cashier").
		Where("StoreID = ?", storeID).
		Order("random()"). // Use the random function specific to your database (e.g., RAND() for MySQL, RANDOM() for SQLite)
		Limit(1).
		Scan(&cashier).
		Error; err != nil {
		return 0, err
	}

	return cashier.EmployeeID, nil
}
func NewTransactionController(g *gin.Engine, db *gorm.DB) {
	router := g.Group("/transaction")
	{
		router.POST("/", createBill(db))
		router.GET("/:id", getBillByID(db))
		router.POST("/items/", createInclude(db))
		router.GET("/items/:id", getAllItemsByBillID(db))
	}
}
func getAllItemsByBillID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var products []entity.Product
		if err := db.Table("include").Joins("JOIN at ON product.ProductID = include.ProductID").
			Where("include.TransactionID = ?", id).Find(&products).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}
func createInclude(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data entity.IncludeCreation
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
			"new_include": data.TransactionID,
		})
	}
}
func createBill(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data entity.BillCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		cashierID, err := getRandomCashierID(db, data.StoreID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to get random cashier ID"})
			return
		}
		bill := entity.Bill{
			TransactionID: data.TransactionID,
			PaymentMethod: data.PaymentMethod,
			DateAndTime:   data.DateAndTime,
			CustomerID:    data.CustomerID,
			StoreID:       data.StoreID,
			CashierID:     cashierID,
		}
		if err := db.Create(&bill).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"new_bill": bill.TransactionID,
		})
	}
}
func getBillByID(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var bill entity.Bill
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Table("bill").Where("TransactionID = ?", id).First(&bill).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, bill)
	}
}
