package main

import (
	"fmt"
	"os"
	"strings"
)

func Splitter(link string, wanted string) string {

	SecondStepURL := ""
	SSL := ""
	Subdomain := ""
	DomainName := ""
	TLD := ""
	ReturnedValue := ""

	MainURL := link

	// Extracting SSL
	if strings.Contains(MainURL, ":") {
		SSL1 := strings.Split(MainURL, "://")
		SSL = SSL1[0]
		FULLSSL := SSL + "://"
		SecondStepMainURL := strings.ReplaceAll(MainURL, FULLSSL, "")
		if strings.Contains(SecondStepMainURL, "/") {
			SecondStepMainURL = strings.ReplaceAll(SecondStepMainURL, "/", "")
		} else {
		}
		SecondStepURL = SecondStepMainURL
	} else {
		SecondStepURL = strings.ReplaceAll(MainURL, "/", "")
	}
	// Extracting Subdomain / Domain
	if strings.Contains(SecondStepURL, ".") {
		DotsCount := strings.Count(SecondStepURL, ".")
		if DotsCount == 1 {
			Splitting1 := strings.Split(SecondStepURL, ".")
			DomainName = Splitting1[0]
			TLD := Splitting1[1]
			fmt.Println("DomainName:", DomainName)
			fmt.Println("TLD:", TLD)
		} else {
			Splitting1 := strings.Split(SecondStepURL, ".")
			Subdomain = Splitting1[0]
			DomainName = Splitting1[1]

			MyAllText := ""
			for i, s := range Splitting1 {
				MyAllText += "." + s
				_ = i // Handle Declaring Err
			}
			TLD = strings.ReplaceAll(MyAllText, "."+Subdomain, "")
			TLD = strings.ReplaceAll(TLD, "."+DomainName, "")
		}
	}

	switch wanted {
	case "Full", "full", "FULL":
		ReturnedValue = MainURL
	case "SSL", "ssl", "Ssl":
		ReturnedValue = SSL
	case "Sub", "sub", "SUB", "Subdomain", "SUBDOMAIN", "SubDomain", "subdomain":
		ReturnedValue = Subdomain
	case "Name", "name", "NAME", "Domain", "domain", "DOMAIN", "DomainName", "DOMAINNAME", "domainname":
		ReturnedValue = DomainName
	case "TLD", "Tld", "tld":
		ReturnedValue = TLD
	}

	return ReturnedValue
}

func main() {
	if os.Args[1] != "" && os.Args[2] != "" {
		// fmt.Println("We Have Smth To Do!", os.Args[1], os.Args[2])
		fmt.Println(Splitter(os.Args[1], os.Args[2]))
	}
}
