package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func getNumber(wg *sync.WaitGroup) {

	defer wg.Done()
	var err error
	var str string
	var input, z, input2, input1 int64

	file, err := os.OpenFile("balance1.txt", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)

	str = string(b)
	i2, err := strconv.ParseInt(str, 10, 64)
	os.Remove("balance1.txt")

	fmt.Println("The current balance is ::", i2)

	fmt.Println("\n Enter the operation that you want to perform \n 1)Deposit \n 2)Withdraw")
	fmt.Scan(&input)
	f, err := os.Create("balance1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	switch input {
	case 1:

		fmt.Print("Enter the amount you want to deposit::")
		fmt.Scanln()
		fmt.Scan(&input1)
		z = int64(input1)
		o := i2 + z
		fmt.Println(o)
		i3 := strconv.FormatInt(int64(o), 10)
		f.WriteString(i3)
	case 2:
		fmt.Print("Enter the amount you want to withdraw::")
		fmt.Scanln()
		fmt.Scan(&input2)
		if i2 > input2 {
			r := i2 - input2
			i4 := strconv.FormatInt(int64(r), 10)

			f.WriteString(i4)
		} else {
			fmt.Println("Insufficent Balance")
		}
	default:
		fmt.Println("wrong choice")
	}
	wg.Done()

}

func main() {
	t := time.Now()
	var wg sync.WaitGroup

	fmt.Println("\nCurrent Time is:", t)
	fmt.Println("\nWelcome to the small banking.")
	wg.Add(2)
	go getNumber(&wg)
	wg.Wait()
	wg.Add(2)
	go getNumber(&wg)
	wg.Wait()
	wg.Add(2)
	go getNumber(&wg)
	wg.Wait()

}
