# IP Tracking
This application using API from https://findip.net, please register first into their website to get API KEY

# How To Use
1. Run command
```bash
go mod install
go run main.go
```
2. Input IP Address you want to track, for example
```aiignore
Hello, this is ip tracking application. Please enter ip address you want to track:
202.56.160.230
```
3. Result of tracking
```aiignore
Result checking: 
- IP Address: 202.56.160.230
- Continent: Asia
- Country: Indonesia
- City: Jakarta
- SubDevision: Jakarta
- Location: 
  - Lat: -6.23865 
  - Long: 106.824 
  - Timezone: Asia/Jakarta
- Intelligence: 
  - Organization: 
  - ISP: Varnion Technology Semesta, PT
  - Verdict: high_risk
  - Summary: This IP is associated with web attack signals.
  - Risk Score: 70 
  - Level: high 
  - Is Malicious: true 
  - Threat type: [web_attack] 
  - Confidence Level: 70 
```