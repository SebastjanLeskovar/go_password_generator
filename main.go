package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pool = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789"

// Generate a random string of characters of lenth l
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		randPos := rand.Intn(len(pool))
		bytes[i] = pool[randPos]
	}
	return string(bytes)
}

// Generate a random array of characters of lenth l
func randomArray(l int) []string {
	a := make([]string, l)
	for i := 0; i <= l-1; i++ {
		randChar := randomString(1)
		a[i] = randChar
	}
	return a
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// randPass := randomString(12)
	randArray := randomArray(10)
	fmt.Println(randArray)
}
