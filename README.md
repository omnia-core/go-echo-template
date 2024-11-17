# Go REST API Template with Echo, GORM, and Fx

This repository provides a boilerplate for building RESTful APIs using the following technologies:

- [Echo](https://echo.labstack.com/): A lightweight and fast Go web framework for HTTP routing and middleware support.
- [GORM](https://gorm.io/): The most popular Go ORM library for database interactions.
- [Fx](https://github.com/uber-go/fx): A dependency injection framework for modular and scalable applications.

It offers a structured project layout to facilitate scalable and maintainable application development.

---

## Features

- **Domain-Driven Design**: The project separates interfaces and implementations by organizing interfaces in the `domain/` directory and their respective implementations in the `internal/` directory.
- **Echo Framework**: Efficient HTTP routing and middleware support for fast REST API development.
- **GORM Integration**: Simplifies database access and ORM capabilities.
- **Fx for Dependency Injection**: Provides modular DI, making it easier to manage complex dependencies.

---

## Project Structure

The project is organized as follows:

```
go-echo-template/
├── cmd/
│   └── app/            # Main application entry point
├── config/             # Configuration files and settings
├── domain/             # Interfaces for each domain
│   ├── user.go         # Example: User-related interfaces
│   ├── auth.go         # Example: Authentication interfaces
│   └── ...             # Additional domain interfaces
├── internal/           # Implementations of the interfaces
│   └── user/           # Example: User-related implementation
│       ├── router/     # HTTP router implementation
│       ├── usecase/    # Business logic implementations
│       ├── store/      # Data store (e.g., DB) implementations
│       └── ...         # Additional domain-specific implementations
├── pkg/                # Reusable utility packages
├── Makefile            # Build and run commands
├── README.md           # Project documentation
├── go.mod              # Go module file
├── go.sum              # Go dependencies checksum
```

---

## Getting Started

### Prerequisites

- **Go 1.20+** installed on your machine.
- A tool like `Make` for running the commands in the `Makefile`.

---

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/omnia-core/go-echo-template.git
   cd go-echo-template
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Build the application**:

   ```bash
   make build
   ```

4. **Run the application**:

   ```bash
   make run
   ```

The server will start, and you can access it at `http://localhost:8080`.

---

## Directory Overview

### `domain/`
Contains all the interfaces for the various domains of the application. These interfaces define the contract between different layers, ensuring a clear separation of concerns.

Examples:
- `user.go`: Defines methods related to user operations (e.g., `GetUser`, `CreateUser`).
- `auth.go`: Defines methods for authentication and authorization.

### `internal/`
Holds the implementation of the interfaces defined in the `domain/` directory.

Subdirectories:
- **`router/`**: Implements HTTP routes using the Echo framework.
- **`usecase/`**: Contains business logic implementations.
- **`store/`**: Implements data access (e.g., database queries) using GORM.

---

## Contributing

Contributions are welcome! Please fork the repository, create a feature branch, and submit a pull request. Make sure your code adheres to the project's coding standards and is well-tested.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgments

Special thanks to the [Echo](https://echo.labstack.com/), [GORM](https://gorm.io/), and [Fx](https://github.com/uber-go/fx) communities for their excellent work and contributions.

