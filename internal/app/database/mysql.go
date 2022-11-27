package database

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Function to connect golang to online mysql server at freesqldatabase.com
func NewConnDatabase() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	dsn := "sql6580832:zwTxNEJj5r@tcp(sql6.freesqldatabase.com:3306)/sql6580832?charset=utf8mb4&parseTime=True&loc=Local"
	/* dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DATABASE"),
	) */
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
