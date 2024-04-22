package db

import (
	"fmt"
	"sync"
)

type MemoryDB struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		data: make(map[string]string),
	}
}

func (db *MemoryDB) Set(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}

func (db *MemoryDB) Get(key string) (string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	value, exists := db.data[key]
	return value, exists
}

func (db *MemoryDB) Del(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.data, key)
}

func (db *MemoryDB) Exists(key string) bool {
	db.mu.RLock()
	defer db.mu.RUnlock()
	_, exists := db.data[key]
	return exists
}

func (db *MemoryDB) Incr(key string) (int, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	value, exists := db.data[key]
	if !exists {
		return 0, fmt.Errorf("key not found")
	}
	// Assuming the value is an integer
	intValue := 0
	_, err := fmt.Sscanf(value, "%d", &intValue)
	if err != nil {
		return 0, fmt.Errorf("invalid value")
	}
	intValue++
	db.data[key] = fmt.Sprintf("%d", intValue)
	return intValue, nil
}

func (db *MemoryDB) Decr(key string) (int, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	value, exists := db.data[key]
	if !exists {
		return 0, fmt.Errorf("key not found")
	}
	// Assuming the value is an integer
	intValue := 0
	_, err := fmt.Sscanf(value, "%d", &intValue)
	if err != nil {
		return 0, fmt.Errorf("invalid value")
	}
	intValue--
	db.data[key] = fmt.Sprintf("%d", intValue)
	return intValue, nil
}
