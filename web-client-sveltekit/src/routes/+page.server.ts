import { listRooms } from '$lib/api/room';
import { listUsers } from '$lib/api/user';
import type { ListFilter } from '$lib/types/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => {
    const filter: ListFilter = { forward: true, cursor: 0, size: 5 };
    const room = await listRooms(filter, fetch);
    const user = await listUsers(fetch);
    return { 
        users: user.data,
        rooms: room.data,
        paging: room.paging
    } 
}