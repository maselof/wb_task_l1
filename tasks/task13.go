package tasks

import "log"

func Task13() {
	a := 1
	b := 10

	log.Printf("Beginner value a = %d b = %d", a, b)
	a, b = b, a

	log.Printf("a = %d b = %d", a, b)
}
