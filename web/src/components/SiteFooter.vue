<template>
  <footer class="mt-20 border-t border-slate-200 bg-ink-800 text-slate-300">
    <div class="wrap grid gap-10 py-14 lg:grid-cols-[1.4fr_2fr]">
      <div>
        <div class="flex items-center gap-2.5">
          <svg class="h-8 w-8" viewBox="0 0 40 40" fill="none" aria-hidden="true">
            <rect x="1" y="1" width="38" height="38" rx="11" fill="url(#footerGradient)" />
            <path d="M9 21h5.5l3-8.5 4.5 16 3.5-9.5H31" stroke="white" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round" />
            <defs>
              <linearGradient id="footerGradient" x1="0" y1="0" x2="40" y2="40">
                <stop stop-color="#3763f5" />
                <stop offset="1" stop-color="#12b892" />
              </linearGradient>
            </defs>
          </svg>
          <span class="text-[17px] font-semibold tracking-tight text-white">auto_sjz</span>
        </div>
        <p class="mt-4 max-w-sm text-sm leading-6 text-slate-400">
          {{ site?.tagline || 'Android 设备自动化管理平台' }}
        </p>
        <p v-if="contact?.email" class="mt-4 text-sm text-slate-400">
          <a :href="`mailto:${contact.email}`" class="underline decoration-slate-600 underline-offset-4 hover:text-white">{{ contact.email }}</a>
        </p>
        <p v-if="contact?.working" class="mt-1 text-xs text-slate-500">{{ contact.working }}</p>
      </div>

      <div class="grid grid-cols-2 gap-8 sm:grid-cols-3">
        <div v-for="col in columns" :key="col.title">
          <p class="text-xs font-semibold uppercase tracking-[0.16em] text-slate-500">{{ col.title }}</p>
          <ul class="mt-3 space-y-2">
            <li v-for="link in col.links" :key="link.href + link.label">
              <router-link v-if="!isExternal(link.href)" :to="link.href" class="text-sm text-slate-400 transition hover:text-white">
                {{ link.label }}
              </router-link>
              <a v-else :href="link.href" class="text-sm text-slate-400 transition hover:text-white" rel="noopener">
                {{ link.label }}
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>

    <div class="border-t border-white/10">
      <div class="wrap flex flex-col gap-3 py-6 text-xs text-slate-500 sm:flex-row sm:items-center sm:justify-between">
        <p>{{ site?.copyright || '© 2026 auto_sjz.' }}</p>
        <div class="flex flex-wrap items-center gap-4">
          <span>官网仅做产品介绍，登录与购买在 Manager 完成</span>
          <a v-if="site?.icp" href="https://beian.miit.gov.cn/" target="_blank" rel="noopener" class="hover:text-slate-300">
            {{ site.icp }}
          </a>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { siteState } from '@/lib/site'

const site = computed(() => siteState.data?.site ?? null)
const contact = computed(() => site.value?.contact ?? null)
const columns = computed(() => site.value?.footer ?? [
  { title: '产品', links: [{ label: '功能总览', href: '/features' }, { label: '按号收费', href: '/pricing' }] },
  { title: '接入', links: [{ label: '部署接入', href: '/deploy' }, { label: '文档', href: '/docs' }] },
  { title: '支持', links: [{ label: '常见问题', href: '/faq' }, { label: '联系与试用', href: '/contact' }] },
])

function isExternal(href: string): boolean {
  return /^https?:/.test(href)
}
</script>
