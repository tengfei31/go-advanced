package distribute

import (
	"fmt"
	"log"
	"testing"
)

func TestSonyflake(t *testing.T) {
	id, err := generateSonyflake()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("id:", id)
	fmt.Print("id:", id)
}
