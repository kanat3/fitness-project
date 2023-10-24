## Features

> `Vera`
- Add a frontend

> `Anna`
- Add triggers for the database
- Add registration of new users and their login to the service
- Add password recovery by e-mail
- Implement full functionality for interaction with the frontend
- Implement the generation of a nutrition and training plan with all the possibilities

## Tech

- [Golang] - version 1.20.7
- [PostgreSQL] - image 15.2-alpine
- [Docker] - version 24.0.5
- [Docker-Compose] - version 1.29.2

## Backend installation

> Note: project isn't ready for building

Set environments
```
export CONFIG_PATH=backend/config/local.yaml
```
Up database. You can connect to db using 5432 port
```
cd backend/database || docker-compose up --build -d
```
Run code
```
cd backend || go run .
```