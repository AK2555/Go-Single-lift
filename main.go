package main

import "fmt"

func main(){
	fmt.Println("Ayush")
	ls1 := NewLiftSystem(30,3,2)
	fmt.Println(ls1.RequestPickup(0,20))
	fmt.Println(ls1.RequestPickup(10,15))
	fmt.Println(ls1.RequestPickup(12,18))

	ls2 := NewLiftSystem(15,2,5)
	fmt.Println(ls2.RequestPickup(0,12))
	fmt.Println(ls2.RequestPickup(1,11))
	fmt.Println(ls2.RequestPickup(2,10))
	fmt.Println(ls2.RequestPickup(14,3))

}