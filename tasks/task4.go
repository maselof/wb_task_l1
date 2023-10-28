package tasks

import (
	"fmt"
	"log"
	"math/rand"
)

func Task4() {
	var n int // Создаем переменную для количества воркеров
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&n)

	ch := make(chan int) // Создаем канал, по которому мы будем передавать значения

	for i := 0; i < n; i++ { // Запускаем заданное количество горутин
		go worker(ch, i)
	}

	val := 0

	for {
		val = rand.Int()
		ch <- val // Посылаем рандомные значение по каналу
	}
}

func worker(ch chan int, ind int) {
	for {
		s := <-ch
		log.Printf("worker %d write value: %d", ind, s)
	}
}
