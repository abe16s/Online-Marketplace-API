package router

import (
	"log"
	"os"

	"github.com/abe16s/Online-Marketplace-API/controllers"
	"github.com/abe16s/Online-Marketplace-API/infrastructures"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRouter(userController *controllers.AuthController) *gin.Engine {
    router := gin.Default()
    err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// get jwt secret from env
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	jwtservice := &infrastructures.JwtService{JwtSecret: jwtSecret}


	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	router.GET("/refreshtoken", infrastructures.RefreshMiddleware(jwtservice) ,userController.RefreshToken)
	router.GET("/activate", userController.ActivateAccount)



	// router.GET("/tasks", infrastructure.AuthMiddleware(jwtservice, false), taskController.GetTasks)
    // router.GET("/tasks/:id", infrastructure.AuthMiddleware(jwtservice, false), taskController.GetTaskById)
    // router.PUT("/tasks/:id", infrastructure.AuthMiddleware(jwtservice, true), taskController.UpdateTaskByID)
    // router.DELETE("/tasks/:id", infrastructure.AuthMiddleware(jwtservice, true), taskController.DeleteTask)
    // router.POST("/tasks", infrastructure.AuthMiddleware(jwtservice, true), taskController.AddTask)

    return router
}