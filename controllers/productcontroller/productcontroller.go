package productcontroller

import (
	"github.com/dickysetiawan031000/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {

	// Mendeklarasikan slice untuk menyimpan data produk
	var products []models.Product

	// Mengambil semua data produk dari database menggunakan GORM
	models.DB.Find(&products)

	// Mengembalikan response JSON dengan status 200 OK, berisi data produk
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func Show(c *gin.Context) {
	// Mendeklarasikan variabel product untuk menyimpan data produk yang akan dicari
	var product models.Product

	// Mengambil parameter "id" dari URL request
	id := c.Param("id")

	// Mencari produk berdasarkan id menggunakan GORM
	if err := models.DB.First(&product, id).Error; err != nil {
		// Menangani error jika produk tidak ditemukan atau terjadi kesalahan lain
		switch err {
		case gorm.ErrRecordNotFound:
			// Jika produk tidak ditemukan, kembalikan response 404 Not Found
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return
		default:
			// Jika terjadi error lain (contoh: masalah koneksi database), kembalikan response 500 Internal Server Error
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	// Jika produk ditemukan, kembalikan response 200 OK dengan data produk dalam format JSON
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
