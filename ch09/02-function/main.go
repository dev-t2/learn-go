package main

import "fmt"

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
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)

	carFuel += toGallons(Liters(40.0))
	busFuel += toLiters(Gallons(30.0))

	fmt.Printf("Gallons: %.1f, Liters: %.1f\n", carFuel, busFuel)
}