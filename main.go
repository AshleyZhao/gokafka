package main

import (
	"fmt"
	"log"
	"time"
	"rsc.io/quote"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	fmt.Println(quote.Go())
}
