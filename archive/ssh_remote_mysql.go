package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Dial(addr string) (net.Conn, error) {
	return self.client.Dial("tcp", addr)
}

func main() {
	hostKey, err := knownhosts.New("/Users/huajiezhang/.ssh/known_hosts")
	if err != nil {
		log.Fatal(err)
	}

	key, err := ioutil.ReadFile("/Users/huajiezhang/.ssh/id_rsa")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "pi",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKey,
	}

	client, err := ssh.Dial("tcp", "192.168.1.137:22", config)
	if err != nil {
		log.Println("failed", err)
	}

	// session, err := client.NewSession()
	// if err != nil {
	// 	log.Println("session", err)
	// }
	// defer session.Close()

	mysql.RegisterDial("mysql+tcp", (&ViaSSHDialer{client}).Dial)

	dbUser := "hzhang"
	dbPass := "12345"
	dbHost := "localhost:3306"
	dbName := "stock"

	if db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@mysql+tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)); err == nil {

		defer db.Close()
		fmt.Printf("Successfully connected to the db\n")
		if rows, err := db.Query("SELECT count(*) FROM pdd"); err == nil {
			for rows.Next() {
				var count int64
				rows.Scan(&count)
				fmt.Printf("count: %v", count)
			}
			rows.Close()
		} else {
			fmt.Printf("Failure: %s", err.Error())
		}

		db.Close()
	}
	// var b bytes.Buffer
	// session.Stdout = &b
	// if err := session.Run("/usr/bin/whoami"); err != nil {
	// 	log.Fatal("Failed to run: " + err.Error())
	// }
	// fmt.Println(b.String())

}
