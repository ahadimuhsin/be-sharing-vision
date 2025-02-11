
## Prerequisites
Ensure you have the following installed before running the project:
- [Go](https://golang.org/doc/install) (Version 1.XX+)
- Database MySQL version 8.0 +

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/ahadimuhsin/be-sharing-vision.git
   cd be-sharing-vision
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage
To run the project, use the following command:
```sh
go run main.go
```

For building an executable:
```sh
go build -o myapp
./myapp
```

## Configuration
- Create Environment File (.env) with contents
- Example:
  ```
    DB_NAME = "YOUR DB_NAME"
    DB_USER = "YOUR DB_USER"
    DB_HOST = 127.0.0.1
    DB_PORT = 3306
    DB_PASSWORD = ""
  ```

## Documentation
See the documentation [here](https://documenter.getpostman.com/view/10026548/2sAYXBEyS6)


