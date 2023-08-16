package main

import (
	"musichain/pkg"
)

const privateKey = "b1afd54b5ddf0192ac029641caef5f92bdcb0c04a5e38b422382772d5c91320d"

func main() {
	pkg.Deploy(privateKey)
}
