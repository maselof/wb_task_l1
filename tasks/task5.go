package tasks

import (
	"fmt"
	"log"
	"time"
)

func Task5() {
	ch := make(chan int, 1)                  // Создаем бууферизированный канал для того, чтобы там было только одно значение
	var n int                                // Создаем переменную для определения количества секунд
	fmt.Print("Введите количество секунд: ") // Вводим с клавиатуры сколько секунд будет работать программа
	fmt.Scan(&n)                             // Записываем данные
	go readChan(ch, 1)                       // Запускаем горутину для считывания значений из канала
	for i := 0; i < n+1; i++ {
		time.Sleep(time.Second)      // Запускаем time.Sleep() для того, чтобы программа работала указанное время
		ch <- i                      // Закидываем в канал значение
		log.Printf("Seconds: %d", i) // Для определения сколько времени работала программа
	}
	close(ch)
	time.Sleep(time.Second) // Для того, чтобы успелось считать с последней горутины

}

func readChan(ch chan int, ind int) {
	for {
		val, ok := <-ch //Считывыем данные из канала
		if !ok {
			log.Printf("Goroutine %d stopped ...", ind) // Останавливаем горутину
			break
		}
		log.Printf("Goroutine %d val: %d", ind, val) // Выводим данные
	}
}
