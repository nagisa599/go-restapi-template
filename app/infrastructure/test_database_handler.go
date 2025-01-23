package infrastructure

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


func NewTestDatabase() *gorm.DB {
	dsn := "root:@tcp(localhost:3307)/mysql_template?parseTime=true&loc=Local&charset=utf8"
	fmt.Print(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Print(dsn)
	if err != nil {
		log.Fatal("Failed to connect to database",err)
	}
	if os.Getenv("ENV") == "development" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} 

	return db}

// func (h *DatabaseHandler) Close() {
// 	sqlDB, _ := h.DB.DB()
// 	sqlDB.Close()
// }

// func (h *DatabaseHandler) Conn(ctx context.Context) *gorm.DB {
// 	return h.DB.WithContext(ctx)
// }