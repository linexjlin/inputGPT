package main

import (
	"time"

	"github.com/micmonay/keybd_event"
	"golang.design/x/clipboard"
	"golang.design/x/hotkey"
)

func escKey() {
	kb, _ := keybd_event.NewKeyBonding()
	kb.SetKeys(keybd_event.VK_ESC)
	if err := kb.Launching(); err != nil {
		panic(err)
	}
	time.Sleep(time.Millisecond * 10)
}

func paste() {
	kb, _ := keybd_event.NewKeyBonding()
	kb.HasCTRL(true)
	kb.SetKeys(keybd_event.VK_V)
	if err := kb.Launching(); err != nil {
		panic(err)
	}
	time.Sleep(time.Millisecond * 10)
}

func backspace() {
	kb, _ := keybd_event.NewKeyBonding()
	kb.SetKeys(keybd_event.VK_BACKSPACE)
	if err := kb.Launching(); err != nil {
		panic(err)
	}
	time.Sleep(time.Millisecond * 10)
}

func enter() {
	kb, _ := keybd_event.NewKeyBonding()
	kb.SetKeys(keybd_event.VK_ENTER)
	if err := kb.Launching(); err != nil {
		panic(err)
	}
	time.Sleep(time.Millisecond * 10)
}

func TypeStr(s string) {
	clipboard.Write(clipboard.FmtText, []byte(s))
	paste()
}

func TypeEnter() {
	enter()
}

func TypeBackspace() {
	backspace()
}

func keyNameToModifier(kn string) hotkey.Modifier {
	switch kn {
	case "shift":
		return hotkey.ModShift
	case "ctrl":
		return hotkey.ModCtrl
	case "alt":
		return hotkey.ModAlt
	case "windows":
		return hotkey.ModWin
	default:
		return hotkey.ModShift
	}
}

func keyNamesToKey(kn string) hotkey.Key {
	switch kn {
	case "space":
		return hotkey.KeySpace
	case "1":
		return hotkey.Key1
	case "2":
		return hotkey.Key2
	case "3":
		return hotkey.Key3
	case "4":
		return hotkey.Key4
	case "5":
		return hotkey.Key5
	case "6":
		return hotkey.Key6
	case "7":
		return hotkey.Key7
	case "8":
		return hotkey.Key8
	case "9":
		return hotkey.Key9
	case "0":
		return hotkey.Key0
	case "a":
		return hotkey.KeyA
	case "b":
		return hotkey.KeyB
	case "c":
		return hotkey.KeyC
	case "d":
		return hotkey.KeyD
	case "e":
		return hotkey.KeyE
	case "f":
		return hotkey.KeyF
	case "g":
		return hotkey.KeyG
	case "h":
		return hotkey.KeyH
	case "i":
		return hotkey.KeyI
	case "j":
		return hotkey.KeyJ
	case "k":
		return hotkey.KeyK
	case "l":
		return hotkey.KeyL
	case "m":
		return hotkey.KeyM
	case "n":
		return hotkey.KeyN
	case "o":
		return hotkey.KeyO
	case "p":
		return hotkey.KeyP
	case "q":
		return hotkey.KeyQ
	case "r":
		return hotkey.KeyR
	case "s":
		return hotkey.KeyS
	case "t":
		return hotkey.KeyT
	case "u":
		return hotkey.KeyU
	case "v":
		return hotkey.KeyV
	case "w":
		return hotkey.KeyW
	case "x":
		return hotkey.KeyX
	case "y":
		return hotkey.KeyY
	case "z":
		return hotkey.KeyZ
	case "enter":
		return hotkey.KeyReturn
	case "esc":
		return hotkey.KeyEscape
	case "delete":
		return hotkey.KeyDelete
	case "tab":
		return hotkey.KeyTab
	case "left":
		return hotkey.KeyLeft
	case "right":
		return hotkey.KeyRight
	case "up":
		return hotkey.KeyUp
	case "down":
		return hotkey.KeyDown
	case "f1":
		return hotkey.KeyF1
	case "f2":
		return hotkey.KeyF2
	case "f3":
		return hotkey.KeyF3
	case "f4":
		return hotkey.KeyF4
	case "f5":
		return hotkey.KeyF5
	case "f6":
		return hotkey.KeyF6
	case "f7":
		return hotkey.KeyF7
	case "f8":
		return hotkey.KeyF8
	case "f9":
		return hotkey.KeyF9
	case "f10":
		return hotkey.KeyF10
	case "f11":
		return hotkey.KeyF11
	case "f12":
		return hotkey.KeyF12
	case "f13":
		return hotkey.KeyF13
	case "f14":
		return hotkey.KeyF14
	case "f15":
		return hotkey.KeyF15
	case "f16":
		return hotkey.KeyF16
	case "f17":
		return hotkey.KeyF17
	case "f18":
		return hotkey.KeyF18
	case "f19":
		return hotkey.KeyF19
	case "f20":
		return hotkey.KeyF20
	default:
		return hotkey.KeySpace
	}
}
