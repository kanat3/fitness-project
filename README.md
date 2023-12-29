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

- [Golang] - version 1.21.4
- [PostgreSQL] - image 15.2-alpine
- [Docker] - version 24.0.5
- [Docker-Compose] - version 1.29.2

## Backend installation
1. Use docker-compose:
```
make all-compose
```
2. If database in docker-compose and project in Dockerfile:
```
make docker
```
It creates project/ directory to use as volume in container.
So if you want to kill containers and delete all trash (like project/) 
```
make clean
```
If you changed only backend/frontend code then don't use 'make docker' again (it will remove all your containers).
Just use
```
make update
```
It saves your container with database

3. If local:
```
make all-local
```
Up database using Docker. You can connect to db using 5432 port. Check config/local,yaml for more info
```
cd ./backend/database || docker-compose up --build -d
```
Run code
```
cd ./backend
./backend
```
You can compile backend faster. Just use it from ./backend:
```
chmod +x up.sh
./up.sh
```