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
14. Run full tests suites
15. Generate mocks for repos

## Section 4 - Protobuf
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
8. Flesh out us+-+----ers `Destroy` route
9. Build initial users tests file
10. Build `Create` tests
11. Build `FindById` tests
12. Build `FindByEmail` tests
13. Build `Update` tests
14. Build `Destroy` tests
15. Run full test suite for users
16. Flesh out auth `Login` route
17. Build initial auth tests file
18. Build `Login` tests
19. Run full test suite for auth
20. Build api file to integrate routes
21. Create a main.go file to run the server

## Section 6 - gRPC command line testing
1. Install grpc_cli
2. Run grpc_cli ls command to see all available services
3. Run grpc_cli call command to create a user
4. Run grpc_cli call command to login with that user
