# Prompt templates

## ChatGPT
```json
{
  "name": "ChatGPT",
  "model": "gpt-3.5-turbo-0613",
  "headMessages": [
    {
      "role": "system",
      "content": "You are a helpful assistant."
    }
  ],
  "maxContext": 20
}
```
## 翻译 Translate
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

```json
{
    "name":"translate to english",
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
## 自动回复 Automatic reply
```json
{
    "name":"自动回复",
    "model": "gpt-3.5-turbo-0613",
    "headMessages": [
      {
        "role": "system",
        "content": "下面是我收的消息，帮我用礼貌的方式回复它"
      }
    ],
    "maxContext": 6
  }
```

## 总结文章　Summary of the article
```json
{
    "name":"总结文章",
    "model": "gpt-3.5-turbo-16k",
    "headMessages": [
      {
        "role": "system",
        "content": "以中文markdown格式总结下面这篇文章"
      }
    ],
    "maxContext": 1
  }
```

## 完型填空 Cloze
Not only for reading comprehension, it can also be used for code completion and poetry completion.

```json
{
    "name": "完型填空",
    "model": "gpt-3.5-turbo-0613",
    "headMessages": [
      {
        "role": "system",
        "content": "Cloze Test. Do not explain."
      },
      {
        "role": "user",
        "content": "填上合适的句子\n君不见黄河之水天上来，_\n君不见高堂明镜悲白发，\n朝如青丝暮成雪。"
      },
      {
        "role": "assistant",
        "content": "奔流到海不复回。"
      },
      {
        "role": "user",
        "content": "陈王昔时宴平乐，斗酒十千恣欢谑。\n_，径须沽取对君酌。"
      },
      {
        "role": "assistant",
        "content": "主人何为言少钱。"
      }
    ],
    "maxContext": 1
  }
```