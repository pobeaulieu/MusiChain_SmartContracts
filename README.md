
# Go service 
install go: https://go.dev/dl/

#### Dependencies
- `go get github.com/gofiber/fiber/v2`
- `go get -u gorm.io/gorm`
- `go get -u gorm.io/driver/sqlite`
- `go get github.com/gofiber/fiber/v2/middleware/cors`
- `go get github.com/mattn/go-isatty@v0.0.17` 
- `go get github.com/dgrijalva/jwt-go`
- `go get golang.org/x/crypto/bcrypt` 
- `go get github.com/joho/godotenv`

or 

`go mod tidy`

## React Webapp
See MusiChain_UI repo https://github.com/pobeaulieu/MusiChain_UI

# Database
The database is SQLite and is stored in the file system for ease of use (MusiChain.db). 

# API doc: Postman
import MusiChain.postman_collection.json in Postman and add rootCa.crt in the settings. 

