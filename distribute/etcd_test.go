package distribute

import (
	"log"
	"testing"
)

func TestEtcdGet(t *testing.T) {
	res, err := getConfig("test", true)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)
}

func TestEtcdPut(t *testing.T) {
	for key, val := range config() {
		err := putConfig(key, val)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}


func TestEtcdWatch(t *testing.T) {
	watchConfig("test")
}
