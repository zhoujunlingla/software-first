// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Delivery struct {
	gorm.Model
	TrackingNumber string
	Status         string
	Name           string
}

var db *gorm.DB

func main() {
	// 连接到MySQL数据库
	dsn := "root:Zjl7758258@tcp(localhost:3306)/kuai_system?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// 迁移数据模型到数据库
	db.AutoMigrate(&Delivery{})

	// 创建Gin实例
	r := gin.Default()

	// 设置静态文件目录
	r.Static("/static", "./static")

	// 设置前端页面路由
	r.LoadHTMLFiles("templates/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 处理录入快递信息的请求
	r.GET("/api/deliveries", func(c *gin.Context) {
		name := c.Query("name")

		var deliveries []Delivery
		result := db.Where("name = ?", name).Find(&deliveries)

		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "无法查找到信息",
			})
			return
		}

		if len(deliveries) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "快递查找不到",
			})
			return
		}

		c.JSON(http.StatusOK, deliveries)
	})
	r.POST("/api/deliveries", func(c *gin.Context) {
		name := c.Query("name")

		var deliveries []Delivery
		result := db.Where("name = ?", name).Find(&deliveries)

		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "无法查找到信息",
			})
			return
		}

		if len(deliveries) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "快递查找不到",
			})
			return
		}

		c.JSON(http.StatusOK, deliveries)
	})

	// 处理查询快递信息的请求
	// 处理根据名称查询快递信息的请求
	// 处理根据运单号查询快递信息的请求
	// 处理根据运单号查询快递信息的请求
	r.GET("/api/deliveries/:trackingNumber", func(c *gin.Context) {
		trackingNumber := c.Param("trackingNumber")

		var deliveries []Delivery
		result := db.Where("tracking_number = ?", trackingNumber).Find(&deliveries)

		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "无法查找到信息",
			})
			return
		}

		if len(deliveries) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "快递查找不到",
			})
			return
		}

		c.JSON(http.StatusOK, deliveries)
	})

	// 启动Gin服务器
	r.Run(":8000")
}
