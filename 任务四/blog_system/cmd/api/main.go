package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"blog_system/internal/app/handler"
	"blog_system/internal/app/middleware"
	"blog_system/internal/app/service"

	// "blog_system/internal/domain/repository"
	"blog_system/internal/infrastructure/config"
	"blog_system/internal/infrastructure/database"
	"blog_system/pkg/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	// 加载配置
	cfg := config.MustLoad()
	// 初始化数据库
	db, err := database.NewDB(&cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	defer db.Close()
	// 自动迁移
	if err := db.AutoMigrate(); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 初始化依赖
	jwtService := auth.NewJWTService(cfg.JWT.Secret, cfg.JWT.Expiration)

	// 初始化仓储
	userRepo := database.NewUserRepository(db.DB)
	postRepo := database.NewPostRepository(db.DB)
	commentRopo := database.NewCommentRepository(db.DB)

	// 初始化服务
	userService := service.NewUserService(userRepo, jwtService)
	postService := service.NewPostService(postRepo)
	commentService := service.NewCommentService(commentRopo)
	// 初始化处理器
	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService)
	commentHandler := handler.NewCommentHandler(commentService)

	// 设置Gin模式
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建路由
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CORS中间件
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORS.AllowOrigins,
		AllowMethods:     cfg.CORS.AllowMethods,
		AllowHeaders:     cfg.CORS.AllowHeaders,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 注册路由
	setupRoutes(router, userHandler, postHandler, commentHandler, jwtService)

	// 启动服务器
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.App.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 优雅关闭
	go func() {
		log.Printf("Server is running on port %d", cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server:", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// 给服务器5秒时间完成当前请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}

func setupRoutes(router *gin.Engine, userHandler *handler.UserHandler, PostHandler *handler.PostHandler, CommentHandler *handler.CommentHandler, jwtService *auth.JWTService) {
	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "timestamp": time.Now().Unix()})
	})

	// API路由组
	api := router.Group("/api")
	{
		// 认证路由（公开）
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/register", userHandler.Register)
			authGroup.POST("/login", userHandler.Login)
		}

		// 需要认证的路由
		protected := api.Group("/users")
		protected.Use(middleware.AuthMiddleware(jwtService))
		{
			protected.GET("/profile", userHandler.GetProfile)
		}
		// 需要认证的路由
		posts := api.Group("/posts")
		posts.Use(middleware.AuthMiddleware(jwtService))
		{

			posts.POST("/create", PostHandler.CreatePost)
			posts.GET("/:id", PostHandler.GetPostByID)
			posts.POST("/ListPosts", PostHandler.ListPosts)
			posts.POST("/UpdatePost", PostHandler.UpdatePost)
			posts.DELETE("/:id", PostHandler.DeletePost)
		}
		// 需要认证的路由
		comments := api.Group("/comments")
		comments.Use(middleware.AuthMiddleware(jwtService))
		{

			comments.POST("/create", CommentHandler.CreateComment)
			comments.GET("/:id", CommentHandler.GetCommentsByPostID)
		}
	}

	// 根路径
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Gin Project API",
			"version": "1.0.0",
		})
	})
}
