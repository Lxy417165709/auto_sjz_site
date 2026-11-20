<template>
  <router-link to="/" class="inline-flex items-center gap-2.5" aria-label="炎火 首页">
    <svg class="h-8 w-8" viewBox="0 0 40 40" fill="none" aria-hidden="true">
      <rect x="1" y="1" width="38" height="38" rx="10" fill="url(#brandGradient)" />
      <path
        d="M20 9.5c1.5 3.9-.2 6.1-1.6 7.8-1.4 1.7-2.3 3.2-2.3 5.1a3.9 3.9 0 0 0 7.8 0c0-1.3-.5-2.5-1.5-3.5-.1 1.4-.8 2.4-1.8 3 .9-3.9-.2-7.6-1.6-11.4z"
        fill="white"
      />
      <path d="M20 20.6c.7 1.7-.3 2.7-1.1 3.6.7.4 1.4.4 2 .1-.3-1.2-.4-2.5-.9-3.7z" fill="#CFE0FF" opacity="0.9" />
      <defs>
        <linearGradient id="brandGradient" x1="4" y1="3" x2="37" y2="38">
          <stop stop-color="#5c86ff" />
          <stop offset="1" stop-color="#2447d6" />
        </linearGradient>
      </defs>
    </svg>
    <span class="text-[17px] font-semibold tracking-[-0.03em]" :class="light ? 'text-white' : 'text-[#07152e]'">炎火</span>
  </router-link>
</template>

<script setup lang="ts">
withDefaults(defineProps<{ light?: boolean }>(), { light: false })
</script>
