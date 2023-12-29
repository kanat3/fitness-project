current_path=$(shell pwd)
SHELL := /bin/bash

front: clean
	if [ ! -d "$(current_path)/backend/static" ]; then \
		mkdir -p "$(current_path)/backend/static"; \
	fi
	
	cp $(current_path)/frontend/*.html $(current_path)/backend/static/
	
	# bad path stucture
	####
	cp $(current_path)/frontend/*.jpg $(current_path)/backend/static/
	####

	if [ ! -d "$(current_path)/backend/assets" ]; then \
		mkdir -p "$(current_path)/backend/assets"; \
	fi

	cp -r $(current_path)/frontend/assets $(current_path)/backend

all: clean front
	cd $(current_path)/backend && go build

docker-down:
	source ./scripts/all-down.sh

clean: docker-down

	if [ -d "$(current_path)/backend/database/cache" ]; then \
		sudo rm -rf "$(current_path)/backend/database/cache"; \
	fi

	if [ -d "$(current_path)/backend/static" ]; then \
		sudo rm -rf "$(current_path)/backend/static"; \
	fi

	if [ -d "$(current_path)/backend/assets" ]; then \
		sudo rm -rf "$(current_path)/backend/assets"; \
	fi

	if [ -d "$(current_path)/project" ]; then \
		sudo rm -rf "$(current_path)/project"; \
	fi

volume:
	
	if [ -d "$(current_path)/backend/database/cache" ]; then \
		sudo rm -rf "$(current_path)/backend/database/cache"; \
	fi

	sudo rm -rf $(current_path)/project
	cd $(current_path)
	mkdir -p $(current_path)/project
	cp -r $(current_path)/backend $(current_path)/project
	cp -r $(current_path)/frontend $(current_path)/project
	cp $(current_path)/go.mod $(current_path)/project
	cp -r $(current_path)/go.sum $(current_path)/project
	cp -r $(current_path)/Makefile $(current_path)/project
	cp -r $(current_path)/scripts $(current_path)/project

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

	cd $(current_path)/backend/database && sudo docker-compose up --build -d
	sudo docker build -t fitness-project:latest .
	# healthcheck

	sudo docker run -p 8080:8080  -v ./project:/etc/project -P --name fitness-project --link fitness_db:fitness_db --net database_default fitness-project:latest
	sudo rm -rf $(current_path)/project

update:
	source ./scripts/service-down.sh

	make volume
	sudo docker run -p 8080:8080 -v ./project:/etc/project -P --name fitness-project --link fitness_db:fitness_db --net database_default fitness-project:latest