import * as grpc from "@grpc/grpc-js";
import { 
    RoomServiceClient, 
    RoomListRequest, 
    RoomListResponse,
    RoomCreateRequest,
    RoomResponseWithUser,
    RoomId
} from "../proto/room";

const roomClient = new RoomServiceClient(
    process.env.GRPC_SERVER_ADDR ?? "localhost:50051",
    grpc.credentials.createInsecure(),
)

export function getListRoom(param: RoomListRequest): Promise<RoomListResponse> {
    return new Promise((resolve, reject) => {
        roomClient.getListRoom(param, (err, res) => (err ? reject(err) : resolve(res)))
    });
}

export function createRoom(param: RoomCreateRequest): Promise<RoomResponseWithUser> {
    return new Promise((resolve, reject) => {
        roomClient.createRoom(param, (err, res) => (err ? reject(err) : resolve(res)))
    })
}

export function getRoomById(param: RoomId): Promise<RoomResponseWithUser> {
    return new Promise((resolve, reject) => {
        roomClient.getRoomById(param, (err, res) => (err ? reject(err) : resolve(res)))
    })
}