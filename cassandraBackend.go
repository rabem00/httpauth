package httpauth

import (
	"log"

	"github.com/gocql/gocql"
)

type CassandraAuthBackend struct {
	cluster  string
	keyspace string
  session *gocql.Session
}

func (b CassandraAuthBackend) connect() *gocql.CollectionType {
	session := b.session.Copy()
	return session.DB(b.keyspace).C("goauth")
}

//     backend = httpauth.CassandraAuthBackend("mongodb://127.0.0.1/", "auth")
//     defer backend.Close()
func NewCassandraBackend(cluster string, keyspace string) (b CassandraAuthBackend, e error) {
	// Set up connection to database
	b.cluster = cluster
	b.keyspace = keyspace

	cluster := gocql.NewCluster(b.cluster)
	cluster.Keyspace = b.keyspace
	b.session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
  return
}

// User returns the user with the given username. Error is set to
// ErrMissingUser if user is not found.
func (b CassandraAuthBackend) User(username string) (user UserData, e error) {
	var result UserData

	c := b.connect()
	defer c.Database.Session.Close()

	err := c.Find(bson.M{"Username": username}).One(&result)
	if err != nil {
		return result, ErrMissingUser
	}
	return result, nil
}

// Close cleans up the backend once done with. This should be called before
// program exit.
func (b CassandraAuthBackend) Close() {
	if b.session != nil {
		b.session.Close()
	}
}
