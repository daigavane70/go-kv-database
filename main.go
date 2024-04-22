package main

import (
	"ginrev/db"
)

func main() {
	redis := db.CreateDB()

	key1 := "key1"
	intVal1 := "intVal1"

	redis.Set(key1, db.CreateDataObject("Message 1"))
	redis.Set(intVal1, db.CreateDataObject(1))
	redis.Incr(intVal1)
	redis.Incr(intVal1)
	redis.Incr(intVal1)
	redis.Decr(intVal1)

	obj1, exists := redis.Get(key1)
	if exists {
		obj1.PrintObject()
	}
	obj2, exists := redis.Get(intVal1)
	if exists {
		obj2.PrintObject()
	}
}
