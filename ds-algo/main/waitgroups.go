package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	time.Sleep(time.Second * 3)
	for i := 0; i < 5; i++ {
		fmt.Println("Sending job", i)
		jobs <- i
	}
	close(jobs)

	<-results

	fmt.Println("Wait groups")

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go workerForWg(i, wg)
	}
	wg.Wait()

}

func workerForWg(i int, group sync.WaitGroup) {

	defer group.Done()
	fmt.Printf("Worker %d starting\n", i)

	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", i)

}
func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Println("Job started", id, j)
		time.Sleep(time.Second)
		fmt.Println("Done Job", id, j)
		results <- j * 2
	}
}
