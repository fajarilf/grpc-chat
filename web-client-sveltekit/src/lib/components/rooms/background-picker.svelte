<script lang="ts">
  import { Check } from "@lucide/svelte";
  import { cn } from "$lib/utils";
  import { presetColors, presetImages } from "$lib/data/backgrounds";
  import type { RoomBackground } from "$lib/types/room";

  let {
    value = $bindable(),
  }: { value: RoomBackground } = $props();

  function selectColor(c: string) {
    value = { kind: "color", value: c };
  }

  function selectImage(src: string) {
    value = { kind: "image", src };
  }

  function isSelectedColor(c: string) {
    return value.kind === "color" && value.value === c;
  }

  function isSelectedImage(src: string) {
    return value.kind === "image" && value.src === src;
  }
</script>

<div class="space-y-3">
  <div>
    <p class="mb-2 text-xs font-medium text-muted-foreground">Color</p>
    <div class="flex flex-wrap gap-2">
      {#each presetColors as c (c.value)}
        <button
          type="button"
          aria-label={c.label}
          title={c.label}
          onclick={() => selectColor(c.value)}
          class={cn(
            "relative h-8 w-8 rounded-md border border-border transition hover:scale-105",
            isSelectedColor(c.value) && "ring-2 ring-ring ring-offset-2 ring-offset-background"
          )}
          style="background-color: {c.value};"
        >
          {#if isSelectedColor(c.value)}
            <Check class="absolute inset-0 m-auto size-4 text-white drop-shadow" />
          {/if}
        </button>
      {/each}
    </div>
  </div>

  <div>
    <p class="mb-2 text-xs font-medium text-muted-foreground">Image</p>
    <div class="grid grid-cols-3 gap-2">
      {#each presetImages as img (img.src)}
        <button
          type="button"
          aria-label={img.label}
          title={img.label}
          onclick={() => selectImage(img.src)}
          class={cn(
            "relative aspect-video overflow-hidden rounded-md border border-border transition hover:scale-[1.02]",
            isSelectedImage(img.src) && "ring-2 ring-ring ring-offset-2 ring-offset-background"
          )}
        >
          <img src={img.src} alt="" loading="lazy" class="h-full w-full object-cover" />
          {#if isSelectedImage(img.src)}
            <span class="absolute inset-0 flex items-center justify-center bg-black/40">
              <Check class="size-5 text-white" />
            </span>
          {/if}
        </button>
      {/each}
    </div>
  </div>
</div>
