<template>
  <div>
    <section class="relative overflow-hidden bg-ink-900 text-white">
      <div class="grid-pattern pointer-events-none absolute inset-0 opacity-70" aria-hidden="true"></div>
      <div class="pointer-events-none absolute -left-32 top-10 h-72 w-72 rounded-full bg-brand-600/30 blur-3xl" aria-hidden="true"></div>
      <div class="pointer-events-none absolute -right-20 bottom-0 h-72 w-72 rounded-full bg-accent-500/20 blur-3xl" aria-hidden="true"></div>

      <div class="wrap relative grid gap-12 py-16 lg:grid-cols-[1.05fr_0.95fr] lg:items-center lg:py-24">
        <div>
          <p class="text-[11px] font-semibold uppercase tracking-[0.24em] text-accent-400">Android Device Automation</p>
          <h1 class="mt-4 text-[32px] leading-[1.12] font-semibold tracking-tight sm:text-[42px] lg:text-[48px]">
            让每一台 Android 设备，<br class="hidden sm:block" />
            都按同一份配置运行
          </h1>
          <p class="mt-5 max-w-xl text-[15px] leading-7 text-slate-300">
            auto_sjz 把设备、账号与自动化任务收拢到一套系统里：Manager 统一维护配置，discovery 负责地址解析，设备本地验签授权并按绝对到期时间执行。
          </p>
          <div class="mt-8 flex flex-wrap gap-3">
            <router-link to="/contact" class="btn-primary px-6">预约演示</router-link>
            <router-link to="/deploy" class="btn-dark px-6">查看部署接入</router-link>
          </div>
          <dl class="mt-10 grid max-w-lg grid-cols-2 gap-x-8 gap-y-4 border-t border-white/10 pt-6 text-sm">
            <div>
              <dt class="text-slate-400">计费方式</dt>
              <dd class="mt-1 font-medium text-white">按账号实例</dd>
            </div>
            <div>
              <dt class="text-slate-400">授权判定</dt>
              <dd class="mt-1 font-medium text-white">绝对到期时间</dd>
            </div>
            <div>
              <dt class="text-slate-400">设备寻址</dt>
              <dd class="mt-1 font-medium text-white">discovery 解析</dd>
            </div>
            <div>
              <dt class="text-slate-400">部署方式</dt>
              <dd class="mt-1 font-medium text-white">Docker 独立编排</dd>
            </div>
          </dl>
        </div>

        <div class="relative">
          <div class="rounded-2xl border border-white/10 bg-ink-800/80 p-4 shadow-2xl backdrop-blur">
            <div class="flex items-center gap-1.5">
              <span class="h-2.5 w-2.5 rounded-full bg-slate-500/70"></span>
              <span class="h-2.5 w-2.5 rounded-full bg-slate-500/70"></span>
              <span class="h-2.5 w-2.5 rounded-full bg-accent-500/80"></span>
              <span class="ml-2 font-mono text-[11px] text-slate-500">device-boot.log</span>
            </div>
            <div class="mt-4 space-y-2.5 font-mono text-[12px] leading-6">
              <p class="text-slate-400"><span class="text-accent-400">$</span> discovery resolve <span class="text-slate-500">→</span> manager 10.0.0.12:5001</p>
              <p class="text-slate-400"><span class="text-accent-400">$</span> snapshot pull <span class="text-slate-500">→</span> 200 OK, 12 devices</p>
              <p class="text-slate-400"><span class="text-accent-400">$</span> local config.json written</p>
              <p class="text-slate-400"><span class="text-accent-400">$</span> license verify <span class="text-slate-500">→</span> signature ok</p>
              <p class="text-slate-400"><span class="text-accent-400">$</span> expire_at 2026-12-31T00:00:00Z</p>
              <p class="text-accent-400">business running · 3 accounts active</p>
            </div>
            <div class="mt-4 border-t border-white/10 pt-3 text-[11px] text-slate-500">本地验签 + 快照落盘的示意输出，不展示真实账号数据</div>
          </div>

          <div class="mt-4 grid grid-cols-3 gap-3 text-center">
            <div class="rounded-xl border border-white/10 bg-white/5 px-3 py-3">
              <p class="font-mono text-lg text-white">5003</p>
              <p class="mt-0.5 text-[11px] text-slate-400">discovery</p>
            </div>
            <div class="rounded-xl border border-white/10 bg-white/5 px-3 py-3">
              <p class="font-mono text-lg text-white">5001</p>
              <p class="mt-0.5 text-[11px] text-slate-400">Manager Root</p>
            </div>
            <div class="rounded-xl border border-white/10 bg-white/5 px-3 py-3">
              <p class="font-mono text-lg text-white">5xxx</p>
              <p class="mt-0.5 text-[11px] text-slate-400">子节点</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <div class="wrap">
      <section class="py-16 sm:py-20">
        <div class="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
          <div>
            <p class="eyebrow">核心能力</p>
            <h2 class="h2 mt-2">一套系统管住设备、配置与授权</h2>
          </div>
          <router-link to="/features" class="text-sm font-medium text-brand-600 hover:text-brand-700">查看全部功能 →</router-link>
        </div>

        <div class="mt-8 grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
          <article v-for="item in features" :key="item.title" class="card">
            <div class="flex h-9 w-9 items-center justify-center rounded-xl bg-brand-50 text-brand-600">
              <span v-html="item.icon" class="[&>svg]:h-[18px] [&>svg]:w-[18px]"></span>
            </div>
            <h3 class="mt-4 text-[15px] font-semibold text-slate-900">{{ item.title }}</h3>
            <p class="mt-2 text-[13px] leading-6 text-slate-500">{{ item.desc }}</p>
          </article>
        </div>
      </section>

      <section class="rounded-3xl border border-slate-200 bg-slate-50/60 px-5 py-12 sm:px-10">
        <div class="max-w-2xl">
          <p class="eyebrow">系统关系</p>
          <h2 class="h2 mt-2">设备、discovery、Manager 各管一段</h2>
          <p class="lead mt-3">
            设备启动后先向 discovery 解析 Manager 地址，再拉取快照写入本地 JSON 执行业务。Manager 是配置唯一权威源，设备离线期间不允许改配置。
          </p>
        </div>
        <ArchitectureDiagram class="mt-6" />
      </section>

      <section class="grid gap-4 py-16 sm:py-20 lg:grid-cols-[1.1fr_0.9fr]">
        <div class="rounded-3xl border border-slate-200 p-7 sm:p-10">
          <p class="eyebrow">按号收费</p>
          <h2 class="mt-2 text-2xl font-semibold tracking-tight text-slate-900">一台设备 5 个账号，就是 5 份授权</h2>
          <p class="mt-3 text-sm leading-7 text-slate-600">
            授权以账号实例为计量单位，设备台数不额外计费。授权时长由 Auth 签发租约，设备本地验签后按绝对到期时间执行，到期不删账号、不登出、不改配置。
          </p>
          <div class="mt-6 rounded-2xl bg-slate-900 px-5 py-5 text-white">
            <p class="font-mono text-[12px] text-slate-400">授权实例数</p>
            <p class="mt-2 font-mono text-lg">
              3 台设备 <span class="text-slate-500">×</span> 5 个账号 <span class="text-slate-500">=</span>
              <span class="text-accent-400">15 份授权</span>
            </p>
          </div>
          <router-link to="/pricing" class="mt-6 inline-flex items-center gap-1 text-sm font-medium text-brand-600 hover:text-brand-700">
            了解计费规则与示例 →
          </router-link>
        </div>

        <div class="rounded-3xl border border-slate-200 bg-white p-7 sm:p-10">
          <p class="eyebrow">配置权威源</p>
          <h2 class="mt-2 text-2xl font-semibold tracking-tight text-slate-900">为什么配置只允许在 Manager 改</h2>
          <ul class="mt-5 space-y-4">
            <li v-for="(point, i) in authority" :key="point.title" class="flex gap-3">
              <span class="mt-0.5 flex h-6 w-6 flex-none items-center justify-center rounded-full bg-brand-50 font-mono text-[11px] text-brand-600">{{ i + 1 }}</span>
              <div>
                <p class="text-sm font-medium text-slate-900">{{ point.title }}</p>
                <p class="mt-1 text-[13px] leading-6 text-slate-500">{{ point.desc }}</p>
              </div>
            </li>
          </ul>
        </div>
      </section>

      <CtaBand
        title="先跑通一套最小环境，再按设备规模扩展"
        desc="官网只做介绍与接入说明，登录、下单、授权管理都在 Manager 完成。需要报价或私有化方案，可以直接联系我们。"
        primary-label="申请试用"
        primary-to="/contact"
        secondary-label="查看文档"
        secondary-to="/docs"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import ArchitectureDiagram from '@/components/ArchitectureDiagram.vue'
