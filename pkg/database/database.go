package database

import (
	"github.com/gocql/gocql"
	"log"
	"os"
	"strconv"
)

var DbSession *gocql.Session

func ConnectDb() {
	portStr := os.Getenv("DB_PORT")
	portInt, errConv := strconv.Atoi(portStr)
	if errConv != nil {
		log.Fatalf("Error in DB connection: Converting Port: %v", errConv)
	}

	log.Println("Connecting to Cassandra DB at port " + portStr)
	var err error
	cluster := gocql.NewCluster(os.Getenv("CLUSTER"))
	cluster.Keyspace = os.Getenv("KEYSPACE")
	cluster.Consistency = gocql.Quorum
	cluster.Port = portInt
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: os.Getenv("USERNAME"), Password: os.Getenv("PASSWORD"), AllowedAuthenticators: []string{"com.instaclustr.cassandra.auth.InstaclustrPasswordAuthenticator"}}
	DbSession, err = cluster.CreateSession()
	if err != nil {
		log.Fatalf("Error connecting to Cassandra DB: %v", err)
	}
	log.Println("Connected to Cassandra DB")
}
