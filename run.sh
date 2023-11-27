#!/bin/bash

rm -rf ./backend/static

export CONFIG_PATH=config/local.yaml

cp -r frontend/static ./backend/

cd ./backend && go build

./backend
