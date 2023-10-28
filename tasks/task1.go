package tasks

import "fmt"

//ПЕРВЫЙ СПОСОБ(ВИД) КОМПОЗИЦИИ!!!

/* // Создаем структуру Human
type Human struct {
	Age      int    // Возраст человека
	Name     string // Имя человека
	LastName string // Фамилия человека
}

// Для структуры Human пишем метод NameHuman, который выводит имя человека
func (h Human) NameHuman() string {
	return h.Name
}

// Создаем структуру Action со встроенным типом Human
type Action struct {
	Eat   bool // Переменная для определения ест ли он сейчас или нет
	Human      // Внедряемый тип Human
}

// Для структуры Action пишем метод EatAction для вывода ответа на вопрос "Кушает ли сейчас человек?"
func (a Action) EatAction() string {
	if a.Eat {
		return "is eating now" // Вывод, если ест
	} else {
		return "is not eating right now" // Вывод, если не ест
	}
} */

//ВТОРОЙ СПОСОБ(ВИД) КОМПОЗИЦИИ!!!

// Создаем структуру Human
type Human struct {
	Age      int    // Возраст человека
	Name     string // Имя человека
	LastName string // Фамилия человека
}

// Для структуры Human пишем метод NameHuman, который выводит имя человека
func (h Human) NameHuman() string {
	return h.Name
}

// Создаем структуру Action со встроенным типом Human
type Action struct {
	Eat   bool  // Переменная для определения ест ли он сейчас или нет
	Human Human // Поле является структурой типа Human
}

// Метод структуры Action, встраиваемая имплементацию на основе метода Human.NameHuman()
func (a Action) NameHuman() string {
	return a.Human.NameHuman()
}

// Для структуры Action пишем метод EatAction для вывода ответа на вопрос "Кушает ли сейчас человек?"
func (a Action) EatAction() string {
	if a.Eat {
		return "is eating now" // Вывод, если ест
	} else {
		return "is not eating right now" // Вывод, если не ест
	}
}

func Task1() {
	action := Action{true, Human{20, "Steve", "Jefferson"}} // Создаем структуру Action, вписываем значения для полей
	fmt.Println(action.NameHuman(), action.EatAction())     // Используем метод Human, не обращаясь к этому полю в Action
}

/*При выполнении данного задания, можно прийти к выводу, что первый способ удобнее, чем второй.
Так как используя первый метод композиции, мы сокращаем количество строчек кода.*/
