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

all-local: clean front
	cd $(current_path)/backend && go build

docker-down:
	source ./scripts/all-down.sh

service-down:
	source ./scripts/service-down.sh

service-up: clean front
	cd ${current_path}/backend && go build
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

	# set config for container
	cd $(current_path)/project/backend/config && sed -i 's/host: "local"/host: "fitness_db"/g' local.yaml

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

	# set config for docker-container
	cd $(current_path)/project/backend/config && sed -i 's/host: "local"/host: "fitness_db"/g' local.yaml

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
	cd $(current_path)/project
	sudo docker-compose up --build
