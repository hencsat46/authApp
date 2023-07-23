package database

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	psx "github.com/jackc/pgx/v5"
)

const DB_URL = "postgresql://postgres:forstudy@localhost:5432/AuthDB"

func AddEncodeData(data map[string]string) {

	newJson, _ := json.Marshal(data)
	err := os.WriteFile("/home/ilya/prog_go/src/authApp/authServ/shitDB/database.json", newJson, 0666)

	fmt.Println(err)
}

// func GetData(data map[string]string) bool {

// }

func ConnectDB() {
	conn, err := psx.Connect(context.Background(), os.Getenv(DB_URL))

	defer conn.Close(context.Background())

	fmt.Println(err)
}
