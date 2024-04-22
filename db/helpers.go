package db

import "log"

func ShowKeyNotFoundError(key string) {
	log.Fatalln("No value exists with key: ", key)
}
