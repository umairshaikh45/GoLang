# Example of the race condition in the go lang.

The structure of the project that contains the various Golang files are described below :

• balance1.txt

• Demo.go

• Main.go

• Wait.go

• Mutex.go

# Balance1.txt

The text file is used to store the balance of the user’s account. The text file can be accessed by the Golang program Demo.go to perform the withdrawal or deposit of the user’s money into the account. The file is updated with the money deposited to or withdrawn from the account in accordance to the user’s deposit and withdrawal.

# Demo.go

This file is predominantly written in Golang. The Demo.go demonstrates to the user as to how the project works for a single user deposits or withdraws to and from the account’s balance at a given point of time. It gives a better understanding of how the project works without a data race condition, for a given balance at any given point of time in the program execution.

# Main.go

The file shows how the program behaves when multiple goroutines (the light weight threads) tries to access the same set of shared memory resources and make changes to the them at the same time, which potentially causes a data race problem as two or more goroutines try to access the same memory location of the compiler concurrently, and of the operations is a write operation to the shared resource. The Main.go reveals, 18 goroutines to access the shared memory resource at a given point of time, trying to access the critical section of the program at the same time in order to make changes to the text file, without any synchronization. The goroutines aren’t at par with each other in terms of synchronization, which leads to a data race condition in the program, when the user tries to access the go routines. This enables us to see prevalently that a data race condition exists in the program, without using the data race detector during debugging the program for stack traces. The data race condition can be double confirmed by employing the inbuilt data race detector while running the program. The command to detect the data races using the inbuilt data race condition is as below : $ go run -race Main.go In order to see if the data race condition exists in the whole package, we can employ the following command. $ go test -race mypackage When the data race detector detects data races in the program, it prints a report that contains the stack traces for conflicting accesses, as well as the stacks where the involved goroutines were created in the first place. The brief description of the other Golang programs are given below.

# Wait.go

The Wait.go implements one of the prevention mechanisms to prevent the data race condition in Golang. This mechanism enables the threads to wait for resources, in order to make write operations in order to avoid conflicting writes to the resource at the critical section of the program. When the wait groups are utilized in the program, the goroutines are required to wait for the other routines to finish, and unlock the thread before it can make a write operation to the resource.

# Mutex.go

Mutexes are data race prevention mechanism in Golang that enables only one goroutine to enter t=into the critical section at one point of time. Mutexes can be utilized to provide locking mechanisms to goroutines, which will enable them to lock and unlock the critical section of the code, which ensures that not more than one goroutine can access the code at the same time.The critical section of the code is locked using mu.Lock() and mu.Unlock() to restrict access and enables only one goroutine to enter into the critical section of the code.

# Experimental Structure

The file Demo.go demonstrates the basic operations of deposit and withdrawal to the user, without the potential risk of a data race for a single user. The Main.go makes changes in the initial program of Demo.go in order to portray how the program works without implementing the data race prevention mechanisms and how the data race works without employing the Golang’s inbuilt data race detector.

# Instructions to run the program.

Prerequisites

The go libraries are downloaded from the below link, which are essential to execute the programs of our project : https://golang.org/dl/ Platform to run Golang programs – Visual Studio Code.

• Set the path after you install go libraries in another drive apart from C drive. The path is as follows

• Path = ”D:\Program Files (x86)\Go\bin”;

• To run the program, use the following command “go run filename.go”.

• To run the Demo.go or any other Golang program file, the command is “go run Demo.go”.

• To run the file using inbuild race detector, the command is “go run -race Main.go”