# GoAppGenerator

Welcome to GoAppGenerator! This tool helps you effortlessly generate the ideal application scaffold for your Go web application or API following clean architecture principles and using the Gin framework for routing.

## English
### Description
GoAppGenerator is a tool designed to help developers quickly scaffold Go web applications or APIs with clean architecture principles and the Gin framework for routing. It automates the creation of a structured project layout, including predefined layers for the application logic, HTTP handlers, repositories, and use cases. With GoAppGenerator, developers can jumpstart their projects with a robust and maintainable codebase, saving time and effort in setting up the initial project structure.

## Русский
### Описание
GoAppGenerator — это инструмент, предназначенный для помощи разработчикам в быстром создании каркаса веб-приложений или API на Go, с использованием принципов чистой архитектуры и фреймворка Gin для маршрутизации. Он автоматизирует создание структурированного проекта, включающего заранее определенные слои для логики приложения, HTTP-обработчиков, репозиториев и вариантов использования. С помощью GoAppGenerator разработчики могут быстро начать свои проекты с надежной и поддерживаемой кодовой базой, экономя время и усилия на настройку начальной структуры проекта.
## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Generated Code](#generated-code)
- [Contributing](#contributing)
- [License](#license)

## Introduction

GoAppGenerator is designed to provide a robust starting point for your Go web applications or APIs. It automates the creation of a clean architecture scaffold with predefined layers and uses the Gin framework for routing.

## Installation

To install and set up the GoAppGenerator, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/turkprogrammer/GoAppGenerator.git
    cd GoAppGenerator
    ```

2. Build the generator:
    ```sh
    go build -o generator main.go
    ```

## Usage

To generate a new Go web application scaffold, run the generator and follow the prompts:

```sh
./generator
```

## Project Structure
### The generated project follows a clean architecture structure:
```
projectname/
├── cmd/
│   └── main.go
├── internal/
│   ├── app/
│   │   └── app.go
│   ├── handler/
│   │   └── handler.go
│   ├── repository/
│   │   └── repository.go
│   └── usecase/
│       └── usecase.go
├── go.mod
├── go.sum
```
```
cmd/: Contains the main entry point of the application.
```
```
internal/: Contains the core application code organized by clean architecture layers.
```
```
app/: Contains the application logic and setup.
```
```
handler/: Contains the HTTP handlers.
```
```
repository/: Contains the repository interfaces and implementations.
```
```
usecase/: Contains the business logic.
```
## Generated Code
### The generated scaffold includes a simple example to get you started. Here is an overview of the generated files:

#### cmd/main.go: The main entry point of the application.
#### internal/app/app.go: Contains the application logic, including setup and graceful shutdown.
#### internal/handler/handler.go: Contains the HTTP handlers using Gin.
#### internal/repository/repository.go: Defines the repository interfaces.
#### internal/usecase/usecase.go: Contains the business logic.

### Example endpoint:
```
GET /: Returns a welcome message.
```
Example response:
```
{
"message": "Hello, Clean Architecture with Gin!"
}
```
### Contributing:

Contributions are welcome! Please fork the repository and create a pull request to contribute.

Fork the repository
Create a new branch (git checkout -b feature-branch)
Commit your changes (git commit -am 'Add some feature')
Push to the branch (git push origin feature-branch)
Create a new Pull Request

### License
This project is licensed under the MIT License - see the LICENSE file for details.