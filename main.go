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

func main() {

	connectCmd := Command{Token: 1, MsgId: 1, ParamSize: 1}

	fmt.Println("Sending command: ", connectCmd)
	conn, err := net.Dial("tcp", "localhost:7878")
	if err != nil {
		fmt.Println("Error connecting to camera: ", err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(conn)
	SendCommand(conn, reader, connectCmd)

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
