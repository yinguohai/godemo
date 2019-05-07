package netstat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func connClient(c net.Conn) {
	defer c.Close()

	reader := bufio.NewReader(os.Stdin)

	buf := make([]byte,1024)

	for {
		//readerString 是reader句柄用来接收键盘的输入connent
		input , _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		if input == "quit" {
			return
		}

		//往socket通道中写数据
		c.Write([]byte(input))


		cnt ,err := c.Read(buf)

		if err != nil {
			fmt.Println("Fail to read data",err)
			continue
		}

		fmt.Println(string(buf[0:cnt]))
	}
}

func Client () {
	conn , err := net.Dial("tcp","localhost:1028")

	if err != nil {
		fmt.Println("Fail to connect",err)
		return
	}

	connClient(conn)
}