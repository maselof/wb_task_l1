package tasks

import (
	"log"
	"sync"
)

// Создаем структуру с двумя полями, поле mu нужно для отслеживания блокировки Mutex'a
type SomeStruct struct {
	mu    *sync.Mutex
	count int
}

// Прибавляем +1 в поле count
func (ss *SomeStruct) addOne() {
	ss.mu.Lock()
	ss.count++
	ss.mu.Unlock()
}

// Выводим результат поля count
func (ss *SomeStruct) Count() int {
	return ss.count
}

func Task18() {
	var mu sync.Mutex                   // Создаем Mutex для блокировки горутин, когда идет инкрементация
	var wg sync.WaitGroup               // Создаем WaitGroup для ожидания остановки горутин
	ss := SomeStruct{mu: &mu, count: 0} // Создаем структуру
	for i := 0; i < 5; i++ {            // Запускаем 5 горутин
		wg.Add(1)
		go func() {
			ss.addOne()
			wg.Done()
		}()
	}
	wg.Wait()
	log.Print(ss.Count()) // Выводим результат
}
