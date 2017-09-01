package main 

import (
	"os"
	"fmt"
	//"net/http"
	//"encoding/json"
	//"time"
)

var (
	kPackage  string = "com.eonsun.myreader"
	kAdSlotId string = "729d94d23cf8ecba91ab40fcc776dbfa"
	kDeviceInfoTestIp string = "111.200.53.10"
	kIpReuseLimit int = 5
	urls = []string {"/anjian", "AnjianHandler", "/ad","Request", "/get_task", "GetTask"}
	resultDict string
)

type anJianHandler struct {


}

type deviceGenerator struct {
	data interface
	info string
	loadCheckPoint func()
	 
}
func loadCheckPoint() {
	locationInt := 0
	context, err := os.Open(config.inputDeviceLocationPath)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		deviceGenerator.data
	}

}

func main() {
	fmt.Print("back")
}