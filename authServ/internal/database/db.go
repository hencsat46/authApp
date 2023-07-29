package database

import (
	"context"
	"errors"
	"fmt"
	"log"

	psx "github.com/jackc/pgx/v5"
	crypto "golang.org/x/crypto/bcrypt"
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

	hashedPassword, _ := crypto.GenerateFromPassword([]byte(structData.password), 5)
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
		return nil
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

func Authorization(data map[string]string) error {

	conn := ConnectDB()
	defer CloseConnection(conn)

	//

	row, err := conn.Query(context.Background(), fmt.Sprintf("SELECT * FROM returnPassword('%s')", data["username"]))

	if err != nil {
		log.Fatal("Cannot authorizate", err)
		return err
	}

	var hashedPassword string
	row.Next()
	row.Scan(&hashedPassword)

	fmt.Println(hashedPassword)

	status := crypto.CompareHashAndPassword([]byte(hashedPassword), []byte(data["password"]))

	if status == nil {
		fmt.Println("Authorization complete")
		return nil
	}

	return errors.New("Invalid login or password")

}
