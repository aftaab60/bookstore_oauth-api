package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	//docker run -dt --name cassandra --hostname 127.0.0.1 -p 9042:9042 cassandra
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	cluster.DisableInitialHostLookup = true

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
