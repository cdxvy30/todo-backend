# To-do list backend
This is a backend service written in Golang along with PostgreSQL. The goal of this project is to let users manipulate a todo-list with RESTful API, including the functionalities below:
- Get all todo items.
- Get a specific todo item with ID.
- Create a todo item.
- Update a specific todo item with ID.
- Delete a specific todo item with ID. 

## Getting Started
### Requirements
- Go 1.20
- PostgreSQL 14.7
- (Optional) An operating system with **Docker Engine**

### Installation
1. Clone the repository.
2. Run `go mod download` to download the dependencies.
3. Run `go run main.go` to start the service.
4. Access the application at `http://localhost:8080/`

### Docker (Optional)
In the project root, run:
```
docker-compose up
```
The web application will be accessible at `http://localhost:8080`.

### API Routes
| Method | Url                 | Description                         |
| ------ | ------------------- | ----------------------------------- |
| GET    | /todos              | Get all todo items                  | 
| GET    | /todos/:id          | Get a specific todo item with ID    |
| POST   | /todos              | Create a todo item                  |
| PUT    | /todos/:id          | Update a specific todo item with ID |
| DELETE | /todos/:id          | Delete a specific todo item with ID |

You can test APIs using software like [Postman](https://www.postman.com/).

### Request body format of **POST /todos** and **PUT /todos/:id**
```
{
  "title": "Title of the todo",
  "description": "Description of the todo",
  "completed": true or false
}
```

Other request **do not** need a request body.

## License
Currently has no license.