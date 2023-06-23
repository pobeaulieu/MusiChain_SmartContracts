
# Go service 
install go: https://go.dev/dl/

#### Dependencies
- `go get github.com/gofiber/fiber/v2`
- `go get -u gorm.io/gorm`
- `go get -u gorm.io/driver/sqlite`
- `go get github.com/gofiber/fiber/v2/middleware/cors`
- `go get github.com/mattn/go-isatty@v0.0.17` 
- `go get github.com/dgrijalva/jwt-go`
- `go get golang.org/x/crypto/bcrypt` 
- `go get github.com/joho/godotenv`

or 

`go mod tidy`

## React Webapp
See MusiChain_UI repo https://github.com/pobeaulieu/MusiChain_UI

# Database
The database is SQLite and is stored in the file system for ease of use (MusiChain.db). 

# API doc: Postman
import `MusiChain.postman_collection.json` in Postman and add `rootCA.crt` in the settings. 


# Contract Management
1. Le fichier .sol dans /contracts est le squelette du fichier qu'on voit déployer. Il faut modifier le chemin pour zeppelin pour le chemin absolu sur votre machine.
2. Ensuite run :
   solcjs  --optimize --bin contracts/musiChain.sol -o build
   solcjs  --optimize --abi contracts/musiChain.sol -o build
   abigen --abi=build/contracts_musiChain_sol_MusiChain.abi --bin=build/contracts_musiChain_sol_MusiChain.bin --pkg=api --out=api/musiChain.go
3. 3 fichiers devraient avoir été générés, 2 dans /build et 1 dans /api
4. Il faut par la suite aller dans contractManager.go et la fonction DeployNewContract()
5. Il faut lancer ganache
6. Puis remplacer la private key par une clé privée de Ganache
7. Puis run contrat_test.go

1. Mêmes étapes pour le contrat des transactions (pas encore testé)
solcjs  --optimize --abi contracts/tokenSale.sol -o build
solcjs  --optimize --bin contracts/tokenSale.sol -o build
abigen --abi=build/contracts_tokenSale_sol_TokenSale.abi --bin=build/contracts_tokenSale_sol_TokenSale.bin --pkg=tokenSale
--out=sale/tokenSale.go

