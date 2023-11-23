package main

import (
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

const (
	natsURL   = "nats://localhost:4222"
	clusterID = "test-cluster"
	clientID  = "service-client1"
	subject   = "new-orders"
)

var sc stan.Conn

func main() {
	initNATS()
	defer closeNATS()
	jsonData := `
{
  "order_uid": "zxcvbnm12345test",
  "track_number": "XYZ123ABC",
  "entry": "XYZ",
  "delivery": {
    "name": "John Doe",
    "phone": "+1234567890",
    "zip": "54321",
    "city": "Metropolis",
    "address": "456 Main Street",
    "region": "Metro Region",
    "email": "john.doe@example.com"
  },
  "payment": {
    "transaction": "zxcvbnm12345test",
    "request_id": "qwerty789",
    "currency": "EUR",
    "provider": "bankpay",
    "amount": 2000,
    "payment_dt": 1638300000,
    "bank": "citybank",
    "delivery_cost": 120,
    "goods_total": 1880,
    "custom_fee": 50
  },
  "items": [
    {
      "chrt_id": 9876543,
      "track_number": "XYZ123ABC",
      "price": 800,
      "rid": "laptop0987654321",
      "name": "Laptop",
      "sale": 10,
      "size": "15 inches",
      "total_price": 720,
      "nm_id": 543210,
      "brand": "TechMaster",
      "status": 205
    },
    {
      "chrt_id": 8765432,
      "track_number": "XYZ123ABC",
      "price": 500,
      "rid": "headphones567890",
      "name": "Wireless Headphones",
      "sale": 15,
      "size": "Over-ear",
      "total_price": 425,
      "nm_id": 654321,
      "brand": "AudioBliss",
      "status": 200
    }
  ],
  "locale": "en",
  "internal_signature": "secure123",
  "customer_id": "johnd",
  "delivery_service": "expressdelivery",
  "shardkey": "5",
  "sm_id": 55,
  "date_created": "2023-04-15T12:30:00Z",
  "oof_shard": "2"
}
`
	err := publishJSONToNATS(jsonData)
	if err != nil {
		// Обработка ошибки
		log.Println("Failed to publish JSON to NATS:", err)
	}
}
func initNATS() {
	var err error
	for {
		sc, err = stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
		if err == nil {
			break
		}
		log.Printf("Не удалось подключиться к NATS: %v, повторная попытка через 5 секунд", err)
		time.Sleep(5 * time.Second)
	}
}

// Функция для отправки сообщения в NATS
// Функция для отправки JSON-строки в NATS
func publishJSONToNATS(jsonData string) error {
	// Публикация JSON-строки в NATS
	err := sc.Publish(subject, []byte(jsonData))
	if err != nil {
		log.Println("Error publishing JSON to NATS:", err)
		return err
	}

	log.Println("JSON published successfully to NATS:", jsonData)
	return nil
}
func closeNATS() {
	if sc != nil {
		err := sc.Close()
		if err != nil {
			log.Println(err)
		}
	}
}
