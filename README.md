# InputGPT 
A program let you query GPT from any input area. 

## What the app can do?
* Call LLM without disconnecting from the workspace.
* Call LLM in the fastest way possible.

## Features
*  Cross-platform [win,mac]
*  Multilingual  [en,zh]
*  Customize prompts

# Usage
## Quick start
1. Copy the text to clipboard
1. Click `shift + space` to query GPT

## User define HotKey 

Add a new line to .env file like this `GPT_HOTKEYS=space+shift`
the keycode [reference](https://github.com/vcaesar/keycode/blob/main/keycode.go):

## Import Prompt
Just Copy the json like below copy one of them:
```json
{
  "name": "ChatGPT",
  "model": "gpt-3.5-turbo-0613",
  "headMessages": [
    {
      "role": "system",
      "content": "You are a help assistant."
    }
  ],
  "maxContext": 20
}
```

```json
{
  "name": "翻译成中文",
  "model": "gpt-3.5-turbo-0613",
  "headMessages": [
    {
      "role": "system",
      "content": "Your a translator you translate any text I give you into Chinese. Just give me the result, do not explain."
    }
  ],
  "maxContext": 1
}
```
then click then import menu of the app．

![](https://ipfs.ee/ipfs/QmPW2FcmLvfZLbT5Ak6FYWRSc9FWJ5p3waQ4PrCPEzeH5R/6d498736-0911-460a-8fe2-8e91c8ca3340.png)

[For more templates](./prompts)

# DEMO
![](https://ipfs.ee/ipfs/QmNcQVdbLMm9WwjyHce4vvPL1mQhi1VJdAkc1B6sy69GdJ/9aee063e-5898-4429-81e0-ef7ba20521d3.png)

[![IMAGE ALT TEXT](http://img.youtube.com/vi/2EpdfYILbgQ/0.jpg)](https://www.youtube.com/watch?v=2EpdfYILbgQ "InputGTP DEMO")

[translate demo](https://ipfs.ee/ipfs/QmepH3EbP71zaXxaLAfQt2domXZxnb7HuaAkxT4jzhajmk/7c5ec8d0-a3d2-4d06-b649-316456390599.mp4)


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