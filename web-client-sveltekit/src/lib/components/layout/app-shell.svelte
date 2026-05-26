<script lang="ts">
  import { page } from "$app/state";
  import type { HTMLAttributes } from "svelte/elements";
  import Separator from "../ui/separator/separator.svelte";
  import SidebarInset from "../ui/sidebar/sidebar-inset.svelte";
  import SidebarProvider from "../ui/sidebar/sidebar-provider.svelte";
  import SidebarTrigger from "../ui/sidebar/sidebar-trigger.svelte";
  import AppSidebar from "./app-sidebar.svelte";
  import type { WithElementRef } from "$lib/utils";

  let { children } : WithElementRef<HTMLAttributes<HTMLElement>> = $props();

  let now = $state(new Date());

  $effect(() => {
    const interval = setInterval(() => {
        now = new Date()
    }, 1000)

    return () => clearInterval(interval);
  })

  function pageTitleFor(path: string): string {
    return "Main"
  }
</script>

<SidebarProvider style="--sidebar-width-icon: 4rem;">
    <AppSidebar/>
    <SidebarInset class="min-w-0">
        <header class="glass-header sticky top-0 z-40 flex h-14 shrink-0 items-center gap-2 border-b px-4">
            <SidebarTrigger/>
            <Separator orientation="vertical" class="mx-1 h-5"/>
            <h1 class="text-sm font-semibold tracking-normal">
                {pageTitleFor(page.url.pathname)}
            </h1>
            <div class="ml-auto flex items-center gap-1">
                <h1 class="text-sm tracking-wide">
                    {now.toLocaleString()}
                </h1>
            </div>
        </header>
        <div class="flex min-w-0 flex-1 flex-col p-4 md:p-6">
            {@render children?.()}
        </div>
    </SidebarInset>
</SidebarProvider>