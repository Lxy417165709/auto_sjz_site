<template>
  <section class="mt-16">
    <div class="relative overflow-hidden rounded-[28px] bg-brand-500 px-6 py-12 sm:px-12 sm:py-14">
      <div class="pointer-events-none absolute -right-24 -top-32 h-[380px] w-[380px] rounded-full bg-accent-400/30 blur-3xl" aria-hidden="true"></div>
      <div class="pointer-events-none absolute -bottom-40 left-1/3 h-[320px] w-[520px] rounded-full bg-brand-700/25 blur-3xl" aria-hidden="true"></div>
      <div class="relative flex flex-col gap-7 lg:flex-row lg:items-center lg:justify-between">
        <div class="max-w-xl">
          <p class="text-[11px] font-semibold uppercase tracking-[0.2em] text-white/60">{{ eyebrow }}</p>
          <h2 class="mt-3 text-[26px] font-semibold leading-[1.15] tracking-[-0.03em] text-white sm:text-[34px]">{{ title }}</h2>
          <p class="mt-3 text-[13px] leading-7 text-white/75 sm:text-sm">{{ desc }}</p>
        </div>
        <div class="flex flex-wrap gap-3">
          <router-link :to="primaryTo" class="btn-light">{{ primaryLabel }}</router-link>
          <router-link v-if="secondaryLabel" :to="secondaryTo" class="btn-dark border-white/25 bg-white/10">{{ secondaryLabel }}</router-link>
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
