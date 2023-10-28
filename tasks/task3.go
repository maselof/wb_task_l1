package tasks

import (
	"log"
	"sync"
)

/*
	Первый способ с использованием функции time.Sleep() для того, чтобы горутины успели выполнить

действия в функции
*/
/*func Task3() {
	slice := [5]int{2, 4, 6, 8, 10}        // Создаем срез из условия
	ch := make(chan int, 2)                // Создаем буфферизированный канал
	res := 0                               // Создаем переменнюу для записи результата
	countGoroutines := 3                   // Указываем количество горутин
	for i := 0; i < countGoroutines; i++ { // По очереди запуска горутины
		go squareTask3(i, ch, &res)
	}

	for i := 0; i < len(slice); i++ {
		ch <- slice[i] // Значения из среза закидываем в канал
	}

	close(ch)                     // Закрываем канал
	time.Sleep(time.Second)       // Даем время для прекращения работы всех горутин
	log.Printf("Result: %d", res) // Вывод результата
} */

func Task3() {
	slice := [5]int{2, 4, 6, 8, 10} // Второй способ включает в себя использование WaitGroup из библиотеки sync
	ch := make(chan int, 2)         // Создаем буфферизированный канал
	res := 0                        // Создаем переменнюу для записи результата
	countGoroutines := 3            // Указываем количество горутин
	var wg sync.WaitGroup           // Создаем переменную для отслеживания работы горутин
	for i := 0; i < countGoroutines; i++ {
		wg.Add(1)          // Добавляем количество ожидаемых горутин
		go func(ind int) { // Создаем анонимную функцию
			squareTask3(ind, ch, &res)
			wg.Done() // Показываем, что горутина выполнила свою работу
		}(i)
	}

	for _, val := range slice {
		ch <- val // Значения из среза закидываем в канал
	}

	close(ch)                     // Закрываем канал
	wg.Wait()                     // Ждем, когда все горутины прекратят работу
	log.Printf("Result: %d", res) //Выводим результат
}

func squareTask3(ind int, ch chan int, res *int) {
	for {
		val, ok := <-ch // Записываем значение и смотрим состояние канала
		if !ok {
			log.Printf("Goroutines %d stopped...", ind) // Если канал закрыт, то мы это выводим и прекращаем работу цикла
			break
		} else {
			*res += val * val                                  //Записываем в res квадрат значения
			log.Printf("Goroutines %d has res: %d", ind, *res) // Отслеживанием какая горутина и какой при нем был res
		}
	}
}
