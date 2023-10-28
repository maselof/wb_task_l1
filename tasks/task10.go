package tasks

import "log"

func Task10() {
	slice := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // Создаем срез с элементами
	valMap := make(map[int][]float32)                                     // Создаем map для группировки значений
	helperVal := 0
	for i := 0; i < len(slice); i++ { // Проходимся по срезу и добавляем каждый элемент в map
		helperVal = int(slice[i]) / 10 * 10
		valMap[helperVal] = append(valMap[helperVal], slice[i])
	}
	log.Print(valMap) // Выводим значения
}
