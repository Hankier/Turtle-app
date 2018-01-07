package main


//Client for testing adding to DAserver

import (
         "fmt"
         "net"
         "log"
         "encoding/json"
    )


type Person struct {
	Type   string
	Data    string
}

type msgMessage struct{
    Type    string
    Content string
    Content2 string
}

type response struct{
    Status string
    Name string
}
func main() {
    fmt.Println("Hello!")
    d, err := net.Dial("tcp", "0.0.0.0:12345")
    if err != nil {
        log.Fatal(err)
    }
    enc := json.NewEncoder(d)

    masnun := msgMessage{
		Type:   "ADD",
        Content:    "10.10.10.10:1300",
        Content2: `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA0dkp8n9H3c7/OlaYKpi8Fu7SWsqF3eeRw3C8gpLPkMhTDFRm
/ZpytL/ImYAh1mQhvvH5kHPHc7unARSDlaa/OzDcMqwSlXmzIWw1k81a3HQj+1wB
anbIftyo/Pc1WwKQ35ppCtwJWytCipeqD3+ACEbhVmb+4TTHX3/MO5E1mWFNISJW
VNTMxIibqBPG2cDjRAu4kGNaVM1pd1/AZrid6LtvmUaXajLhodywtopBGUana4qR
5ou9OPYc3gELMkMQCM4AUqWM04iJ0+620rmoeFHRu3L4miekU+bwZiDiNXG4wXDp
HgmbStRVUhLHntYGs+4bZBfghW5WZqDxNsIMPwIDAQABAoIBAQCXftDyuYLXlgXa
RwPJ1MQNRlLkqsrj/bbUwsHE/loNKyIRh6lmsqbW6JHYh5FmJpnaMPS7nWpDmhii
Bf5M/rmV8Ns3VdSAxwBUQ7uWPa2387y6TZzUEHcEZyc0oP+K+Zo/Y0ksRtgWUm/S
gFWMpL54uzsY1nhxe1noDuoRou5wEGkmatlXr0FvME4OmwOe1S24bQg97A9dqngh
PELJ+pdQKcSqWAgW8Y3PMxx88DI08DOF1IxyV8VvTrWhbzJvTZnpQqxy3pDFr5TO
S287PicRAeCyDM5Rxi6tzJaLuh19RT4YtKZ1qgJRQUEOtM6Pf0rT7/YAjiFdpoT0
MNNboGCBAoGBAOdqwI+KEvPf5RXaTeiZcymS9fwGCyObDfOvtrOhKAc+cv5wNwJG
g6LsK+i/5EWi9INEd8O1WcYVaqjSG6r36zxjcWOKrRNU1gyvE8ONg+DuQ34nZz0V
rjgnzSIzK4OlOdPpfU/wKuVipC0quY2HGODxrV7KtaU4++mr4bIefEPBAoGBAOgj
3RK6t2jCABmGG1iruY2/YEhT+blOJ9A6YKMptfWMwuewL5A11vH1j00JKm6N1twQ
kHKSdUVurZArwArQ7LxLNa2VhTzjMIviMtnTQc+0CdkUKNGeCvx2JVQbzr8Av2eg
r/Dt7weSBLerWFq2cAwvrRZ665sx6cRU3QTckk//AoGAe/kXgY4hix6F1kgl9pbG
OB5vwvzl2MRHHCYlBWQvUnolFqO9BG4MNSq6dyzduGSNAwmZ83Fiz5hHlHtCsTux
fJ91bjMrdzC6nv7n4pocbVKXO60WRIYp2BGSdmDdTeAk856hMELkaBCJDV1XHDek
n1U5YI/N8d5uLgeTmF12isECgYEAq1F8X8woe0lhJXURTXk+cVvhRL+ktpr1Svkq
RIAN52/Aj5g5IeZ6AQtGfIXdKMXI4ZPf5o4rudgagyGmktTpQXUH4llMgUjxlOqU
uKjuEsk901TLYxeN6A+RMOdsxw1YNLQj5FzUYPPkQ2BSzm+BdZzh0otYwaouaVRv
4Jyf5iUCgYAFAzx2WWGhDGnyMweiEfKmwoKDaOAShDadvFmwy5irTKiceflQRhHJ
obORGt2A/HtVMpiG+4fjHnAMaKWahaTJ5FKgaRHWahYjVsPRRJwrqbLILEo9Q90N
CuHDJhgWQwJl6/9NXxbFTrCcBRTd5SHvptaUh+ZponwS/ltLdi6bPQ==
-----END RSA PRIVATE KEY-----`,
	}
    enc.Encode(masnun)
    dec := json.NewDecoder(d)

    var odp response

    err = dec.Decode(&odp)

    log.Printf("Odp Status: %v", odp.Status)
    log.Printf("Odp Name: %v", odp.Name)
}
