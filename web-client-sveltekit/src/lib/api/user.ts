import type { SuccessResponse } from "$lib/types/api";
import type { User } from "$lib/types/user";

const BASE = import.meta.env.VITE_API_BASE; 

export async function listUsers (fetchFn = fetch): Promise<SuccessResponse<User[]>> {
    const res = await fetchFn(`${BASE}/users`);
    if (!res.ok) throw new Error(`listUsers ${res.status}`);
    return res.json();
}