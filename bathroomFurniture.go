package zones

import (
 "fmt"
 "goProjects/home/gear"
 "math/rand"
)

type LivingZone struct {
 Name    string
 Length  float32
 Width   float32
 Height  float32
 Surface float32
}

func (lz *LivingZone) computeSurface() {
 lz.Surface = lz.Length * lz.Width
}

func DesignZone(name string) LivingZone {
 zone := LivingZone{
  Name:   name,
  Length: rand.Float32()*(12-3) + 3,
  Width:  rand.Float32()*(8-2) + 2,
  Height: 2.9,
 }
 zone.computeSurface()

 fmt.Println(zone)
 return zone
}

type CulinaryZone struct {
 Zone    LivingZone
 Utensils gear.KitchenGear
}

type RestingPlace struct {
 Zone    LivingZone
 Furnishings gear.LoungeFurniture
}

type PassageArea struct {
 Zone    LivingZone
 Decor gear.HallDecor
}

type LavatoryZone struct {
 Zone    LivingZone
 Accessories gear.BathroomAccessories
}
