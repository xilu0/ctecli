# ctecli 

A command line interface for the coze

## Installation 
    go get github.com/xilu0/ctecli


## config
    ctecli.exe config --token=

go to [https://www.coze.com/token](https://www.coze.com/token) to get token

##  Usage
```
ctecli.exe -c "hi, how is it going?"
```

## example 

1. chat 
```
ctecli.exe -c "hi, how is it going?"
Hi! It's going well, thanks. How about you? 😊 
```

2. fix grammer 
```
ctecli.exe -c "I likes you."        
I like you too! However, it seems you've made a tiny grammar error. 😊

Could it be that you meant to say "I **like** you"?  Remember, with "I" we use the base form of the verb.

Let me know if you have any other phrases you'd like to practice! I'm always happy to help. 😄
```

3. translate
```
ctecli.exe -c "翻译：明天是星期几？"
What day is tomorrow?
```