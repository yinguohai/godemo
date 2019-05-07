package netstat

import (
	"fmt"
	"net"
	"strings"
)

func connHandler(c net.Conn) {
	if c == nil {
		return
	}

	buf := make([]byte,4096)

	for {
		//读取数据
		cnt , err := c.Read(buf)

		if err != nil || cnt == 0 {
			c.Close()
			break
		}

		inStr := strings.TrimSpace(string(buf[0:cnt]))

		inputs := strings.Split(inStr," ")

		switch inputs[0] {
		case "ping":
			c.Write([]byte("pong\n"))
		case "echo":
			echoStr := strings.Join(inputs[1:]," ") + "\n"
			c.Write([]byte(echoStr))
		case "quit":
			c.Close()
			break
		default:
			fmt.Println("Unsupported command: \n",inputs[0])
		}

		fmt.Println("Connection form ",c.RemoteAddr(),"closed")
	}
}

func Server(){
	server , err := net.Listen("tcp","127.0.0.1:1208")

	if err != nil {
		fmt.Println("Fail to start server , ",err)
	}

	fmt.Println("Server Started ....")

	for {
		//这个地方会一直等待，客户端来连接，并返回连接句柄
		//连接局部可以用来读，也可以用来写，或者关闭

		conn , err := server.Accept()

		if err != nil {
			fmt.Println("Fail to connect",err)
		}
		go connHandler(conn)
	}


}