package family

import "fmt"

// Member представляет информацию о члене семьи.
type Member struct {
	Gender       bool   // Пол (true - мужской, false - женский)
	Age          int    // Возраст
	Name         string // Имя
	Married      bool   // Семейное положение (true - в браке, false - не в браке)
	NumberOfKids int    // Количество детей
}


func (m Member) DisplayInfo() {
	
	fmt.Printf("Персональные данные:\n\tИмя: %s\n\tВозраст: %d\n", m.Name, m.Age)

	
	if m.Gender {
		// Если мужчина, выводим информацию о семейном положении.
		fmt.Print("Пол: Мужской\n")
		if m.Married {
			fmt.Print("Семейное положение: Женат\n")
		} else {
			fmt.Print("Семейное положение: Не женат\n")
		}
	} else {
		
		fmt.Print("Пол: Женский\n")
		if m.Married {
			fmt.Print("Семейное положение: Замужем\n")
		} else {
			fmt.Print("Семейное положение: Не замужем\n")
		}
	}

	
	if m.NumberOfKids == 0 {
		fmt.Print("Дети: Нет\n")
	} else {
		fmt.Print("Дети: ", m.NumberOfKids, "\n")
	}
}
