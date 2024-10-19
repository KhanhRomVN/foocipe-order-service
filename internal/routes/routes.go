package routes

import (
	"foocipe-recipe-service/internal/handlers"
	"foocipe-recipe-service/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(r *gin.Engine, db *pgxpool.Pool) {
	v1 := r.Group("/v1")
	v1.Use(middleware.AuthToken())

	setupCartRoutes(v1, db)
}

func setupCartRoutes(rg *gin.RouterGroup, db *pgxpool.Pool) {
	carts := rg.Group("/carts")
	{
		carts.POST("", handlers.CreateCart(db))
		carts.GET("/:id", handlers.GetCartsByUserID(db))
		carts.PUT("/:id", handlers.UpdateQuantityCart(db))
		carts.DELETE("/:id", handlers.DeleteCartItem(db))
		carts.DELETE("/clear", handlers.DeleteCarts(db))
	}
}
