<script lang="ts">
  import { Users } from "@lucide/svelte";
  import { ImageCard } from "$lib/components/ui/image-card";
  import type { Room } from "$lib/types/room";

  let { room, href }: { room: Room; href?: string } = $props();
</script>

<ImageCard
  href={href ?? `/rooms/${room.id}`}
  class="aspect-video w-72 shrink-0"
>
  {#snippet bg()}
    {#if room.background.kind === "color"}
      <div class="absolute inset-0" style="background-color: {room.background.value};"></div>
    {:else}
      <img
        src={room.background.src}
        alt=""
        loading="lazy"
        class="absolute inset-0 h-full w-full object-cover"
      />
    {/if}
  {/snippet}

  <div class="flex items-end justify-between gap-2">
    <h3 class="text-base font-semibold leading-tight">Room {room.number}</h3>
    <div class="flex items-center gap-1 text-xs opacity-90">
      <Users class="size-3.5" />
      <span>{room.memberIds.length}</span>
    </div>
  </div>
  {#if room.description}
    <p class="mt-1 line-clamp-1 text-xs opacity-80">{room.description}</p>
  {/if}
</ImageCard>
