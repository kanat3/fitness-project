#!/bin/bash

res=$(sudo docker ps -a | awk '{print $(NF)}' | grep fitness_db)
if [ ! -z $res ]; then \
    sudo docker stop fitness_db; \
    sudo docker rm fitness_db; \
fi