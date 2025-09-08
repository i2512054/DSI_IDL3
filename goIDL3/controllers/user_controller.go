package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "equisd.com/bichito/robotapp/models"
)

type UserController struct {
    DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
    return &UserController{DB: db}
}

//Método para Listar todas las Categorías
func (rc *UserController) Index(c *gin.Context) {
    var users []models.User
    rc.DB.Order("created_at desc").Find(&users)
    c.HTML(http.StatusOK, "user_index.html", gin.H{
        "users": users,
    })
}

func (rc *UserController) New(c *gin.Context) {
    c.HTML(http.StatusOK, "user_add.html", gin.H{})
}
//Método para Crear una Categoría
func (rc *UserController) Create(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	ind_status := c.PostForm("ind_status")
    user := models.User{Name: name, Password: password, Email: email, IndStatus: ind_status}
    if err := rc.DB.Create(&user).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "user_add.html", gin.H{
            "error": err.Error(),
        })
        return
    }
    c.Redirect(http.StatusFound, "/users")
}
//Método para Mostrar una Categoría por ID
func (rc *UserController) Show(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var user models.User
    if err := rc.DB.First(&user, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    c.HTML(http.StatusOK, "user_show.html", gin.H{
        "user": user,
    })
}

func (rc *UserController) Edit(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var user models.User
    if err := rc.DB.First(&user, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    c.HTML(http.StatusOK, "user_edit.html", gin.H{
        "user": user,
    })
}
//Método para Actualizar una Categoría
func (rc *UserController) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var user models.User
    if err := rc.DB.First(&user, id).Error; err != nil {
        c.String(http.StatusNotFound, "Not found")
        return
    }
    user.Name = c.PostForm("name")
    user.Password = c.PostForm("password")
	user.Email = c.PostForm("email")
	user.IndStatus = c.PostForm("ind_status")
    if err := rc.DB.Save(&user).Error; err != nil {
        c.HTML(http.StatusInternalServerError, "user_edit.html", gin.H{
            "error": err.Error(),
            "user": user,
        })
        return
    }
    c.Redirect(http.StatusFound, "/users")
}
//Método para Eliminar una Categoría
func (rc *UserController) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := rc.DB.Delete(&models.User{}, id).Error; err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
    c.Redirect(http.StatusFound, "/users")
}
