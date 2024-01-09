#!/bin/bash

cd .. && make all-in-container
cd ./backend && go build
./backend