import CtaBand from '@/components/CtaBand.vue'

const features = [
  {
    title: '统一配置下发',
    desc: 'Manager 维护全部设备配置，快照下发后设备落盘执行，避免多点修改造成状态漂移。',
    icon: '<svg viewBox="0 0 20 20" fill="none"><rect x="3" y="3" width="6" height="6" rx="1.5" stroke="currentColor" stroke-width="1.5"/><rect x="11" y="11" width="6" height="6" rx="1.5" stroke="currentColor" stroke-width="1.5"/><path d="M9 6h3.5A1.5 1.5 0 0 1 14 7.5V11" stroke="currentColor" stroke-width="1.5"/></svg>',
  },
  {
    title: '自动寻址接入',
    desc: '设备通过 discovery 解析 Manager 地址，换机房、换 IP 只改登记，不用刷设备。',
    icon: '<svg viewBox="0 0 20 20" fill="none"><circle cx="10" cy="10" r="2.5" stroke="currentColor" stroke-width="1.5"/><circle cx="10" cy="10" r="6.5" stroke="currentColor" stroke-width="1.5" opacity="0.5"/><path d="M10 1.5v3M10 15.5v3M1.5 10h3M15.5 10h3" stroke="currentColor" stroke-width="1.5"/></svg>',
  },
  {
    title: '本地验签授权',
    desc: 'Auth 签发租约，设备侧验签后按绝对时间过期；离线、关机同样按点失效。',
    icon: '<svg viewBox="0 0 20 20" fill="none"><rect x="4" y="8" width="12" height="9" rx="2" stroke="currentColor" stroke-width="1.5"/><path d="M7 8V6a3 3 0 1 1 6 0v2" stroke="currentColor" stroke-width="1.5"/><path d="M10 12v2" stroke="currentColor" stroke-width="1.5"/></svg>',
  },
  {
    title: '多节点扩展',
    desc: 'Manager 支持根节点 + 子节点，按设备规模拆分接入，配置仍由根节点统一管理。',
    icon: '<svg viewBox="0 0 20 20" fill="none"><rect x="7" y="2.5" width="6" height="4" rx="1" stroke="currentColor" stroke-width="1.5"/><rect x="2.5" y="13.5" width="6" height="4" rx="1" stroke="currentColor" stroke-width="1.5"/><rect x="11.5" y="13.5" width="6" height="4" rx="1" stroke="currentColor" stroke-width="1.5"/><path d="M10 6.5v3.5M5.5 13.5v-3.5h9v3.5" stroke="currentColor" stroke-width="1.5"/></svg>',
  },
]

const authority = [
  { title: '单一权威源', desc: '配置只在 Manager 修改，设备端不提供本地改配置入口。' },
  { title: '离线只读', desc: '断网期间设备沿用最后一次成功下发的快照，恢复后自动对齐。' },
  { title: '按节点拆分', desc: '根节点与子节点独立编排部署，Manager 与 discovery 不必同机。' },
]
</script>
