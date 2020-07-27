package whatmac

func GetProductBuildVersion() string {
	return macOSInfo.ProductBuildVersion
}

func GetProductCopyright() string {
	return macOSInfo.ProductCopyright
}

func GetProductName() string {
	return macOSInfo.ProductName
}

func GetProductUserVisibleVersion() string {
	return macOSInfo.ProductUserVisibleVersion
}

func GetProductVersion() string {
	return macOSInfo.ProductVersion
}

func GetiOSSupportVersion() string {
	return macOSInfo.IOSSuportVersion
}

func GetProductUserVisibleVersionAsFloat() float64 {
	return toFloat(macOSInfo.ProductUserVisibleVersion)
}

func GetiOSSupportVersionAsFloat() float64 {
	return toFloat(macOSInfo.IOSSuportVersion)
}

func GetProductVersionAsFloat() float64 {
	return toFloat(macOSInfo.ProductVersion)
}
