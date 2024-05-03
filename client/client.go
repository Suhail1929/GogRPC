package main

import (
	"Bakri-Souhail/GoGrpcClient/data"
	"fmt"
	"time"

)


func main() {
	for i := 1; i <= 5; i++ {
		data.ShowDevices(fmt.Sprintf("%d", i))
	    time.Sleep(3 * time.Second)
	}

	for true {
		time.Sleep(3 * time.Second)
	}
	// os.Exit(0)
}


