package main

import (
	"fmt"
	"time"
	"strconv"
)

func serveBeer(ch chan string){
	for i:=1; i < 16 ; i++ {
		time.Sleep(1e9)
		beerNumber := "beer number " + strconv.Itoa(i)
		fmt.Println("Giving beer to grant...", beerNumber)
		ch <- beerNumber
	}
}

func drinkBeer(ch chan string, done chan string){
	for i := 1; i < 16 ; i++{
		time.Sleep(1e9)
		beerNumber := <-ch
		fmt.Println("grant is now drinking a beer...",beerNumber)
	}
	done <- "grant is very drunk now"
}



func main(){
	ch := make(chan string)
	done := make(chan string)
	//producer
	go serveBeer(ch)
	//consumber
	go drinkBeer(ch,done)
	waitForFinish  := <-done
	fmt.Println("Grant is done now....", waitForFinish)
	
	//time.Sleep(10 * 1e9)
	
}