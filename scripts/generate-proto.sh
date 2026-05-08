#!/bin/bash

GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}Generating gRPC code for all language...${NC}\n"

#generate to go
echo -e "${GREEN}Generating Go code...${NC}"
protoc --go_out=server-go --go_opt=paths=source_relative \
       --go-grpc_out=server-go --go-grpc_opt=paths=sourc_relative \
       proto/chat.proto

#copy to client-go
cp -r server-go/proto client-go/

#generate to Node.js
echo -e"${GREEN} Generating Node.js code...${NC}"
grpc_tools_node_protoc --js_out=import_style=commonjs,binary:client-node/proto \
                       --grpc_out=grpc_js:client-node/proto \
                       proto/chat.proto

echo -e "${BLUE}All code generation complete!${NC}"