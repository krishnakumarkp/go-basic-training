package main

import (
	"fmt"
	"sync"
	"time"
)

func getUsersFromSite(site string) []string {
	// Simulating 1 second delay to fetch users
	time.Sleep(1 * time.Second)
	return []string{site + "-user1", site + "-user2", site + "-user3"} // Simulated users
}

func main() {
	sites := []string{"site1", "site2", "site3", "site4", "site5" /* Add more sites here */}

	// Create a channel to collect results
	results := make(chan []string)

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	start := time.Now()

	// Launch a goroutine for each site
	for _, site := range sites {
		wg.Add(1)
		go func(site string) {
			defer wg.Done()
			users := getUsersFromSite(site)
			results <- users
		}(site)
	}

	// Start a goroutine to close the results channel after all goroutines finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and print results from the channel
	for users := range results {
		fmt.Println("Users from site:", users)
	}
	elapsed := time.Since(start)
	fmt.Println("Total time:", elapsed)
}
