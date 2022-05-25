# GoTLD in GOlang

GoTLD is a importable module to split domain names parts

**Note:** You can import by `GoTLD "github.com/SirBugs/GoTLD"` or `GoTLD "GoTld"` or `"GoTld"`
**Note:** By depending on: https://pkg.go.dev/golang.org/x/net/publicsuffix

**Note: Go Version: 1.18**

### Installation (Linux, Mac, Windows)
```
go get github.com/SirBugs/GoTLD
```
```
Download file directly to C://Program Files/Go/Src/
```

### Usage1 (Return)
```go
package main

import (
	// GoTLD "GoTLD"
	GoTLD "github.com/SirBugs/GoTLD"
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

### What Can I Get?:
- **Outputs**
  - For getting Full Link/Domain: **Full, full, FULL**
  - For getting ssl use: **SSL, ssl, Ssl**
  - For getting subdomain use: **Sub, sub, SUB, Subdomain, SUBDOMAIN, SubDomain, subdomain**
  - For getting domainname use: **Name, name, NAME, Domain, domain, DOMAIN, DomainName, DOMAINNAME, domainname**
  - For getting tld use: **TLD, Tld, tld**

