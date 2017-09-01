package main

import (
	"fmt"
	"os"
	//"flag"
	"bufio"
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"
)

//var value = map[string]string{}
type u struct {
	InputDeviceFilePath     string `json:inputdevicefilepath`
	InputDeviceLocationPath string `json:inputdevicelocationpath`
}

type n struct {
	f io.Reader
	s string
}

/*type v struct {
	data interface
	load func()
	info string
}*/

/*
if context, err := os.Open("test.txt"); err != nil {
	fmt.Println(err.Error())
	os.Exit(-1)
}
var deviceGenerator = v{context,}
func (deviceGenerator)saveCheckPoint() {
	locationInt := deviceGenerator.data
	locationStr := strconv.Itoa(locationInt)

	if f, err := os.Open(pathConfig.InputDeviceLocationPath); err != nil {
		fmt.Println(err.Error())
	}
}

func (deviceGenerator)loadCheckPoint() {
	if f, err := os.Open(pathConfig.InputDeviceLocationPath); err != nil {
		fmt.Println(err.Error())
	}
	locationStr=f.Read()
	locationInt=strconv.Atoi(locationStr)
	deviceGenerator.data.seek(locationInt)

}

func (deviceGenerator)moveNext() {

}
*/

var pathConfig u
var pos int = 0

func readLine(f io.Reader) string {
	rd := bufio.NewReader(f)
	fmt.Println(rd.Buffered())
	line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
	if err != nil || io.EOF == err {
		return ""
	}
	pos += len(line)
	fmt.Println(rd.Buffered(), len(line))
	return line
}

/*func readFile0(filename string) interface {
	if f, err := os.Open(filename); err != nil {
		fmt.Println(err.Error())
		return err.Error()
	} else {
		return f
	}
}*/
func saveCheckPoint() {
	location := []byte(strconv.Itoa(pos))
	err := ioutil.WriteFile("test.txt", location, 755)
	if err != nil {
		fmt.Printf(err.Error())
	}
	//defer f.Close()
	//f.WriteString(strconv.Itoa(location))
}

func main() {
	filename := "config.json"
	context, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	json.Unmarshal(context, &pathConfig)
	fmt.Println(pathConfig)
	f, _ := os.Open("ImeiJsonMi.txt")
	//f.Seek(584, 0)
	defer f.Close()
	fmt.Println(readLine(f))
	saveCheckPoint()
}
