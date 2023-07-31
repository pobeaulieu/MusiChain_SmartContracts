package main

import (
	"musichain/pkg"
)

const privateKey = "bab108fa9ea5eb0a23204e2827a8ca011b6c4ae864e1621e0e634fe2a13ba640"

func main() {
	pkg.Deploy(privateKey)
}
