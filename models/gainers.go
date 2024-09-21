package models

import "gorm.io/gorm"

type Gainer struct {
    gorm.Model
    Name      string
    Symbol    string
    Price     float64
    Change    float64
    Percent   float64
}
