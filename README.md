# CloudSSH

:cloud: SSH key cloud management tool inspired by Bitwarden.

## Usage

### Server

```shell
go run cmd/server/main.go
```

### Client

```shell
go run cmd/client/main.go -h
go run cmd/client/main.go signup -s http://localhost -u hi@example.com -p password
go run cmd/client/main.go server -h
go run cmd/client/main.go server list
```

## TODO

- [ ] account
  - [x] sign up
  - [x] login in, keep user status at client
  - [x] logout
  - [ ] change password
- [x] server
  - [x] create server
  - [x] list server
  - [x] connect to server
  - [x] delete server
  - [x] update server
- [ ] organization
  - [x] create organization
  - [ ] edit organization
  - [ ] delete organization
  - [ ] add user
  - [ ] edit user permission
  - [ ] delete user
  - [x] add server
  - [x] edit server
  - [x] delete server

Thanks for those awesome work:

```go
https://github.com/jcs/rubywarden/blob/master/API.md
https://github.com/philhug/bitwarden-client-go/blob/master/bitwarden/authentication.go
https://github.com/VictorNine/bitwarden-go/blob/master/internal/auth/auth.go
https://github.com/bitwarden/server/issues/26
```
