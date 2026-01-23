package myutil

import "fmt"

func Maps() {
	// MAPS
	domainToIp := map[string]string{}
	domainToIp["https://www.google.com"] = "8.8.8.8"
	domainToIp["https://www.amazon.com"] = "98.87.174.74"
	domainToIp["https://www.microsoft.com"] = "13.107.213.68"
	url := "https://www.amazon.com"
	fmt.Printf("Ip of %s is %s\n", url, domainToIp[url])
	fmt.Println()

	delete(domainToIp, "https://www.microsoft.com")
	for k, v := range domainToIp {
		fmt.Printf("Ip of %s is %s\n", k, v)
	}
	fmt.Println()
}
