package main

import (
	"fmt"
	api "github.com/Tambarie/movie-api/internal/adapters/api/movie"
	"github.com/Tambarie/movie-api/internal/adapters/repository/postgresdb"
	"github.com/Tambarie/movie-api/internal/adapters/repository/redisdb"
	"github.com/Tambarie/movie-api/internal/core/helper"
	services "github.com/Tambarie/movie-api/internal/core/services/movie"
	"github.com/Tambarie/movie-api/internal/core/shared"
	"github.com/Tambarie/movie-api/internal/ports"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	helper.InitializeLogDir()
	_, postgresdb_pass, service_address, service_port, _, postgresdb_host, postgresdb_mode, postgresdb_name, postgresdb_user, postgresdb_port, postgresdb_timezone, redis_host, redis_port, _ := helper.LoadConfig()
	Addr := fmt.Sprintf("%s:%s", redis_host, redis_port)

	dbRepository := ConnectToPostgres(postgresdb_user, postgresdb_pass, postgresdb_host, postgresdb_name, postgresdb_port, postgresdb_mode, postgresdb_timezone)
	log.Println(postgresdb_name)
	redisRepository := ConnectToRedis(Addr)
	service := services.New(dbRepository)
	redisService := services.NewRedisService(redisRepository)
	handler := api.NewHTTPHandler(service, redisService)

	router := gin.Default()
	//
	router.Use(helper.LogRequest)

	router.GET("/api/movies", handler.GetMovies())
	router.GET("/api/movies/:movieID/characters", handler.GetMoviesCharacters())
	router.POST("/api/movies/:movieID/comments", handler.AddCommentToMovies())
	router.GET("/api/movies/:movieID/comments", handler.GetCommentsInMovie())

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404,
			helper.PrintErrorMessage("404", shared.NoResourceFound))
	})

	fmt.Println("service running on " + service_address + ":" + service_port)
	helper.LogEvent("info", fmt.Sprintf("started movie service on "+service_address+":"+service_port+" in "+time.Since(time.Now()).String()))
	_ = router.Run(":" + service_port)
}

func ConnectToPostgres(DBUser, DBPass, DBHost, DBName, DBPort, DBTimezone, DBMode string) ports.MovieRepository {
	repo := postgresdb.NewPostgresRepository(DBUser, DBPass, DBHost, DBName, DBPort, DBMode, DBTimezone)
	return services.New(repo)
}

func ConnectToRedis(addr string) ports.RedisRepository {
	redisRe := redisdb.NewRedisClient(addr, 0, 15)
	return services.NewRedisService(redisRe)
}
