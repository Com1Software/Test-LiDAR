package main

import (
	"fmt"
	"log"
	"strings"

	"go.bug.st/serial"
)

func main() {
	fmt.Println("hello world")

	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	// Print the list of detected ports
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	// Open the first serial port detected at 9600bps N81
	mode := &serial.Mode{
		// BaudRate: 57600,
		BaudRate: 230400,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(ports[0], mode)
	if err != nil {
		log.Fatal(err)
	}

	//	line := ""
	buff := make([]byte, 1)
	on := true
	max := 10000
	cl := 0
	for on != false {
		for {
			n, err := port.Read(buff)
			if err != nil {
				log.Fatal(err)
			}
			if n == 0 {
				fmt.Println("---EOF \nEOF")
				break
			}

			fmt.Printf("%v", buff[:n])
			if max > 0 {
				cl++
				if cl > max {
					on = false
				}

			}
			// If we receive a newline stop reading
			if strings.Contains(string(buff[:n]), "\n") {
				break
			}
		}
	}
}
