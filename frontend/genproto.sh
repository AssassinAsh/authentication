PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"
PROTOC_GEN_WEB_PATH="./node_modules/.bin/protoc-gen-grpc-web"
PROTOC_OUT_DIR="./proto/generated"
mkdir -p ${PROTOC_OUT_DIR}
protoc \
       --plugin="protoc-gen-grpc-web=${PROTOC_GEN_WEB_PATH}" \
       --grpc-web_out="import_style=typescript,mode=grpcwebtext:$PROTOC_OUT_DIR" \
       proto/auth-messages.proto
protoc \
       --plugin="protoc-gen-grpc-web=${PROTOC_GEN_WEB_PATH}" \
       --grpc-web_out="import_style=typescript,mode=grpcwebtext:$PROTOC_OUT_DIR" \
       proto/auth-services.proto