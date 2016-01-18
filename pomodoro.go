package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var currentRound = 0

func main() {

	fmt.Println("Pomodoro Timer!")
	wTime, err := getWorkTime()

	if err != nil {
		usageMessage()
		return
	}
	
	for {

		startWork(wTime)
		currentRound++
		printTomatoes()

		if currentRound < 4 {

			takeBreak()
		} else {

			takeLongBreak()
			currentRound = 0
		}

		fmt.Println()
	}
}

func getWorkTime() (time.Duration, error) {

	var err error
	var wTime time.Duration = 25

	if len(os.Args) > 1 {
		t, err := strconv.Atoi(os.Args[1])
		if err != nil {
			return 0, err
		}

		wTime = time.Duration(t)
	}

	return wTime, err
}

func usageMessage() {

	fmt.Println("usage: pomo [minutes]")

}

func startWork(wTime time.Duration) {

	fmt.Println("Start work!")
	

	time.Sleep(wTime * time.Second)

}

func takeBreak() {

	fmt.Println("Take a break!")
	time.Sleep(3 * time.Second)
}

func printTomatoes(){
	
	for i := 0; i < currentRound; i++ {
		fmt.Print("🍅  ")
	}
	fmt.Println()
}

func takeLongBreak() {

	fmt.Println("Take a longer break!")
	time.Sleep(15 * time.Second)
}
