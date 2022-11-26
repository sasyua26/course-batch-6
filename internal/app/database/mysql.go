package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Function to connect golang to online mysql server at freesqldatabase.com
func NewConnDatabase() *gorm.DB {
	dsn := "sql6580832:zwTxNEJj5r@tcp(sql6.freesqldatabase.com:3306)/sql6580832?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
