# gRPC-Mit-GO

## Project Overview

A microservice project that uses gRPC API to request and get response from the server to the client. gRPC is a remote procedure call (RPC) framework from Google. It uses Protocol Buffers as a serialization format and uses HTTP2 as the transport medium. In gRPC, a client application can directly call methods on a server application on a different machine as if it was a local object; this process makes it easier to create distributed applications and services.

### Project Tasks

#####   The project is Language-independent gRPC-based microservice:
###### Client Goland
###### Sever- Goland
* Create a gRPC server (in your preferred language) that serves the time-based electricity consumption data found in the attached file: meterusage.csv.
* Use whatever tool you prefer to create an http server that will request the data from the gRPC server and deliver the consumption data as JSON. 
*  Create a single page html document that requests the JSON from the http server and displays it. Please supply a link to your work in a public repository and a short README with a description of what you did."

## Implementation code directory

change the code struct：
![image](https://user-images.githubusercontent.com/50584494/86209089-13bc6100-bb72-11ea-8cb0-41f4416a5e07.png)


####  Run Server.
- change the csv file to cli paramters. so please run command ` go run . -f YOURCSVFILE.csv` to start the server

### Run Client:
- run command ' go run . ' under client folder.

about other GO language details, ,you can check the GO official website golang.org
