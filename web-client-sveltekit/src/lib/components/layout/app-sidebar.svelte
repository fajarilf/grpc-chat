<script lang="ts">
  import { LayoutDashboard, Monitor, Users, History, MessageCircleCode } from '@lucide/svelte/icons'  
  import { Sidebar, useSidebar } from "../ui/sidebar";
  import SidebarHeader from '../ui/sidebar/sidebar-header.svelte';
  import SidebarContent from '../ui/sidebar/sidebar-content.svelte';
  import SidebarGroup from '../ui/sidebar/sidebar-group.svelte';
  import SidebarGroupContent from '../ui/sidebar/sidebar-group-content.svelte';
  import SidebarMenu from '../ui/sidebar/sidebar-menu.svelte';
  import SidebarMenuItem from '../ui/sidebar/sidebar-menu-item.svelte';
  import SidebarMenuButton from '../ui/sidebar/sidebar-menu-button.svelte';
  import { page } from '$app/state';

  const navItems = [
    { label: 'Main', href: '/', icon: LayoutDashboard },
    // { label: "Machine Monitoring", href: "/monitoring", icon: Monitor },
    // { label: "History", href: "/history", icon: History },
    // { label: "User Management", href: "/user-management", icon: Users },
  ]

  const state = useSidebar();
  const pathName = $derived(page.url.pathname);
</script>

<Sidebar collapsible="icon">
    <SidebarHeader>
        <div class="flex items-center gap-2 px-2 py-1.5">
            <div class="flex size-8 shrink-0 items-center justify-center rounded-md text-warm-white bg-blue-600">
                <MessageCircleCode class="size-4"/>
            </div>
            {#if state.open}
                <div class="flex flex-col leading-tight">
                    <span class="text-sm font-semibold text-center">gRPC Chat</span>
                </div>
            {/if}
        </div>
    </SidebarHeader>
    <SidebarContent class="mt-5">
        <SidebarGroup>
            <SidebarGroupContent>
                <SidebarMenu>
                    {#each navItems as { label, href, icon: Icon } (href)}
                        <SidebarMenuItem>
                            <SidebarMenuButton
                                class="mb-2 rounded-sm text-[14px] mx-auto"
                                isActive={pathName === href || pathName.startsWith(href + "/")}
                                tooltipContent={label}
                            >
                                {#snippet child({props})}
                                    <a {href} {...props}>
                                        <Icon class="size-4" />
                                        <span>{label}</span>
                                    </a>
                                {/snippet}
                            </SidebarMenuButton>
                        </SidebarMenuItem>
                    {/each}
                </SidebarMenu>
            </SidebarGroupContent>
        </SidebarGroup>
    </SidebarContent>
</Sidebar>