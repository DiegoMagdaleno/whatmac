package whatmac

/*GetProductBuildVersion gets the build version
of the current running macOS instance. Returns a string*/
func GetProductBuildVersion() string {
	return macOSInfo.ProductBuildVersion
}

/*GetProductCopyright gets the copyright that was acssigned
to the current running macOS instance. Returns a string*/
func GetProductCopyright() string {
	return macOSInfo.ProductCopyright
}

/*GetProductName gets the current Product name
of the current running macOS instance. It should always return macOS,
unless running on another Darwin-like operating system. Returns a string*/
func GetProductName() string {
	return macOSInfo.ProductName
}

/*GetProductUserVisibleVersion gets the user visible version
of the current running macOS instance. This is differnt from the real
product version that the macOS instance might be running. So it is suggested to
only use this if the program displays information to the user, if you want to compare versions
this isnt recommended, it is recommended to check GetProductVersionAsFloat. Returns a string*/
func GetProductUserVisibleVersion() string {
	return macOSInfo.ProductUserVisibleVersion
}

/*GetProductVersion gets the real product version of the current
macOS instance, emphasis in real macOS version, however, this function is still
not recommened if you want to do a comparison, if you wish to do a comparison, it
is recommeded to use GetProductVersionAsFloat since strings are not allowed to perform
comparisons in Golang. Returns a string*/
func GetProductVersion() string {
	return macOSInfo.ProductVersion
}

/*GetiOSSupportVersion gets the current iOSSuppoertVersion of the macOS instance, this is also called
Catalyst or Marzipan, as a fair warning, versions before macOS Mojave (10.14) do not have
an iOSSupportVersion so this string might be empty, is recommended to handle this in your
program. Returns a string*/
func GetiOSSupportVersion() string {
	return macOSInfo.IOSSuportVersion
}

/*GetProductUserVisibleVersionAsFloat allows you to get the version that user sees in multiple menus
as a float64 variable, however, it is not recommended to use this for precise checks, as the user visible
version might defer from the real version the macOS instance is running. Returns a float64*/
func GetProductUserVisibleVersionAsFloat() float64 {
	return toFloat(macOSInfo.ProductUserVisibleVersion)
}

/*GetiOSSupportVersionAsFloat returns a float64variable of the current iOS support in the running macOS instace
allowing you to perform comparison operations. Returns a float64*/
func GetiOSSupportVersionAsFloat() float64 {
	return toFloat(macOSInfo.IOSSuportVersion)
}

/*GetProductVersionAsFloat returns a float64 variable of the real macOS version, this is stronly recommended
to use if your program is performining comparions, however, it is not suggested to execute functions like "to round"
to the returned value, as this could make the program get a wrong macOS version*/
func GetProductVersionAsFloat() float64 {
	return toFloat(macOSInfo.ProductVersion)
}
