1. Le fichier .sol dans /contracts est le squelette du fichier qu'on voit déployer. Il faut modifier le chemin pour zeppelin pour le chemin absolu sur votre machine.
2. Ensuite run : 
solcjs  --optimize --bin contracts/musiChain.sol -o build
solcjs  --optimize --abi contracts/musiChain.sol -o build
abigen --abi=build/contracts_musiChain_sol_MusiChain.abi --bin=build/contracts_musiChain_sol_MusiChain.bin --pkg=api --out=api/musiChain.go
3. 3 fichiers devraient avoir été générés, 2 dans /build et 1 dans /api
4. Il faut par la suite aller dans playground.go et la fonction DeployNewContract()
5. Il faut lancer ganache
6. Puis remplacer la private key par une clé privée de Ganache
7. Puis run web3_test.go
