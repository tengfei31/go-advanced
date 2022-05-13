package distribute

import (
	"log"
	"math/rand"
	"testing"
)

func TestSonyflake(t *testing.T) {
	id, err := generateSonyflake()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("id:", id)
}


func TestRand(t *testing.T) {
	slice := rand.Perm(7)
	log.Print(slice)
}
