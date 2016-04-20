package main

import "github.com/rabem00/httpauth"

var (
	backend     httpauth.CassandraAuthBackend
	aaa         httpauth.Authorizer
	roles       map[string]httpauth.Role
	port        = 8009
	backendfile = "auth.leveldb"
)

func main() {
	var err error

	backend, err = httpauth.NewCassandraBackend("192.168.1.4", "demo")
	if err != nil {
		panic(err)
	}
	backend.User("Jones")
	backend.Close()
}
