## Overview

Saw this idea and thought it sounded fun. Made a helloworld bruteforcer, with unicode support.  

Requires [pixel](https://github.com/faiface/pixel) for drawing the window:  

```
go get github.com/faiface/pixel
go get github.com/faiface/glhf
go get github.com/go-gl/glfw/v3.3/glfw
``` 

&nbsp;  

## Usage  

```
  -w string
        A string to bruteforce (default "hello world")  
  -b    Enables benchmark mode if set (Use same random number generator each time)
```

### Example

```
go run main.go -w 'Hello world 👾🐈🏀🎸🍩🤯'
```
