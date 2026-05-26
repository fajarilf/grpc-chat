<script lang="ts">
  import { Plus } from "@lucide/svelte";
  import * as Dialog from "$lib/components/ui/dialog";
  import Button from "$lib/components/ui/button/button.svelte";
  import Input from "$lib/components/ui/input/input.svelte";
  import Label from "$lib/components/ui/label/label.svelte";
  import BackgroundPicker from "./background-picker.svelte";
  import { mockUsers, currentUserId } from "$lib/data/users";
  import { defaultBackground } from "$lib/data/backgrounds";
  import { createRoom } from "$lib/data/rooms.svelte";
  import type { RoomBackground } from "$lib/types/room";

  let open = $state(false);
  let number = $state<number | null>(null);
  let description = $state("");
  let selectedMemberIds = $state<string[]>([]);
  let background = $state<RoomBackground>(defaultBackground);

  const availableUsers = mockUsers.filter((u) => u.id !== currentUserId);
  const canSubmit = $derived(number !== null && number > 0 && selectedMemberIds.length > 0);

  function reset() {
    number = null;
    description = "";
    selectedMemberIds = [];
    background = defaultBackground;
  }

  function toggleMember(id: string) {
    selectedMemberIds = selectedMemberIds.includes(id)
      ? selectedMemberIds.filter((m) => m !== id)
      : [...selectedMemberIds, id];
  }

  function handleSubmit(event: SubmitEvent) {
    event.preventDefault();
    if (!canSubmit || number === null) return;
    createRoom({
      number,
      description: description.trim(),
      memberIds: selectedMemberIds,
      background,
    });
    open = false;
    reset();
  }
</script>

<Dialog.Root bind:open>
  <Dialog.Trigger>
    {#snippet child({ props })}
      <button
        {...props}
        type="button"
        class="flex aspect-video w-72 shrink-0 flex-col items-center justify-center gap-2 rounded-lg border-2 border-dashed border-border bg-muted/30 text-muted-foreground transition hover:border-foreground/40 hover:bg-muted/50 hover:text-foreground"
      >
        <Plus class="size-8" />
        <span class="text-sm font-medium">Create new room</span>
      </button>
    {/snippet}
  </Dialog.Trigger>

  <Dialog.Content class="sm:max-w-lg">
    <Dialog.Header>
      <Dialog.Title>Create new room</Dialog.Title>
      <Dialog.Description>
        Set a room number, a short description, and pick who can join.
      </Dialog.Description>
    </Dialog.Header>

    <form onsubmit={handleSubmit} class="space-y-4">
      <div class="space-y-1.5">
        <Label for="room-number">Room number</Label>
        <Input
          id="room-number"
          type="number"
          min="1"
          required
          bind:value={number}
          placeholder="e.g. 101"
        />
      </div>

      <div class="space-y-1.5">
        <Label for="room-description">Description <span class="text-gray-400">(optional)</span></Label>
        <Input
          id="room-description"
          bind:value={description}
          placeholder="What is this room for?"
        />
      </div>

      <div class="space-y-1.5">
        <Label>Members</Label>
        <div class="max-h-40 overflow-y-auto rounded-md border border-border">
          <ul class="divide-y divide-border">
            {#each availableUsers as user (user.id)}
              <li>
                <label class="flex cursor-pointer items-center gap-3 px-3 py-2 hover:bg-muted/50">
                  <input
                    type="checkbox"
                    class="size-4 rounded border-border"
                    checked={selectedMemberIds.includes(user.id)}
                    onchange={() => toggleMember(user.id)}
                  />
                  <span class="text-sm">{user.name}</span>
                </label>
              </li>
            {/each}
          </ul>
        </div>
        <p class="text-xs text-muted-foreground">
          {selectedMemberIds.length} selected · you'll be added automatically
        </p>
      </div>

      <div class="space-y-1.5">
        <Label>Background</Label>
        <BackgroundPicker bind:value={background} />
      </div>

      <Dialog.Footer>
        <Dialog.Close>
          {#snippet child({ props })}
            <Button {...props} type="button" variant="outline">Cancel</Button>
          {/snippet}
        </Dialog.Close>
        <Button type="submit" disabled={!canSubmit}>Create room</Button>
      </Dialog.Footer>
    </form>
  </Dialog.Content>
</Dialog.Root>
