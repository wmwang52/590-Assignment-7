// Team Members: Dylan Nicks, Coralee Rogers-Vickers, William Wang

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numChairs    = 2  
	numCustomers = 10 
	barberTime   = 4  
)

var (
	mutex       sync.Mutex  
	waitingRoom = make(chan int, numChairs) 
	wg          sync.WaitGroup
)

func barber() {
	for {
		mutex.Lock()
		if len(waitingRoom) == 0 {
			fmt.Println("Barber is sleeping as there are no customers.\n")
			mutex.Unlock()
			time.Sleep(2 * time.Second) 
		} else {
			customerID := <-waitingRoom
			mutex.Unlock()
			fmt.Printf("Barber is cutting hair of customer %d.\n\n", customerID)
			time.Sleep(time.Duration(barberTime) * time.Second) 
			fmt.Printf("Barber finished cutting hair of customer %d.\n\n", customerID)
			wg.Done()
		}
	}
}

func customer(customerID int) {
	mutex.Lock()
	if len(waitingRoom) == numChairs {
		fmt.Printf("Customer %d leaves as the waiting room is full.\n\n", customerID)
		mutex.Unlock()
		wg.Done()
		return
	}

	fmt.Printf("Customer %d is waiting in the waiting room.\n\n", customerID)
	waitingRoom <- customerID
	mutex.Unlock()
}

func main() {
	go barber()

	for i := 1; i <= numCustomers; i++ {
		wg.Add(1) 
		go customer(i)
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second) 
	}

	wg.Wait()
	fmt.Println("All customers have been served or left.\n")
}