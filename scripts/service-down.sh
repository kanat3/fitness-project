#!/bin/bash

res=$(sudo docker ps -a | awk '{print $(NF)}' | grep fitness-project)
if [ ! -z $res ]; then \
    sudo docker stop fitness-project; \
    sudo docker rm fitness-project; \
fi