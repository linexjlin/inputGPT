# InputGPT 
A program let you query GPT from any input area. 

## What the app can do?
* Call LLM without disconnecting from the workspace.
* The fastest way to get answer from GPT.

## Features
*  Cross-platform [win,mac]
*  Multilingual  [en,zh,jp,es,az,fr]
*  Customize prompts

# Usage

## Quick start
1. Open InputGPT click "Set API KEY" to provide OpenAI Key. 
1. Select the text for query, copy the text to clipboard
1. Click HotKey `shift + space` to query GPT
1. One click `ESC` key to stop generate
1. Triple click `ESC` key to quick clear context

## User define HotKey 

Click "Set API KEY" InputGPT will open `env.txt` file

Add a new line to the file like this `GPT_HOTKEYS=space+shift` then save and close the file.

the keycode [reference](https://github.com/vcaesar/keycode/blob/main/keycode.go):

![](https://ipfs.ee/ipfs/QmTusxyAgEg8cFWU7dvtPcU5R7tY3mxpZMd36dBkixTQhT/3b12e13c-19cf-4392-867b-74b35e030e9c.png)

## Import Prompt

Just Copy the json like below copy one of them:

Act as ChatGPT
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

Translate

```json
{
  "name": "Translate",
  "model": "gpt-3.5-turbo",
  "headMessages": [
    {
      "role": "system",
      "content": "Your are a translator, you translate any text I give you into {{.mylang}}. Just give me the result, do not explain."
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

![](https://ipfs.ee/ipfs/QmTGUWr8TzEurjv3MiUuhuzuex5Pb4Lb2SRiBsxmqGErnx/08a6b1b8-6436-4556-b5a8-87b305f577b7.png)

[For more templates](https://inputgpt.vercel.app/examples)

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
