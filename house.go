
package main

import (
 "fmt"
 "goProjects/home/animals"
 "goProjects/home/family"
 "goProjects/home/furniture"
 "goProjects/home/rooms"
)

type DreamHouse struct {
 TotalArea float32
 Kitchen   kitchenFurniture.go
 MasterBed bedroomFurniture.go
 LivingRoom hallFurniture.go
 Bath      bathroomFurniture
 Fam       family
 Pets      Animals
}

func main() {
 dreamHome := DreamHouse{
  Kitchen:    kitchenFurniture.go{Furniture: furniture.SetupKitchen(), Area: rooms.EstablishRoom("Kitchen")},
  MasterBed:  bedroomFurniture{Area: rooms.EstablishRoom("Master Bedroom"), Furniture: furniture.SetupMasterBedroom()},
  LivingRoom: hallFurniture.go{Area: rooms.EstablishRoom("Living Room"), Furniture: furniture.SetupLivingRoom()},
  Bath:       bathroomFurniture{Area: rooms.EstablishRoom("Bathroom"), Furniture: furniture.SetupBathroom()},
  Fam:        family.InitFamily(),
  Pets:       animals.InitAnimals(),
 }

 fmt.Println("\n\n\n\n", dreamHome)
}
