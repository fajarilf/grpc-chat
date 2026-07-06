import { listRooms } from '$lib/api/room';
import type { ListFilter } from '$lib/types/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ fetch }) => {
    const filter: ListFilter = { forward: true, cursor: 0, size: 5 };
    const result = await listRooms(filter, fetch);
    return { 
        rooms: result.data,
        paging: result.paging
    } 
}