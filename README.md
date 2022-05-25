# GoTLD in GOlang

GoTLD is a importable module to split domain names parts

**Note:** You can import by `GoTLD "GoTld"` or `"GoTld"`
**Note:** By depending on: https://pkg.go.dev/golang.org/x/net/publicsuffix

### Installation (Linux, Mac)
```
go get github.com/SirBugs/GoTLD
```
### Installation (Windows)
```
Download file directly to C://Program Files/Go/Src/
```

### Usage1 (Return)
```go
package main

import (
	GoTLD "GoTLD"
	"fmt"
)

func main() {
	fmt.Println("Your https://maps.google.com/ TLD is:", GoTLD.Splitter("https://maps.google.com/", "TLD"))
	fmt.Println("Your https://maps.google.com/ SSL is:", GoTLD.Splitter("https://maps.google.com/", "SSL"))
	fmt.Println("Your https://maps.google.com/ DomainName is:", GoTLD.Splitter("https://maps.google.com/", "Name"))
	fmt.Println("Your https://maps.google.com/ Subdomain is:", GoTLD.Splitter("https://maps.google.com/", "Sub"))
}
```
```
Your https://maps.google.com/ TLD is: .com
Your https://maps.google.com/ SSL is: https
Your https://maps.google.com/ DomainName is: google
Your https://maps.google.com/ Subdomain is: maps
```

### Usage2 (Args)
```
go run main.go https://maps.google.com/ SSL
```
```
https
```
