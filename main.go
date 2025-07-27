package main

import (
	"time"

	"github.com/lxn/walk"
)

func main() {

	if time.Now().Hour() < 8 {
		walk.MsgBox(nil, "GOOD MORNING BOZO", "Walk for a bit before working, would ya?!", walk.MsgBoxIconExclamation)
	}
	
}

