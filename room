package spaces

import (
 "fmt"
 "goProjects/home/stuff"
 "math/rand"
)

type LivingSpace struct {
 Name   string
 Length float32
 Width  float32
 Height float32
 Square float32
}

func (ls *LivingSpace) computeArea() {
 ls.Square = ls.Length * ls.Width
}

func GenerateSpace(name string) LivingSpace {
 space := LivingSpace{
  Name:   name,
  Length: rand.Float32()*(10-2) + 2,
  Width:  rand.Float32()*(7-2) + 2,
  Height: 2.8,
 }
 space.computeArea()

 fmt.Println(space)
 return space
}

type CookingArea struct {
 Space     LivingSpace
 Inventory stuff.KitchenTools
}

type SleepingChamber struct {
 Space     LivingSpace
 Inventory stuff.BedroomItems
}

type CommonArea struct {
 Space     LivingSpace
 Inventory stuff.HallFurnishings
}

type SanitaryRoom struct {
 Space     LivingSpace
 Inventory stuff.BathroomAccessories
}
