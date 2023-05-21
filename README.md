# ChAMP-Backend-Final-Project

By Tuchtham Sungkameka

I plan to use Golang, which I have zero experience with.

These are the keywords I will try to adapt into my project.

1. Gin: rest api
2. Gorm: orm
3. Testify, GoMock: testing
4. Swagger: api documentation
5. PostgreSQL: database (use eleplantSQL)

Initial Setup

1. `go mod init ChAMP-Backend-Final-Project`
2. `go get github.com/githubnemo/CompileDaemon` to make go build every time files change
3. `go get -u github.com/gin-gonic/gin` to install Gin framework
4. `go get github.com/joho/godotenv` for easy env
5. `go get -u gorm.io/gorm` our ORM
6. `go get -u gorm.io/driver/postgres` our DB

Terminal Commands

1. `CompileDaemon -command="./ChAMP-Backend-Final-Project"` to make CompileDaemon report to terminal in realtime
