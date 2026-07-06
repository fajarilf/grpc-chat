<script lang="ts">
  import RoomCarousel from "$lib/components/rooms/room-carousel.svelte";
  import { listRooms } from "$lib/api/room";
  import type { ListFilter } from "$lib/types/api";
  import type { Room } from "$lib/types/room";

  let { data } = $props();
  // svelte-ignore state_referenced_locally
  let rooms = $state<Room[]>(data.rooms ?? []);
  // svelte-ignore state_referenced_locally
  let paging = $state(data.paging);

  type snippetParam = {
    title: string,
    desc: string,
    rooms: Room[],
    createRoom?: boolean,
    onLoadMore?: () => Promise<void>
  }

  async function loadMore() {
    if (!paging?.hasMore) return;
    const filter: ListFilter = { forward: true, cursor: paging.nextCursor, size: 5 };
    const result = await listRooms(filter);
    rooms = [...rooms, ...result.data];
    paging = result.paging;
  }
</script>

{#snippet roomGroup({title, desc, rooms, createRoom, onLoadMore}: snippetParam)}
<div class="flex flex-col gap-6">
  <div class="flex items-end justify-between">
    <div>
      <h1 class="text-lg font-semibold">{title}</h1>
      <p class="text-sm text-muted-foreground">{desc}</p>
    </div>
  </div>

  <RoomCarousel rooms={rooms} createRoom={createRoom} onLoadMore={onLoadMore}/>
</div>
{/snippet}

<div class="flex flex-col gap-14">
  
  {@render roomGroup({
    title: "Chat Room List",
    desc: "Pick a room to join or create a new one.",
    rooms: rooms,
    createRoom: true,
    onLoadMore: loadMore
  })}

  {@render roomGroup({
    title: "Favorite Rooms",
    desc: "List of room that you frequently visited",
    rooms: [],
  })}

  {@render roomGroup({
    title: "Currently Visited Room",
    desc: "List of room you have visited recently",
    rooms: [],
  })}

</div>
