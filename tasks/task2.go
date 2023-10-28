package tasks

import (
	"log"
	"sync"
)

/* Первый способ с использованием функции time.Sleep() для того, чтобы горутины успели выполнить
действия в функции*/

/*func Task2() {
	slice := [5]int{2, 4, 6, 8, 10} // Создаем срез из условия
	ch := make(chan int, 2) // Создаем буфферизированный канал
	countGoroutines := 3 // Указываем количество горутин
	for i := 0; i < countGoroutines; i++ {
		go squareTask2(i+1, ch) // Поочередно запускаем каждую горутину
	}

	for _, val := range slice {
		ch <- val // Значения из среза закидываем в канал
	}

	close(ch) // Закрываем канал
	time.Sleep(time.Second) // Ожидаем, когда все горутины прекратят работу
}
*/

// Второй способ включает в себя использование WaitGroup из библиотеки sync
func Task2() {
	slice := [5]int{2, 4, 6, 8, 10} // Создаем срез из условия
	ch := make(chan int, 2)         // Создаем буфферизированный канал
	var wg sync.WaitGroup           // Создаем переменную для отслеживания работы горутин
	countGoroutines := 3            // Указываем количество горутин
	for i := 0; i < countGoroutines; i++ {
		wg.Add(1)        // Добавляем количество ожидаемых горутин
		go func(i int) { // Создаем анонимную функцию
			squareTask2(i+1, ch)
			wg.Done() // Показываем, что горутина выполнила свою работу
		}(i)
	}

	for _, val := range slice {
		ch <- val // Значения из среза закидываем в канал
	}

	close(ch) // Закрываем канал
	wg.Wait() // Ждем, когда все горутины прекратят работу
}

func squareTask2(ind int, ch chan int) {
	for {
		val, ok := <-ch // Записываем значение и смотрим состояние канала
		if !ok {
			log.Printf("Goroutines %d stopped ...", ind) // Если канал закрыт, то мы это выводим и прекращаем работу цикла
			break
		}
		log.Printf("Goroutines %d value: %d", ind, val*val) // Выводи значение
	}
}
