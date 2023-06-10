// Copyright 2016 The go-vgo Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-vgo/robotgo/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package main

import (
	"context"
	"fmt"

	hook "github.com/robotn/gohook"
)

func addEvent() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("--- Please press ctrl + shift + q ---")
	ok := hook.AddEvents(ctx, "q", "ctrl", "shift")
	if ok {
		fmt.Println("add events...")
	}

	fmt.Println("--- Please press w---")
	ok = hook.AddEvents(ctx, "w")
	if ok {
		fmt.Println("add events")
	}

	// start hook
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	s := hook.Start(ctx, hook.BothHookEnabled)

	for ev := range s {
		fmt.Println("hook: ", ev)
	}
}

func addMouse() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("--- Please press left mouse button ---")
	ok := hook.AddMouse(ctx, "left")
	if ok {
		fmt.Println("add mouse...")
	}

	fmt.Println("--- Please press left mouse button and move mosue to 100,100 ---")
	ok = hook.AddMouse(ctx, "left", 100, 100)
	if ok {
		fmt.Println("add mouse and move to 100,100 ...")
	}

	fmt.Println("--- Please move mosue to 100,100 ---")
	ok = hook.AddMousePos(ctx, 100, 100)
	if ok {
		fmt.Println(" move mouse to 100,100 ...")
	}
}

func add() {
	fmt.Println("--- Please press v---")
	eve := hook.AddEvent("v")

	if eve {
		fmt.Println("--- You press v---", "v")
	}

	fmt.Println("--- Please press k---")
	keve := hook.AddEvent("k")
	if keve {
		fmt.Println("--- You press k---", "k")
	}

	fmt.Println("--- Please press f1---")
	feve := hook.AddEvent("f1")
	if feve {
		fmt.Println("You press...", "f1")
	}
}

func event() {
	////////////////////////////////////////////////////////////////////////////////
	// Global event listener
	////////////////////////////////////////////////////////////////////////////////

	add()

	fmt.Println("--- Please press left mouse button---")
	mleft := hook.AddEvent("mleft")
	if mleft {
		fmt.Println("--- You press left mouse button---", "mleft")
	}

	mright := hook.AddEvent("mright")
	if mright {
		fmt.Println("--- You press right mouse button---", "mright")
	}

	// stop AddEvent
	// hook.StopEvent()
}

func main() {
	fmt.Println("test begin...")

	addEvent()

	addMouse()

	event()
}
