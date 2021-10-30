
<p align="center">
<img src="https://cdn.osxdaily.com/wp-content/uploads/2018/07/classic-mac-finder-icon.jpg" width=25% height=25%/>
</p>

## Table of contents
- [Whatmac](#whatmac)
- [Why](#why)
- [Example](#example)
- [License](#license)

# Whatmac
Library written in Golang to dump information about a macOS operating system at runtime, such as version is running on, codename of the operating system, etc.


### Why? 

- I initially made this library for fun.
- Allows getting macOS version in an easy way, without having to call 
  an external command.
  
### Getting started

- To get started download the module into your repository

`$ go get github.com/diegomagdaleno/whatmac`

- You are ready to Go!

### Example

```go
package main

import (
	"fmt"

	"github.com/diegomagdaleno/whatmac"
)

func main() {
	fmt.Println(whatmac.GetProductVersionAsFloat())
}
```

This in my case returns `10.16` that can be used for comparing.

### License 

MIT
