package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"ip-tracking/configs"
	"net/http"
	"os"
	"strings"
	"time"
)

type ResponseStruct struct {
	City         City          `json:"city"`
	Continent    Continent     `json:"continent"`
	Country      Country       `json:"country"`
	Location     Location      `json:"location"`
	SubDivisions []SubDivision `json:"subdivisions"`
	Traits       Traits        `json:"traits"`
	Intelligence Intelligence  `json:"intelligence"`
}

type City struct {
	GeoNameID int   `json:"geoname_id"`
	Name      Names `json:"names"`
}

type Continent struct {
	Code      string `json:"code"`
	GeoNameID int    `json:"geoname_id"`
	Name      Names  `json:"names"`
}

type Country struct {
	GeoNameID         int    `json:"geoname_id"`
	IsInEuropeanUnion bool   `json:"is_in_european_union"`
	IsoCode           string `json:"iso_code"`
	Name              Names  `json:"names"`
}

type SubDivision struct {
	GeoNameID int    `json:"geoname_id"`
	IsoCode   string `json:"iso_code"`
	Name      Names  `json:"names"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"time_zone"`
}

type Traits struct {
	AutonomousSystemNumber       int    `json:"autonomous_system_number"`
	AutonomousSystemOgranization string `json:"autonomous_system_ogranization"`
	ConnectionType               string `json:"connection_type"`
	ISP                          string `json:"isp"`
	Organization                 string `json:"organization"`
	UserType                     string `json:"user_type"`
}

type Intelligence struct {
	Summary  Summary  `json:"summary"`
	Flags    Flags    `json:"flags"`
	Risk     Risk     `json:"risk"`
	Security Security `json:"security"`
}

type Summary struct {
	Verdict     string `json:"verdict"`
	Description string `json:"description"`
}

type Flags struct {
	IsAnonymous      bool `json:"is_anonymous"`
	IsProxy          bool `json:"is_proxy"`
	IsVPN            bool `json:"is_vpn"`
	IsRelay          bool `json:"is_relay"`
	IsMalicious      bool `json:"is_malicious"`
	IsHosting        bool `json:"is_hosting"`
	IsNetworkService bool `json:"is_network_service"`
}

type Risk struct {
	Score int    `json:"score"`
	Level string `json:"level"`
}

type Security struct {
	IsMalicious bool     `json:"is_malicious"`
	ThreatScore int      `json:"threat_score"`
	ThreatLevel string   `json:"threat_level"`
	ThreatType  []string `json:"threat_types"`
	Confidence  int      `json:"confidence"`
	LastSeen    string   `json:"last_seen"`
}

type Names struct {
	De   string `json:"de"`
	En   string `json:"en"`
	Es   string `json:"es"`
	Fr   string `json:"fr"`
	Ja   string `json:"ja"`
	Ko   string `json:"ko"`
	PtBR string `json:"pt-BR"`
	Ru   string `json:"ru"`
	ZhCN string `json:"zh-CN"`
}

func main() {
	fmt.Println("Hello, this is ip tracking application. Please enter ip address you want to track (or type exit or quit if you want to terminate): ")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if strings.EqualFold(input, "exit") || strings.EqualFold(input, "quit") {
			fmt.Println("Bye!")
			break
		} else {
			fmt.Println("Checking IP Address: ", stripProtocol(input))
			checkResult := checkIPAddress(stripProtocol(input))

			fmt.Println("")
			fmt.Println("Result checking: ")
			fmt.Println("- IP Address: " + input)
			fmt.Println("- Continent: " + checkResult.Continent.Name.En)
			fmt.Println("- Country: " + checkResult.Country.Name.En)
			fmt.Println("- City: " + checkResult.City.Name.En)
			fmt.Println("- Location: ")
			fmt.Printf("  - Lat: %v \n", checkResult.Location.Latitude)
			fmt.Printf("  - Long: %v \n", checkResult.Location.Longitude)
			fmt.Println("  - Timezone: " + checkResult.Location.Timezone)
			fmt.Println("- Intelligence: ")
			fmt.Println("  - Organization: " + checkResult.Traits.AutonomousSystemOgranization)
			fmt.Println("  - ISP: " + checkResult.Traits.ISP)
			fmt.Println("  - Verdict: " + checkResult.Intelligence.Summary.Verdict)
			fmt.Println("  - Summary: " + checkResult.Intelligence.Summary.Description)
			fmt.Printf("  - Risk Score: %v \n", checkResult.Intelligence.Risk.Score)
			fmt.Printf("  - Level: %v \n", checkResult.Intelligence.Risk.Level)
			fmt.Printf("  - Is Malicious: %v \n", checkResult.Intelligence.Security.IsMalicious)
			fmt.Printf("  - Threat type: %v \n", checkResult.Intelligence.Security.ThreatType)
			fmt.Printf("  - Confidence Level of Threat: %v \n", checkResult.Intelligence.Security.Confidence)
			fmt.Println("")
			fmt.Println("Want to search again? please enter ip address you want to track (or type exit or quit if you want to terminate): ")
		}

		if scanner.Err() != nil {
			fmt.Println(scanner.Err())
		}
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

func checkIPAddress(ipAddress string) ResponseStruct {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := configs.ApiEndPoint() + "/" + ipAddress + "/?token=" + configs.ApiKey()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("Request timed out")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response: ", err)
	}

	var responseStruct ResponseStruct
	errUnmarshal := json.Unmarshal(body, &responseStruct)
	if errUnmarshal != nil {
		fmt.Println("Error unmarshalling response: ", errUnmarshal)
	}

	return responseStruct
}
