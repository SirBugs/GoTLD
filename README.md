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

### What Can I Get?
- **Outputs**
  - For getting Full Link/Domain: **Full, full, FULL**
  - For getting ssl use: **SSL, ssl, Ssl**
  - For getting subdomain use: **Sub, sub, SUB, Subdomain, SUBDOMAIN, SubDomain, subdomain**
  - For getting domainname use: **Name, name, NAME, Domain, domain, DOMAIN, DomainName, DOMAINNAME, domainname**
  - For getting tld use: **TLD, Tld, tld**

### Contact
**Facebook: SirBugs**

### LICENSE
Copyright © 2022 SirBugs

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
