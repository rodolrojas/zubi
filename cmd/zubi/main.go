package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"rodolrojas.com/zubi/internal/gateway"
)

func main() {	
	logrus.SetLevel(logrus.DebugLevel)
	fmt.Println("Loading Zubi Gateway...")
	var instance = gateway.NewGateway()
	if err := instance.Start(); err != nil {
		fmt.Printf("Error starting gateway: %v\n", err)
	}
}