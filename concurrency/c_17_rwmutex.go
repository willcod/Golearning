package main

import (
	"fmt"
	"sync"
	"time"
)

func introduction() {
	fmt.Println("This code shows how to use RWmutex")
}

var Password = secret{password: "123456"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

func change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("Change-lock")
	time.Sleep(2 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}

func show(c *secret) string {
	c.RWM.RLock()
	fmt.Print("show")
	time.Sleep(time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

func main() {
	introduction()

	var showFunc = func(c *secret) string { return "" }
	showFunc = show
	var wg sync.WaitGroup

	fmt.Println("Password:", showFunc(&Password))
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Password:", showFunc(&Password))
		}()
	}

	go func() {
		wg.Add(1)
		defer wg.Done()
		change(&Password, "abcdef")
	}()

	wg.Wait()
	fmt.Println("Password: ", showFunc(&Password))
}
