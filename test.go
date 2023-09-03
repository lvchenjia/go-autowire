package main

import (
	"math/rand"
	"time"
)

type user struct {
	id   int
	name string
	age  int
}

type admin struct {
	id    int
	name  string
	age   int
	level int
}

// Autowire Test
func test1() {
	execQuery(createInsertQueryAuto(user{
		id:   100,
		name: "Tom",
		age:  rand.Intn(100),
	}))
	showTable("user")

	execQuery("delete from user where name = 'Tom'")
	showTable("user")
}

// Auto Create Table
func test2() {
	execQuery(createTable("admin", admin{}))

	execQuery(createInsertQueryAuto(admin{
		id:    100,
		name:  "Tom",
		age:   rand.Intn(100),
		level: 100,
	}))
	showTable("admin")
	deleteTable("admin")
}

func test() {
	rand.Seed(time.Now().UnixNano())
	printlnGreen("Test1: Autowire test")
	test1()
	printlnGreen("Test2: Create table test")
	test2()
}
