# PhoneBook-API

PhoneBook API
Crud make phone book to add , delete , edit and add to favorites numbers.

## Tech Stack
1. PostgreSQL
2. Golang

## Prerequisites
Enter the following command to install sql driver for mysql, gin and gorm in the project.
```
cd /go-PhonebookAPI
go mod tidy
```

## Running the REST API
```
go run main.go
```

## Setting Port & Database
```
file - .env
DB_HOST = localhost
DB_PORT = 5432
DB_PASSWORD = dbpassword
DB_USER = dbuser
DB_SSLMODE = disable
DB_NAME = dbname
PORT = "4000"
```


## API Endpoints
| Route                      | HTTP          | Description                    |
| ---------------------------|:-------------:| ------------------------------:|
| /api/                      | GET           | Get all Numbers                |
| /api/add                   | POST          | Create Number                  |
| /api/delete/:id            | DELETE        | Delete Number                  |
| /api/detail/:id            | GET           | Detail Number                  |
| /api/favorites/            | GET           | Get all favorites Numbers      |
| /api/favorites/detail/:id  | GET           | Get Detail Favorites Number    |
| /api/favorites/add/:id     | POST          | Add Favorites Number           |
| /api/favorites/delete/:id  | DELETE        | Delete Detail Favorites Number |



