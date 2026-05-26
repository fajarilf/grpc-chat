import type { Room, RoomBackground } from "$lib/types/room";
import { currentUserId } from "./users";
import { presetColors, presetImages } from "./backgrounds";

function makeId() {
  return Math.random().toString(36).slice(2, 10);
}

function getRandomRooms(array: Room[], count: number = 2): Room[] {
  const shuffled = [...array].sort(() => Math.random() - 0.5);
  return shuffled.slice(0, count);
}

const seedRooms: Room[] = [
  {
    id: makeId(),
    number: 101,
    description: "General engineering chat",
    memberIds: [currentUserId, "u1", "u3", "u8"],
    background: { kind: "color", value: presetColors[1].value },
    createdAt: new Date(),
  },
  {
    id: makeId(),
    number: 202,
    description: "Production line monitoring",
    memberIds: [currentUserId, "u2", "u5"],
    background: { kind: "image", src: presetImages[1].src },
    createdAt: new Date(),
  },
  {
    id: makeId(),
    number: 303,
    description: "Off-topic lounge",
    memberIds: [currentUserId, "u4", "u7", "u1", "u3"],
    background: { kind: "color", value: presetColors[5].value },
    createdAt: new Date(),
  },
];

const currentRoom: Room[] = getRandomRooms(seedRooms);
const favRoom: Room[] = getRandomRooms(seedRooms, 1);

type roomExportType = {
  rooms: Room[],
  currentRoom: Room[], 
  favRoom: Room[]
}

export const roomStore = $state<roomExportType>({ rooms: seedRooms, currentRoom, favRoom });

export function createRoom(input: {
  number: number;
  description: string;
  memberIds: string[];
  background: RoomBackground;
}): Room {
  const room: Room = {
    id: makeId(),
    number: input.number,
    description: input.description,
    memberIds: Array.from(new Set([currentUserId, ...input.memberIds])),
    background: input.background,
    createdAt: new Date(),
  };
  roomStore.rooms = [...roomStore.rooms, room];
  return room;
}
