package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

var str string
var input, z, userinput2, userinput1, l, add, subt int64
var err error

func getNumber(num int) {

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

	f, err := os.Create("balance1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	switch num {
	case 1:

		fmt.Print("\nEnter the amount you want to deposit::")
		fmt.Scanln()
		fmt.Scan(&userinput1)
		z = int64(userinput1)
		add = i2 + z
		fmt.Println(add)
		i3 := strconv.FormatInt(int64(add), 10)
		f.WriteString(i3)
	case 2:
		fmt.Print("\nEnter the amount you want to withdraw::")
		fmt.Scanln()
		fmt.Scan(&userinput2)
		if i2 > userinput2 {
			subt = i2 - userinput2
			i4 := strconv.FormatInt(int64(subt), 10)

			f.WriteString(i4)
		} else {
			fmt.Println("Insufficent Balance")
		}
	default:
		fmt.Println("wrong choice")
	}

}

func main() {
	var start time.Time
	t := time.Now()

	fmt.Println("\nCurrent Time is:", t)
	fmt.Println("\nWelcome to the small banking.")
	fmt.Println("main execution started at time", time.Since(start))
	for i := 0; i < 7; i++ {
		go getNumber(1)
		go getNumber(2)
		go getNumber(4)
		time.Sleep(time.Millisecond * 1000)
	}
	fmt.Println("\nmain execution stopped at time", time.Since(start))

}
