
### Cassandra used

https://academy.datastax.com/resources/getting-started-apache-cassandra-and-go

http://planetcassandra.org/create-a-keyspace-and-table/

Set Cassandra config:
vi /etc/cassandra/default.conf/cassandra.yaml
search rpc_address and set to 0.0.0.0 (only on a test host)

systemctl restart cassandra
systemctl stop firewalld.service

go get github.com/gocql/gocql

Try it out:

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "demo"
	session, _ := cluster.CreateSession()

### TODO

- User roles - modification
- SMTP email validation (key based)
- Possible remove dependance on bcrypt
