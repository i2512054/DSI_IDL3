package models

import (
    "time"

    "gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
    DB = db
}

// Robot model
type Robot struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    Nombre      string         `gorm:"size:255;not null" json:"nombre"`
    Descripcion string         `gorm:"type:text" json:"descripcion"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Category Model
type Category struct {
    ID              int             `gorm:"primaryKey" json:"id"`
    CategoryId      string          `gorm:"size:150;not null" json:"category_id"`
    Name            string          `gorm:"size:150;not null" json:"name"`
    CreatedAt       time.Time       `json:"created_at"`
    UpdatedAt       time.Time       `json:"updated_at"`
    DeletedAt       gorm.DeletedAt  `gorm:"index" json:"-"`
}

// Product Model
type Product struct {
    ID              int             `gorm:"primaryKey" json:"id"`
    CategoryId      string          `gorm:"size:150;not null" json:"category_id"`
    Name            string          `gorm:"size:150;not null" json:"name"`
    Description     string          `gorm:"size:255;not null" json:"description"`
    Stock           string          `gorm:"size:6;not null" json:"stock"`
    CreatedAt       time.Time       `json:"created_at"`
    UpdatedAt       time.Time       `json:"updated_at"`
    DeletedAt       gorm.DeletedAt  `gorm:"index" json:"-"`
}

type Customer struct {
    ID              int             `gorm:"primaryKey" json:"id"`
    FirstName       string          `gorm:"size:80;not null" json:"first_name"`
    LastName        string          `gorm:"size:80;not null" json:"last_name"`
    DocumentNumber  string          `gorm:"size:14;not null" json:"document_number"`
    Email           string          `gorm:"size:255;not null" json:"email"`
    CreatedAt       time.Time       `json:"created_at"`
    UpdatedAt       time.Time       `json:"updated_at"`
    DeletedAt       gorm.DeletedAt  `gorm:"index" json:"-"`
}

type User struct {
    ID              int             `gorm:"primaryKey" json:"id"`
    Name            string          `gorm:"size:80;not null" json:"name"`
    Password        string          `gorm:"size:30;not null" json:"password"`
    Email           string          `gorm:"size:255not null" json:"email"`
    IndStatus      string          `gorm:"size:2;not null" json:"ind_status"`
    CreatedAt       time.Time       `json:"created_at"`
    UpdatedAt       time.Time       `json:"updated_at"`
    DeletedAt       gorm.DeletedAt  `gorm:"index" json:"-"`
}