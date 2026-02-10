package test

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go func() {
		chan1 <- 1
		time.Sleep(5 * time.Second)
	}()
	go func() {
		chan2 <- 1
		time.Sleep(5 * time.Second)
	}()

	time.Sleep(2 * time.Second)
	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	default:
		fmt.Println("default")
	}
	fmt.Println("main exit.")
}

func TestSelect2(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go func() {
		close(chan1)
	}()
	go func() {
		close(chan2)
	}()
	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	}
	fmt.Println("main exit.")
}
