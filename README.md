## Overview

Saw this idea on Reddit and thought it was pretty funny. Made a helloworld bruteforce, but with full unicode support.  

&nbsp;  

## Usage  

```
  -w string
        A string to bruteforce (default "hello world")
```

Best ran after resetting terminal, since it uses terminal control characters to update the display:  

```
reset && time go run main.go -w 'Hello world 👾🐈🏀🎸🍩🤯'
```