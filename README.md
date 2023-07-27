# InputGPT 
A program let you query GPT from any input area. 

# Usage
1. Copy the text to clipboard
1. Click `shift + space` to query GPT

## User define HotKey 
Add a new line to .env file like this `GPT_HOTKEYS=space+shift`
the keycode reference:
https://github.com/vcaesar/keycode/blob/main/keycode.go


## Add new new prompts 
The default mode is continue writing with the given text. 
Create a json file like below the add it to the prompts folder
```json
{
    "model": "gpt-3.5-turbo-0613",
    "headMessages": [
      {
        "role": "system",
        "content": "Translate all the messages I give you into English"
      }
    ],
    "maxContext": 1
  }
``` 

# DEMO
![](https://ipfs.ee/ipfs/QmNcQVdbLMm9WwjyHce4vvPL1mQhi1VJdAkc1B6sy69GdJ/9aee063e-5898-4429-81e0-ef7ba20521d3.png)

[![IMAGE ALT TEXT](http://img.youtube.com/vi/2EpdfYILbgQ/0.jpg)](https://www.youtube.com/watch?v=2EpdfYILbgQ "InputGTP DEMO")

# Build 
windows 
```cmd
build_win.bat
```

# Credit

https://github.com/getlantern/systray

https://github.com/go-vgo/robotgo

https://github.com/robotn/gohook

https://github.com/hanyuancheung/gpt-go