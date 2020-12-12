package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sablev/go-clean-architecture-std/internal/auth"
	"github.com/sablev/go-clean-architecture-std/internal/bookmark"
	authhttp "github.com/sablev/go-clean-architecture-std/internal/auth/delivery/http"
	authmongo "github.com/sablev/go-clean-architecture-std/internal/auth/repository/mongo"
	authuc "github.com/sablev/go-clean-architecture-std/internal/auth/usecase"
	bmhttp "github.com/sablev/go-clean-architecture-std/internal/bookmark/delivery/http"
	bmmongo "github.com/sablev/go-clean-architecture-std/internal/bookmark/repository/mongo"
	bmuc "github.com/sablev/go-clean-architecture-std/internal/bookmark/usecase"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	httpServer *http.Server

	bookmarkUC bookmark.UseCase
	authUC     auth.UseCase
}

func New() *App {
	db := initDB()

	userRepo := authmongo.New(db, viper.GetString("mongo.user_collection"))
	bmRepo := bmmongo.New(db, viper.GetString("mongo.bookmark_collection"))

	return &App{
		bookmarkUC: bmuc.New(bmRepo),
		authUC: authuc.New(
			userRepo,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetDuration("auth.token_ttl"),
		),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Set up http handlers
	// SignUp/SignIn endpoints
	authhttp.RegisterEndpoints(router, a.authUC)

	// API endpoints
	authMiddleware := authhttp.NewMiddleware(a.authUC)
	api := router.Group("/api", authMiddleware)

	bmhttp.RegisterEndpoints(api, a.bookmarkUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.uri")))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(viper.GetString("mongo.name"))
}
