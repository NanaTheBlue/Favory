package main

import (
	"net/http"

	authapi "github.com/nanagoboiler/internal/api/auth"
	favorapi "github.com/nanagoboiler/internal/api/favors"
	"github.com/nanagoboiler/internal/auth"
	"github.com/nanagoboiler/internal/bootstrap"
	"github.com/nanagoboiler/internal/favors"
	authrepo "github.com/nanagoboiler/internal/repository/auth"
	favorsrepo "github.com/nanagoboiler/internal/repository/favors"

	"context"
)

func main() {
	router := http.NewServeMux()
	ctx := context.Background()
	pool, err := bootstrap.NewPostgresPool(ctx)
	if err != nil {
		panic(err)
	}

	// Repositories
	authRepo := authrepo.NewUserRepository(pool)
	tokenRepo := authrepo.NewTokensRepository(pool)
	favorRepo := favorsrepo.NewFavorsRepository(pool)

	// Services
	authService := auth.NewAuthService(authRepo, tokenRepo)
	favorService := favors.NewFavorService(authRepo, favorRepo)

	// Handlers
	authRegister := authapi.Register(authService)
	authLogin := authapi.Login(authService)
	bingus := authapi.Health()
	renew := authapi.Renew(authService)
	createFavor := favorapi.Create(favorService)

	//Routes
	router.HandleFunc("POST /create", auth.AuthMiddleware(createFavor))
	router.HandleFunc("POST /register/", authRegister)
	router.HandleFunc("POST /login/", authLogin)
	router.HandleFunc("POST /health/", bingus)
	router.HandleFunc("GET /renew/", renew)

	println("Server Listening on Port 8085")
	http.ListenAndServe(":8085", router)
}
