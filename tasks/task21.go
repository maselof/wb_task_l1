package tasks

import (
	"log"
)

// Создаем структуру Wolf, у которой есть метод Wooo
type Wolf struct{}

func (w *Wolf) Wooo() {
	log.Print("i'm wolf woooooo")
}

// Создаем структуру Sheep, у которой есть метод Beee
type Sheep struct{}

func (s *Sheep) Beee() {
	log.Print("i'm sheep beeeeee")
}

// Создаем интерфейс классов адаптеров
type Adapter interface {
	Operation()
}

// Создаем структуру WolfAdapter, которая указывает на структуру Wolf и
// указываем два метода, в которой адаптер создается и вызывается метод Wooooo, также все у Sheep
type WolfAdapter struct {
	*Wolf
}

func NewWolfAdapter(w *Wolf) Adapter {
	return &WolfAdapter{w}
}

func (wa *WolfAdapter) Operation() {
	wa.Wooo()
}

type SheepAdapter struct {
	*Sheep
}

func NewSheepAdapter(s *Sheep) Adapter {
	return &SheepAdapter{s}
}

func (sa *SheepAdapter) Operation() {
	sa.Beee()
}

// Вывод результатов
func Task21() {
	adapter := [2]Adapter{NewWolfAdapter(&Wolf{}), NewSheepAdapter(&Sheep{})}
	for i := range adapter {
		adapter[i].Operation()
	}
}
