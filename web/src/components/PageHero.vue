<template>
  <section class="border-b border-[#f0dfd3] bg-[#fff8f1]">
    <div class="wrap py-12 sm:py-16">
      <p class="eyebrow">{{ eyebrow }}</p>
      <h1 class="h1 mt-3 max-w-4xl">{{ title }}</h1>
      <p class="lead mt-4 max-w-2xl">{{ desc }}</p>
      <div v-if="$slots.default" class="mt-6">
        <slot />
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
defineProps<{ eyebrow: string; title: string; desc: string }>()
</script>
