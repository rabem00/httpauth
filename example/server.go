package main

import "github.com/apexskier/httpauth"

var (
	backend     httpauth.CassandraAuthBackend
	aaa         httpauth.Authorizer
	roles       map[string]httpauth.Role
	port        = 8009
	backendfile = "auth.leveldb"
)

func main() {
	var err error

	backend, err = httpauth.New

}
