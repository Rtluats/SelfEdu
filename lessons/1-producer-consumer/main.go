//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream) <-chan *Tweet {
	c := make(chan *Tweet)
	go func() {
		for {
			tweet, err := stream.Next()
			if err == ErrEOF {
				break
			}

			c <- tweet
		}
		close(c)
	}()
	return c
}

func consumer(tweets <-chan *Tweet) chan struct{} {
	c := make(chan struct{})
	go func() {
		for t := range tweets {
			if t.IsTalkingAboutGo() {
				fmt.Println(t.Username, "\ttweets about golang")
			} else {
				fmt.Println(t.Username, "\tdoes not tweet about golang")
			}
		}
		close(c)
	}()
	return c
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// Producer
	tweets := producer(stream)

	// Consumer
	c := consumer(tweets)
	<-c
	fmt.Printf("Process took %s\n", time.Since(start))
}
