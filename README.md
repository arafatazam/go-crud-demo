# go-crud-demo
This is a demo CRUD project using Fiber golang framework. It has implemented a REST api for User model.

# How to run the demo
* You need to have go1.14, docker and docker-compose installed in your system.
* Go to the project directory and run: `docker-compose up -d`
* Copy the .env.example file as .env file: `cp .env.example .env`
* Run `go mod download` for installing the dependency module.
* Run `go run .`
* Now you can test the endpoints using an HTTP client such as postman, insomnia etc.

# Endpoints
* Default url would be: http://localhost:3000
* You can change the port from .env file

| Verb          | Endpoint      | Description  |
| :------------ |:--------------| :------------|
| GET           | /users        | List Users   |
| POST          | /users        | Create User  |
| GET           | /users/:id    | View User    |
| PUT           | /users/:id    | Update User  |
| DELETE        | /users/:id    | Delete User  |

# Sample User Data

```json
{
	"email": "me@belabose.com",
	"first_name": "Bela",
	"last_name": "Bose",
	"address": "Kolkata",
	"phone": "2441139"
}
```

* Consult the docker-compose file for viewing the data using adminer sql admin.
