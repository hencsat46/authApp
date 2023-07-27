package database

import (
	"context"
	"fmt"
	"log"

	psx "github.com/jackc/pgx/v5"
	cryp "golang.org/x/crypto/bcrypt"
)

type userData struct {
	username string
	password string
}

const DB_URL = "postgresql://postgres:forstudy@localhost:5432/AuthDB"

func AddEncodeData(data map[string]string) {

	for _, value := range data {
		if len(value) == 0 {
			log.Fatal("Username or password is invalid")
			return
		}
	}

	var structData userData = userData{data["username"], data["password"]}

	hashedPassword, _ := cryp.GenerateFromPassword([]byte(structData.password), 5)
	conn := ConnectDB()
	defer CloseConnection(conn)

	_, err := conn.Query(context.Background(), fmt.Sprintf("INSERT INTO users(username, passwd) VALUES ('%s', '%s');", structData.username, string(hashedPassword)))

	fmt.Println(err)

}

func ConnectDB() *psx.Conn {
	config, _ := psx.ParseConfig(DB_URL)

	conn, err := psx.ConnectConfig(context.Background(), config)

	if err != nil {
		log.Fatal("Cannot to connect to database", err)
	}

	return conn

}

func ShowUsers(connection *psx.Conn) {
	rows, err := connection.Query(context.Background(), "SELECT * FROM users;")

	if err != nil {
		log.Fatal("Cannot execute query", err)
	}

	for rows.Next() {
		var username string
		var password string

		rows.Scan(&username, &password)

		fmt.Printf("Username: %v\nPassword: %v\n", username, password)
	}
}

func CloseConnection(connection *psx.Conn) {
	connection.Close(context.Background())
}
