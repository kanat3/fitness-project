#!/bin/bash

#test for signup - OK
curl -i -X POST -H 'Content-Type: application/json' -d '{"first_name": "Anna", "second_name": "K", "last_name": "V", "email": "test@mail.ru", "phone": "87505197421", "password": "12334"}' http://localhost:8080/signup

#test login - ERROR
curl -i -X POST -H 'Content-Type: application/json' -d '{"name": "Anna", "email": "b@mail.ru", "password": "1234"}' http://localhost:8080/login

#test login - OK
curl -i -X POST -H 'Content-Type: application/json' -d '{"name": "NoMatter", "email": "test@mail.ru", "password": "12334"}' http://localhost:8080/login

#test reset email - ERROR
curl -i -X POST -H 'Content-Type: application/json' -d '{"first_name": "Anna", "second_name": "K", "last_name": "V", "email": "a3@mail.ru", "phone": "87505197429", "password": "test1234"}' http://localhost:8080/reset-email

#test reset email - OK
curl -i -X POST -H 'Content-Type: application/json' -d '{"first_name": "Anna", "second_name": "K", "last_name": "V", "email": "new-test@mail.ru", "phone": "87505197421", "password": "test1234"}' http://localhost:8080/reset-email
