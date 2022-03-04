package postgresdb

import (
	"fmt"
	"github.com/Tambarie/movie-api/internal/core/helper"
	"github.com/Tambarie/movie-api/internal/ports"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgresRepository struct {
	DB *gorm.DB
}

func NewPostgresClient(DBUser, DBPass, DBHost, DBName, DBPort, DBTimezone, DBMode string) *gorm.DB {
	var dsn string
	//fmt.Println(DBUser, DBPass, DBHost, DBName, DBPort, DBTimezone, DBMode)
	dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", DBHost, DBUser, DBPass, DBName, DBPort, DBMode, DBTimezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//fmt.Println(*db)
	if err != nil {
		log.Fatal(helper.PrintErrorMessage("500", "failed to connect to database"))
	}
	return db
}

func NewPostgresRepository(DBUser, DBPass, DBHost, DBName, DBPort, DBTimezone, DBMode string) ports.MovieRepository {
	db := NewPostgresClient(DBUser, DBPass, DBHost, DBName, DBPort, DBTimezone, DBMode)
	return &PostgresRepository{
		DB: db,
	}
}
