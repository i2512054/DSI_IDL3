package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "equisd.com/bichito/robotapp/models"
)

type CategoryController struct {
    DB *gorm.DB
}

func NewCategoryController(db *gorm.DB) *CategoryController {
    return &CategoryController{DB: db}
}

//Método para Listar todas las Categorías
func (rc *CategoryController) Index(c *gin.Context) {
    var categories []models.Category
    rc.DB.Order("created_at desc").Find(&categories)
    c.HTML(http.StatusOK, "category_index.html", gin.H{
        "categories": categories,
    })
}

func (rc *CategoryController) New(c *gin.Context) {
    c.HTML(http.StatusOK, "category_add.html", gin.H{})
}
//Método para Crear una Categoría
func (rc *CategoryController) Create(c *gin.Context) {
	categoryId := c.PostForm("category_id")
	name := c.PostForm("name")
    category := models.Category{CategoryId: categoryId, Name: name}
    if err := rc.DB.Create(&category).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "category_add.html", gin.H{
            "error": err.Error(),
        })
        return
    }
    c.Redirect(http.StatusFound, "/categories")
}
//Método para Mostrar una Categoría por ID
func (rc *CategoryController) Show(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var category models.Category
    if err := rc.DB.First(&category, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    c.HTML(http.StatusOK, "category_show.html", gin.H{
        "category": category,
    })
}

func (rc *CategoryController) Edit(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var category models.Category
    if err := rc.DB.First(&category, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    c.HTML(http.StatusOK, "category_edit.html", gin.H{
        "category": category,
    })
}
//Método para Actualizar una Categoría
func (rc *CategoryController) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var category models.Category
    if err := rc.DB.First(&category, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    category.CategoryId = c.PostForm("category_id")
    category.Name = c.PostForm("name")
    if err := rc.DB.Save(&category).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "category_edit.html", gin.H{
            "error": err.Error(),
            "category": category,
        })
        return
    }
    c.Redirect(http.StatusFound, "/categories")
}
//Método para Eliminar una Categoría
func (rc *CategoryController) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := rc.DB.Delete(&models.Category{}, id).Error; err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
    c.Redirect(http.StatusFound, "/categories")
}
