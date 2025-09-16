package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	//Declare host variable and take user input for hostname
	host := flag.String("h")
	fmt.Scan(&host)
	//Use LookupHost to get IP and print
	addr, err := net.LookupHost(host)
	if err != nil {
		log.Fatalf("Error looking up host %s: %v", host, err)
	}
	print("The IP for ", host, " is: \n")
	for _, x := range addr {
		fmt.Println(x)
	}
}
