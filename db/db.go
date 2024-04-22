package db

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type DataObject struct {
	data       interface{}
	ttl        *int // time to live for any value in secs (Eg. 10 for 10 seconds)
	created_at int  // creation time in the form of epoch in seconds
}

type KVDatabase struct {
	values map[string]DataObject
	mu     sync.RWMutex
}

func CreateDB() *KVDatabase {
	return &KVDatabase{values: make(map[string]DataObject)}
}

func CreateDataObject(value interface{}) DataObject {
	return DataObject{data: value, ttl: nil, created_at: int(time.Now().Unix())}
}

func (db *KVDatabase) Exists(key string) bool {
	_, exists := db.values[key]
	return exists
}

func (data DataObject) PrintObject() {
	fmt.Printf("data: %v, ttl: %v, created_at: %d \n", data.data, data.ttl, data.created_at)
}

func (db *KVDatabase) Set(key string, value DataObject) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	_, exists := db.values[key]
	if exists {
		log.Fatalf("Value with key '%s' already exists", key)
		return false
	}
	db.values[key] = value
	return true
}

func (db *KVDatabase) Get(key string) (DataObject, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	value, exists := db.values[key]
	if !exists {
		ShowKeyNotFoundError(key)
	}
	return value, exists
}

func (db *KVDatabase) Incr(key string) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	value, exists := db.values[key]
	if !exists {
		ShowKeyNotFoundError(key)
	}
	if intValue, ok := value.data.(int); ok {
		value.data = intValue + 1
		db.values[key] = value
	} else {
		log.Fatalf("The value of key: %s is not an integer.", key)
	}
	return true
}

func (db *KVDatabase) Decr(key string) bool {
	db.mu.Lock()
	defer db.mu.Unlock()
	value, exists := db.values[key]
	if !exists {
		ShowKeyNotFoundError(key)
	}
	if intValue, ok := value.data.(int); ok {
		value.data = intValue - 1
		db.values[key] = value
	} else {
		log.Fatalf("The value of key: %s is not an integer.", key)
	}
	return true
}
