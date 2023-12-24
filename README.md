# InputGPT 
A program let you query GPT from any input area. 

## What the app can do?
* Call LLM without disconnecting from the workspace.
* Call LLM in the fastest way possible.

## Features
*  Cross-platform [win,mac]
*  Multilingual  [en,zh,jp,es,az]
*  Customize prompts

# Usage
## Quick start
1. Open InputGPT click "Set API KEY" to provide OpenAI Key. 
1. Copy the text to clipboard
1. Click `shift + space` to query GPT
1. Press `ESC` key to stop generate
1. Triple click `ESC` key to quick clear context

## User define HotKey 
Click "Set API KEY" InputGPT will open `env.txt` file 
Add a new line to the file like this `GPT_HOTKEYS=space+shift` then save and close the file.
the keycode [reference](https://github.com/vcaesar/keycode/blob/main/keycode.go):
![](https://ipfs.ee/ipfs/QmaBtanJEmt8krtLLAL2zE9QYyNodQ7bvkRofNuWABaZmn/d6636a7b-cb75-494f-84ac-3935382544d8.png)

## Import Prompt
Just Copy the json like below copy one of them:
```json
{
  "name": "ChatGPT",
  "model": "gpt-3.5-turbo",
  "headMessages": [
    {
      "role": "system",
      "content": "You are a helpful assistant."
    }
  ],
  "maxContext": 20
}
```

```json
{
  "name": "Translate",
  "model": "gpt-3.5-turbo",
  "headMessages": [
    {
      "role": "system",
      "content": "Your a translator you translate any text I give you into {{.mylang}}. Just give me the result, do not explain."
    },
    {
      "role": "user",
      "content": "{{.msg}}"
    }
  ],
  "maxContext": 0
}
```
then click then import menu of the app．

![](https://ipfs.ee/ipfs/QmPW2FcmLvfZLbT5Ak6FYWRSc9FWJ5p3waQ4PrCPEzeH5R/6d498736-0911-460a-8fe2-8e91c8ca3340.png)

[For more templates](https://inputgpt.vercel.app/)

# DEMO

## ChatGPT

![ChatGPT](https://ipfs.ee/ipfs/QmdQetjhkFgNDGf5HhSgbML1rRcYPWQsexxiPggATZ3qLm/d0d6c03a-b0cc-40f3-952c-cb81ef88f6f6.gif)

## Bidirectional Translation

![](https://ipfs.ee/ipfs/QmfJUmAURswjtncxk94KE9RKJUpgH72tcsN9Mq6FkGUiZp/c1fe75b6-eb44-47dd-b138-4056045e57d9.gif)

Work like github copilot complete the codes in the middle of code file.
[code cloze](https://ipfs.ee/ipfs/QmRp351kZ9fB1y1k9vWCHJq3egG8wZT39LYeVr9RhzbkVU/a159ab5f-e308-4d02-8d64-9c02ea0fc48e.mp4)

# Build 
```cmd
build_win.bat
```
or
```bash
build.sh
```

# Credit

https://github.com/getlantern/systray

https://github.com/go-vgo/robotgo

https://github.com/robotn/gohook

https://github.com/hanyuancheung/gpt-go