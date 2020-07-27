package WhatMac

import (
	"bytes"
	"io/ioutil"

	"howett.net/plist"
)

var MacOSInfo macOSVersion

func Decode() error {
	initialFile, err := ioutil.ReadFile("/System/Library/CoreServices/SystemVersion.plist")
	fileToReader := bytes.NewReader(initialFile)
	if err != nil {
		return err
	}
	decoder := plist.NewDecoder(fileToReader)
	err = decoder.Decode(&MacOSInfo)
	if err != nil {
		return err
	}
	return nil
}
