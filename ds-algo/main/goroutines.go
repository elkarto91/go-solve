package main

import (
	"fmt"
	"time"
)

func main() {

	mainChannel := make(chan string)
	go syncReader(mainChannel)
	mainChannel <- "Initialized Channel"
	time.Sleep(time.Second)

	bufferedMainChannel := make(chan string, 4)
	go bufferedSyncReader(bufferedMainChannel)
	time.Sleep(time.Second)

	bufferedMainChannel <- "Initialized Buffered Channel"
	bufferedMainChannel <- "Second Item in Buffered Channel"

	sum(50)
	fmt.Println(" Go routine tests")
	numArray := []int{234, 34, 3, 434}
	fmt.Println(numArray)

	for _, v := range numArray {
		go sum(v)
	}

	bufferedMainChannel <- "Third Item in Buffered Channel"
	time.Sleep(time.Second)

	waiterChannel := make(chan string)
	go workerFunction(waiterChannel)

	fmt.Println(" Got response, can close now", <-waiterChannel)

	chandlerChannel := make(chan string, 1)
	monicaChannel := make(chan string, 1)

	sendOnlyChannel(chandlerChannel, "chandler......")
	receiveOnlyChanel(chandlerChannel, monicaChannel)

	fmt.Println(<-monicaChannel)

	fmt.Println("select channels")

	chanOne := make(chan string, 1)
	chanTwo := make(chan string, 2)

	go func() {
		time.Sleep(time.Second)
		chanOne <- "oneeeeee"
	}()

	go func() {
		time.Sleep(time.Second)

		chanTwo <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {

		case <-chanTwo:
			fmt.Println(" Chan two came")
		case <-chanOne:
			fmt.Println("chan one came")

		}
	}

	timeOutChannel := make(chan int, 1)

	go func() {
		time.Sleep(3 * time.Second)
		timeOutChannel <- 100
	}()

	select {
	case <-timeOutChannel:
		fmt.Println("Received")
	case <-time.After(4 * time.Second):
		fmt.Println("Timed out")
	}

	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}

func workerFunction(waiterChannel chan string) {
	fmt.Println("Sleeping a while")
	time.Sleep(time.Second)
	fmt.Println("Done sleeping")
	waiterChannel <- "Done Here"

}
func sum(num int) {
	sum := num + 9999
	fmt.Println("Sum of your number and 9999 is ", sum)
	return
}
func bufferedSyncReader(cc chan string) {
	fmt.Println(" buffered reader", <-cc)
	fmt.Println(" buffered reader", <-cc)
	fmt.Println(" buffered reader", <-cc)
}
func syncReader(cc chan string) {
	now := <-cc
	fmt.Println("Now channel says", now)
}
func sendOnlyChannel(hitMe chan<- string, message string) {
	hitMe <- message
}

func receiveOnlyChanel(hitMe <-chan string, pongs chan<- string) {

	msg := <-hitMe
	pongs <- msg
}
