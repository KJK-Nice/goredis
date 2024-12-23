package client

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {

	client, err := New("localhost:6969")
    if err != nil {
        log.Fatal(err)
    }
	for i := 0; i < 10; i++ {
        fmt.Println("SET =>", fmt.Sprintf("foo_%d", i))
		if err := client.Set(context.Background(), fmt.Sprintf("foo_%d", i), fmt.Sprintf("bar_%d", i)); err != nil {
			log.Fatal(err)
		}
		val, err := client.Get(context.Background(), fmt.Sprintf("foo_%d", i));
        if err != nil {
			log.Fatal(err)
		}
        fmt.Println("GET =>", val)
	}
}
