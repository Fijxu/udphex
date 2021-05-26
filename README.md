# udphex
This Is A UDP DoS Method Which Sends A Hex String.

# Use without compile
**Requires golang installed**\
`go run udphex.go HOST PORT`

# Compile :)
**Also requires golang installed, but you can use it in multiple machines without install golang in every machine**\
`go build`\
Then search for the udphex binary file

# Use the compiled file
`./udphex HOST PORT`\
Example: `./udphex 127.0.0.1 666`

# How to change the hex message (recommended) if you want a custom message
Why should you change the hex value?\
The program sends this message **(in this fork the message is replaced by "abc")** via UDP.\
![image](https://user-images.githubusercontent.com/61166695/116642459-f2b4d000-a93c-11eb-87fb-8794b8956823.png)\
You need modify the source code, the udphex.go file.
Go to line 18 and you will see something like this (in this fork)
```go
conn.Write([]byte("\x61\x73\x64"))
```
`\x61\x73\x64` equals to "abc"\
So if we want it to say "cba" we have to change it for `\x63\x62\x61`\
And now if you want a custom message you need a text to hex convertor, also you can use HxD\
![e9ez2DCXWl](https://user-images.githubusercontent.com/61166695/116642777-9bfbc600-a93d-11eb-8cab-a4858446f704.gif)\
Write what do you want and convert it to hex, in the gif i get `64 69 73 63 6F 72 64 2E 67 67 2F 61 61 61` as output\
So you need to delete the spaces and put `\x` every 2 numbers/letters and you should get something like this `\x64\x69\x73\x63\x6F\x72\x64\x2E\x67\x67\x2F\x61\x61\x61`\
Then copy that here and compile it or watever you want lol.\
![nERS00ObPL](https://user-images.githubusercontent.com/61166695/116643195-7de29580-a93e-11eb-9e94-bfa3f2a7a157.gif)


# Stop Attack
Ctrl+C

# Edited by 
The one that forked this.
The original repository was deleted OMEGALUL
