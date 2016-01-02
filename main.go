package main

import "fmt"
import "net"
import "os"
import "bufio"
import "encoding/json"

type Command struct {
	Token     int `json:"token"`
	MsgId     int `json:"msg_id"`
	ParamSize int `json:"param_size"`
}

var CONNECT = Command{Token: 1, MsgId: 1, ParamSize: 1}
var UNKNOWN = Command{Token: 1, MsgId: 2, ParamSize: 0}
var CAPTURE = Command{Token: 1, MsgId: 5, ParamSize: 0}
var START_RECORDING = Command{Token: 1, MsgId: 3, ParamSize: 0}
var STOP_RECORDING = Command{Token: 1, MsgId: 4, ParamSize: 0}

func main() {

	fmt.Println("Sending command: ", CONNECT)
	conn, err := net.Dial("tcp", "localhost:7878")
	if err != nil {
		fmt.Println("Error connecting to camera: ", err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(conn)
	SendCommand(conn, reader, CONNECT)

	fmt.Println("Exiting...")
}

func SendCommand(connection net.Conn, reader *bufio.Reader, command Command) {

	jsonBlob, err := json.Marshal(command)
	if err != nil {
		fmt.Printf("Error marshalling json for command %v: %v\n", command, err)
		return // TODO return err
	}

	_, errSend := connection.Write(jsonBlob)
	if err != nil {
		fmt.Printf("Error sending command: %v\n", errSend)
	}

	response, errRead := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading response: %v\n", errRead)
	}
	fmt.Printf("Response: %v\n", response)

}
