## Golang, Gin framework, GORM, Postgres, JWT auth and CRUD Application

### Used Packages:
1. Gin Framework
2. Postgres
3. GORM
4. Golang JWT (https://github.com/golang-jwt/jwt)
5. Godotenv (https://github.com/joho/godotenv)
6. Validation (https://github.com/go-playground/validator)
7. Slug (https://github.com/gosimple/slug)

### Steps to follow
1. Clone the repo
2. Run the command `go mod download`
3. Rename the .env.example file to .env 
4. Create a database in postgres 
5. Change the DNS value in .env file 
6. Run the command `go run migrate/migrate.go` (Drop existing tables and recreate those)
7. Check your database, tables should be available
8. Run the project using the command `go run main.go`
9. Test the application in Postman

#### Routes
1. http://localhost:3000/signup (Signup)
```json
{
   "name": "John Doe",
   "email": "john@doe.com",
   "password": "123456"
}
```
2. http://localhost:3000/login (Login)
```json 
{
  "email": "john@doe.com",
  "password": "123456"
}
```
3. http://localhost:3000/logout (Logout)
4. http://localhost:3000/categories/create (Create category)
```json
{
  "name": "National"
}
```
5. http://localhost:3000/categories (Get all category)
6. http://localhost:3000/categories/1/edit (Edit category)
7. http://localhost:3000/categories/1/update (Update category)
```json
{
  "name": "Sports"
}
```
8. http://localhost:3000/categories/1/delete (Soft delete a category)
9. http://localhost:3000/categories/all-trash (Get all trashed category)
10. http://localhost:3000/categories/delete-permanent/1 (Delete a trashed category permanently)
11. http://localhost:3000/posts/create (Create post)
```json
{
  "title": "Awesome post",
  "body": "This is the awesome post details",
  "categoryId": 1
}
```
12. http://localhost:3000/posts (Get all post)
13. http://localhost:3000/posts/1/show (Show a single post)
14. http://localhost:3000/posts/1/edit (Edit post)
15. http://localhost:3000/posts/1/update (Update post)
```json
{
  "title": "Hello World",
  "body": "This is the hello world post details",
  "categoryId": 1
}
```
16. http://localhost:3000/posts/1/delete (Soft delete a post)
17. http://localhost:3000/posts/all-trash (Get all trashed post)
18. http://localhost:3000/posts/delete-permanent/1 (Delete a trashed post permanently)
19. http://localhost:3000/posts/1/comment/store (Comment on a post)
```json
{
  "postId": 1,
  "body": "This is a comment"
}
```
20. http://localhost:3000/posts/1/comment/1/edit (Edit a comment)
21. http://localhost:3000/posts/1/comment/1/update (Update a comment)
```json
{
  "postId": 1,
  "body": "This is a updated comment"
}
```
22. http://localhost:3000/posts/1/comment/1/delete (Delete a comment)