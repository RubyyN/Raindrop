/*
 *  Description: 	The main class of the program
 *	Author:			RubyyN
 *	Last change:	24-07-2017
 */

package main

import (
	"Helper/Webhelper"
	"Objects/Settings"
	"fmt"
	"Objects/Information"
	"time"
)

func main(){
	Settings.LoadSettings()
	Information.Initialize()
	fmt.Println("Listening on port 9822")
	Webhelper.CreateWebhelper()
	for 1 == 1 {
		time.Sleep(time.Second * 10)
	}
}


