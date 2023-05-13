# go-token-service


Curl İstekleri:

curl --location --request POST 'http://localhost:8080/balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "contractAddress": "0xced2C48a463F9607d2D91ba30891dc011Cda6b92",
    "walletAddress": "0x78147638E581acC89B0d7E311e480beC7a473044"
}'



curl --location --request POST 'http://localhost:8080/transfer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "privateKey": "f983127c3bcda9a6d4e5edd9c3614007983e8f9766db623d4afcaffd1e2c115b",
    "contractAddress":"0xced2C48a463F9607d2D91ba30891dc011Cda6b92",
    "toAddress":"0xd54804e0E63f3Dade1576087FFF6b4a16cDb81Ba",
    "amount":"5"
}'
