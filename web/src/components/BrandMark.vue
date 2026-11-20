<template>
  <router-link to="/" class="inline-flex items-center gap-2.5" aria-label="炎火 首页">
    <svg class="h-8 w-8" viewBox="0 0 40 40" fill="none" aria-hidden="true">
      <rect x="1" y="1" width="38" height="38" rx="11" fill="url(#brandGradient)" />
      <path
        d="M9 21h5.5l3-8.5 4.5 16 3.5-9.5H31"
        stroke="white"
        stroke-width="2.2"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
      <defs>
        <linearGradient id="brandGradient" x1="0" y1="0" x2="40" y2="40">
          <stop stop-color="#3763f5" />
          <stop offset="1" stop-color="#12b892" />
        </linearGradient>
      </defs>
    </svg>
    <span class="text-[17px] font-semibold tracking-tight text-slate-900">炎火</span>
  </router-link>
</template>
