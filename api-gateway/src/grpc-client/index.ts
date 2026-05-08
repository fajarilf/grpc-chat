import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import path from "path";

const PROTO_PATH = path.join(__dirname, "../../proto/chat.proto");

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});

const chatProto: any = grpc.loadPackageDefinition(packageDefinition).chat;

export const chatClient = new chatProto.ChatService(
    'localhost:50051',
    grpc.credentials.createInsecure()
)