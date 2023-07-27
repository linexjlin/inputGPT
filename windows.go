//go:build windows
// +build windows

package main

import "github.com/lxn/win"

func OSDepCheck() {
	if !isDebug() {
		win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)
	}
}
