package main

import "fmt"

func main() {

	var numberOfRifles int = 4
	var sniperRifles [4]string = [4]string{"L42A1 Enfield", "Heckler & Koch PSG1", "AS50", "AWM"}
	var gunShot [4]int = [4]int{2473, 800, 1500, 800}
	var capacity [4]int = [4]int{10, 12, 14, 15}
	for i := 0; i < numberOfRifles; i++ {
		if gunShot[i] >= 1500 {
			if capacity[i] >= 12 {
				fmt.Println(sniperRifles[i], "Purchased")
			} else {
				fmt.Println(sniperRifles[i], "Capacity is not enough")
			}
		} else {
			fmt.Println(sniperRifles[i], "Range is not enough")
		}
	}
	fmt.Println("Shopping completed")
}
