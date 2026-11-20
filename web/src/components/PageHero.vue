<template>
  <section class="relative overflow-hidden border-b border-[#e7e5e2] bg-[#fafaf9]">
    <div class="pointer-events-none absolute -right-24 -top-32 h-[420px] w-[420px] rounded-full bg-brand-500/[0.07] blur-3xl" aria-hidden="true"></div>
    <div class="wrap relative py-14 sm:py-20">
      <p class="eyebrow">{{ eyebrow }}</p>
      <h1 class="h1 mt-4 max-w-4xl">{{ title }}</h1>
      <p class="lead mt-5 max-w-2xl">{{ desc }}</p>
      <div v-if="$slots.default" class="mt-7">
        <slot />
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
defineProps<{ eyebrow: string; title: string; desc: string }>()
</script>
