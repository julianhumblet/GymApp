package main

import (
	"GymApp/logfile"
	"fmt"
	"os"
)

func init() {
	pathLogfile := "./logs.log"
	err := logfile.LoadLogfile(pathLogfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

}
