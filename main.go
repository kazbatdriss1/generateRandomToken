package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generateRandomString generates a random string of length 20
func generateRandomString() string {
	const charset = " ‎ⵔⵕⵖⵙⵚⵛ ‌‎​ⵎⵏⵓⵜⵟⵡⵢⵣⵥⵯⵍⵊЦЩЪЫЬДйлЮяюфⵉⵇⵅⵄⵃⵀⴽⴼⴻБГПФЦЩЪЫЬДйлЮяюфЖщⴹⴷⴳⴱⴰⴲⴴⴵⴶⴸⴺⴿⵁⵒⵝⵞⵠ"
	var p *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 20)
	for i := range b {
		b[i] = charset[p.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	// Generate a random string and print it
	randomString := generateRandomString()
	fmt.Println("Generated Random String:", randomString)
}
