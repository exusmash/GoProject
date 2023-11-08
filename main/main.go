package main

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "mydb"
)

type Order struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

//func (o Order) lalka() (int, string) {
//	return o.ID, o.Name
//}

var db *sql.DB

func main() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/addOrder", addOrder)
	http.HandleFunc("/getOrders", getOrders)
	http.HandleFunc("/getOrder", getOrder)
	http.HandleFunc("/deleteOrder", deleteOrder)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addOrder(w http.ResponseWriter, r *http.Request) {

}

func getOrders(w http.ResponseWriter, r *http.Request) {
	// Обработка вывода всех заказов
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	// Обработка вывода заказа по ID
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	// Обработка удаления заказа по ID
}
