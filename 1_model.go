package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

//Model
type PrimeNumber struct {
	Id int `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Numbers int `gorm:"not null" form:"numbers" json:"numbers"`
	NameNumber  string `gorm:"not null" form:"name_number" json:"name_number"`
}

type EvenOddNumber struct {
	Id int `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	NumberEvens int `gorm:"not null" form:"number_evens" json:"number_evens"`
	NumberOdds int `gorm:"not null" form:"number_odds" json:"number_odds"`
	Total int `gorm:"not null" form:"total" json:"total"`
	NameTotal  string `gorm:"not null" form:"name_total" json:"name_total"`
}

type NominalAmount struct {
	Id int `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Nominal float64 `gorm:"not null" form:"nominal" json:"nominal"`
	AdditionalNumber int `gorm:"not null" form:"additional_number" json:"additional_number"`
	Total float64 `gorm:"not null" form:"total" json:"total"`
}

//Migration
func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("mysql", "root:@/tunaiku")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&PrimeNumber{}) {
		db.CreateTable(&PrimeNumber{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&PrimeNumber{})
	}
	if !db.HasTable(&EvenOddNumber{}) {
		db.CreateTable(&EvenOddNumber{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&EvenOddNumber{})
	}
	if !db.HasTable(&NominalAmount{}) {
		db.CreateTable(&NominalAmount{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&NominalAmount{})
	}
	return db
}