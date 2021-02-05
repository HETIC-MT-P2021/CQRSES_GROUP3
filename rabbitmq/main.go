package main

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/rabbitmq/consummer"
)

func main() {
	consummer.ReceiveFromRabbit()
}