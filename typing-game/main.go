package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var words = []string{
	"apple",
	"banana",
	"cherry",
	"rice",
	"noodle",
	"bread",
}

func run(ctx context.Context, wg *sync.WaitGroup) {
	total_count := 0
	correct_count := 0

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time is up!")
			fmt.Printf("Total: %d, Miss Rate: %.2f%%\n", total_count, 100.0-100.0*float64(correct_count)/float64(total_count))
			wg.Done()
			return
		default:
			idx := rand.Intn(len(words))
			answer := words[idx]
			fmt.Println("Problem:", answer)
			var input string
			fmt.Print("Input  : ")
			fmt.Scanln(&input)
			if input == answer {
				fmt.Println("Correct!")
				correct_count++
			} else {
				fmt.Println("Wrong!")
			}
			total_count++
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	go run(ctx, &wg)
	wg.Wait()
}
