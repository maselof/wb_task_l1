package tasks

import (
	"log"
	"sync"
)

func Task7() {
	const ( // Заводим константы для количества горутин и количества значений в канале
		countGoroutines   = 3
		countValInChannel = 2
	)
	valMap := make(map[int]int)             // Создаем отображение для записи в него данных
	ch := make(chan int, countValInChannel) // Создаем буфферизированный канал
	var wg sync.WaitGroup                   // Используем WaitGroup для ожидания закрытия всех горутин и Mutex для устранения гонок
	var mx sync.Mutex

	for i := 0; i < countGoroutines; i++ { // Запускаем горутины
		wg.Add(1)
		go func(ind int, mx *sync.Mutex) {
			writeDataToMap(ch, ind, valMap, mx)
			wg.Done()
		}(i, &mx)
	}

	for i := 0; i < 20; i++ { // Посылаем значения в канал
		ch <- i
	}
	close(ch)           // Закрываем канал
	wg.Wait()           // Дожидаемся окончания горутин
	log.Println(valMap) // Выводим  данные

}

func writeDataToMap(ch chan int, ind int, valMap map[int]int, mx *sync.Mutex) {
	for {
		val, ok := <-ch // Получаем данные из канала
		if !ok {
			log.Printf("Goroutine %d stopped ...", ind) // Выводим данные о остановке горутины
			break
		} else {
			mx.Lock() // Блокируем Mutex
			valMap[val] = val
			log.Printf("Goroutine %d write %d in map", ind, val)
			mx.Unlock() // Разблокируем Mutex
		}

	}
}
