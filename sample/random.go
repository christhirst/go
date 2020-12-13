package sample

import (
	"math/rand"
)

func randomStringFormat(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomUser() string { return randomStringFormat("Jim", "John") }

func randomId() uint32 { return 2 }

func randomBool() bool { return rand.Intn(2) == 1 }
