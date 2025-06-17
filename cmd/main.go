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

	config := loadConfig()

	db := connectDatabase(config)
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
	migrateDatabase(db)

	repoUser := initializeUserRepository(db)
	svcUser := initializeUserService(config, repoUser)
	ctrlUser := initializeUserController(svcUser)

	repoTran := initializeTransactionRepository(db)
	svcTran := initializeTransactionService(repoTran)
	ctrlTran := initializeTransactionController(svcTran)

	startServer(ctrlUser, ctrlTran)
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

func initializeUserRepository(db *gorm.DB) repository.UserRepository {
	log.Println("Initializing user repository...")
	repo, err := repository.NewUserRepository(db)
	if err != nil {
		log.Fatalf("Error creating user repository: %v", err)
	}
	log.Println("User repository initialized successfully")
	return repo
}

func initializeUserService(config *configuration.Config, repo repository.UserRepository) service.UserService {
	log.Println("Initializing user service...")
	svc, err := service.NewUserServiceImpl(config, repo)
	if err != nil {
		log.Fatalf("Error creating user service: %v", err)
	}
	log.Println("User service initialized successfully")
	return svc
}

func initializeUserController(svc service.UserService) controller.UserController {
	log.Println("Initializing user controller...")
	ctrl, err := controller.NewUserController(svc)
	if err != nil {
		log.Fatalf("Error creating user controller: %v", err)
	}
	log.Println("User controller initialized successfully")
	return ctrl
}
func initializeTransactionController(svc service.TransactionService) controller.TransactionController {
	log.Println("Initializing transaction controller...")
	ctrl, err := controller.NewTransactionController(svc)
	if err != nil {
		log.Fatalf("Error creating transaction controller: %v", err)
	}
	log.Println("Transaction controller initialized successfully")
	return ctrl
}

func initializeTransactionService(repo repository.TransactionRepository) service.TransactionService {
	svc, err := service.NewTransactionServiceImpl(repo)
	if err != nil {
		log.Fatalf("Error creating transaction service: %v", err)
	}
	log.Println("Transaction service initialized successfully")
	return svc
}

func initializeTransactionRepository(db *gorm.DB) repository.TransactionRepository {
	log.Println("Initializing transaction service...")
	repo, err := repository.NewTransactionRepositoryImpl(db)
	if err != nil {
		log.Fatalf("Error creating transaction repository: %v", err)
	}
	log.Println("Transaction repository initialized successfully")
	return repo

}

func startServer(user controller.UserController, transfer controller.TransactionController) {
	log.Println("Starting server on port 8080...")
	rest.StartServer(user, transfer)
	log.Println("Server started successfully on port 8080")
}
