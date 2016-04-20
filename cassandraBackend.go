package httpauth

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

type CassandraAuthBackend struct {
	clusterURL string
	keyspace   string
	session    *gocql.Session
}

//func (b CassandraAuthBackend) connect() *gocql.CollectionType {
//session := b.session.
//return session.DB(b.keyspace).C("goauth")
//}

func NewCassandraBackend(clusterURL string, keyspace string) (b CassandraAuthBackend, e error) {
	cluster := gocql.NewCluster(clusterURL)
	cluster.Keyspace = keyspace
	session, err := cluster.CreateSession()
	b.session = session
	if err != nil {
		log.Fatal(err)
	}
	return
}

// User returns the user with the given username. Error is set to
// ErrMissingUser if user is not found.
func (b CassandraAuthBackend) User(username string) (user UserData, e error) {
	var result UserData
	var firstname string
	var age int

	c := b.session
	defer c.Close()

	if err := c.Query("SELECT firstname, age FROM users WHERE lastname='Smith'").Scan(&firstname, &age); err != nil {
		log.Fatal(err)
	}

	result.Username = firstname
	fmt.Println(result.Username)
	return result, nil
}

// Close cleans up the backend once done with. This should be called before
// program exit.
func (b CassandraAuthBackend) Close() {
	if b.session != nil {
		b.session.Close()
	}
}
