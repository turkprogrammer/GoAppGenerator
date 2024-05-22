package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	fmt.Println("Welcome to GoAppGenerator!")
	fmt.Println("Effortlessly generate the ideal application scaffold for your Go web application or API.")

	// Получаем название проекта от пользователя
	fmt.Print("Enter your project name: ")
	var projectName string
	fmt.Scanln(&projectName)

	// Создаем директорию для проекта
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		fmt.Println("Error creating project directory:", err)
		return
	}

	// Переходим в созданную директорию
	err = os.Chdir(projectName)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	// Инициализируем модуль Go
	err = initGoModule(projectName)
	if err != nil {
		fmt.Println("Error initializing Go module:", err)
		return
	}

	// Создаем директории для слоев чистой архитектуры
	err = createDirectories([]string{"cmd", "internal/app", "internal/handler", "internal/repository", "internal/usecase"})
	if err != nil {
		fmt.Println("Error creating directories:", err)
		return
	}

	// Генерируем файл main.go в директории cmd
	err = generateMainFile("cmd", projectName)
	if err != nil {
		fmt.Println("Error generating main.go file:", err)
		return
	}

	// Генерируем файлы внутри internal
	err = generateInternalFiles()
	if err != nil {
		fmt.Println("Error generating internal files:", err)
		return
	}

	// Генерируем файл app.go в директории internal/app
	err = generateAppFile("internal/app")
	if err != nil {
		fmt.Println("Error generating app.go file:", err)
		return
	}

	// Создаем файл go.mod
	err = generateGoModFile()
	if err != nil {
		fmt.Println("Error generating go.mod file:", err)
		return
	}

	fmt.Println("Your Go web application scaffold with clean architecture and Gin routing has been generated successfully!")
}

func createDirectories(dirs []string) error {
	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		fmt.Println("Created directory:", dir)
	}
	return nil
}

func initGoModule(projectName string) error {
	cmd := exec.Command("go", "mod", "init", "github.com/example/"+projectName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func generateMainFile(dir, projectName string) error {
	fileContent := fmt.Sprintf(`
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/example/project/internal/app"
	"github.com/example/project/internal/handler"
)

func main() {
	fmt.Println("Hello, Clean Architecture with Gin!")

	// Setup your application components here
	app := app.NewApp()
	h := handler.NewHandler(app)
	h.SetupRoutes()

	// Handle graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(":8080", h.Router); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-signals
	fmt.Println("Shutting down gracefully...")
	app.Shutdown()
	log.Println("Server gracefully stopped")
}
`, projectName)
	fileName := fmt.Sprintf("%s/main.go", dir)

	// Создаем файл и записываем в него содержимое
	err := os.WriteFile(fileName, []byte(fileContent), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Generated", fileName)
	return nil
}

func generateInternalFiles() error {
	// Генерируем файлы внутри internal
	internalFiles := map[string]string{
		"internal/app/app.go": `
package app

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

// App struct represents your application
type App struct {
	mu    sync.Mutex
	quit  chan struct{}
}

// NewApp creates a new instance of the App
func NewApp() *App {
	return &App{
		quit: make(chan struct{}),
	}
}

// Run starts the application
func (a *App) Run() {
	log.Println("Application is running...")
	<-a.quit
}

// Shutdown gracefully shuts down the application
func (a *App) Shutdown() {
	log.Println("Shutting down...")
	close(a.quit)
	a.mu.Lock()
	defer a.mu.Unlock()
	select {
	case <-time.After(5 * time.Second):
		log.Println("Forced shutdown after 5 seconds timeout")
	case <-a.quit:
		log.Println("Shutdown completed")
	}
}
`,
		"internal/handler/handler.go": `
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/example/project/internal/app"
	"github.com/example/project/internal/usecase"
)

// Handler struct handles HTTP requests
type Handler struct {
	App    *app.App
	Usecase *usecase.Usecase
	Router *gin.Engine
}

// NewHandler creates a new instance of the Handler
func NewHandler(app *app.App) *Handler {
	router := gin.Default()
	h := &Handler{
		App:    app,
		Usecase: usecase.NewUsecase(),
		Router: router,
	}
	return h
}

// SetupRoutes configures the HTTP routes
func (h *Handler) SetupRoutes() {
	h.Router.GET("/", h.handleHello)
}

func (h *Handler) handleHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Clean Architecture with Gin!",
	})
}
`,
		"internal/repository/repository.go": `
package repository

// Repository interface defines the methods for data storage/retrieval
type Repository interface {
	// Add your repository methods here
}
`,
		"internal/usecase/usecase.go": `
package usecase

import (
	"log"

	"github.com/example/project/internal/repository"
)

// Usecase struct contains the business logic
type Usecase struct {
	repo repository.Repository
}

// NewUsecase creates a new instance of the Usecase
func NewUsecase() *Usecase {
	return &Usecase{}
}

// StartUsecase starts the business logic
func (u *Usecase) StartUsecase() {
	// Add your business logic here
	log.Println("Business logic started...")
}
`,
	}

	for filePath, fileContent := range internalFiles {
		err := os.MkdirAll(filepath.Dir(filePath), 0755)
		if err != nil {
			return err
		}

		err = os.WriteFile(filePath, []byte(fileContent), 0644)
		if err != nil {
			return err
		}

		fmt.Println("Generated", filePath)
	}

	return nil
}
func generateAppFile(dir string) error {
	fileContent := `
package app

import (
	"log"
	"sync"
	"time"
)

// App struct represents your application
type App struct {
	mu    sync.Mutex
	quit  chan struct{}
}

// NewApp creates a new instance of the App
func NewApp() *App {
	return &App{
		quit: make(chan struct{}),
	}
}

// Run starts the application
func (a *App) Run() {
	log.Println("Application is running...")
	<-a.quit
}

// Shutdown gracefully shuts down the application
func (a *App) Shutdown() {
	log.Println("Shutting down...")
	close(a.quit)
	a.mu.Lock()
	defer a.mu.Unlock()
	select {
	case <-time.After(5 * time.Second):
		log.Println("Forced shutdown after 5 seconds timeout")
	case <-a.quit:
		log.Println("Shutdown completed")
	}
}
`
	fileName := fmt.Sprintf("%s/app.go", dir)

	// Создаем файл и записываем в него содержимое
	err := os.WriteFile(fileName, []byte(fileContent), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Generated", fileName)
	return nil
}
func getDependencies() error {
	cmd := exec.Command("go", "get", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func generateGoModFile() error {
	fileContent := `
module github.com/example/project

go 1.17

require (
	github.com/gin-gonic/gin v1.7.4
)
`
	fileName := "go.mod"

	// Создаем файл и записываем в него содержимое
	err := os.WriteFile(fileName, []byte(fileContent), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Generated", fileName)

	// Получаем зависимости проекта
	err = getDependencies()
	if err != nil {
		return err
	}

	return nil
}
