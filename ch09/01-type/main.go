package main

import "fmt"

type Title string

type Gallons float64

type Liters float64

func main() {
	fmt.Println(Title("ATitle") == Title("ATitle"))
	fmt.Println(Title("ATitle") < Title("ZTitle"))
	fmt.Println(Title("ATitle") > Title("ZTitle"))
	fmt.Println(Title("ATitle") + "s")

	fmt.Println()

	fmt.Println(Liters(1.2) + Liters(3.4))
	fmt.Println(Gallons(5.5) - Gallons(2.2))
	fmt.Println(Liters(2.2) / Liters(1.1))
	fmt.Println(Gallons(1.2) == Gallons(1.2))
	fmt.Println(Liters(1.2) < Liters(3.4))
	fmt.Println(Gallons(1.2) > Gallons(3.4))

	fmt.Println()
	
	fmt.Println(Liters(1.2) + 3.4)
	fmt.Println(Gallons(5.5) - 2.2)
	fmt.Println(Liters(1.2) == 1.2)
	fmt.Println(Gallons(1.2) < 3.4)

	fmt.Println()

	carFuel := Gallons(10.0)
	busFuel := Liters(240.0)

	fmt.Printf("Gallons: %.1f, Liters: %.1f\n", carFuel, busFuel)

	carFuel = Gallons(Liters(40.0))
	busFuel = Liters(Gallons(63.0))

	fmt.Printf("Gallons: %.1f, Liters: %.1f\n", carFuel, busFuel)

	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(Gallons(63.0) * 3.785)

	fmt.Printf("Gallons: %.1f, Liters: %.1f\n", carFuel, busFuel)
}