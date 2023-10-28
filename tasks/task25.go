package tasks

import (
	"log"
	"time"
)

func Task25() {
	log.Println("Start sleeping...")
	doSleep(500) // Запускаем функцию
	log.Println("Wake up!")
}

func doSleep(s int) {
	<-time.After(time.Duration(s) * time.Millisecond) // Вспомнил про функцию after, которую использовали ранее
}
