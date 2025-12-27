package main

import (
	"fmt"
	"maps"
)

// map -> hash, object, dict(python)
func main() {
	// creating map
	// syntax: map[key]value
	m := make(map[string]string)
	m["name"] = "golang"
	m["area"] = "backend"
	// get an element
	fmt.Println(m["name"])
	fmt.Println(m["area"])
	// NOTE: if key does not exists in the map then it returns zero(for int it gives zero, for string it gives empty string)
	fmt.Println(m["phone"]) //empty value(empty string)

	m2 := make(map[string]int)
	m2["age"] = 25
	fmt.Println(m2["age"])
	fmt.Println(m2["phone"]) // empty value(0)

	fmt.Println("map 1 length", len(m))
	fmt.Println("map 2 length", len(m2))

	// delete a key from map
	fmt.Println("before deleting area from map:", m)
	delete(m, "area")
	fmt.Println("after deleting area from map:", m)

	// empty the map
	clear(m)

	m3 := map[string]int{"price": 40, "phones": 3}
	fmt.Println("map3 values:", m3)
	//IMPORTANT: ok is an idiomatic way in golang
	v, ok := m3["price"]
	// NOTE: here map returns multiple value in first 'v' we are getting value of the key and in 'ok' we are getting true/false for value exists or not
	// fmt.Println("ok", ok)
	fmt.Println(v)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	// checking equality between different maps
	m4 := map[string]int{"price": 50, "phone": 1}
	m5 := map[string]int{"price": 50, "phone": 1}

	// fmt.Println("m4 == m5", m4 == m5) //NOT ALLOWED
	
	fmt.Println("m4 == m5:", maps.Equal(m4, m5))
}
