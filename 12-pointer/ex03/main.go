package main

import "fmt"

type Data struct {
	data  int
	datas [100]int
}

func UpdateData(d Data) {
	d.data = 99
	d.datas[99] = 99
}

func UpdatePointData(d *Data) {
	d.data = 99
	d.datas[99] = 99
}

func main() {
	var d Data

	UpdateData(d)

	fmt.Printf("data = %d\n", d.data)
	fmt.Printf("datas[99] = %d\n", d.datas[99])

	UpdatePointData(&d)
	
	fmt.Printf("data = %d\n", d.data)
	fmt.Printf("datas[99] = %d\n", d.datas[99])
}