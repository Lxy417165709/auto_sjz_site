<template>
  <header class="sticky top-0 z-40 border-b border-slate-200/70 bg-white/85 backdrop-blur">
    <div class="wrap flex h-16 items-center justify-between gap-4">
      <BrandMark />

      <nav class="hidden items-center gap-0.5 lg:flex" aria-label="主导航">
        <router-link
          v-for="item in nav"
          :key="item.href"
          :to="item.href"
          class="nav-link"
          active-class="nav-link-active"
        >
          {{ item.label }}
        </router-link>
      </nav>

      <div class="hidden items-center gap-3 lg:flex">
        <router-link :to="secondary.href" class="btn-secondary">{{ secondary.label }}</router-link>
        <a :href="managerLink" class="btn-primary" rel="noopener">
          {{ cta.label }}
          <svg class="h-3.5 w-3.5" viewBox="0 0 16 16" fill="none" aria-hidden="true">
            <path d="M3 8h9M9 4l4 4-4 4" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" />
          </svg>
        </a>
      </div>

      <button
        class="inline-flex h-10 w-10 items-center justify-center rounded-full border border-slate-200 text-slate-700 lg:hidden"
        :aria-expanded="open"
        aria-controls="mobile-nav"
        aria-label="打开菜单"
        @click="open = !open"
      >
        <svg v-if="!open" class="h-5 w-5" viewBox="0 0 20 20" fill="none" aria-hidden="true">
          <path d="M3 5h14M3 10h14M3 15h14" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" />
        </svg>
        <svg v-else class="h-5 w-5" viewBox="0 0 20 20" fill="none" aria-hidden="true">
          <path d="M5 5l10 10M15 5L5 15" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" />
        </svg>
      </button>
    </div>

    <div v-if="open" id="mobile-nav" class="border-t border-slate-200 bg-white lg:hidden">
      <div class="wrap flex flex-col gap-1 py-4">
        <router-link
          v-for="item in nav"
          :key="item.href"
          :to="item.href"
          class="rounded-xl px-3 py-2.5 text-sm text-slate-700 hover:bg-slate-50"
          active-class="bg-slate-100 text-slate-900"
        >
          {{ item.label }}
        </router-link>
        <div class="mt-2 flex flex-col gap-2 border-t border-slate-100 pt-3">
          <router-link :to="secondary.href" class="btn-secondary w-full">{{ secondary.label }}</router-link>
          <a :href="managerLink" class="btn-primary w-full" rel="noopener">{{ cta.label }}</a>
        </div>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import BrandMark from './BrandMark.vue'
import { loadSite, managerEntry, siteState } from '@/lib/site'
import type { NavItem, LinkItem } from '@/lib/types'

const route = useRoute()
const open = ref(false)

const fallbackNav: NavItem[] = [
  { label: '功能', href: '/features' },
  { label: '按号收费', href: '/pricing' },
  { label: '部署接入', href: '/deploy' },
  { label: '文档', href: '/docs' },
  { label: 'FAQ', href: '/faq' },
]

const nav = computed<NavItem[]>(() => siteState.data?.site.nav ?? fallbackNav)
const cta = computed<LinkItem>(() => siteState.data?.site.cta ?? { label: '登录 / 购买', href: '' })
const secondary = computed<LinkItem>(() => siteState.data?.site.secondaryCta ?? { label: '预约演示', href: '/contact' })
const managerLink = computed(() => managerEntry())

watch(() => route.fullPath, () => { open.value = false })
void loadSite()
</script>
