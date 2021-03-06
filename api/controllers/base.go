package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/technodeguy/real-estate/api/config"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"

	"github.com/technodeguy/real-estate/api/services"
)

type Server struct {
	cnf          *config.Config
	db           *sql.DB
	redisClient  *redis.Client
	router       *mux.Router
	tokenService services.ITokenService
	s3Service    services.S3ServiceInterface
}

func NewServer(cnf *config.Config) *Server {
	return &Server{cnf: cnf}
}

func (server *Server) Initialize() {
	var err error
	server.db, err = sql.Open("mysql", server.cnf.Db.Uri)

	if err != nil {
		log.Fatalf("Unable to connect to db %v", err.Error())
	}

	if err = server.db.Ping(); err != nil {
		log.Fatalf("Unable to ping to db %v", err.Error())
	}

	log.Println("DB connected successfully")

	rc := redis.NewClient(&redis.Options{
		Addr: server.cnf.Redis.Uri,
		DB:   0,
	})

	_, err = rc.Ping().Result()

	if err != nil {
		log.Fatalf("Unable to connect to redis %v", err.Error())
	}

	server.redisClient = rc

	log.Println("Redis connected successfully")

	server.router = mux.NewRouter()

	awsS3Config := server.cnf.Aws

	awsS3Service := services.NewAwsS3Service(
		awsS3Config.AccessKeyId,
		awsS3Config.SecretAccesKey,
		awsS3Config.BucketName,
	)

	awsS3Service.Initialize()

	server.s3Service = awsS3Service

	server.tokenService = services.NewTokenService(&server.cnf.Jwt, server.redisClient)

	log.Println("Services initialized successfully")

	server.initializeRoutes()
}

func (server *Server) RunServer() {
	host := server.cnf.Server.Host
	port := strconv.Itoa(server.cnf.Server.Port)
	appUrl := host + ":" + port

	fmt.Println(fmt.Sprintf("Listening on %v", appUrl))
	log.Fatal(http.ListenAndServe(":"+port, server.router))
}
