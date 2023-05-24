# ChAMP-Backend-Final-Project

By Tuchtham Sungkameka

I plan to use Golang, which I have zero experience with.

### These are the keywords I will try to adapt into my project.

1. Gin: rest api
2. Gorm: orm
3. Testify, GoMock: testing
4. Swagger: api documentation
5. PostgreSQL: database (use eleplantSQL)
6. TablePlus: database management tool

### Initial Setup

1. `go mod init ChAMP-Backend-Final-Project`
2. `go get github.com/githubnemo/CompileDaemon` to make go build every time files change
3. `go get -u github.com/gin-gonic/gin` to install Gin framework
4. `go get github.com/joho/godotenv` for easy env
5. `go get -u gorm.io/gorm` our ORM
6. `go get -u gorm.io/driver/postgres` our DB

### Terminal Commands

1. `CompileDaemon -command="./ChAMP-Backend-Final-Project"` to make CompileDaemon report to terminal in realtime
2. `go build .\migrate\; go run .\migrate\` to migrate (update table schema) of our model files to the database

### Tips I learned

1. Please use GORM's logger and set the mode to "info" there're info, warn, error, silent.
   logger.Silent: No logs are generated.
   logger.Error: Only error messages are logged.
   logger.Warn: Error and warning messages are logged.
   logger.Info: Error, warning, and informational messages are logged.
2. Use UNIX time or ISO 8601 time as the standard.
   To test on Postman, use this format `"DueDate":"2023-10-30T17:00:00.000Z"`
3. The attribute name "Order" messes up everything related to SQL command line. Must cover with "" or ''
   For example, `"order" desc`

### Design

- Create a task
  Auto-set order

- Update a task
  If update order to X, update order of every task from X

- Delete a task
  If deleted order X, update order of every task after X

- Create a list
  Auto-set order

- Update a list
  If update order to X, update order of every task after X

- Move a task to another list
  Place at last order

- Reorder a task in a list
  Works like update task

- Reorder a list
  Works like update list

- Delete a list, also every tasks in it
  OnDelete CASCADE
