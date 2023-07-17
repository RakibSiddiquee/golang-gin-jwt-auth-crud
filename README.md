## Golang, Gin framework, GORM, Postgres, JWT auth and CRUD Application

### Used Packages:
1. Gin Framework
2. Postgres
3. GORM
4. Golang JWT (https://github.com/golang-jwt/jwt)
5. Godotenv (https://github.com/joho/godotenv)

### Steps to follow
1. Rename the .env.example file to .env
2. Create a database in postgres
3. Change the DNS value in .env file
4. Run the command `go run migrate/migrate.go`
5. Check your database, tables should be availabe