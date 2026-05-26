<script lang="ts">
  import { cn, type WithElementRef } from "$lib/utils.js";
  import type { HTMLAnchorAttributes, HTMLAttributes } from "svelte/elements";
  import type { Snippet } from "svelte";

  type BaseProps = {
    overlayClass?: string;
    contentClass?: string;
    bg?: Snippet<[]>;
    overlay?: Snippet<[]>;
    children?: Snippet<[]>;
  };

  type Props = BaseProps &
    (
      | ({ href: string } & WithElementRef<HTMLAnchorAttributes, HTMLAnchorElement>)
      | ({ href?: undefined } & WithElementRef<HTMLAttributes<HTMLDivElement>, HTMLDivElement>)
    );

  let {
    ref = $bindable(null),
    href,
    class: className,
    overlayClass,
    contentClass,
    bg,
    overlay,
    children,
    ...restProps
  }: Props = $props();
</script>

{#snippet body()}
  {#if bg}
    {@render bg()}
  {/if}
  {#if overlay}
    {@render overlay()}
  {:else}
    <div
      aria-hidden="true"
      class={cn(
        "absolute inset-0 bg-linear-to-t from-black/80 via-black/30 to-transparent",
        overlayClass
      )}
    ></div>
  {/if}
  <div class={cn("relative flex h-full flex-col justify-end p-4 text-white", contentClass)}>
    {@render children?.()}
  </div>
{/snippet}

{#if href}
  <a
    bind:this={ref as HTMLAnchorElement | null}
    {href}
    data-slot="image-card"
    class={cn(
      "group relative block aspect-4/3 w-full overflow-hidden rounded-lg",
      "transition-transform duration-200 hover:-translate-y-0.5",
      className
    )}
    {...restProps as HTMLAnchorAttributes}
  >
    {@render body()}
  </a>
{:else}
  <div
    bind:this={ref as HTMLDivElement | null}
    data-slot="image-card"
    class={cn(
      "relative block aspect-4/3 w-full overflow-hidden rounded-lg",
      className
    )}
    {...restProps as HTMLAttributes<HTMLDivElement>}
  >
    {@render body()}
  </div>
{/if}
