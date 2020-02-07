package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Dial(addr string) (net.Conn, error) {
	return self.client.Dial("tcp", addr)
}

const (
	dbUser      string = "hzhang"
	dbPass      string = "12345"
	dbHost      string = "localhost:3306"
	dbName      string = "for_test"
	known_host  string = "/Users/huajiezhang/.ssh/known_hosts"
	private_key string = "/Users/huajiezhang/.ssh/id_rsa"
)

var db *sql.DB

var phone = `1234567890
123 456 7891
(123) 456 7892
(123) 456-7893
123-456-7894
123-456-7890
1234567892
(123)456-7892`

func main() {

	hostKey, err := knownhosts.New(known_host)
	if err != nil {
		log.Println(err)
	}

	key, err := ioutil.ReadFile(private_key)
	if err != nil {
		log.Println(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Println("unable to parse private key: %v", err)
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

	mysql.RegisterDial("mysql+tcp", (&ViaSSHDialer{client}).Dial)

	if db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@mysql+tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)); err == nil {
		defer db.Close()

		fmt.Printf("---Successfully connected to the db---\n")

		// drop_tb()

		// create_tb()

		// insert_one()

		// mult_insert()

		// use_tx()

		// select_query()

	}

	useGorm()
}

func insert_one() {
	insertStmt := `insert into phone (phone)
		values('02345678900')
		;`
	_, err := db.Exec(insertStmt)
	if err != nil {
		log.Println(err)
	}
}

func drop_tb() {
	Stmt := "drop table for_test.phone"
	_, err := db.Exec(Stmt)
	if err != nil {
		log.Println(err)
	}
}

func create_tb() {
	createStmt := `
		create table for_test.phone (
		id int primary key AUTO_INCREMENT,
		phone varchar(10))
		`
	_, err := db.Exec(createStmt)
	if err != nil {
		log.Println(err)
	}
}

func mult_insert() {
	stmt, err := db.Prepare("insert into phone (phone) values (?)")
	if err != nil {
		log.Println(err)
	}

	for _, v := range strings.Split(phone, "\n") {
		_, err := stmt.Exec(numOnly(v))
		if err != nil {
			log.Println(err)
		}
	}
}

func use_tx() {
	tx, err := db.Begin()
	defer tx.Rollback()

	if err != nil {
		log.Println(err)
	}

	stmt, err := db.Prepare("insert into phone (phone) values (?)")
	if err != nil {
		log.Println(err)
	}

	for _, v := range strings.Split(phone, "\n") {
		_, err := stmt.Exec(numOnly(v))
		if err != nil {
			log.Println(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func select_query() {
	selectStmt := `select id, phone from phone`
	rows, err := db.Query(selectStmt)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var (
		id    int32
		phone string
	)

	for rows.Next() {
		err = rows.Scan(&id, &phone)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(id, phone)
	}
	rows.Close()
}

func numOnly(s string) string {
	var ret string
	for _, v := range s {
		if v >= '0' && v <= '9' {
			ret += string(v)
		}
	}
	return ret
}

type Phone struct {
	// gorm.Model
	Id    int    `sql:"AUTO_INCREMENT"`
	Phone string `sql:"varchar(10)"`
}

func useGorm() {
	db, err := gorm.Open("sqlite3", "phone.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Phone{})

	tx := db.Begin()
	// defer tx.Rollback()

	for _, v := range strings.Split(phone, "\n") {
		if err = tx.Create(&Phone{Phone: numOnly(v)}).Error; err != nil {
			tx.Rollback()
			log.Println(err)
		}
	}
	tx.Commit()
}

func testUseGorm() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Phone{}) // this is needed???

	// p := Phone{Phone: "125"}
	// db.Create(&p)

	var p0 Phone
	db.First(&p0)
	fmt.Println(p0)

	var p []Phone
	db.Find(&p)
	fmt.Println(p)
}
