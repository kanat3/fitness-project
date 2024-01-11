#!/bin/bash

#test for signup - OK
curl -i -X POST -H 'Content-Type: application/json' -d '{"first_name": "Anna", "second_name": "K", "last_name": "V", "email": "test3@mail.ru", "phone": "87505197423", "password": "12334"}' http://localhost:8080/signup
echo ""
#test login - OK
curl -i -X POST -H 'Content-Type: application/json' -d '{"email": "test3@mail.ru", "password": "12334"}' http://localhost:8080/login
echo ""
#test login - ERROR
curl -i -X POST -H 'Content-Type: application/json' -d '{"email": "b@mail.ru", "password": "1234"}' http://localhost:8080/login
echo ""
#test login - ERROR
curl -i -X POST -H 'Content-Type: application/json' -d '{"email": "test3@mail.ru", "password": "wrong-password"}' http://localhost:8080/login
echo ""
#test reset email - ERROR
curl -i -X POST -H 'Content-Type: application/json' -d '{"email": "a3@mail.ru", "phone": "87505197429", "password": "test1234"}' http://localhost:8080/reset-email
echo ""
#test reset email - OK
curl -i -X POST -H 'Content-Type: application/json' -d '{"email": "new-test@mail.ru", "phone": "87505197423", "password": "12334"}' http://localhost:8080/reset-email
echo ""
#test reset password - ERROR
curl -i -X POST -H 'Content-Type: application/json' -d '{"email": "wrong@mail.ru", "password": "12334"}' http://localhost:8080/reset-password
echo ""
#test reset password - OK - test without jwt
curl -i -X POST -H 'Content-Type: application/json' -d '{"email": "new-test@mail.ru", "password": "wrong"}' http://localhost:8080/reset-password
echo ""
#get jwt
token=$(curl -i -X POST -H 'Content-Type: application/json' -d '{"email": "new-test@mail.ru", "password": "12334"}' http://localhost:8080/login | grep -oP '(?<=token=)[^;]*')
echo ""
echo "token: token=$token"
echo ""
#test reset password - OK - test without jwt
curl -i -X POST -H 'Content-Type: application/json' -b token=$token -d '{"email": "new-test@mail.ru", "password": "12334"}' http://localhost:8080/reset-password
echo ""
# test account - OK
curl -i -X POST -H 'Content-Type: application/json' http://localhost:8080/account/99
echo ""
# test set bid - OK
curl -i -X POST -H 'Content-Type: application/json' -d '{"optional_goal": "my test goal", "optional_message": "hello tinder"}' http://localhost:8080/account/99/set_bid
echo ""
# test get bid - OK
curl -i -X GET -H 'Content-Type: application/json' http://localhost:8080/account/99/get_bid
echo ""