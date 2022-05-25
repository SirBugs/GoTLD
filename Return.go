package GoTLD

import (
	"fmt"
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
	// fmt.Println("Your Main URL Is:", MainURL)

	// Extracting SSL
	if strings.Contains(MainURL, ":") {
		SSL1 := strings.Split(MainURL, "://")
		SSL = SSL1[0]
		FULLSSL := SSL + "://"
		// fmt.Println("URL SSL:", SSL)
		SecondStepMainURL := strings.ReplaceAll(MainURL, FULLSSL, "")
		if strings.Contains(SecondStepMainURL, "/") {
			SecondStepMainURL = strings.ReplaceAll(SecondStepMainURL, "/", "")
			// fmt.Println("Clean URL Step1:", SecondStepMainURL)
		} else {
			// fmt.Println("Clean URL Step1:", SecondStepMainURL)
		}
		SecondStepURL = SecondStepMainURL
		// fmt.Println(SecondStepURL)
	} else {
		SecondStepURL = strings.ReplaceAll(MainURL, "/", "")
	}
	// Extracting Subdomain / Domain
	if strings.Contains(SecondStepURL, ".") {
		DotsCount := strings.Count(SecondStepURL, ".")
		// fmt.Println("Dots In Your Domain Are:", strconv.Itoa(DotsCount))
		if DotsCount == 1 {
			Splitting1 := strings.Split(SecondStepURL, ".")
			DomainName = Splitting1[0]
			TLD := Splitting1[1]
			fmt.Println("DomainName:", DomainName)
			fmt.Println("TLD:", TLD)
		} else {
			Splitting1 := strings.Split(SecondStepURL, ".")
			// fmt.Println(Splitting1)
			// fmt.Println(strconv.Itoa(len(Splitting1)))
			Subdomain = Splitting1[0]
			DomainName = Splitting1[1]

			MyAllText := ""
			for i, s := range Splitting1 {
				// fmt.Println(i, s)
				MyAllText += "." + s
				_ = i // Handle Declaring Err
			}
			// TLD := Splitting[]
			// fmt.Println(MyAllText)
			TLD = strings.ReplaceAll(MyAllText, "."+Subdomain, "")
			TLD = strings.ReplaceAll(TLD, "."+DomainName, "")

			// fmt.Println("Your Link SSL:", SSL)
			// fmt.Println("Your Subdomain Is:", Subdomain)
			// fmt.Println("Your DomainName Is:", DomainName)
			// fmt.Println("Your TLD Is:", TLD)
		}
	}

	switch wanted {
	case "Full", "full", "FULL":
		ReturnedValue = MainURL
	case "SSL", "ssl", "Ssl":
		ReturnedValue = SSL
	case "Sub", "sub", "SUB", "Subdomain", "SUBDOMAIN", "SubDomain":
		ReturnedValue = Subdomain
	case "Name", "name", "NAME", "Domain", "domain", "DOMAIN", "DomainName", "DOMAINNAME":
		ReturnedValue = DomainName
	case "TLD", "Tld", "tld":
		ReturnedValue = TLD
	}

	return ReturnedValue
}
