<template>
  <div>
    <section class="relative overflow-hidden bg-ink-900 text-white">
      <div class="grid-pattern pointer-events-none absolute inset-0 opacity-70" aria-hidden="true"></div>
      <div class="pointer-events-none absolute -left-32 top-10 h-72 w-72 rounded-full bg-brand-600/30 blur-3xl" aria-hidden="true"></div>
      <div class="pointer-events-none absolute -right-20 bottom-0 h-72 w-72 rounded-full bg-accent-500/20 blur-3xl" aria-hidden="true"></div>

      <div class="wrap relative grid gap-12 py-16 lg:grid-cols-[1.05fr_0.95fr] lg:items-center lg:py-24">
        <div>
          <p class="text-[11px] font-semibold uppercase tracking-[0.24em] text-accent-400">按号收费 · 设备不按台计费</p>
          <h1 class="mt-4 text-[32px] leading-[1.12] font-semibold tracking-tight sm:text-[42px] lg:text-[48px]">
            一个号 3 元/月，<br />
            设备不再按台收费
          </h1>
          <p class="mt-5 max-w-xl text-[15px] leading-7 text-slate-300">
            auto_sjz 用一个管理中台管住设备与账号：按账号结算、批量下发配置、异常自动通知。设备多少不额外收钱，账号增减才计费。
          </p>
          <div class="mt-8 flex flex-wrap gap-3">
            <router-link to="/contact" class="btn-primary px-6">预约演示</router-link>
            <router-link to="/pricing" class="btn-dark px-6">查看按号收费</router-link>
          </div>
          <dl class="mt-10 grid max-w-lg grid-cols-2 gap-x-8 gap-y-4 border-t border-white/10 pt-6 text-sm">
            <div>
              <dt class="text-slate-400">统一单价</dt>
              <dd class="mt-1 font-medium text-white">3 元 / 号 / 月</dd>
            </div>
            <div>
              <dt class="text-slate-400">设备费用</dt>
              <dd class="mt-1 font-medium text-white">不计费</dd>
            </div>
            <div>
              <dt class="text-slate-400">管理方式</dt>
              <dd class="mt-1 font-medium text-white">中台批量管理</dd>
            </div>
            <div>
              <dt class="text-slate-400">异常提醒</dt>
              <dd class="mt-1 font-medium text-white">监控通知</dd>
            </div>
          </dl>
        </div>

        <div class="relative">
          <div class="rounded-2xl border border-white/10 bg-ink-800/80 p-5 shadow-2xl backdrop-blur">
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium text-white">管理台 · 设备概览</span>
              <span class="rounded-full border border-white/15 px-2 py-0.5 text-[10px] text-slate-400">示意</span>
            </div>
            <div class="mt-4 grid grid-cols-2 gap-3">
              <div class="rounded-xl border border-white/10 bg-white/5 p-4">
                <p class="text-[11px] text-slate-400">在线设备</p>
                <p class="mt-1 font-mono text-xl text-white">12 / 12</p>
              </div>
              <div class="rounded-xl border border-white/10 bg-white/5 p-4">
                <p class="text-[11px] text-slate-400">账号实例</p>
                <p class="mt-1 font-mono text-xl text-white">60</p>
              </div>
              <div class="rounded-xl border border-white/10 bg-white/5 p-4">
                <p class="text-[11px] text-slate-400">本月账单</p>
                <p class="mt-1 font-mono text-xl text-accent-400">¥180</p>
              </div>
              <div class="rounded-xl border border-white/10 bg-white/5 p-4">
                <p class="text-[11px] text-slate-400">待处理通知</p>
                <p class="mt-1 font-mono text-xl text-white">2</p>
              </div>
            </div>
            <div class="mt-3 rounded-xl border border-white/10 bg-white/5 p-4">
              <div class="flex items-center justify-between text-[11px] text-slate-400">
                <span>授权周期</span>
                <span class="font-mono text-slate-300">到期 2026-12-31</span>
              </div>
              <div class="mt-2 h-1.5 overflow-hidden rounded-full bg-white/10">
                <div class="h-full w-[68%] rounded-full bg-accent-400"></div>
              </div>
            </div>
            <p class="mt-3 text-[11px] text-slate-500">示意数据，不含真实设备与账号信息</p>
          </div>
        </div>
      </div>
    </section>

    <div class="wrap">
      <section class="py-16 sm:py-20">
        <div class="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
          <div>
            <p class="eyebrow">我们的特点</p>
            <h2 class="h2 mt-2">不按设备卖容量，按号卖服务</h2>
          </div>
          <router-link to="/features" class="text-sm font-medium text-brand-600 hover:text-brand-700">查看业务能力 →</router-link>
        </div>

        <div class="mt-8 grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
          <article v-for="item in highlights" :key="item.title" class="card">
            <div class="flex h-9 w-9 items-center justify-center rounded-xl bg-brand-50 text-brand-600">
              <span v-html="item.icon" class="[&>svg]:h-[18px] [&>svg]:w-[18px]"></span>
            </div>
            <h3 class="mt-4 text-[15px] font-semibold text-slate-900">{{ item.title }}</h3>
            <p class="mt-2 text-[13px] leading-6 text-slate-500">{{ item.desc }}</p>
          </article>
        </div>
      </section>

      <section class="rounded-3xl border border-slate-200 bg-slate-50/70 px-5 py-10 sm:px-10">
        <div class="grid gap-8 lg:grid-cols-[1fr_0.85fr] lg:items-center">
          <div>
            <p class="eyebrow">算一笔账</p>
            <h2 class="mt-2 text-2xl font-semibold tracking-tight text-slate-900">10 台设备、每台 5 个号，一个月 150 元</h2>
            <p class="mt-3 text-sm leading-7 text-slate-600">
              按号收费的好处是账目简单：设备数量不影响费用，账号加一个算一个，减一个省一个。传统按设备收费，设备闲着也要付整套的钱。
            </p>
            <div class="mt-6 rounded-2xl bg-ink-900 px-5 py-5 text-white">
              <p class="font-mono text-[12px] text-slate-400">10 台 × 5 个号</p>
              <p class="mt-2 font-mono text-lg">
                50 个号 <span class="text-slate-500">×</span> 3 元 <span class="text-slate-500">=</span>
                <span class="text-accent-400">150 元 / 月</span>
              </p>
            </div>
            <router-link to="/pricing" class="mt-6 inline-flex items-center gap-1 text-sm font-medium text-brand-600 hover:text-brand-700">
              用计算器算你的月费 →
            </router-link>
          </div>

          <div class="rounded-2xl border border-slate-200 bg-white p-6">
            <p class="text-xs font-semibold uppercase tracking-[0.16em] text-slate-500">谁适合用</p>
            <ul class="mt-4 space-y-4">
              <li v-for="scene in scenes" :key="scene.title">
                <p class="text-sm font-medium text-slate-900">{{ scene.title }}</p>
                <p class="mt-1 text-[13px] leading-6 text-slate-500">{{ scene.desc }}</p>
              </li>
            </ul>
          </div>
        </div>
      </section>

      <section class="mt-20 grid gap-4 rounded-3xl border border-slate-200 p-7 sm:p-10 lg:grid-cols-[1fr_auto] lg:items-center">
        <div>
          <p class="eyebrow">给技术同学</p>
          <h2 class="mt-2 text-2xl font-semibold tracking-tight text-slate-900">架构、端口、授权链路、部署形态</h2>
          <p class="mt-3 max-w-2xl text-sm leading-7 text-slate-600">
            设备如何找到 Manager、快照怎么落盘、授权如何本地验签、根节点与子节点怎么拆，这些内容集中在技术架构与部署接入里，不和业务介绍混在一起。
          </p>
        </div>
        <div class="flex flex-wrap gap-3">
          <router-link to="/tech" class="btn-primary">进入技术架构</router-link>
          <router-link to="/deploy" class="btn-secondary">部署接入</router-link>
        </div>
      </section>

      <CtaBand
        title="先按 3 元/号/月 算，再决定规模"
        desc="告诉我们设备台数与每台账号数，我们给出对应的账号实例数、月费和部署建议。官网不承接下单，开通与续费在 Manager 完成。"
        primary-label="申请试用"
        primary-to="/contact"
        secondary-label="了解计费规则"
        secondary-to="/pricing"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import CtaBand from '@/components/CtaBand.vue'

