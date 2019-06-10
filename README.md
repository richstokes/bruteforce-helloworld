## Overview

Saw this idea on Reddit and thought it sounded fun. Made a helloworld bruteforcer, with full unicode support.  

![example gif](https://github.com/richstokes/bruteforce-helloworld/raw/master/example.gif)

&nbsp;  

## Usage  

```
  -w string
        A string to bruteforce (default "hello world")  
  -b    Enables benchmark mode if set (Use same random number generator each time)
```

Best ran in a clear terminal, since it uses control characters to update the console in place:  

```
reset && go run main.go -w 'Hello world 👾🐈🏀🎸🍩🤯'
```