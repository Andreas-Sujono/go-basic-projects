package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andreas-sujono/go-basic-projects/rssAggregators/handler"
	"github.com/andreas-sujono/go-basic-projects/rssAggregators/internal/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	if portString == "" || dbUrl == "" {
		log.Fatal("PORT is not found")
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Database connection cannot be opened")
	}

	dbQueries := database.New(db)
	apiConfig := handler.ApiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/test", handler.TestSuccessHandler)
	v1Router.Get("/error", handler.TestErrHandler)

	v1Router.Get("/users", apiConfig.AuthMiddleware(apiConfig.GetUser))
	v1Router.Post("/users", apiConfig.CreateUser)

	v1Router.Get("/feeds", (apiConfig.GetFeeds))
	v1Router.Get("/posts", apiConfig.AuthMiddleware(apiConfig.GetPosts))
	v1Router.Post("/feeds", apiConfig.AuthMiddleware(apiConfig.CreateFeed))

	v1Router.Get("/feeds-follows", apiConfig.AuthMiddleware(apiConfig.GetFeedFollows))
	v1Router.Post("/feeds-follows", apiConfig.AuthMiddleware(apiConfig.CreateFeedFollow))
	v1Router.Delete("/feed-follows/{feedFollowID}", apiConfig.AuthMiddleware(apiConfig.DeleteFeedFollow))

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    "localhost:" + portString,
	}

	go handler.StartScraping(
		dbQueries,
		10,
		time.Minute,
	)

	log.Printf("Server is starting on port %v", portString)
	err2 := server.ListenAndServe()
	if err2 != nil {
		log.Fatal(err2)
	}
}
