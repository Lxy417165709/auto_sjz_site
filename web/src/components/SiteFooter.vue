<template>
  <footer class="mt-24 border-t border-[#dfe6f3] bg-[#f3f7ff] text-[#55647c]">
    <div class="wrap grid gap-10 py-14 lg:grid-cols-[1.4fr_2fr]">
      <div>
        <div class="flex items-center gap-2.5">
          <svg class="h-8 w-8" viewBox="0 0 40 40" fill="none" aria-hidden="true">
            <rect x="1" y="1" width="38" height="38" rx="10" fill="url(#footerGradient)" />
            <path d="M20 9.5c1.5 3.9-.2 6.1-1.6 7.8-1.4 1.7-2.3 3.2-2.3 5.1a3.9 3.9 0 0 0 7.8 0c0-1.3-.5-2.5-1.5-3.5-.1 1.4-.8 2.4-1.8 3 .9-3.9-.2-7.6-1.6-11.4z" fill="white" />
            <defs>
              <linearGradient id="footerGradient" x1="0" y1="0" x2="40" y2="40">
                <stop stop-color="#5c86ff" />
                <stop offset="1" stop-color="#2447d6" />
              </linearGradient>
            </defs>
          </svg>
          <span class="text-[20px] font-semibold tracking-[-0.03em] text-[#07152e]">炎火</span>
        </div>
        <p class="mt-4 max-w-sm text-sm leading-6 text-[#66748c]">
          {{ site?.tagline || '三角洲业务自动化管理平台' }}
        </p>
        <p v-if="contact?.email" class="mt-4 text-sm text-[#66748c]">
          <a :href="`mailto:${contact.email}`" class="underline decoration-[#c3cfe4] underline-offset-4 hover:text-brand-700">{{ contact.email }}</a>
        </p>
        <p v-if="contact?.working" class="mt-1 text-xs text-[#8490a5]">{{ contact.working }}</p>
      </div>

      <div class="grid grid-cols-2 gap-8 sm:grid-cols-3">
        <div v-for="col in columns" :key="col.title">
          <p class="text-xs font-semibold uppercase tracking-[0.16em] text-[#8490a5]">{{ col.title }}</p>
          <ul class="mt-3 space-y-2">
            <li v-for="link in col.links" :key="link.href + link.label">
              <router-link v-if="!isExternal(link.href)" :to="link.href" class="text-sm text-[#66748c] transition hover:text-brand-700">
                {{ link.label }}
              </router-link>
              <a v-else :href="link.href" class="text-sm text-[#66748c] transition hover:text-brand-700" rel="noopener">
                {{ link.label }}
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>

    <div class="border-t border-[#dfe6f3]">
      <div class="wrap flex flex-col gap-3 py-6 text-xs text-[#8490a5] sm:flex-row sm:items-center sm:justify-between">
        <p>{{ site?.copyright || '© 2026 炎火.' }}</p>
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
