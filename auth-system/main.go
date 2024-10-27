package main

import (
    "log"
    "net/http"
    "project/config"
    "project/db"
    "project/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    // Load environment variables
    config.LoadEnv()

    // Initialize Redis
    if err := db.InitRedis(); err != nil {
        log.Fatal("Failed to connect to Redis:", err)
    }
    defer db.Rdb.Close()

    // Initialize PostgreSQL
    if err := db.InitPostgres(); err != nil {
        log.Fatal("Failed to connect to PostgreSQL:", err)
    }
    defer db.Db.Close()

    // Setup Gin router
    router := gin.Default()
    routes.SetupRoutes(router)

    // Start the server
    port := config.GetEnv("PORT", "8080")
    if err := router.Run(":" + port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}