import * as grpc from "@grpc/grpc-js"
import { 
    UserServiceClient,
    UserLoginRequest,
    UserLoginResponse,
    UserCreateRequest,
    UserResponse
} from "../proto/user"

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