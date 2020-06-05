# authentication
It's an authentication application written in Go.

Golang - Basic Login/Register services implemented using
- gRPC
- Database (MySQL)

### Do Insert an application.yaml file with the structure shown in properties dir

### Make MySQL DB named `authentication` and run 

> CREATE TABLE `authentication.user_credentials`
>          ( `id` INT unsigned NOT NULL AUTO_INCREMENT,
>            `username` VARCHAR(30) NOT NULL,
>            `password` VARCHAR(100) NOT NULL,
>            `phone` VARCHAR(13) NOT NULL,
>            `created_at` timestamp NOT NULL DEFAULT current_timestamp,
>            `updated_at` timestamp NOT NULL DEFAULT current_timestamp,
>            `deleted_at` varchar(45) NULL default '0000-00-00 00:00:00',
>             PRIMARY KEY (username, id)); 

Set the go path by running

> export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))

After installing all the dependencies

> go run AuthenticationApp.go

It will start the grpc server at the address mentioned in your application.yaml file

To run the Unit tests

> go test server/GrpcServer_test.go -v
