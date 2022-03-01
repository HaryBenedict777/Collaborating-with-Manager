package main

import "fmt"

func main() {
	c := cron.New()
	c.AddFunc("0 3 * * * *", func() { fmt.Println("Feed spot") })
	c.Start()
	for {
	}

}
