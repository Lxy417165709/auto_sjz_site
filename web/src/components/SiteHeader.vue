<template>
  <header class="sticky top-0 z-40 border-b border-white/10 bg-[#08080a]/95 backdrop-blur-xl">
    <div class="wrap flex h-16 items-center justify-between gap-4">
      <BrandMark light />

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

      <div class="hidden items-center gap-2 lg:flex">
        <router-link :to="secondary.href" class="rounded-full px-4 py-2 text-sm text-white/55 transition hover:bg-white/[0.06] hover:text-white">
          {{ secondary.label }}
        </router-link>
        <a :href="managerLink" class="btn-primary" rel="noopener">
          {{ cta.label }}
          <svg class="h-3.5 w-3.5" viewBox="0 0 16 16" fill="none" aria-hidden="true">
            <path d="M3 8h9M9 4l4 4-4 4" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" />
          </svg>
        </a>
      </div>

      <button
        class="inline-flex h-10 w-10 items-center justify-center rounded-full border border-white/15 text-white/80 lg:hidden"
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

    <div v-if="open" id="mobile-nav" class="border-t border-white/10 bg-[#08080a] lg:hidden">
      <div class="wrap flex flex-col gap-1 py-4">
        <router-link
          v-for="item in nav"
          :key="item.href"
          :to="item.href"
          class="rounded-xl px-3 py-2.5 text-sm text-white/65 hover:bg-white/[0.06]"
          active-class="bg-white/[0.08] text-white"
        >
          {{ item.label }}
        </router-link>
        <div class="mt-2 flex flex-col gap-2 border-t border-white/10 pt-3">
          <router-link :to="secondary.href" class="btn-dark w-full">{{ secondary.label }}</router-link>
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
  { label: '业务能力', href: '/features' },
  { label: '按号收费', href: '/pricing' },
  { label: '接入流程', href: '/start' },
  { label: '技术架构', href: '/tech' },
  { label: '文档', href: '/docs' },
  { label: 'FAQ', href: '/faq' },
]

const nav = computed<NavItem[]>(() => siteState.data?.site.nav ?? fallbackNav)
const cta = computed<LinkItem>(() => siteState.data?.site.cta ?? { label: '登录 / 购买', href: '' })
const secondary = computed<LinkItem>(() => siteState.data?.site.secondaryCta ?? { label: '界面演示', href: '/console' })
const managerLink = computed(() => managerEntry())

watch(() => route.fullPath, () => { open.value = false })
void loadSite()
</script>
