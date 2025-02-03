package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib" // Import PostgreSQL driver
	"github.com/jmoiron/sqlx"
	"github.com/kritpi/learn-anywhere-services/configs"
	"github.com/kritpi/learn-anywhere-services/internal/adapters/middlewares"
	_ "github.com/lib/pq"
)

func main() {
	// Load Config
	cfg := configs.NewConfig()

	// Setup Postgres Connection
	pgDb, pgErr := configs.InitDatabase(cfg)
	if pgErr != nil {
		log.Fatalf("⛔️ Failed to connect to Postgres: %v", pgErr)
	}
	defer pgDb.Close()

	minDb, minErr := configs.InitMinio(cfg)
	if minErr != nil {
		log.Fatalf("⛔️ Failed to connect to Minio: %v", minErr)
	}
	log.Printf("🚰 MinIO is ready! Bucket: %s, Endpoint: %s", minDb.Bucket, minDb.Endpoint)

	// Setup Fiber
	app := fiber.New()
	middlewares.SetUpMiddleware(app)
	setupRoutes(app, pgDb, cfg)
	startServer(app)
}

func setupRoutes(app *fiber.App, db *sqlx.DB, cfg *configs.Config) {
	// Repositories, Services, Handlers Setup can go here

	// Example Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Fiber 🌈")
	})
}

func startServer(app *fiber.App) {
	// Run Server
	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("⛔️ Error starting server: %v", err)
		}
	}()
	log.Println("🚀 🌈 Server running on port 8080")

	// Graceful Shutdown Handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("⛔️ Shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("⛔️ Server forced to shutdown: %v", err)
	}
	log.Println("✅ Server exited gracefully")
}
