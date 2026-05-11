import { NextFunction, Request, Response } from "express";
import { broadcastToRoom, rooms, sendToUser } from "../../store/rooms";

export class RoomController {

    static async Get (req: Request, res: Response, next: NextFunction){
        try {
            const out = [...rooms.entries()].map(([roomId, members]) => ({
                roomId,
                members: [...members],
            }));
            res.json({ rooms: out });
        } catch (error) {
            next(error);
        }
    }

    static async Create (req: Request, res: Response, next: NextFunction) {
        try {
            const roomId = req.body?.roomId;
            if (!roomId || typeof roomId !== 'string') {
                return res.status(400).json({ error: 'roomId required' });
            }
            if (rooms.has(roomId)) {
                return res.status(409).json({ error: 'room already exists' });
            }
            rooms.set(roomId, new Set());
            res.status(201).json({ roomId });
        } catch (error) {
            next(error);
        }
    }

    static async Delete (req: Request, res: Response, next: NextFunction) {
        try {
            const roomId = String(req.params.roomId);
            const members = rooms.get(roomId);
            if (!members) return res.status(404).json({ error: 'room not found' });
            
            // tell each member they're leaving this room, then drop it
            for (const username of members) {
                sendToUser(username, { type: 'user_leaving', room: roomId, username });
            }
            rooms.delete(roomId);
            res.json({ ok: true });
        } catch (error) {
            next(error)
        }
    }

    static async Join (req: Request, res: Response, next: NextFunction) {
        try {
            const roomId = String(req.params.roomId);
            const username = req.body?.username;
            if (!username || typeof username !== 'string') {
                return res.status(400).json({ error: 'username required' });
            }
            const members = rooms.get(roomId);
            if (!members) return res.status(404).json({ error: 'room not found' });
            if (members.has(username)) {
                return res.status(409).json({ error: 'already joined' });
            }
            
            // notify existing members first, then add the new one
            broadcastToRoom(roomId, { type: 'user_joined', room: roomId, username });
            members.add(username);
            
            res.json({ roomId, members: [...members] });
        } catch (error) {
            next(error);
        }
    }

}