const highlights = [
  {
    title: '按号收费，设备不计费',
    desc: '一个号 3 元/月，一台设备 5 个号就是 15 元/月。加设备不加钱，账号增减直接影响成本。',
    icon: '<svg viewBox="0 0 20 20" fill="none"><path d="M10 2.5v15" stroke="currentColor" stroke-width="1.5"/><path d="M5.5 6.5h7a2 2 0 0 1 0 4H7a2 2 0 0 0 0 4h7" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/></svg>',
  },
  {
    title: '一个中台，批量管理',
    desc: '设备、分组、账号与配置集中在一个后台，批量下发、集中查看，不用逐台登录操作。',
    icon: '<svg viewBox="0 0 20 20" fill="none"><rect x="2.5" y="3" width="15" height="11" rx="2" stroke="currentColor" stroke-width="1.5"/><path d="M6 17h8M10 14v3M6.5 8.5h4M6.5 11h7" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/></svg>',
  },
  {
    title: '监控通知，异常不漏',
    desc: '设备在线状态、执行异常、授权到期集中监控，触发后按配置的通知渠道提醒。',
    icon: '<svg viewBox="0 0 20 20" fill="none"><path d="M6 8a4 4 0 0 1 8 0c0 3 1 4 1.5 4.5H4.5C5 12 6 11 6 8Z" stroke="currentColor" stroke-width="1.5"/><path d="M8.5 15a1.5 1.5 0 0 0 3 0" stroke="currentColor" stroke-width="1.5"/></svg>',
  },
  {
    title: '配置统一，不用逐台改',
    desc: '配置只在中台维护，变更批量下发到设备。设备离线只读，恢复后自动对齐最新配置。',
    icon: '<svg viewBox="0 0 20 20" fill="none"><rect x="3" y="3" width="6" height="6" rx="1.5" stroke="currentColor" stroke-width="1.5"/><rect x="11" y="11" width="6" height="6" rx="1.5" stroke="currentColor" stroke-width="1.5"/><path d="M9 6h3.5A1.5 1.5 0 0 1 14 7.5V11" stroke="currentColor" stroke-width="1.5"/></svg>',
  },
  {
    title: '到期不删号、不登出',
    desc: '授权到期只停执行，账号与设备配置原样保留，续费后接着跑，不用重新登录。',
    icon: '<svg viewBox="0 0 20 20" fill="none"><rect x="4" y="8" width="12" height="9" rx="2" stroke="currentColor" stroke-width="1.5"/><path d="M7 8V6a3 3 0 1 1 6 0v2M10 12v2" stroke="currentColor" stroke-width="1.5"/></svg>',
  },
  {
    title: '独立部署，数据自持',
    desc: 'Docker 私有化部署，Manager 与 discovery 可以分开，账号与业务数据留在自己的机器上。',
    icon: '<svg viewBox="0 0 20 20" fill="none"><rect x="3" y="4" width="14" height="5" rx="1.5" stroke="currentColor" stroke-width="1.5"/><rect x="3" y="11" width="14" height="5" rx="1.5" stroke="currentColor" stroke-width="1.5"/><circle cx="6.5" cy="6.5" r="1" fill="currentColor"/><circle cx="6.5" cy="13.5" r="1" fill="currentColor"/></svg>',
  },
]

const scenes = [
  { title: '多账号运营团队', desc: '一台设备挂多个号，按号结算不会为设备容量多付钱。' },
  { title: '设备规模持续扩张', desc: '加设备不加授权费，按实际账号数扩展，预算可控。' },
  { title: '需要长期稳定执行', desc: '到期不删号、离线按同一时间失效，续费即恢复，不用重建环境。' },
]
</script>
