/*
 *  Description: 	Contains the Settings and communicates with the database
 *	Author:			RubyyN
 *	Last change:	24-07-2017
 */

package Settings

var InterfaceID string
var Enable_Upload bool
var Max_Uploads_To_Server int
var Max_Downloads_From_Server int
var Upload_Speed uint16
var Download_Speed uint16

const (
	CURRENT_VERSION = 001
)

func LoadSettings() {
	Max_Downloads_From_Server = 6
	// sql database connection
}

func SaveSettings() {
	// sql database connection
}

