# To-do list backend
This is a backend service written in Golang along with database like PostgreSQL. The goal of this project is to let users manipulate a todo list with RESTful API, including the functionalities below:
- Get all todo items.
- Get a specific todo item with ID.
- Create a todo item.
- Update a specific todo item with ID.
- Delete a specific todo item with ID. 

## Getting Started
### Prerequisites
- An operating system installed with `docker`

### Start the service
In the project root, run:
```
docker-compose up
```

### API Routes
| Method | Url                 | Description                         |
| ------ | ------------------- | ----------------------------------- |
| GET    | /todos              | Get all todo items                  | 
| GET    | /todos/:id          | Get a specific todo item with ID    |
| POST   | /todos              | Create a todo item                  |
| PUT    | /todos/:id          | Update a specific todo item with ID |
| DELETE | /todos/:id          | Delete a specific todo item with ID |

You can test APIs using software like [Postman](https://www.postman.com/).

### Request body format of `POST /todos` and `PUT /todos/:id`
```
{
  "title": "Title of the todo",
  "description": "Description of the todo",
  "completed": true or false
}
```
