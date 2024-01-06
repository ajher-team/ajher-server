package main

import (
	"log"

	"ajher-server/api/controllers"
	"ajher-server/api/middleware"
	"ajher-server/database"
	"ajher-server/docs"
	"ajher-server/internal/otp"
	"ajher-server/internal/participation"
	"ajher-server/internal/question"
	"ajher-server/internal/quiz"
	"ajher-server/internal/quizCategory"
	"ajher-server/internal/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	docs.SwaggerInfo.Title = "Ajher API"
	docs.SwaggerInfo.Description = "Ajher Backend API documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.ConnectDB()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	userRoute := api.Group("user")
	quizCategoryRoute := api.Group("quizCategory")
	quizRoute := api.Group("quiz")
	questionRoute := api.Group("question")

	// repositories
	userRepository := user.NewRepository(db)
	otpRepository := otp.NewRepository(db)
	quizCategoryRepository := quizCategory.NewRepository(db)
	quizRepository := quiz.NewRepository(db)
	participationRepository := participation.NewRepository(db)
	questionRepository := question.NewRepository(db)

	// services
	userService := user.NewService(userRepository, otpRepository)
	otpService := otp.NewService(otpRepository)
	quizCategoryService := quizCategory.NewService(quizCategoryRepository)
	quizService := quiz.NewService(quizRepository, participationRepository)
	questionService := question.NewService(questionRepository)

	// controllers
	userHandler := controllers.NewUserHandler(userService, otpService)
	quizCategoryHandler := controllers.NewQuizCategoryHandler(quizCategoryService)
	quizHandler := controllers.NewQuizHandler(quizService)
	questionHandler := controllers.NewQuestionHandler(questionService)

	// auth middleware
	authMiddleware := middleware.NewAuthMiddleware(userService)

	// user routes
	userRoute.POST("/register", userHandler.Register)
	userRoute.POST("/login", userHandler.Login)
	userRoute.GET("/profile", authMiddleware.AuthMiddleware, userHandler.GetProfile)
	userRoute.GET("/validateToken", authMiddleware.AuthMiddleware, userHandler.ValidateToken)
	userRoute.POST("/refreshToken", authMiddleware.RefreshTokenMiddleware, userHandler.RefreshToken)
	userRoute.POST("/googleAuth", userHandler.GoogleAuth)
	userRoute.POST("/resetPassword", userHandler.ResetPassword)
	userRoute.POST("/verifyOtp", userHandler.VerifyOtp)
	userRoute.POST("/changePassword", userHandler.ChangePassword)

	// quizCategoryRoutes
	quizCategoryRoute.GET("/", authMiddleware.AuthMiddleware, quizCategoryHandler.GetAll)
	quizCategoryRoute.GET("/:id", authMiddleware.AuthMiddleware, quizCategoryHandler.GetById)
	quizCategoryRoute.POST("/save", authMiddleware.AuthMiddleware, quizCategoryHandler.Save)

	// quiz
	quizRoute.POST("/save", authMiddleware.AuthMiddleware, quizHandler.Save)

	// question
	questionRoute.POST("/save", authMiddleware.AuthMiddleware, questionHandler.Save)

	router.Run(":5000")

}
