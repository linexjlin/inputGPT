# InputGPT 
A program let you query GPT from any input area. 

## What the app can do?
* Call LLM without disconnecting from the workspace.
* The fastest way to get answer from GPT.

## Features
*  Cross-platform [win,mac]
*  Multilingual  [en,zh,jp,es,az,fr]
*  Customize prompts, temperature

# Usage

## Quick start
1. Open InputGPT click "Set API KEY" to provide OpenAI Key. 
1. Select the text for query, copy the text to clipboard
1. Click HotKey `space + shift ` to query GPT
1. One click `ESC` key to stop generate
1. Triple click `ESC` key to quick clear context

## User define HotKey 

Click "Set API KEY" InputGPT will open `env.txt` file

Add a new line to the file like this `GPT_HOTKEYS=shift+space` then save and close the file.

the keycode [reference](https://github.com/linexjlin/inputGPT/blob/0f3c44b3afe93745321f38a093d9393e5c6d496a/utils.go#L37):

![](https://ipfs.ee/ipfs/QmenMRoLGro1DVQE5pF74toqrZzGsjsUjAksgpTSNsKZeP/8dhMn.png)

## Import Prompt

Just Copy the json like below copy one of them:

Act as ChatGPT
```json
{
  "name": "ChatGPT",
  "model": "gpt-4o",
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
  "model": "gpt-4o",
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

## Muti model

![](https://ipfs.ee/ipfs/QmeMPt6jMC3UeZdeAKihCBhK3yrWcuS2r7sLxbEKurkoT6/IiuLp.gif)

## [Chat](https://inputgpt.vercel.app/examples#act-as-chatgpt)

![Chat](https://ipfs.ee/ipfs/QmdQetjhkFgNDGf5HhSgbML1rRcYPWQsexxiPggATZ3qLm/d0d6c03a-b0cc-40f3-952c-cb81ef88f6f6.gif)

## [Bidirectional Translation](https://inputgpt.vercel.app/examples#bi-translator)

![](https://ipfs.ee/ipfs/QmfJUmAURswjtncxk94KE9RKJUpgH72tcsN9Mq6FkGUiZp/c1fe75b6-eb44-47dd-b138-4056045e57d9.gif)


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

https://github.com/hanyuancheung/gpt-go
