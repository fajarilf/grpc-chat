import type { Room, RoomBackground } from "$lib/types/room";

const BASE = import.meta.env.VITE_API_BASE;

export type CreateRoomInput = {
    number: number,
    description: string,
    memberIds: string[],
    background: RoomBackground
}

export async function listRooms(fetchFn = fetch): Promise<Room[]> {
    const res = await fetchFn(`${BASE}/rooms`);
    if (!res.ok) throw new Error(`listRooms ${res.status}`);
    return res.json();
}

export async function createRoom(input: CreateRoomInput, fetchFn = fetch): Promise<Room> {
    const res = await fetchFn(`{BASE}/rooms`, {
        method: "POST",
        headers: {"content-type": "application/json"},
        body: JSON.stringify(input),
    });

    if (!res.ok) throw new Error(`listRooms ${res.status}`);
    return res.json();
}