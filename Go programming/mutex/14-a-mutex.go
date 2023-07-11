package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.RWMutex // defing a mutex 
}

func (sc safeCounter) inc(key string) {
	// mutex in action, it locks the dangerous function and no other function can affect it during its execution
	sc.mu.Lock()
	defer sc.mu.Unlock()
	// this is the dangerous function below 
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	// RWmutex in action, it locks the dangerous function and no other function can affect it during its execution but can read it 
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	// this is the dangerous function below 
	return sc.counts[key]
}

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

type emailTest struct {
	email string
	count int
}

func test(sc safeCounter, emailTests []emailTest) {
	emails := make(map[string]struct{})

	var wg sync.WaitGroup
	for _, emailT := range emailTests {
		emails[emailT.email] = struct{}{}
		for i := 0; i < emailT.count; i++ {
			wg.Add(1)
			go func(emailT emailTest) {
				sc.inc(emailT.email)
				wg.Done()
			}(emailT)
		}
	}
	wg.Wait()

	emailsSorted := make([]string, 0, len(emails))
	for email := range emails {
		emailsSorted = append(emailsSorted, email)
	}
	sort.Strings(emailsSorted)

	for _, email := range emailsSorted {
		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
	}
	fmt.Println("=====================================")
}

func main() {
	sc := safeCounter{
		counts: make(map[string]int),
		mu:     &sync.RWMutex{},
	}
	test(sc, []emailTest{
		{
			email: "john@example.com",
			count: 23,
		},
		{
			email: "john@example.com",
			count: 29,
		},
		{
			email: "jill@example.com",
			count: 31,
		},
		{
			email: "jill@example.com",
			count: 67,
		},
	})
	test(sc, []emailTest{
		{
			email: "kaden@example.com",
			count: 23,
		},
		{
			email: "george@example.com",
			count: 126,
		},
		{
			email: "kaden@example.com",
			count: 31,
		},
		{
			email: "george@example.com",
			count: 453,
		},
	})
}
