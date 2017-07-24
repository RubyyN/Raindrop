/*
 *  Description: 	Contains methods to encode / decode the messages
 *	Author:			RubyyN
 *	Last change:	24-07-2017
 */


package Coder

import (
	"io"
	"compress/flate"
	"bytes"
	"encoding/base64"
)

func Decode(source string) string {
	base64_decoded, _ := base64.StdEncoding.DecodeString(source)
	decompressor := flate.NewReader(bytes.NewReader(base64_decoded))
	destination := new(bytes.Buffer)
	io.Copy(destination, decompressor)
	decompressor.Close()
	return destination.String()
}

func Encode(source string, level int) string{
	destination := new(bytes.Buffer)
	compressor, _ := flate.NewWriter(destination,level)
	compressor.Write([]byte(source))
	compressor.Close()
	return base64.StdEncoding.EncodeToString([]byte(destination.String()))
}