export type RoomBackground =
  | { kind: "color"; value: string }
  | { kind: "image"; src: string };

export type User = {
  id: string;
  name: string;
};

export type Room = {
  id: string;
  number: number;
  description: string;
  memberIds: string[];
  background: RoomBackground;
  createdAt: Date;
};
