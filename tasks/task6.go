package tasks

import (
	"log"
	"sync"
)

// Остановка горутины через закрывание канала
func Task6_1() {
	ch := make(chan int, 1) // Создаем буфферизированный канал
	var wg sync.WaitGroup   // Создаем переменную для отслеживания работы горутины
	wg.Add(1)               // Добавляем одну горутину
	go func() {             // Создаем анонимную функцию для запуска горутины
		worker6_1(ch, 1)
		wg.Done() // Декриментируем количество горутин
	}()
	ch <- 1   // Отправляем значение в канал
	close(ch) // Закрываем канал
	wg.Wait() // Ожидаем окончания горутины
}

func worker6_1(ch chan int, ind int) {
	for {
		val, ok := <-ch // Считываем данные
		if !ok {
			log.Printf("Goroutines %d stopped ...", ind) // Вывод в логи информацию о остановке горутины
			break                                        // Завершаем бесконечный цикл
		}
		log.Printf("Goroutines %d val: %d", ind, val) // Выводим значение, которое получили в канал
	}
}

/*
// Ожидание окончания горутины с помощью time.Sleep()
func Task6_2() {
	go worker6_2(1)
	time.Sleep(time.Second * 2)
}

func worker6_2(ind int) {
	log.Printf("Goroutines %d started", ind)
	time.Sleep(time.Second)
	log.Printf("Goroutines %d stopped ...", ind)
}*/

/*
// Остановка горутины с помощью оператора select
func Task6_3() {
	ch := make(chan bool, 1)
	go worker6_3(1, ch)
	time.Sleep(time.Millisecond * 50)
	ch <- true
	time.Sleep(time.Second)
}

func worker6_3(ind int, ch chan bool) {
	for {
		select {
		case <-ch:
			log.Printf("Goroutines %d stopped ...", ind)
			return
		default:
			log.Printf("Goroutines %d start", ind)
			time.Sleep(time.Second)
		}
	}
}*/

/*
// Остановка горутины с помощью time.After()
func Task6_4() {
	ch := make(chan bool) // Создаем канал
	go worker6_4(ch)      // Запуска горутину
	ch <- false           // Посылаем данные в канал
	time.Sleep(time.Second * 2)
}

func worker6_4(ch chan bool) {
	for {
		timeout := time.After(time.Second) // Создаем переменную timeout для остановки горутины по истечению времени
		select {
		case <-ch:
			log.Printf("Some work ...")

		case <-timeout:
			log.Printf("Goroutines stopped ...")
			return
		}
	}
}*/

/*
// Остановка горутины с помощью time.Tick()
func Task6_5() {
	ch := make(chan bool)       // Создаем канал
	go worker6_5(ch)            // Запускаем горутину
	ch <- false                 // Посылаем значение в канал
	time.Sleep(time.Second * 2) // Чтобы горутина успела завершить работу
}

func worker6_5(ch chan bool) {
	heartbeat := time.Tick(time.Second)
	for {
		select {
		case <-ch:
			log.Printf("Goroutines start and do something")
		case <-heartbeat:
			log.Printf("Goroutines stopped ...")
		}
	}
}*/
