package main

import (
	"musichain/pkg"
)

const privateKey = "9e2f3158a45ce798cb7342a3f6c8e5dafd926347c039b3596b4975aabb22347e"

func main() {
	pkg.Deploy(privateKey)
}
