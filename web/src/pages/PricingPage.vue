<template>
  <div>
    <PageHero
      eyebrow="按号收费"
      title="授权按账号实例计量，设备台数不额外计费"
      desc="一台设备配 5 个账号，就是 5 份授权；10 台设备各 5 个账号，就是 50 份授权。授权时效由 Auth 签发，设备按绝对到期时间执行。"
    >
      <div class="flex flex-wrap items-center gap-3">
        <span class="tag font-mono">1 账号 = 1 授权实例</span>
        <span class="tag">到期不删号、不登出</span>
        <span class="tag">离线不顺延</span>
      </div>
    </PageHero>

    <div class="wrap">
      <div v-if="!pricing" class="py-16">
        <div v-if="siteState.error" class="rounded-2xl border border-red-200 bg-red-50 p-6 text-sm text-red-700">
          计费内容加载失败（{{ siteState.error }}）。
          <button class="ml-2 font-medium underline" @click="reload">重新加载</button>
        </div>
        <div v-else class="space-y-3">
          <div class="h-6 w-56 animate-pulse rounded bg-slate-100"></div>
          <div class="h-24 animate-pulse rounded-2xl bg-slate-100"></div>
        </div>
      </div>

      <template v-else>
        <section class="grid gap-4 py-14 sm:py-16 lg:grid-cols-[1.15fr_0.85fr]">
          <div>
            <p class="eyebrow">计费模型</p>
            <h2 class="h2 mt-2">{{ pricing.headline }}</h2>
            <p class="lead mt-4">{{ pricing.summary }}</p>

            <ol class="mt-8 space-y-4">
              <li v-for="(rule, index) in pricing.billingRules" :key="rule" class="flex gap-3">
                <span class="mt-0.5 flex h-6 w-6 flex-none items-center justify-center rounded-full bg-brand-50 font-mono text-[11px] text-brand-600">
                  {{ String(index + 1).padStart(2, '0') }}
                </span>
                <p class="text-sm leading-7 text-slate-600">{{ rule }}</p>
              </li>
            </ol>
          </div>

          <aside class="h-fit rounded-3xl border border-slate-200 bg-white p-6 shadow-card">
            <p class="eyebrow">计费示例</p>
            <p class="mt-2 text-sm text-slate-500">{{ pricing.calculator.hint }}</p>

            <div class="mt-5 space-y-4">
              <label class="block">
                <span class="text-xs font-medium text-slate-600">设备台数</span>
                <input v-model.number="devices" type="number" min="0" max="100000" class="field mt-1.5" />
              </label>
              <label class="block">
                <span class="text-xs font-medium text-slate-600">每台设备账号数</span>
                <input v-model.number="accountsPerDevice" type="number" min="0" max="1000" class="field mt-1.5" />
              </label>
            </div>

            <div class="mt-5 rounded-2xl bg-ink-900 px-5 py-5 text-white">
              <p class="font-mono text-[11px] text-slate-400">授权实例数</p>
              <p class="mt-1 font-mono text-2xl text-accent-400">{{ totalUnits }}</p>
              <p class="mt-1 font-mono text-[11px] text-slate-400">{{ devices || 0 }} × {{ accountsPerDevice || 0 }}</p>
            </div>

            <div class="mt-5 flex items-baseline justify-between border-t border-slate-100 pt-4">
              <span class="text-xs text-slate-500">参考单价</span>
              <span class="text-sm text-slate-900">
                <template v-if="pricing.unitPriceCny !== null">
                  {{ formatCny(pricing.unitPriceCny) }} <span class="text-xs text-slate-500">/ {{ pricing.unitLabel }}</span>
                </template>
                <template v-else>报价以商务确认为准</template>
              </span>
            </div>
            <div v-if="pricing.unitPriceCny !== null" class="mt-2 flex items-baseline justify-between">
              <span class="text-xs text-slate-500">月度合计</span>
              <span class="font-mono text-lg text-slate-900">{{ formatCny(totalUnits * pricing.unitPriceCny) }}</span>
            </div>

            <router-link to="/contact" class="btn-primary mt-5 w-full">获取正式报价</router-link>
            <p class="mt-3 text-[11px] leading-5 text-slate-400">示例仅做数量换算，最终金额与授权周期以合同为准。</p>
          </aside>
        </section>

        <section class="rounded-3xl border border-slate-200 bg-slate-50/60 px-5 py-10 sm:px-9">
          <div class="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
            <div>
              <p class="eyebrow">授权包含</p>
              <h2 class="h2 mt-2">每一份账号实例授权，包含这些能力</h2>
            </div>
            <span class="tag">不按设备台数叠加</span>
          </div>
          <div class="mt-6 grid gap-4 sm:grid-cols-2">
            <div v-for="item in pricing.includes" :key="item.title" class="rounded-2xl border border-slate-200 bg-white p-5">
              <h3 class="text-sm font-semibold text-slate-900">{{ item.title }}</h3>
              <p class="mt-2 text-[13px] leading-6 text-slate-500">{{ item.desc }}</p>
            </div>
          </div>
        </section>

        <section class="grid gap-4 py-14 sm:py-16 lg:grid-cols-[1fr_1fr]">
          <div class="rounded-3xl border border-slate-200 p-7">
            <h2 class="text-lg font-semibold text-slate-900">为什么按账号，而不是按设备</h2>
            <p class="mt-3 text-sm leading-7 text-slate-600">
              设备是载体，账号才是实际业务单元。按账号实例计量，账号增减能直接反映到授权数量，避免大机器带单账号的浪费，也避免一台设备塞满账号却按一台收费的失真。
            </p>
            <ul class="mt-5 space-y-3 text-sm text-slate-600">
              <li v-for="item in compare" :key="item" class="flex gap-2.5">
                <svg class="mt-1 h-4 w-4 flex-none text-accent-600" viewBox="0 0 16 16" fill="none" aria-hidden="true">
                  <path d="M3 8.5l3 3 7-7" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" />
                </svg>
                <span>{{ item }}</span>
              </li>
            </ul>
          </div>

          <div class="rounded-3xl border border-amber-200 bg-amber-50/60 p-7">
            <h2 class="text-lg font-semibold text-slate-900">购买前需要确认</h2>
            <ul class="mt-4 space-y-3 text-[13px] leading-6 text-slate-600">
              <li v-for="note in pricing.notes" :key="note">{{ note }}</li>
              <li>官网不承接下单与支付，开通、续费、账号实例变更都在 Manager 完成。</li>
            </ul>
            <router-link to="/contact" class="btn-secondary mt-5">联系商务</router-link>
          </div>
        </section>

        <CtaBand
          title="算清授权数量，再决定规模"
          desc="告诉我们设备台数与每台账号数，我们会给出对应的授权实例数与部署建议。"
          primary-label="申请报价"
          primary-to="/contact"
          secondary-label="查看常见问题"
          secondary-to="/faq"
        />
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import PageHero from '@/components/PageHero.vue'
import CtaBand from '@/components/CtaBand.vue'
import { loadSite, siteState } from '@/lib/site'

const pricing = computed(() => siteState.data?.pricing ?? null)

const devices = ref(3)
const accountsPerDevice = ref(5)

watch(
  pricing,
  (value) => {
    if (value) {
      devices.value = value.calculator.defaultDevices
      accountsPerDevice.value = value.calculator.defaultAccountsPerDevice
    }
  },
  { immediate: true },
)

const totalUnits = computed(() => {
  const d = Number.isFinite(devices.value) ? Math.max(0, Math.floor(devices.value)) : 0
  const a = Number.isFinite(accountsPerDevice.value) ? Math.max(0, Math.floor(accountsPerDevice.value)) : 0
  return d * a
})

const compare = [
  '账号增减即授权增减，不浪费空置设备',
  '同一台设备多账号按实例计费，账目可对',
  '过期只停执行，账号与配置原样保留',
]

function formatCny(value: number): string {
  return new Intl.NumberFormat('zh-CN', { style: 'currency', currency: 'CNY', maximumFractionDigits: 2 }).format(value)
}

function reload() {
  void loadSite()
}
</script>
