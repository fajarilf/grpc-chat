import * as grpc from "@grpc/grpc-js"
import { 
    UserServiceClient,
    UserLoginRequest,
    UserLoginResponse,
    UserCreateRequest,
    UserListResponse,
} from "../proto/user"
import { UserResponse } from "../proto/type"
import { Empty } from "../proto/google/protobuf/empty"

const userClient = new UserServiceClient(
    process.env.GRPC_SERVER_ADDR ?? "localhost:50051",
    grpc.credentials.createInsecure(),
)

export function register(param: UserCreateRequest): Promise<UserResponse> {
    return new Promise((resolve, reject) => {
        userClient.register(param, (err, res) => (err ? reject(err) : resolve(res)))
    })
}

export function login(param: UserLoginRequest): Promise<UserLoginResponse> {
    return new Promise((resolve, reject) => {
        userClient.login(param, (err, res) => (err ? reject(err) : resolve(res)))
    })
}

export function getList(empty: Empty): Promise<UserListResponse> {
    return new Promise((resolve, reject) => {
        userClient.getList(empty, (err, res) => (err ? reject(err) : resolve(res)))
    })
}