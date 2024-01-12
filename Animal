package pets

import (
	"fmt"
)

type Animal struct {
	Species string
	Kind    string
	Age     int
	Name    string
}

func (a Animal) DisplayInfo() {
	fmt.Println("\nЖивотное: ", a.Species, "\nВид: ", a.Kind, "\nИмя: ", a.Name, "\nВозраст: ", a.Age, "\n")
}

type PetInfoPrinter struct {
	AnimalInfoPrinter
}

func (p PetInfoPrinter) PrintPetInfo(animal Animal) {
	p.PrintInfo(animal)
	fmt.Println("Это домашнее животное.")
}

type WildAnimalInfoPrinter struct {
	AnimalInfoPrinter
}

func (w WildAnimalInfoPrinter) PrintWildAnimalInfo(animal Animal) {
	w.PrintInfo(animal)
	fmt.Println("Это дикий зверь.")
}

type AnimalInfoPrinter struct{}

func (a AnimalInfoPrinter) PrintInfo(animal Animal) {
	fmt.Println("\nИнформация о животном:\nВид: ", animal.Species, "\nИмя: ", animal.Name, "\nВозраст: ", animal.Age, "\n")
}
