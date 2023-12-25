current_path=$(shell pwd)

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

clean:
	if [ -d "$(current_path)/backend/static" ]; then \
		rm -rf "$(current_path)/backend/static"; \
	fi
	if [ -d "$(current_path)/backend/assets" ]; then \
		rm -rf "$(current_path)/backend/assets"; \
	fi