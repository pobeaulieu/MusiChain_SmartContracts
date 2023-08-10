package main

import (
	"musichain/pkg"
)

const privateKey = "44dfab8d648dc9bb6099097366821ebdc1b9364754ab313098aeb666257e901c"

func main() {
	pkg.Deploy(privateKey)
}
