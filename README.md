learning-go-gRPC
====================

## Notes
https://github.com/happilymarrieddad/learning-go-gRPC


## Things to put in your bash
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:/home/comp/go/bin

## Other things to install
# Ubuntu
sudo apt install mysql-server

## Snippets
{
    "Omega Assertion": {
        "prefix": "omega",
        "body": "Ω($1)",
        "description": "Ginkgo Gomega Assertion"
    },
}

## Section 1
1. Intro
2. Setup ENV

## Section 2 - 
1. setup migrations
    - installing
    - https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
    - migrate create -dir migrations -ext sql create_users_table
    - migrate -path ./migrations -database mysql://root:pass@/grpc -verbose up
2. Create `User` types
3. Add validation

## Section 3 - setup repos
1. Create global repo
2. Create users repo
    - https://xorm.io/
3. Add `Create` func
4. Add `FindById` func
5. Add `FindByEmail` func
6. Add `Update` func
8. Discuss testing and install gomega && ginkgo
    - ginkgo bootstrap
    - IF YOU WANT TO USE SQL-MOCK, please take a look at git branch `section3-lesson10`
9. Add `Create` tests
    - ginkgo -v -failFast --focus="UsersRepo"
10. Add `FindById` tests
11. Add `FindByEmail` tests
12. Add `Update` tests
15. Generate mocks for repos

## Section 4 - Protobuf
If you get the wrong version of protobuf use these commands

# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.2.0/protoc-3.2.0-linux-x86_64.zip

# Unzip
unzip protoc-3.2.0-linux-x86_64.zip -d protoc3

# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/

# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/

# Optional: change owner
sudo chown $USER /usr/local/bin/protoc
sudo chown -R $USER /usr/local/include/google

# I did change the last step to:
ln -s /protoc3/bin/protoc /usr/bin/protoc

1. Intro to protobufs
2. Add `types` protobuf
3. Add `users` protobuf
4. Add `auth` protobuf
5. Generate protobuf files

## Section 5 - API
1. Create `utils` file
2. Create globalRepo intercepter
3. Create users handler
4. Flesh out users `Create` route
5. Flesh out users `FindById` route
6. Flesh out users `FindByEmail` route
7. Flesh out users `Update` route
9. Build initial users tests file
10. Build `Create` tests
11. Build `FindById` tests
12. Build `FindByEmail` tests
13. Build `Update` tests
14. Run full test suite for users
15. Create a main.go file to run the server

## Section 6 - gRPC command line testing
1. Install grpc_cli
    - git clone -b $(curl -L https://grpc.io/release) https://github.com/grpc/grpc
    - cd grpc
    - git submodule update --init
    - make grpc_cli
    - sudo make install grpc_cli
    - IF IT FAILS ON YOU DO THE FOLLOWING!!
        - I just found out a workaround. If you add -Wno-unused-variable to the cppflags at line 356 of Makefile at grpc root folder, the complier will ignore the unused variable warning/error.
        - https://github.com/grpc/grpc/issues/16739
        - cp ./bins/opt/grpc_cli ~/go/bin
2. Run grpc_cli ls command to see all available services
3. Run grpc_cli call command to create a user
    - grpc_cli call localhost:8080 users.V1Users.Create "newUser:{firstName:'Nick',lastName:'Doe2',email:'foo@bar.com',password:'1234',confirmPassword:'1234'}"

## Section 7 - JWT
1. Create Auth repo
2. Add `GetNewClaims` func
 - go get -u github.com/pascaldekloe/jwt
3. Add `GetSignedToken` func
4. Add `GetDataFromToken` func
5. Create tests file
6. Build `GetNewClaims` tests
7. Build `GetSignedToken` tests
8. Build `GetDataFromToken` tests
9. Run full test suite for auth
10. Flesh out auth `Login` route 
11. Build initial auth tests file
12. Build `Login` tests
13. Run full test suite for auth
14. Build api file to integrate routes
15. Run grpc_cli call command to login with a user
    - grpc_cli call localhost:8080 auth.V1Auth.Login "email:'foo@bar.com', password: '1234'"
16. Intro to gRPC authentication
17. Add JWT field to routes
18. Add JWT checks to the interceptor
    - grpc_cli call localhost:8080 auth.V1Auth.Login "email:'foo@bar.com', password: '1234'"
    - grpc_cli call localhost:8080 users.V1Users.Create "newUser:{firstName:'Nick2',lastName:'Doe2',email:'foo2@bar.com',password:'1234',confirmPassword:'1234'}"
19. Test routes



## GRPC-WEB
https://github.com/grpc/grpc-web
https://github.com/improbable-eng/grpc-web/tree/master/go/grpcwebproxy

install into our project

npm install grpc grpc-web google-protobuf --save
