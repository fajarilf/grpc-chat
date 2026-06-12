import express from "express"
import { UserController } from "../controllers/user-controller"

export const userRoutes = express.Router();

userRoutes.get("/", UserController.Get);