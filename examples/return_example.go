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
