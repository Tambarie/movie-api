package postgresdb

import (
	"fmt"
	domain "github.com/Tambarie/movie-api/internal/core/domain/movie"
	"github.com/Tambarie/movie-api/internal/core/helper"
	"github.com/Tambarie/movie-api/internal/ports"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type PostgresRepository struct {
	DB *gorm.DB
}

// NewPostgresClient Initializing Postgres Client
func NewPostgresClient(DBUser, DBPass, PostgresDBURL, DBHost, DBName, DBPort, DBTimezone, DBMode string) *gorm.DB {
	var dsn string

	dsn = os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf("host=%v user=%v dbname=%v port=%v sslmode=%v TimeZone=%v", DBHost, DBUser, DBName, DBPort, DBMode, DBTimezone)
	}
	log.Println("am in postgres db")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(helper.PrintErrorMessage("500", "failed to connect to database"))
	}
	err = db.AutoMigrate(&domain.Comment{})
	if err != nil {
		panic(err)
	}
	return db
}

// NewPostgresRepository  Initializing Postgres repository
func NewPostgresRepository(DBUser, DBPass, PostgresDBUrl, DBHost, DBName, DBPort, DBTimezone, DBMode string) ports.MovieRepository {
	db := NewPostgresClient(DBUser, DBPass, PostgresDBUrl, DBHost, DBName, DBPort, DBTimezone, DBMode)
	return &PostgresRepository{
		DB: db,
	}
}
