package main

import (
	"fmt"
	api "github.com/Tambarie/movie-api/internal/adapters/api/movie"
	"github.com/Tambarie/movie-api/internal/adapters/repository/postgresdb"
	"github.com/Tambarie/movie-api/internal/adapters/repository/redisdb"
	"github.com/Tambarie/movie-api/internal/adapters/server"
	"github.com/Tambarie/movie-api/internal/core/helper"
	services "github.com/Tambarie/movie-api/internal/core/services/movie"
	"github.com/Tambarie/movie-api/internal/ports"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

// @title         Movie-API Service
// @version      1
// @description  Repo can be found here:https://github.com/Tambarie/movie-api

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8090
// @BasePath  /
// @securityDefinitions.basic  BasicAuth
func main() {

	helper.InitializeLogDir()

	var Addr string

	//variables gotten from movie-service.env files
	_, postgresdb_pass, postgres_database_url, service_address, service_port, _, postgresdb_host, postgresdb_mode, postgresdb_name, postgresdb_user, postgresdb_port, postgresdb_timezone, redis_host, redis_port, _ := helper.LoadConfig()

	if os.Getenv("REDIS_URL ") == "" {
		Addr = fmt.Sprintf("%s:%s", redis_host, redis_port)
	}

	dbRepository := ConnectToPostgres(postgresdb_user, postgresdb_pass, postgres_database_url, postgresdb_host, postgresdb_name, postgresdb_port, postgresdb_mode, postgresdb_timezone)

	redisRepository := ConnectToRedis(Addr)

	service := services.New(dbRepository)

	redisService := services.NewRedisService(redisRepository)
	handler := api.NewHTTPHandler(service, redisService)

	// Initiating the router
	router := gin.Default()

	server.DefineRouter(handler, router)

	fmt.Println("service running on " + service_address + ":" + service_port)
	helper.LogEvent("info", fmt.Sprintf("started movie service on "+service_address+":"+service_port+" in "+time.Since(time.Now()).String()))
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if PORT == ":" {
		PORT += "8090"
	}
	_ = router.Run(PORT)
}

// ConnectToPostgres Connecting to PostgresDB
func ConnectToPostgres(DBUser, DBPass, PostgresDBUrl, DBHost, DBName, DBPort, DBTimezone, DBMode string) ports.MovieRepository {
	repo := postgresdb.NewPostgresRepository(DBUser, DBPass, PostgresDBUrl, DBHost, DBName, DBPort, DBMode, DBTimezone)
	return services.New(repo)
}

// ConnectToRedis Connecting to RedisDB
func ConnectToRedis(addr string) ports.RedisRepository {
	redisRe := redisdb.NewRedisClient(addr, 0, 15)
	return services.NewRedisService(redisRe)
}
