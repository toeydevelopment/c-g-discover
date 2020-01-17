# Format number

## How does it work

```golang
package main

import (
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	ans := ""

	phoneNumber := args[0] 

    format := args[1]
    
    // retreive string that contain only x
	onlyX := strings.ReplaceAll(format, "-", "")

	if len(onlyX) != len(phoneNumber) {
		panic("Wrong input x character must be equal input number")
	}

    // counting offset
	checkIndex := 0

	for i := 0; i < len(format); i++ {
		if strings.ToLower(string(format[i])) == "x" {
			ans += string(phoneNumber[i-checkIndex])
		} else {
			checkIndex += 1
			ans += "-"
		}
	}

	print(ans)
}

```


## How to run

if 
    you using mac you can run it as binary without compiled 

```bash
./formatter 1111111111 xx-xxxx-xxxx
```

else 
    copy all code and go to https://play.golang.org/ then change 3 variable like this
```golang

    package main

import (
	"strings"
)

func main() {

	ans := ""

	phoneNumber := //put your number

    format := // put your format
    
    // retreive string that contain only x
	onlyX := strings.ReplaceAll(format, "-", "")

	if len(onlyX) != len(phoneNumber) {
		panic("Wrong input x character must be equal input number")
	}

    // counting offset
	checkIndex := 0

	for i := 0; i < len(format); i++ {
		if strings.ToLower(string(format[i])) == "x" {
			ans += string(phoneNumber[i-checkIndex])
		} else {
			checkIndex += 1
			ans += "-"
		}
	}

	print(ans)
}
```
Go Playground not supported parse argument via command line

