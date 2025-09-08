package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "equisd.com/bichito/robotapp/models"
)

type CustomerController struct {
    DB *gorm.DB
}

func NewCustomerController(db *gorm.DB) *CustomerController {
    return &CustomerController{DB: db}
}

//Método para Listar todas las Categorías
func (rc *CustomerController) Index(c *gin.Context) {
    var customers []models.Customer
    rc.DB.Order("created_at desc").Find(&customers)
    c.HTML(http.StatusOK, "customer_index.html", gin.H{
        "customers": customers,
    })
}

func (rc *CustomerController) New(c *gin.Context) {
    c.HTML(http.StatusOK, "customer_add.html", gin.H{})
}
//Método para Crear una Categoría
func (rc *CustomerController) Create(c *gin.Context) {
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	documentNumber := c.PostForm("document_number")
	email := c.PostForm("email")
    customer := models.Customer{FirstName: firstName, LastName: lastName, DocumentNumber: documentNumber, Email: email}
    if err := rc.DB.Create(&customer).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "customer_add.html", gin.H{
            "error": err.Error(),
        })
        return
    }
    c.Redirect(http.StatusFound, "/customers")
}
//Método para Mostrar una Categoría por ID
func (rc *CustomerController) Show(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var customer models.Customer
    if err := rc.DB.First(&customer, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    c.HTML(http.StatusOK, "customer_show.html", gin.H{
        "customer": customer,
    })
}

func (rc *CustomerController) Edit(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var customer models.Customer
    if err := rc.DB.First(&customer, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    c.HTML(http.StatusOK, "customer_edit.html", gin.H{
        "customer": customer,
    })
}
//Método para Actualizar una Categoría
func (rc *CustomerController) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var customer models.Customer
    if err := rc.DB.First(&customer, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    customer.FirstName = c.PostForm("first_name")
    customer.LastName = c.PostForm("last_name")
	customer.DocumentNumber = c.PostForm("document_number")
	customer.Email = c.PostForm("email")
    if err := rc.DB.Save(&customer).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "customer_edit.html", gin.H{
            "error": err.Error(),
            "customer": customer,
        })
        return
    }
    c.Redirect(http.StatusFound, "/customers")
}
//Método para Eliminar una Categoría
func (rc *CustomerController) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := rc.DB.Delete(&models.Customer{}, id).Error; err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
    c.Redirect(http.StatusFound, "/customers")
}
