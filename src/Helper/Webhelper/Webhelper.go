/*
 *  Description: 	Class to handle connection over web socket.
 *	Author:			RubyyN
 *	Last change:	24-07-2017
 */

package Webhelper

import (
	"fmt"
	"net/http"
	"strings"
	"Helper/Coder"
	"Helper/Server_Connector"
	"Helper/Json"
	"io/ioutil"
)

func CreateWebhelper() {
	http.HandleFunc("/",handleWebhelperRequest)
	http.HandleFunc("/Search/",handleSearchRequest)
	http.ListenAndServe(":9822", nil)
}

func handleWebhelperRequest(w http.ResponseWriter, r *http.Request) {
	dat,err := http.Get("https://github.com/RubyyN/Raindrop/src/HTML/interface.html")
	if err != nil {
		panic(err)
	}
	defer dat.Body.Close()
	body, err := ioutil.ReadAll(dat.Body)
	fmt.Fprintf(w,string(body))
}

func handleSearchRequest(w http.ResponseWriter, r *http.Request) {
	var searchResult = strings.Split(r.URL.Path[1:],"/")
	response := Json.ConvertToJson(Server_Connector.Search(searchResult[1]))
	fmt.Fprintf(w,Coder.Encode(response,9))
}



