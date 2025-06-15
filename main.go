package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// github mux, cors handing
	"itops-assignment/backend/internal/api"
	"itops-assignment/backend/internal/repository"
	"itops-assignment/backend/internal/service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// 1. repo 초기화 진행
	issueRepos := repository.NewInMemoryIssueRepository()
	// 2. Service 초기화 진행
	issueService := service.NewIssueService(issueRepos)
	// 3. API 호출 초기화
	issueHandlers := api.NewIssueHandlers(issueService)

	r := mux.NewRouter()
	// 4. 라우터 등록 진행
	r.HandleFunc("/issue", issueHandlers.CreateIssue).Methods("POST")
	r.HandleFunc("/issues", issueHandlers.GetIssues).Methods("GET")
	r.HandleFunc("/issue/{id}", issueHandlers.GetIssueByID).Methods("GET")
	r.HandleFunc("/issue/{id}", issueHandlers.UpdateIssue).Methods("PATCH")
	// 페이지 인증 진행
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		//
		AllowedMethods:   []string{"GET", "POST", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		Debug:            true,
	})

	// 5. 헨들러 등록
	handlers := c.Handler(r)
	// 6. 서버 시작
	port := os.Getenv("POST")
	if port == "" {
		port = "8080"
	}

	serverAddr := fmt.Sprintf(":%s", port)

	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      handlers,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Printf("Server starting:%s\n", serverAddr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", serverAddr, err)
	}
}
