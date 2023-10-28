package tasks

import (
	"log"
	"math"
	"wb_task_l1/point"
)

func Task24() {
	p1 := point.Point{} // Создаем точки с координатами
	p1.NewPoint(5, 9)
	p2 := point.Point{}
	p2.NewPoint(3, 4)

	dist := math.Sqrt(math.Pow(p1.X()-p2.X(), 2) + math.Pow(p1.Y()-p2.Y(), 2)) // Вычисляем расстояние
	log.Printf("Result: %0.10f", dist)                                         // Выводим расстояние
}
