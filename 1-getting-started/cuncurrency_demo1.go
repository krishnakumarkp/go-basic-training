package main

import (
	"fmt"
	"time"
)

func getUsersFromSite(site string) []string {
	// Simulating 1 second delay to fetch users
	time.Sleep(1 * time.Second)
	return []string{site + "-user1", site + "-user2", site + "-user3"} // Simulated users
}

func main() {
	sites := []string{"site1", "site2", "site3", "site4", "site5"}

	start := time.Now()

	// Query users for each site sequentially
	for _, site := range sites {
		users := getUsersFromSite(site)
		fmt.Println("Users from site:", users)
	}

	elapsed := time.Since(start)
	fmt.Println("Total time:", elapsed)
}
