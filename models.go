package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// var db *gorm.DB
// var err error

// Genre of the book
type Genre struct {
	gorm.Model
	Name   string
	BookID uint
}

// Book
type Book struct {
	gorm.Model
	Title            string
	Author           string
	YearPublished    int
	ISBN             string
	Language         string
	Genres           []Genre
	Keyword          string
	YearBought       string
	PricePaid        string
	PlaceWhereBought Place
}

// Place where the book was bought (geographical location, e.g city)
type Place struct {
	gorm.Model
	Name    string
	Country string
	Lat     float64 `gorm:"type:decimal(10,8)"`
	Lng     float64 `gorm:"type:decimal(11,8)"`
	BookID  uint
}
