package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, this is ip tracking application. Please enter ip address you want to track: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		ip := scanner.Text()
		fmt.Println("Checking IP Address: ", stripProtocol(ip))
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}

// Strip protocol and trailing slash from IP Address
func stripProtocol(ipAddress string) string {
	if strings.Contains(ipAddress, "https://") {
		strippedProtocol := strings.ReplaceAll(ipAddress, "https://", "")
		return strings.ReplaceAll(strippedProtocol, "/", "")
	}
	strippedProtocol := strings.ReplaceAll(ipAddress, "http://", "")
	return strings.ReplaceAll(strippedProtocol, "/", "")
}
