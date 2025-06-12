package main

import (
	"fmt"
	rest "hello_bets/pkg/api"
	"hello_bets/pkg/configuration"
	"hello_bets/pkg/controller"
	database "hello_bets/pkg/infrastructure"
	"hello_bets/pkg/repository"
	"hello_bets/pkg/service"
	"log"

	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting Hello Bets...")

	// Cargar configuración
	config := loadConfig()

	// Conectar a la base de datos
	db := connectDatabase(config)

	// Verificar conexión a la base de datos
	if db == nil {
		log.Fatal("Failed to connect to the database")
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error getting database instance: %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
		log.Println("Database connection closed successfully")
	}()

	// Migrar la base de datos
	migrateDatabase(db)

	// Inicializar repositorio, servicio y controlador
	repo := initializeRepository(db)
	svc := initializeService(config, repo)
	ctrl := initializeController(svc)

	// Iniciar el servidor
	startServer(ctrl, svc)
}

func loadConfig() *configuration.Config {
	log.Println("Loading configuration...")
	config, err := configuration.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	log.Println("Configuration loaded successfully")
	return config
}

func connectDatabase(config *configuration.Config) *gorm.DB {
	log.Println("Connecting to the database...")
	db, err := database.Connect(
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	log.Println("Database connection established successfully")
	return db
}

func migrateDatabase(db *gorm.DB) {
	log.Println("Migrating database...")
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	log.Println("Database migration completed successfully")
}

func initializeRepository(db *gorm.DB) repository.UserRepository {
	log.Println("Initializing user repository...")
	repo, err := repository.NewUserRepository(db)
	if err != nil {
		log.Fatalf("Error creating user repository: %v", err)
	}
	log.Println("User repository initialized successfully")
	return repo
}

func initializeService(config *configuration.Config, repo repository.UserRepository) service.UserService {
	log.Println("Initializing user service...")
	svc, err := service.NewUserServiceImpl(config, repo)
	if err != nil {
		log.Fatalf("Error creating user service: %v", err)
	}
	log.Println("User service initialized successfully")
	return svc
}

func initializeController(svc service.UserService) controller.UserController {
	log.Println("Initializing user controller...")
	ctrl, err := controller.NewUserController(svc)
	if err != nil {
		log.Fatalf("Error creating user controller: %v", err)
	}
	log.Println("User controller initialized successfully")
	return ctrl
}

func startServer(ctrl controller.UserController, svc service.UserService) {
	log.Println("Starting server on port 8080...")
	rest.StartServer(ctrl, svc)
	log.Println("Server started successfully on port 8080")
}
