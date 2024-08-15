package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/abe16s/Online-Marketplace-API/controllers"
	"github.com/abe16s/Online-Marketplace-API/infrastructures"
	"github.com/abe16s/Online-Marketplace-API/repositories"
	"github.com/abe16s/Online-Marketplace-API/router"
	"github.com/abe16s/Online-Marketplace-API/usecases"
	"github.com/abe16s/Online-Marketplace-API/usecases/interfaces"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}

	dbName := "Online-Marketplace"
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	PasswordService := &infrastructures.PasswordService{}
	JwtService := &infrastructures.JwtService{JwtSecret: jwtSecret}
	EmailService := &infrastructures.EmailService{}

	var userRepository interfaces.IUserRepo = repositories.NewUserRepository(client, dbName, "users")
	var authUseCase = usecases.AuthUseCase{UserRepository: userRepository, PwdService: PasswordService, JwtService: JwtService, EmailService: EmailService}
	authController := controllers.AuthController{AuthUseCase: &authUseCase}

	r := router.SetupRouter(&authController)
	r.Run("localhost:" + os.Getenv("SERVER_PORT"))
}