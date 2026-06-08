import express from "express";
import { RoomController } from "../controllers/room-controller";

export const roomRoutes = express.Router();

roomRoutes.get("/", RoomController.Get);
roomRoutes.post("/", RoomController.Create);
