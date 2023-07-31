package pkg

import (
	"log"
	"os"
	"testing"
)

const privateKey = "bab108fa9ea5eb0a23204e2827a8ca011b6c4ae864e1621e0e634fe2a13ba640"

func TestDeploy(t *testing.T) {
	metaAddress := deployNewMetaDataContract(privateKey)
	baseAddress := deployNewBaseContract(privateKey, metaAddress)
	saleAddress := deployNewSaleContract(privateKey, baseAddress)

	f, err := os.Create(".env")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString("REACT_APP_METADATA_ADDRESS = " + metaAddress + "\n")
	_, err = f.WriteString("REACT_APP_BASE_ADDRESS = " + baseAddress + "\n")
	_, err = f.WriteString("REACT_APP_SALE_ADDRESS = " + saleAddress + "\n")

	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	f.Sync()
}
