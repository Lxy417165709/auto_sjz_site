<template>
  <section class="mt-16">
    <div class="relative overflow-hidden rounded-3xl bg-ink-900 px-6 py-12 sm:px-12">
      <div class="grid-pattern pointer-events-none absolute inset-0 opacity-60" aria-hidden="true"></div>
      <div class="absolute -right-24 -top-24 h-64 w-64 rounded-full bg-brand-600/25 blur-3xl" aria-hidden="true"></div>
      <div class="relative flex flex-col gap-6 lg:flex-row lg:items-center lg:justify-between">
        <div class="max-w-xl">
          <p class="text-[11px] font-semibold uppercase tracking-[0.22em] text-accent-400">{{ eyebrow }}</p>
          <h2 class="mt-3 text-2xl font-semibold tracking-tight text-white sm:text-3xl">{{ title }}</h2>
          <p class="mt-3 text-sm leading-7 text-slate-300">{{ desc }}</p>
        </div>
        <div class="flex flex-wrap gap-3">
          <router-link :to="primaryTo" class="btn-primary">{{ primaryLabel }}</router-link>
          <router-link v-if="secondaryLabel" :to="secondaryTo" class="btn-dark">{{ secondaryLabel }}</router-link>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
withDefaults(
  defineProps<{
    eyebrow?: string
    title: string
    desc: string
    primaryLabel: string
    primaryTo: string
    secondaryLabel?: string
    secondaryTo?: string
  }>(),
  {
    eyebrow: '开始接入',
    secondaryLabel: '',
    secondaryTo: '/features',
  },
)
</script>
