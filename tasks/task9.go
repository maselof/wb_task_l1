package tasks

import (
	"fmt"
	"log"
	"time"
)

func Task9() {
	chIn := make(chan int) // Создаем каналы
	chOut := make(chan int)
	var val int // Создаем переменную, которую будет возводить в квадрат
	log.Print("Введите число: ")
	fmt.Scan(&val)
	go valInput(chIn, chOut) // Запускаем горутины
	go squareVal(chOut)
	chIn <- val             // Закидываем число в канал
	time.Sleep(time.Second) // Ожидаем завершения горутин
}

func valInput(in, out chan int) {
	val := <-in // Считываем и посылаем в другую горутину
	out <- val
}

func squareVal(in chan int) {
	val := <-in                              // Считываем и возводим в квадрат
	log.Printf("Квадрат числа: %d", val*val) // Выводим результат
}
