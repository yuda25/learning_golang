<!-- create modul -->
go mod init learning_golang

<!-- auto run -->
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon

<!-- to auto run -->
CompileDaemon -command="./learning_golang"

<!-- env -->
go get github.com/joho/godotenv

<!-- framework gin -->
go get -u github.com/gin-gonic/gin

<!-- orm for GO -->
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

<!-- migrate -->
go run migrate/migrate.go

<!-- bcrypt -->
go get -u golang.org/x/crypto/bcrypt

<!-- jwt -->
go get -u github.com/golang-jwt/jwt/v4

<!-- gomail -->
go get gopkg.in/gomail.v2

<!-- excel -->
go get github.com/xuri/excelize/v2

<!-- uuid -->
go get github.com/google/uuid