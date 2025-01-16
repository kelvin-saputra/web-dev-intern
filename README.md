# Web Development Technical Interview

## Introduction
This repository contains all the necessary information and resources for the technical interview process. This is my work for the technical interview.

## Table of Contents
- [Introduction](#introduction)
- [Project Structure](#project-structure)
- [Setup Instructions](#setup-instructions)
- [Usage](#usage)
- [Api Endpoint](#api-endpoint)

## Project Structure
```
ROOT PROJECT
├── config/
│   ├── config.go
├── controller/
│   ├── comment_controller.go
├── dto/
│   ├── comment_dto.go
├── migrations/
│   ├── 0001_create_comments_table.down.sql
│   └── 0001_create_comments_table.down.sql
├── model/
│   ├── comment.go
├── routes/
│   ├── comment_route.go
├── service/
│   ├── comment_service.go
├── main.go
├── go.mod
├── go.sum
├── README.md
├── sample.json
└── .gitignore
```

## Setup Instructions
1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/Primeskill-Technical-Interview.git
    ```
2. Navigate to the project directory:
    ```sh
    cd `<project_directory>`
    ```
3. Install the necessary dependencies:
    ```sh
    go mod tidy
    ```

## Usage
To run the project, use the following command:
```sh
go run main.go
```

## API Endpoint
### CREATE COMMENT
---
**Description:** Create a new comment. <br>
**Method:** `POST` <br>
URL : [http://localhost:8080/api/v0/comment/create](http://localhost:8080/api/v0/comment/create)

**Request Body:**
```json
{
    "userId": 12,
    "title": "eum et est occaecati",
    "body": "ullam et saepe reiciendis voluptatem adipisci\nsit amet autem assumenda provident rerum culpa\nquis hic commodi nesciunt rem tenetur doloremque ipsam iure\nquis sunt voluptatem rerum illo velit"
}
```

**Success Response:**
```json
{
    "userId": 12,
    "id": 307,
    "title": "eum et est occaecati",
    "body": "ullam et saepe reiciendis voluptatem adipisci\nsit amet autem assumenda provident rerum culpa\nquis hic commodi nesciunt rem tenetur doloremque ipsam iure\nquis sunt voluptatem rerum illo velit"
}
```
---
### CREATE COMMENT BULK
---
**Description:** Create a new bulk comment.<br>
**Method:** `POST` <br>
URL : [http://localhost:8080/api/v0/comment/create/bulk](http://localhost:8080/api/v0/comment/create/bulk)

**Request Body:**
```json
[
    {
    "userId": 1,
    "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
    "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
    },
    {
    "userId": 1,
    "title": "qui est esse",
    "body": "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla"
    }
]
```

**Success Response:**
```json
[
    {
        "userId": 1,
        "id": 1508,
        "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
        "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
    },
    {
        "userId": 1,
        "id": 1509,
        "title": "qui est esse",
        "body": "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla"
    }
]
```
---
### GET ALL COMMENT
---
**Description:** Retrieve All Comments in Database.<br>
**Method:** `GET` <br>
URL : [http://localhost:8080/api/v0/comment](http://localhost:8080/api/v0/comment)

**Success Response:**
```json
[
    {
        "userId": 1,
        "id": 1508,
        "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
        "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
    },
    {
        "userId": 1,
        "id": 1509,
        "title": "qui est esse",
        "body": "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla"
    }
]
```
---
### GET COMMENT BY ID
---
**Description:** Retrieve Comment Filtered by ID in Database.<br>
**Method:** `GET` <br>
URL : [http://localhost:8080/api/v0/comment/:id](http://localhost:8080/api/v0/comment/:id)
EXAMPLE URL : [http://localhost:8080/api/v0/comment/1508](http://localhost:8080/api/v0/comment/1508)

**Success Response:**
```json
{
    "userId": 1,
    "id": 1508,
    "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
    "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
}
```
---
### GET COMMENT BY USER ID
---
**Description:** Retrieve All Comments That Belong to a Specific User by Id in Database.<br>
**Method:** `GET` <br>
URL : [http://localhost:8080/api/v0/comment/user/:id](http://localhost:8080/api/v0/comment/user/:id)
EXAMPLE URL : [http://localhost:8080/api/v0/comment/user/1](http://localhost:8080/api/v0/comment/user/1)

**Success Response:**
```json
[
    {
        "userId": 1,
        "id": 1508,
        "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
        "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
    },
    {
        "userId": 1,
        "id": 1509,
        "title": "qui est esse",
        "body": "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla"
    }
]
```
---
### DELETE COMMENT BY ID
---
**Description:** Delete Comments in Database by Id.<br>
**Method:** `DELETE` <br>
URL : [http://localhost:8080/api/v0/comment/delete/:id](http://localhost:8080/api/v0/comment/delete/:id)
EXAPMLE URL : [http://localhost:8080/api/v0/comment/delete/1509](http://localhost:8080/api/v0/comment/delete/1509)

**Success Response:**
```json
{
    "message": "Comment with ID 1509 deleted"
}
```
---
### DELETE COMMENT BY USER ID
---
**Description:** Delete Comments in Database by User Id.<br>
**Method:** `DELETE` <br>
URL : [http://localhost:8080/api/v0/comment/delete/user/:id](http://localhost:8080/api/v0/comment/delete/user/:id)
EXAPMLE URL : [http://localhost:8080/api/v0/comment/delete/user/1](http://localhost:8080/api/v0/comment/delete/user/1)

**Success Response:**
```json
{
    "message": "Comment Deleted Successfully"
}
```
---
### DELETE COMMENT BULK
---
**Description:** Delete multiple comments in the database.<br>
**Method:** `DELETE` <br>
URL : [http://localhost:8080/api/v0/comment/delete/bulk](http://localhost:8080/api/v0/comment/delete/bulk)
EXAPMLE URL : [http://localhost:8080/api/v0/comment/delete/bulk](http://localhost:8080/api/v0/comment/delete/bulk)

**Request Body:**
```json
{
    "id": [1510,1511,1512]
}
```

**Success Response:**
```json
{
    "message": "Comments Deleted Successfully"
}
```
---
### DELETE ALL COMMENT
---
**Description:** Delete All comments in the database.<br>
**Method:** `DELETE` <br>
URL : [http://localhost:8080/api/v0/comment/delete/all](http://localhost:8080/api/v0/comment/delete/all)
EXAPMLE URL : [http://localhost:8080/api/v0/comment/delete/all](http://localhost:8080/api/v0/comment/delete/all)

**Success Response:**
```json
{
    "message": "All Comments Deleted Successfully"
}
```