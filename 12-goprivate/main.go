package main

import (
	"fmt"

	"github.com/orodrigosouzadev/goexpert-fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
