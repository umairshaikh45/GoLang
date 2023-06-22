package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

var input, z, n, b, m int64
var str string



func getNumber() {
	
	func deposit (d chan int)
	{
		if deposit !=0 {
			i2 += int64(deposit)
		}
		d <- i2
	}
	func withdraw(w chan int)
	{
		
		for withdraw !=0 {
			i2 -= int64(withdraw)
		}
		w <- i2
	}

	var err error

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
	switch input; {
	case 1:
		d := make(chan int)
		//w:=make chan(int)
		deposit(d)
		//go withdraw(w)
		x := <-d
		fmt.Println(x)
		//fmt.Println(y)
		i3 := strconv.FormatInt(int64(x), 10)
		f.WriteString(i3)
	case 2:
		w := make(chan int)
		withdraw(w)
		y := <-w
		fmt.Println(y)
		i4 := strconv.FormatInt(int64(y), 10)
		f.WriteString(i4)

	}

	/*switch input {
	case 1:

		fmt.Print("Enter the amount you want to deposit::")
		fmt.Scanln()
		fmt.Scan(&m)
		z = int64(m)
		o := i2 + z
		fmt.Println(o)
		i3 := strconv.FormatInt(int64(o), 10)
		f.WriteString(i3)
	case 2:
		fmt.Print("Enter the amount you want to withdraw::")
		fmt.Scanln()
		fmt.Scan(&n)
		if i2 > n {
			r := i2 - n
			i4 := strconv.FormatInt(int64(r), 10)

			//e := ioutil.WriteFile("balance.txt", []byte(i4), 0666)
			f.WriteString(i4)
		} else {
			fmt.Println("Insufficent Balance")
		}
	default:
		fmt.Println("wrong choice")
	}*/

}
func main() {
	t := time.Now()

	fmt.Println("\nCurrent Time is:", t)
	fmt.Println("\nWelcome to the small banking.")
	go getNumber()
	go getNumber()

}
