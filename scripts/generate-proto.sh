#!/bin/bash
set -euo pipefail

GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

# Run from the repo root regardless of where the script is called from.
cd "$(dirname "$0")/.."

PROTO_DIR="proto"
PROTO_FILES="type.proto user.proto room.proto"

echo -e "${BLUE}Generating gRPC code for all languages...${NC}\n"

# --- Go (shared) ---
echo -e "${GREEN}Generating Go code (shared)...${NC}"
mkdir -p "proto-go/proto"
protoc -I "${PROTO_DIR}" \
       --go_out=proto-go/proto --go_opt=paths=source_relative \
       --go-grpc_out=proto-go/proto --go-grpc_opt=paths=source_relative \
       ${PROTO_FILES}

# --- Node.js / TypeScript (api-gateway) via ts-proto ---
# On Windows, protoc.exe spawns the plugin via the Win32 loader, so it needs the
# .cmd wrapper with backslashes. We keep the path RELATIVE (cwd is the repo root)
# to avoid the absolute path's parentheses, which cmd.exe mis-parses.
TS_PROTO_PATH="api-gateway/node_modules/.bin/protoc-gen-ts_proto"
case "$(uname -s)" in
  MINGW* | MSYS* | CYGWIN*)
    TS_PROTO_PATH="${TS_PROTO_PATH//\//\\}.cmd"
    ;;
esac
OUT_PATH="api-gateway/src/proto"

echo -e "${GREEN}Generating TypeScript code (ts-proto)...${NC}"
mkdir -p "${OUT_PATH}"
protoc -I "${PROTO_DIR}" \
       --plugin="protoc-gen-ts_proto=${TS_PROTO_PATH}" \
       --ts_proto_out="${OUT_PATH}" \
       --ts_proto_opt=outputServices=grpc-js,esModuleInterop=true,useExactTypes=false \
       ${PROTO_FILES}

echo -e "${BLUE}All code generation complete!${NC}"