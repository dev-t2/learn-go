package main

import "fmt"

type Title string

type Gallons float64

type Liters float64

type Milliliters float64

func toGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func toLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func LitersToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func MillilitersToGallons(m Milliliters) Gallons {
	return Gallons(m * 0.000264)
}

func GallonsToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func GallonsToMilliliters(g Gallons) Milliliters {
	return Milliliters(g * 3785.41)
}

func main() {
	aTitle := Title("ATitle")
	bTitle := Title("ZTitle")

	fmt.Println(aTitle == aTitle)
	fmt.Println(aTitle < bTitle)
	fmt.Println(aTitle > bTitle)

	fmt.Println()

	fmt.Println(aTitle + "s")

	fmt.Println()

	aLiters := Liters(1.2)
	bLiters := Liters(3.6)
	aGallons := Gallons(1.2)
	bGallons := Gallons(3.4)

	fmt.Println(aLiters + bLiters)
	fmt.Println(bGallons - aGallons)
	fmt.Println(bLiters / aLiters)
	fmt.Println(aGallons == aGallons)
	fmt.Println(aLiters < bLiters)
	fmt.Println(aGallons > bGallons)

	fmt.Println()

	fmt.Println(aLiters + 3.6)
	fmt.Println(bGallons - 1.2)
	fmt.Println(bLiters / 1.2)
	fmt.Println(aGallons == 1.2)
	fmt.Println(aLiters < 3.6)
	fmt.Println(aGallons > 3.4)

	fmt.Println()

	carFuel := Gallons(10.6)
	busFuel := Liters(238.5)

	fmt.Printf("Gallons: %.1f, Liters: %.1f\n", carFuel, busFuel)

	carFuel = Gallons(Liters(40.0))
	busFuel = Liters(Gallons(63.0))

	fmt.Printf("Gallons: %.1f, Liters: %.1f\n", carFuel, busFuel)

	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(Gallons(63.0) * 3.785)

	fmt.Printf("Gallons: %.1f, Liters: %.1f\n", carFuel, busFuel)

	carFuel = toGallons(Liters(40.0))
	busFuel = toLiters(Gallons(63.0))

	fmt.Printf("Gallons: %.1f, Liters: %.1f\n", carFuel, busFuel)
}