current_path=$(shell pwd)
SHELL := /bin/bash

front:

	if [ -d "$(current_path)/backend/frontend" ]; then \
		sudo rm -rf "$(current_path)/backend/frontend"; \
	fi

	if [ ! -d "$(current_path)/backend/frontend" ]; then \
		mkdir -p "$(current_path)/backend/frontend"; \
	fi

	# bad path stucture, copy all files in frontend/
	####
	cp -r $(current_path)/frontend/* $(current_path)/backend/frontend/
	####

all-local: clean front database-up
	# for database -> bad solution
	sleep 30
	cd $(current_path)/backend/config && sed -i 's/host: "fitness_db"/host: "localhost"/g' local.yaml
	cd $(current_path)/backend && go build
	cd $(current_path)/backend && ./backend

all-in-container:
	cd $(current_path)/backend && go build

docker-down:
	source ./scripts/all-down.sh

service-down:
	source ./scripts/service-down.sh

service-up: clean front
	cd ${current_path}/backend && go build

database-up:
	cd ${current_path}/backend/database && sudo docker-compose up --build -d

clean: docker-down

	cd $(current_path)/backend/config && sed -i 's/host: "fitness_db"/host: "localhost"/g' local.yaml

	if [ -d "$(current_path)/backend/database/cache" ]; then \
		sudo rm -rf "$(current_path)/backend/database/cache"; \
	fi

	if [ -d "$(current_path)/backend/frontend" ]; then \
		sudo rm -rf "$(current_path)/backend/frontend"; \
	fi

	if [ -d "$(current_path)/project" ]; then \
		sudo rm -rf "$(current_path)/project"; \
	fi

volume: front
	# set config for container
	cd $(current_path)/backend/config && sed -i 's/host: "localhost"/host: "fitness_db"/g' local.yaml

docker: clean

	echo "y" | sudo docker image prune -a
	cd $(current_path)
	mkdir -p $(current_path)/project
	cp -r $(current_path)/backend $(current_path)/project
	cp -r $(current_path)/frontend $(current_path)/project
	cp $(current_path)/go.mod $(current_path)/project
	cp -r $(current_path)/go.sum $(current_path)/project
	cp -r $(current_path)/Makefile $(current_path)/project
	cp -r $(current_path)/scripts $(current_path)/project

	# set config for docker-container
	cd $(current_path)/project/backend/config && sed -i 's/host: "localhost"/host: "fitness_db"/g' local.yaml

	cd $(current_path)/backend/database && sudo docker-compose up --build -d
	cd $(current_path)/scripts && sudo docker build -t fitness-project:latest .
	cd $(current_path)
	# need healthcheck. use bad solution
	sleep 120
	# try to connect. if database isn't ready need to recreate fitness-project
	sudo docker run -p 8080:8080  -v ./project:/etc/project -P --name fitness-project --link fitness_db:fitness_db --net database_default fitness-project:latest
	sudo rm -rf $(current_path)/project

update:

	source ./scripts/service-down.sh

	make volume
	sudo docker run -p 8080:8080 -v ./project:/etc/project -P --name fitness-project --link fitness_db:fitness_db --net database_default fitness-project:latest

all-compose: clean volume

	cd $(current_path)/backend
	sudo docker-compose up --build
