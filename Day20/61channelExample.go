package main

import (
	"fmt"
	"time"
)

func sendEmail(emailchan chan string,done chan bool){

	defer func ()  {
		done <- true
	}()
	for email:= range emailchan{
		fmt.Println("Sending email to:", email)
		time.Sleep(1*time.Second)
	}
}
func main() {

	emailchan := make(chan string, 5)
	done := make(chan bool)
	go sendEmail(emailchan, done)
	for i := 1; i <= 5; i++ {
		emailchan <- fmt.Sprintf("%d@gmail.com",i)
	}
	fmt.Println("Sending Done")
	close(emailchan)
	<-done
	close(done)


}