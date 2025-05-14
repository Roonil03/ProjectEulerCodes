package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	res, id, a := 0, 1, 0
	for j := 0; j < 12; j++ {
		if j == 0 || j == 2 || j == 4 || j == 6 || j == 7 || j == 9 || j == 11 {
			id = (id + 3) % 7
		} else if j == 3 || j == 5 || j == 8 || j == 10 {
			id = (id + 2) % 7
		} else if (1900%400 == 0) || (1900%100 != 0 && 1900%4 == 0) {
			id = (id + 1) % 7
		} else {
			id = (id + 0) % 7
		}
		if id == 0 {
			a++
		}
	}
	fmt.Println("1900 had", a, "Sundays and 1 Jan begins with", days[id])
	for i := 1901; i <= 2000; i++ {
		for j := 0; j < 12; j++ {
			if j == 0 || j == 2 || j == 4 || j == 6 || j == 7 || j == 9 || j == 11 {
				id = (id + 3) % 7
			} else if j == 3 || j == 5 || j == 8 || j == 10 {
				id = (id + 2) % 7
			} else if (i%400 == 0) || (i%100 != 0 && i%4 == 0) {
				id = (id + 1) % 7
			} else {
				id = (id + 0) % 7
			}
			if id == 0 {
				res++
			}
		}
	}
	fmt.Print("The answer is: ", res)
}
