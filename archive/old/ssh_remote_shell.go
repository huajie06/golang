package archive

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func mainSSHShell() {
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

	session, err := client.NewSession()
	if err != nil {
		log.Println("session", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

}
