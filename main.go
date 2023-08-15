package main

import (
	"musichain/pkg"
)

// We hardcode the privatekey here that we use to deploy our contracts
const privateKey = "44dfab8d648dc9bb6099097366821ebdc1b9364754ab313098aeb666257e901c"

func main() {
	pkg.Deploy(privateKey)
}
