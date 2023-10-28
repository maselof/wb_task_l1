package tasks

import (
	"log"
	"math/big"
)

func Task22() {
	s1 := "567898567"
	s2 := "985678987"

	// Cоздаем числа из строк
	a, flag := new(big.Float).SetString(s1)
	if !flag {
		log.Printf("Не может взять число из строки\n")
		return
	}
	b, flag := new(big.Float).SetString(s2)
	if !flag {
		log.Printf("Не может взять число из строки\n")
		return
	}

	val := new(big.Float)
	val.Add(a, b)
	log.Println("Сумма: ", val)

	val.Sub(a, b)
	log.Println("Разница: ", val)

	val.Quo(a, b)
	log.Println("Деление: ", val)

	val.Mul(a, b)
	log.Println("Умножение: ", val)
}
