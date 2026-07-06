import { NextFunction, Request, Response } from "express";
import * as roomClient from "../../grpc-client/room"
import { ApiResponse } from "../../types/api-response"
import { RoomCreateRequest, RoomListRequest, RoomResponseWithUser } from "../../proto/room";

export class RoomController {

    static async Get (
        req: Request, 
        res: Response<ApiResponse<RoomResponseWithUser[]>>, 
        next: NextFunction
    ) {
        try {
            const query = req.query;
            const request: RoomListRequest = {
                forward: query.forward !== "false",
                cursor: parseInt(String(query.cursor)) || 0,
                size: parseInt(String(query.size)) || 5
            }
            const response = await roomClient.getListRoom(request);
            res.status(200).json({
                status: "success",
                data: response.rooms,
                paging: {
                    nextCursor: response.nextCursor,
                    prevCursor: response.prevCursor,
                    hasMore: response.hasMore
                }
            });
        } catch (error) {
            next(error);
        }
    }

    static async Create (
        req: Request, 
        res: Response<ApiResponse<RoomResponseWithUser>>, 
        next: NextFunction
    ) {
        try {
            const request = RoomCreateRequest.fromJSON(req.body);
            const response = await roomClient.createRoom(request);
            res.status(200).json({
                status: "success",
                data: response,
            });
        } catch (error) {
            next(error);
        }
    }

    static async GetById (
        req: Request,
        res: Response<ApiResponse<RoomResponseWithUser>>,
        next: NextFunction
    ) {
        try {
            const id: number = parseInt(String(req.params.roomId))
            const response = await roomClient.getRoomById({id})
            res.status(200).json({
                status: "success",
                data: response,
            });
        } catch (error) {
            next(error)
        }
    }

}