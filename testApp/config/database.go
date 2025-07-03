package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// func Koneksi() {
// 	db, err := sql.Open("mysql", "root:@/dbgo")
// 	if err != nil {
// 		panic(err)
// 	}

// 	log.Println("Koneksi berhasil")
// 	DB = db
// }

func Koneksi() {
	dsn := "root:@tcp(127.0.0.1:3306)/dbgo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal Koneksi")
	}

	log.Println("Koneksi berhasil")
	DB = db
}
