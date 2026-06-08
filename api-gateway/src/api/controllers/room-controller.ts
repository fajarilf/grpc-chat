import { NextFunction, Request, Response } from "express";
import * as roomClient from "../../grpc-client/room"
import { RoomCreateRequest, RoomListRequest } from "../../proto/room";

export class RoomController {

    static async Get (req: Request, res: Response, next: NextFunction) {
        try {
            const query = req.query;
            const request: RoomListRequest = {
                forward: query.forward !== "false",
                cursor: parseInt(String(query.cursor)) || 1,
                size: parseInt(String(query.size)) || 10
            }
            const response = await roomClient.getListRoom(request);
            res.status(200).json({
                data: response
            });
        } catch (error) {
            next(error);
        }
    }

    static async Create (req: Request, res: Response, next: NextFunction) {
        try {
            const request: RoomCreateRequest = req.body as RoomCreateRequest;
            const response = await roomClient.createRoom(request);
            res.status(200).json({
                data: response
            });
        } catch (error) {
            next(error);
        }
    }

}