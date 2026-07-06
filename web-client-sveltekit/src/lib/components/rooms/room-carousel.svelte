<script lang="ts">
  import { ChevronLeft, ChevronRight } from "@lucide/svelte";
  import Button from "$lib/components/ui/button/button.svelte";
  import RoomCard from "./room-card.svelte";
  import CreateRoomDialog from "./create-room-dialog.svelte";
  import type { Room } from "$lib/types/room";

  let scroller = $state<HTMLDivElement | null>(null);

  let { rooms, createRoom = false, onLoadMore }: { 
    rooms: Room[], 
    createRoom?: boolean,
    onLoadMore?: () => Promise<void>
  } = $props();

  function scrollBy(direction: 1 | -1) {
    if (!scroller) return;
    const amount = scroller.clientWidth * 0.8;
    scroller.scrollBy({ left: direction * amount, behavior: "smooth" });

    if (direction === 1) {
      const { scrollLeft, scrollWidth, clientWidth } = scroller;
      if (scrollLeft + clientWidth >= scrollWidth - clientWidth * 0.2) {
        onLoadMore?.();
      }
    }
  }

</script>

<div class="relative">
  <div class="absolute -top-12 right-0 hidden gap-1 sm:flex">
    <Button variant="outline" size="icon" class="size-8" onclick={() => scrollBy(-1)} aria-label="Scroll left">
      <ChevronLeft class="size-4" />
    </Button>
    <Button variant="outline" size="icon" class="size-8" onclick={() => scrollBy(1)} aria-label="Scroll right">
      <ChevronRight class="size-4" />
    </Button>
  </div>

  <div
    bind:this={scroller}
    class="flex snap-x snap-mandatory gap-4 overflow-x-auto scroll-smooth pb-2 scrollbar-thin"
  >
    {#if createRoom}
      <div class="snap-start">
        <CreateRoomDialog />
      </div>
    {/if}
    {#each rooms as room (room.id)}
      <div class="snap-start">
        <RoomCard {room} />
      </div>
    {/each}
  </div>
</div>
