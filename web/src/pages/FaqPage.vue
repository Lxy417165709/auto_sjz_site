<template>
  <div>
    <PageHero
      eyebrow="FAQ"
      title="先把计费、授权、接入三件事问清楚"
      desc="官网只做介绍，登录与购买在 Manager 完成。下面这些回答来自实际接入中最常被问到的点。"
    />

    <div class="wrap">
      <div v-if="!faq.length" class="py-16">
        <div v-if="siteState.error" class="rounded-2xl border border-red-200 bg-red-50 p-6 text-sm text-red-700">
          内容加载失败（{{ siteState.error }}）。
          <button class="ml-2 font-medium underline" @click="reload">重新加载</button>
        </div>
        <div v-else class="space-y-3">
          <div v-for="i in 5" :key="i" class="h-16 animate-pulse rounded-xl bg-slate-100"></div>
        </div>
      </div>

      <template v-else>
        <div class="flex flex-wrap items-center gap-2 py-10">
          <button
            v-for="group in groups"
            :key="group"
            class="rounded-full border px-4 py-1.5 text-sm transition"
            :class="activeGroup === group
              ? 'border-brand-600 bg-brand-600 text-white'
              : 'border-slate-200 bg-white text-slate-600 hover:border-slate-300'"
            @click="activeGroup = group"
          >
            {{ group }}
          </button>
          <span class="ml-auto text-xs text-slate-400">共 {{ filtered.length }} 个问题</span>
        </div>

        <section class="divide-y divide-slate-200 rounded-2xl border border-slate-200 bg-white">
          <div v-for="item in filtered" :key="item.id" class="px-5 sm:px-7">
            <button
              class="flex w-full items-center justify-between gap-4 py-5 text-left"
              :aria-expanded="openIds.includes(item.id)"
              @click="toggle(item.id)"
            >
              <span class="flex items-start gap-3">
                <span class="mt-0.5 hidden font-mono text-[11px] text-slate-400 sm:block">{{ item.group }}</span>
                <span class="text-sm font-medium text-slate-900">{{ item.question }}</span>
              </span>
              <span
                class="flex h-6 w-6 flex-none items-center justify-center rounded-full border border-slate-200 text-slate-500 transition"
                :class="openIds.includes(item.id) ? 'rotate-45 border-brand-200 text-brand-600' : ''"
                aria-hidden="true"
              >
                <svg class="h-3 w-3" viewBox="0 0 12 12" fill="none">
                  <path d="M6 2v8M2 6h8" stroke="currentColor" stroke-width="1.4" />
                </svg>
              </span>
            </button>
            <div v-show="openIds.includes(item.id)" class="pb-5 pl-0 text-[13px] leading-7 text-slate-600 sm:pl-14">
              {{ item.answer }}
            </div>
          </div>
        </section>

        <section class="mt-12 grid gap-4 sm:grid-cols-2">
          <div class="rounded-2xl border border-[#f0dfd3] bg-[#fffaf4] p-6">
            <h2 class="text-sm font-semibold text-slate-900">没找到答案？</h2>
            <p class="mt-2 text-[13px] leading-6 text-slate-500">
              把设备数量、每台账号数、当前部署形态发给我们，会更快给出结论。
            </p>
            <router-link to="/contact" class="btn-primary mt-4">联系与试用</router-link>
          </div>
          <div class="rounded-2xl border border-slate-200 p-6">
            <h2 class="text-sm font-semibold text-slate-900">想先看规则细节</h2>
            <p class="mt-2 text-[13px] leading-6 text-slate-500">
              计费换算、授权到期判定、discovery 与 Manager 的链路关系，文档里有完整说明。
            </p>
            <router-link to="/docs" class="btn-secondary mt-4">打开文档</router-link>
          </div>
        </section>

        <CtaBand
          title="还有问题，直接说"
          desc="我们不做机器人客服，接入问题由熟悉部署的人直接回复。"
          primary-label="提交问题"
          primary-to="/contact"
          secondary-label=""
          secondary-to="/docs"
        />
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import PageHero from '@/components/PageHero.vue'
import CtaBand from '@/components/CtaBand.vue'
import { loadSite, siteState } from '@/lib/site'

const faq = computed(() => siteState.data?.faq ?? [])
const activeGroup = ref('全部')
const openIds = ref<string[]>([])

const groups = computed(() => ['全部', ...Array.from(new Set(faq.value.map((item) => item.group)))])
const filtered = computed(() =>
  activeGroup.value === '全部' ? faq.value : faq.value.filter((item) => item.group === activeGroup.value),
)

function toggle(id: string) {
  openIds.value = openIds.value.includes(id) ? openIds.value.filter((x) => x !== id) : [...openIds.value, id]
}

function reload() {
  void loadSite()
}
</script>
