#Cryptovote-service

##Server gRPC Link
Server: 64.225.61.39:3007

##Proto Files:
/pkg/pb/voteservice

##Calls with BloomRPC

<img src="https://github.com/rogerio410/cryptovote-service/blob/main/docs/allcrypto.png?raw=true"/>

<img src="https://github.com/rogerio410/cryptovote-service/blob/main/docs/crypto.png?raw=true"/>

<img src="https://github.com/rogerio410/cryptovote-service/blob/main/docs/vote.png?raw=true"/>

<img src="https://github.com/rogerio410/cryptovote-service/blob/main/docs/remove.png?raw=true"/>

##Golang gRPC API

The Technical Challenge consists of creating an API with Golang using gRPC with stream pipes that exposes an Upvote service endpoints. The API will provide the user an interface to upvote or downvote a known list of the main Cryptocurrencies (Bitcoin, ethereum, litecoin, etc..).

##Technical requirements:

Keep the code in Github
The API must have a read, insert, delete and update interfaces.
The API must guarantee the typing of user inputs. If an input is expected as a string, it can only be received as a string.
The API must contain unit test of methods it uses
You can choose the database but the structs used with it should support Marshal/Unmarshal with bson, json and struct

##Extra:

Deliver the whole solution running in some free cloud service
The API have a method that stream a live update of the current sum of the votes from a given Cryptocurrency
