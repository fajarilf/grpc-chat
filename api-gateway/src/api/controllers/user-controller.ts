import { NextFunction, Request, Response } from "express";
import * as userClient from "../../grpc-client/user";
import { ApiResponse } from "../../types/api-response";
import { UserResponse } from "../../proto/type";
import { Empty } from "../../proto/google/protobuf/empty";

export class UserController {
    static async Get(
        req: Request,
        res: Response<ApiResponse<UserResponse[]>>,
        next: NextFunction
    ) {
        try {
            const result = await userClient.getList(Empty.create());
            return res.status(200).json({
                status: "success",
                data: result.users
            })
        } catch (error) {
            next(error)
        }
    }
}