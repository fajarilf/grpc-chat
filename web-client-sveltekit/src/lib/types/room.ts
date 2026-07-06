import type { User } from "./user";

export enum BackgroundType {
  COLOR = "COLOR",
  IMAGE = "IMAGE"
}

export const backgroundTypeToNumber: Record<BackgroundType, number> = {
  [BackgroundType.COLOR]: 0,
  [BackgroundType.IMAGE]: 1,
};

export type RoomBackground = {
  kind: BackgroundType
  value: string;
}

export type Room = {
  id: string;
  name: number;
  description: string;
  users: User[];
  background: string;
  backgroundType: string
  createdAt: Date;
};

export type CreateRoomInput = {
  name: string,
  description: string,
  backgroundType: BackgroundType
  background: string
  user_ids: number[],
}
