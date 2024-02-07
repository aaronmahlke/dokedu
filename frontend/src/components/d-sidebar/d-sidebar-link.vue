<template>
  <component
    :is="to ? 'router-link' : 'a'"
    :to="to"
    :href="href"
    :target="href ? '_blank' : ''"
    class="flex items-center h-9 gap-3 rounded-xl pl-3 pr-1 text-neutral-500 transition-all duration-100 border border-transparent hover:text-neutral-950 relative group"
    active-class=""
    :class="{ 'text-strong': active }"
  >
    <div
      class="absolute w-full h-full bg-black/5 left-0 rounded-xl scale-50 group-hover:scale-100 transition-all opacity-0 group-hover:opacity-100"
      :class="{ '!bg-strong !scale-100 !opacity-100': active }"
    ></div>
    <component :is="icon" class="stroke-neutral-500 z-10" :size="18" :class="{ '!stroke-neutral-900': active }" />
    <div class="text-sm w-full flex-1 z-10">
      <slot />
    </div>
  </component>
</template>
<script setup lang="ts">
import { type RouteLocationRaw } from "vue-router/auto"
import type { Icon } from "@/types/types"

type Props = {
  to?: RouteLocationRaw
  href?: string
  icon: Icon
  active?: boolean
}

withDefaults(defineProps<Props>(), {
  active: false
})
</script>
