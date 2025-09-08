package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "equisd.com/bichito/robotapp/models"
)

type ProductController struct {
    DB *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
    return &ProductController{DB: db}
}

//Método para Listar todas las Categorías
func (rc *ProductController) Index(c *gin.Context) {
    var products []models.Product
    rc.DB.Order("created_at desc").Find(&products)
    c.HTML(http.StatusOK, "product_index.html", gin.H{
        "products": products,
    })
}

func (rc *ProductController) New(c *gin.Context) {
    c.HTML(http.StatusOK, "product_add.html", gin.H{})
}
//Método para Crear una Categoría
func (rc *ProductController) Create(c *gin.Context) {
	categoryId := c.PostForm("category_id")
	name := c.PostForm("name")
	description := c.PostForm("description")
	stock := c.PostForm("stock")
    product := models.Product{CategoryId: categoryId, Name: name, Description: description, Stock: stock}
    if err := rc.DB.Create(&product).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "product_add.html", gin.H{
            "error": err.Error(),
        })
        return
    }
    c.Redirect(http.StatusFound, "/products")
}
//Método para Mostrar una Categoría por ID
func (rc *ProductController) Show(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var product models.Product
    if err := rc.DB.First(&product, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    c.HTML(http.StatusOK, "product_show.html", gin.H{
        "product": product,
    })
}

func (rc *ProductController) Edit(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var product models.Product
    if err := rc.DB.First(&product, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    c.HTML(http.StatusOK, "product_edit.html", gin.H{
        "product": product,
    })
}
//Método para Actualizar una Categoría
func (rc *ProductController) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var product models.Product
    if err := rc.DB.First(&product, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    product.CategoryId = c.PostForm("category_id")
    product.Name = c.PostForm("name")
	product.Description = c.PostForm("description");
	product.Stock = c.PostForm("stock");
    if err := rc.DB.Save(&product).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "product_edit.html", gin.H{
            "error": err.Error(),
            "product": product,
        })
        return
    }
    c.Redirect(http.StatusFound, "/products")
}
//Método para Eliminar una Categoría
func (rc *ProductController) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := rc.DB.Delete(&models.Product{}, id).Error; err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
    c.Redirect(http.StatusFound, "/products")
}
