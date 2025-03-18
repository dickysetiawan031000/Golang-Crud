package productcontroller

import (
	"github.com/dickysetiawan031000/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func Show(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func Create(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBind(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})

}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	// Cari produk berdasarkan ID
	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	// Bind data yang dikirim oleh user
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Update produk
	if err := models.DB.Model(&product).Where("id = ?", id).Updates(product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func Delete(c *gin.Context) {
	idParam := c.Param("id")

	// Convert ID ke uint
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return
	}

	// Cek apakah produk ada
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	// Hapus produk
	if err := models.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
