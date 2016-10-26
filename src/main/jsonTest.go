package main

import (
	"main/net"
)

func main() {
	jstr, _ := net.PersonJson()
	net.JsonPerson(jstr)
}
