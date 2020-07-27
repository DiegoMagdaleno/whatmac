package whatmac

import (
	"bytes"
	"io/ioutil"

	"howett.net/plist"
)

var macOSInfo macOSVersion

func init() {
	initialFile, _ := ioutil.ReadFile("/System/Library/CoreServices/SystemVersion.plist")
	fileToReader := bytes.NewReader(initialFile)
	decoder := plist.NewDecoder(fileToReader)
	decoder.Decode(&macOSInfo)
}
