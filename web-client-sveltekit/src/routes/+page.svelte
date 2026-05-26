<script lang="ts">
  import RoomCarousel from "$lib/components/rooms/room-carousel.svelte";
  import { roomStore } from "$lib/data/rooms.svelte";
  import type { Room } from "$lib/types/room";

  type snippetParam = {
    title: string,
    desc: string,
    rooms: Room[],
    createRoom?: boolean
  }
</script>

{#snippet roomGroup({title, desc, rooms, createRoom}: snippetParam)}
<div class="flex flex-col gap-6">
  <div class="flex items-end justify-between">
    <div>
      <h1 class="text-lg font-semibold">{title}</h1>
      <p class="text-sm text-muted-foreground">{desc}</p>
    </div>
  </div>

  <RoomCarousel rooms={rooms} createRoom={createRoom}/>
</div>
{/snippet}

<div class="flex flex-col gap-14">
  
  {@render roomGroup({
    title: "Chat Room List",
    desc: "Pick a room to join or create a new one.",
    rooms: roomStore.rooms,
    createRoom: true
  })}

  {@render roomGroup({
    title: "Favorite Rooms",
    desc: "List of room that you frequently visited",
    rooms: roomStore.favRoom,
  })}

  {@render roomGroup({
    title: "Currently Visited Room",
    desc: "List of room you have visited recently",
    rooms: roomStore.currentRoom,
  })}

</div>
