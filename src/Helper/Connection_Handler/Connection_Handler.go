/*
 *  Description: 	Class to handle connection between the server and the client
 *	Author:			RubyyN
 *	Last change:	24-07-2017
 */

package Connection_Handler

import (
	"net"
	"fmt"
	"os"
)

const (
	WEBHELPER_HOST = "localhost"
	WEBHELPER_PORT = "9822"
	WEBHELPER_TYPE = "tcp"
)

func CreateWebhelper() {
	l, err := net.Listen(WEBHELPER_TYPE,WEBHELPER_HOST+":"+WEBHELPER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(-1)
	}

	defer l.Close()
	fmt.Println("Listening on " + WEBHELPER_HOST + ":" + WEBHELPER_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue;
		}

		go handleWebhelperRequest(conn)
	}
}

func handleWebhelperRequest(conn net.Conn) {

}

