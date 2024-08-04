package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	arr := []int{1, 2, 3, 4, 5}
	scanner := bufio.NewScanner(os.Stdin)
	x := len(arr)

	fmt.Println("1)At Begining\n2)At End\n3)At Pos\n4)Exit")
	fmt.Println("Enter Choice : ")

	scanner.Scan()
	sw, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	opt := int(sw)

	switch opt {
	case 1:
		fmt.Println("ARRAY : ", arr)
		// fmt.Println("START")
		fmt.Print("Enter Number : ")
		scanner.Scan()
		num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		arr[0] = int(num)
		fmt.Println("ARRAY : ", arr)

	case 2:
		fmt.Println("ARRAY : ", arr)
		// fmt.Println("END")
		fmt.Print("Enter Number : ")
		scanner.Scan()
		num, _ := strconv.ParseInt(scanner.Text(), 10, 10)
		arr[x-1] = int(num)
		fmt.Println("ARRAY : ", arr)

	case 3:
		fmt.Println("ARRAY : ", arr)
		// fmt.Println("POSITION")
		fmt.Print("Total Positions 0-", len(arr))
		fmt.Printf("\nEnter Position : ")
		scanner.Scan()
		z, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		a := int(z)

		if a > 0 && a <= x && a != x {
			fmt.Println("ALLOWED")
			fmt.Print("Enter Number : ")
			scanner.Scan()
			num, _ := strconv.ParseInt(scanner.Text(), 10, 10)
			arr[a] = int(num)
		} else {
			fmt.Println("------ERROR------")
		}
		fmt.Println("ARRAY : ", arr)

	case 4:
		fmt.Println("EXIT")
		break
	}
}