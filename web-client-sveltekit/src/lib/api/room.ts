import type { ListFilter, SuccessResponse } from "$lib/types/api";
import type { CreateRoomInput, Room, RoomBackground } from "$lib/types/room";

const BASE = import.meta.env.VITE_API_BASE;

export async function listRooms(filter: ListFilter, fetchFn = fetch): Promise<SuccessResponse<Room[]>> {
    const params = new URLSearchParams({                                                                                                              
        forward: String(filter.forward),                                                                                                              
        cursor: String(filter.cursor),                                                                                                                
        size: String(filter.size),                                                                                                                    
    });
    
    const res = await fetchFn(`${BASE}/rooms?${params}`);
    if (!res.ok) throw new Error(`listRooms ${res.status}`);
    return res.json();
}

export async function createRoom(input: CreateRoomInput, fetchFn = fetch): Promise<SuccessResponse<Room>> {
    const res = await fetchFn(`${BASE}/rooms`, {
        method: "POST",
        headers: {"content-type": "application/json"},
        body: JSON.stringify(input),
    });

    if (!res.ok) throw new Error(`createRoom ${res.status}`);
    return res.json();
}