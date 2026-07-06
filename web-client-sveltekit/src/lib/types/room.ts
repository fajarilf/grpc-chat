import type { User } from "./user";

export type RoomBackground =
  | { kind: "color"; value: string }
  | { kind: "image"; src: string };

export type Room = {
  id: string;
  name: number;
  description: string;
  users: User[];
  background: string;
  backgroundType: "COLOR" | "IMAGE"
  createdAt: Date;
};

export type CreateRoomInput = {
    number: number,
    description: string,
    memberIds: string[],
    background: RoomBackground
}
