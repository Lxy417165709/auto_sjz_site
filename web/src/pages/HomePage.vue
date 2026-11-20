<template>
  <div>
    <!-- Hero -->
    <section class="hero-bg relative overflow-hidden">
      <div class="hero-grid" aria-hidden="true"></div>
      <div class="hero-beam hero-beam-a" aria-hidden="true"></div>
      <div class="hero-beam hero-beam-b" aria-hidden="true"></div>
      <div class="hero-flare" aria-hidden="true"></div>

      <div class="wrap relative pt-16 text-center sm:pt-20 lg:pt-24">
        <div class="mx-auto inline-flex items-center gap-2 rounded-full border border-[#c9d8ff] bg-white/80 px-3.5 py-1.5 text-[11px] tracking-[0.14em] text-brand-700 shadow-[0_8px_24px_-18px_rgba(36,71,214,0.45)] backdrop-blur">
          <span class="h-1.5 w-1.5 rounded-full bg-brand-500 shadow-[0_0_10px_rgba(55,99,245,0.65)]"></span>
          三角洲业务自动化 · 按号收费
        </div>

        <h1 class="mx-auto mt-7 max-w-5xl text-[38px] font-semibold leading-[1.04] tracking-[-0.045em] text-[#07152e] sm:text-[54px] lg:text-[68px]">
          设备不再按台收费<br />
          <span class="hero-title-fire">一个号 3 元 / 月</span>
        </h1>

        <p class="mx-auto mt-6 max-w-[640px] text-[14px] leading-8 text-[#5f6c82] sm:text-[15px]">
          炎火为三角洲业务提供一套管理中台，统一管住设备与账号：按账号结算、批量下发配置、异常自动通知。设备多少不额外收钱，账号增减才计费。
        </p>

        <div class="mt-9 flex flex-wrap items-center justify-center gap-3">
          <router-link to="/start" class="btn-primary h-11 px-6 text-[15px]">查看接入流程</router-link>
          <router-link to="/console" class="btn-secondary h-11 px-5 text-[15px]">打开管理台演示 →</router-link>
        </div>

        <div class="mt-7 flex flex-wrap items-center justify-center gap-x-5 gap-y-2 font-mono text-[11px] tracking-wide text-[#8490a5]">
          <span>¥3 / 号 / 月</span>
          <span class="hidden h-px w-4 bg-[#cfd9eb] sm:block"></span>
          <span>设备费 ¥0</span>
          <span class="hidden h-px w-4 bg-[#cfd9eb] sm:block"></span>
          <span>私有化部署</span>
        </div>
      </div>

      <div class="wrap relative pb-2 pt-12 sm:pt-16">
        <div class="relative mx-auto max-w-[1040px]">
          <div class="pointer-events-none absolute inset-x-16 -bottom-6 h-20 rounded-full bg-[#5c86ff]/30 blur-3xl" aria-hidden="true"></div>
          <div class="relative overflow-hidden rounded-[18px] border border-[#d7e2f5] bg-white shadow-[0_44px_100px_-46px_rgba(36,71,214,0.38)]">
            <div class="flex h-10 items-center justify-between border-b border-[#e7edf8] bg-[#fbfcff] px-4">
              <div class="flex items-center gap-1.5" aria-hidden="true">
                <span class="h-2 w-2 rounded-full bg-[#c8d4e7]"></span>
                <span class="h-2 w-2 rounded-full bg-[#cfd9eb]"></span>
                <span class="h-2 w-2 rounded-full bg-[#e3e9f3]"></span>
              </div>
              <p class="font-mono text-[10px] uppercase tracking-[0.16em] text-[#8490a5]">Manager / 业务状态</p>
              <span class="inline-flex items-center gap-1.5 text-[10px] text-[#70809a]">
                <span class="h-1.5 w-1.5 rounded-full bg-brand-500"></span>
                演示数据
              </span>
            </div>
            <img
              src="/console-preview-blue.png"
              alt="炎火 Manager 业务状态界面，按设备聚合账号数据"
              class="block w-full"
              width="2272"
              height="1120"
            />
          </div>
          <div class="mt-5 flex flex-col items-center justify-between gap-2 text-[11px] text-[#8490a5] sm:flex-row">
            <p>真实管理台界面，账号与数值为演示数据</p>
            <router-link to="/console" class="text-brand-600 transition hover:text-brand-700">打开可交互演示 →</router-link>
          </div>
        </div>
      </div>
    </section>

    <!-- Stats -->
    <section class="border-b border-[#dfe6f3] bg-white">
      <div class="wrap grid grid-cols-2 sm:grid-cols-4">
        <div
          v-for="(stat, index) in stats"
          :key="stat.label"
          class="py-8 sm:py-9"
          :class="[
            index > 0 ? 'sm:border-l sm:border-[#dfe6f3] sm:pl-8' : '',
            index % 2 === 1 ? 'border-l border-[#dfe6f3] pl-5 sm:pl-8' : '',
          ]"
        >
          <p class="text-[28px] font-semibold leading-none tracking-[-0.04em] text-[#07152e] sm:text-[34px]">{{ stat.value }}</p>
          <p class="mt-2 text-[12px] text-[#77839a]">{{ stat.label }}</p>
        </div>
      </div>
    </section>

    <!-- Capabilities -->
    <section class="bg-[#f8faff] py-20 sm:py-28">
      <div class="wrap">
        <div class="max-w-3xl">
          <p class="eyebrow">业务能力</p>
          <h2 class="mt-3 text-[32px] font-semibold leading-[1.1] tracking-[-0.035em] text-[#07152e] sm:text-[44px]">
            管设备只是手段，<br class="hidden sm:block" />
            <span class="text-brand-600">管住号才是目的</span>
          </h2>
          <p class="mt-5 max-w-xl text-[15px] leading-7 text-[#5f6c82]">
            计费、配置、授权都围绕账号展开。设备只是执行载体，不再决定你每个月付多少钱。
          </p>
        </div>

        <div class="mt-12 grid gap-4 lg:grid-cols-6">
          <article class="relative overflow-hidden rounded-[24px] bg-[#2447d6] p-8 text-white lg:col-span-4 sm:p-10">
            <div class="pointer-events-none absolute -right-24 -top-32 h-[360px] w-[360px] rounded-full bg-[#5c86ff]/30 blur-3xl" aria-hidden="true"></div>
            <div class="relative">
              <p class="font-mono text-[10px] uppercase tracking-[0.18em] text-[#bfd0ff]">按号计费</p>
              <h3 class="mt-4 max-w-xl text-[26px] font-semibold leading-[1.15] tracking-[-0.035em] sm:text-[34px]">设备不收费，账号才计费</h3>
              <p class="mt-4 max-w-lg text-[14px] leading-7 text-white/70">一台设备挂 5 个号，就是 5 份授权。加设备不加钱，账号加一个算一个，减一个省一个。</p>
              <div class="mt-9 flex items-end gap-2">
                <span class="font-mono text-[54px] font-medium leading-none tracking-[-0.06em]">¥3</span>
                <span class="mb-1 text-[13px] text-white/60">/ 号 / 月</span>
              </div>
              <div class="mt-8 grid grid-cols-1 gap-2 border-t border-white/15 pt-5 text-[11px] text-white/55 sm:grid-cols-3">
                <div class="flex items-center justify-between gap-2 sm:block"><span>设备 10 台</span><span class="ml-2 font-mono text-white sm:ml-0 sm:mt-1 sm:block">¥0</span></div>
                <div class="flex items-center justify-between gap-2 sm:block"><span>账号 50 个</span><span class="ml-2 font-mono text-white sm:ml-0 sm:mt-1 sm:block">¥150</span></div>
                <div class="flex items-center justify-between gap-2 sm:block"><span>月费合计</span><span class="ml-2 font-mono text-accent-300 sm:ml-0 sm:mt-1 sm:block">¥150</span></div>
              </div>
            </div>
          </article>

          <article class="rounded-[24px] border border-[#dfe6f3] bg-white p-7 lg:col-span-2">
            <p class="font-mono text-[10px] uppercase tracking-[0.18em] text-[#77839a]">统一配置</p>
            <h3 class="mt-3 text-[21px] font-semibold tracking-[-0.02em] text-[#07152e]">只维护一份配置</h3>
            <p class="mt-2 text-[13px] leading-6 text-[#5f6c82]">变更在中台改一次，批量下发到设备。离线设备只读，恢复后自动对齐。</p>
            <div class="mt-7 rounded-xl border border-[#e2e9f7] bg-[#f5f8ff] p-3 font-mono text-[11px]">
              <div class="flex items-center justify-between border-b border-[#e2e9f7] pb-2"><span class="text-[#7b87a0]">采集频率</span><span class="text-[#15213a]">30s</span></div>
              <div class="flex items-center justify-between border-b border-[#e2e9f7] py-2"><span class="text-[#7b87a0]">生产队列</span><span class="text-[#15213a]">3 / 3</span></div>
              <div class="flex items-center justify-between pt-2"><span class="text-[#7b87a0]">同步状态</span><span class="text-brand-600">已下发</span></div>
            </div>
          </article>

          <article v-for="item in capabilities" :key="item.title" class="rounded-[24px] border border-[#dfe6f3] bg-white p-7 lg:col-span-2">
            <div class="flex h-10 w-10 items-center justify-center rounded-xl bg-brand-50 text-brand-600">
              <span v-html="item.icon" class="[&>svg]:h-5 [&>svg]:w-5"></span>
            </div>
            <h3 class="mt-5 text-[17px] font-semibold tracking-[-0.02em] text-[#07152e]">{{ item.title }}</h3>
            <p class="mt-2 text-[13px] leading-6 text-[#5f6c82]">{{ item.desc }}</p>
          </article>
        </div>
      </div>
    </section>

    <!-- Steps -->
    <section class="border-y border-[#dfe6f3] bg-[#f3f7ff] py-20 sm:py-24">
      <div class="wrap">
        <div class="flex flex-col gap-4 sm:flex-row sm:items-end sm:justify-between">
          <div>
            <p class="eyebrow">接入流程</p>
            <h2 class="mt-3 text-[32px] font-semibold leading-[1.1] tracking-[-0.035em] text-[#07152e] sm:text-[42px]">从申请到运行，五步走完</h2>
          </div>
          <router-link to="/start" class="text-sm font-medium text-brand-600 transition hover:text-brand-700">看完整接力图 →</router-link>
        </div>

        <div class="mt-12 grid gap-x-8 gap-y-10 sm:grid-cols-2 lg:grid-cols-5">
          <router-link v-for="step in startSteps" :key="step.no" to="/start" class="group relative border-t border-[#c9d7eb] pt-6">
            <span class="absolute -top-[5px] left-0 h-2.5 w-2.5 rounded-full" :class="step.owner === '我方' ? 'bg-accent-500' : 'bg-brand-500'"></span>
            <div class="flex items-center justify-between">
              <span class="font-mono text-[11px] text-[#77839a]">{{ step.no }}</span>
              <span class="text-[10px] tracking-wide" :class="step.owner === '我方' ? 'text-accent-600' : 'text-brand-600'">{{ step.owner }}</span>
            </div>
            <h3 class="mt-5 text-[16px] font-semibold tracking-[-0.01em] text-[#15213a] group-hover:text-brand-600">{{ step.title }}</h3>
            <p class="mt-2 text-[12px] leading-6 text-[#77839a]">{{ step.desc }}</p>
          </router-link>
        </div>
      </div>
    </section>

    <!-- Pricing -->
    <section class="bg-[#f8faff] py-20 sm:py-28">
      <div class="wrap">
        <div class="overflow-hidden rounded-[28px] border border-[#dfe6f3] bg-white shadow-[0_30px_80px_-56px_rgba(36,71,214,0.4)] lg:grid lg:grid-cols-[1.08fr_0.92fr]">
          <div class="relative overflow-hidden bg-[#eef3ff] p-8 text-[#07152e] sm:p-12">
            <div class="pointer-events-none absolute -bottom-40 -left-20 h-[420px] w-[420px] rounded-full bg-[#5c86ff]/30 blur-3xl" aria-hidden="true"></div>
            <div class="relative">
              <p class="font-mono text-[10px] uppercase tracking-[0.18em] text-brand-600">算一笔账</p>
              <h2 class="mt-4 max-w-md text-[26px] font-semibold leading-[1.15] tracking-[-0.03em] sm:text-[34px]">10 台设备、每台 5 个号，一个月 150 元</h2>
              <p class="mt-12 font-mono text-[11px] tracking-wide text-[#8490a5]">10 台 × 5 个号 × ¥3</p>
              <div class="mt-3 flex items-end gap-2">
                <span class="text-[60px] font-semibold leading-none tracking-[-0.05em] text-brand-700 sm:text-[72px]">150</span>
                <span class="mb-2 text-[13px] text-[#5f6c82]">元 / 月</span>
              </div>
              <div class="mt-8 space-y-2 border-t border-[#cfd9e8] pt-5 text-[12px] text-[#5f6c82]">
                <div class="flex items-center justify-between"><span>设备费</span><span class="font-mono text-[#07152e]">0 元</span></div>
                <div class="flex items-center justify-between"><span>账号费</span><span class="font-mono text-[#07152e]">150 元</span></div>
                <div class="flex items-center justify-between"><span>设备闲置</span><span class="text-accent-600">不影响费用</span></div>
              </div>
              <router-link to="/pricing" class="mt-8 inline-flex items-center gap-1 text-sm font-medium text-brand-600 transition hover:text-brand-700">用计算器算你的月费 →</router-link>
            </div>
          </div>

          <div class="p-8 sm:p-12">
            <p class="font-mono text-[10px] uppercase tracking-[0.18em] text-[#77839a]">适用场景</p>
            <ul class="mt-6 divide-y divide-[#dfe6f3]">
              <li v-for="scene in scenes" :key="scene.title" class="py-5 first:pt-0 last:pb-0">
                <p class="text-[15px] font-semibold tracking-[-0.01em] text-[#07152e]">{{ scene.title }}</p>
                <p class="mt-1.5 text-[13px] leading-6 text-[#5f6c82]">{{ scene.desc }}</p>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </section>

    <!-- Engineering -->
    <section class="border-t border-[#dfe6f3] bg-white py-14">
      <div class="wrap grid gap-8 lg:grid-cols-[1fr_auto] lg:items-center">
        <div>
          <p class="font-mono text-[10px] uppercase tracking-[0.18em] text-[#77839a]">For engineers</p>
          <h2 class="mt-3 text-[22px] font-semibold tracking-[-0.025em] text-[#07152e]">架构、端口、授权链路、部署形态</h2>
          <p class="mt-3 max-w-2xl text-[13px] leading-7 text-[#5f6c82]">
            设备如何找到 Manager、快照怎么落盘、授权如何本地验签、根节点与子节点怎么拆，这些内容集中在技术架构与部署接入里，不和业务介绍混在一起。
          </p>
          <div class="mt-5 flex flex-wrap gap-2 font-mono text-[10px] text-[#77839a]">
            <span class="rounded-md border border-[#dfe6f3] bg-[#f5f8ff] px-2 py-1">Manager</span>
            <span class="rounded-md border border-[#dfe6f3] bg-[#f5f8ff] px-2 py-1">discovery</span>
            <span class="rounded-md border border-[#dfe6f3] bg-[#f5f8ff] px-2 py-1">Docker</span>
            <span class="rounded-md border border-[#dfe6f3] bg-[#f5f8ff] px-2 py-1">license</span>
          </div>
        </div>
        <div class="flex flex-wrap gap-3">
          <router-link to="/tech" class="btn-primary">进入技术架构</router-link>
          <router-link to="/deploy" class="btn-secondary">部署接入</router-link>
        </div>
      </div>
    </section>

    <div class="wrap">
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

const stats = [
  { value: '¥3', label: '每个账号 / 月' },
  { value: '¥0', label: '设备授权费' },
  { value: '4 类', label: '业务模块' },
  { value: '1 套', label: '管理中台' },
]

const capabilities = [
  {
    title: '异常先通知，不等你发现',
    desc: '设备在线状态、执行异常、授权到期集中监控，触发后按配置的通知渠道提醒。',
    icon: '<svg viewBox="0 0 20 20" fill="none"><path d="M6 8a4 4 0 0 1 8 0c0 3 1 4 1.5 4.5H4.5C5 12 6 11 6 8Z" stroke="currentColor" stroke-width="1.5"/><path d="M8.5 15a1.5 1.5 0 0 0 3 0" stroke="currentColor" stroke-width="1.5"/></svg>',
  },
  {
    title: '到期不删号，也不登出',
    desc: '授权到期只停执行，账号与设备配置原样保留。续费后接着跑，不用重建环境。',
    icon: '<svg viewBox="0 0 20 20" fill="none"><rect x="4" y="8" width="12" height="9" rx="2" stroke="currentColor" stroke-width="1.5"/><path d="M7 8V6a3 3 0 1 1 6 0v2M10 12v2" stroke="currentColor" stroke-width="1.5"/></svg>',
  },
  {
    title: '私有化部署，数据自持',
    desc: 'Docker 私有化部署，Manager 与 discovery 可以分开，账号与业务数据留在自己的机器上。',
    icon: '<svg viewBox="0 0 20 20" fill="none"><rect x="3" y="4" width="14" height="5" rx="1.5" stroke="currentColor" stroke-width="1.5"/><rect x="3" y="11" width="14" height="5" rx="1.5" stroke="currentColor" stroke-width="1.5"/><circle cx="6.5" cy="6.5" r="1" fill="currentColor"/><circle cx="6.5" cy="13.5" r="1" fill="currentColor"/></svg>',
  },
]

const startSteps = [
  { no: '01', owner: '接入方', title: '申请管理中台', desc: '提交设备台数、账号数量与使用场景。' },
  { no: '02', owner: '我方', title: '部署管理中台', desc: '部署 Manager 与 discovery，交付登录地址。' },
  { no: '03', owner: '接入方', title: '配置并申请授权', desc: '登录配置设备、账号、模块，提交授权申请。' },
  { no: '04', owner: '我方', title: '发放令牌', desc: '按账号发放授权令牌，确定账号有效期。' },
  { no: '05', owner: '接入方', title: '设备运行', desc: '设备拉取配置与令牌，开始执行。' },
]

const scenes = [
  { title: '多账号运营团队', desc: '一台设备挂多个号，按号结算不会为设备容量多付钱。' },
  { title: '设备规模持续扩张', desc: '加设备不加授权费，按实际账号数扩展，预算可控。' },
  { title: '需要长期稳定执行', desc: '到期不删号、离线按同一时间失效，续费即恢复，不用重建环境。' },
]
</script>

<style scoped>
.hero-bg {
  background:
    radial-gradient(62% 42% at 50% 0%, rgba(55, 99, 245, 0.14) 0%, rgba(55, 99, 245, 0) 70%),
    radial-gradient(34% 34% at 84% 18%, rgba(18, 184, 146, 0.08) 0%, rgba(18, 184, 146, 0) 72%),
    linear-gradient(180deg, #f5f9ff 0%, #ffffff 72%);
}

.hero-grid {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(to right, rgba(36, 71, 214, 0.055) 1px, transparent 1px),
    linear-gradient(to bottom, rgba(36, 71, 214, 0.055) 1px, transparent 1px);
  background-size: 48px 48px;
  -webkit-mask-image: radial-gradient(62% 58% at 50% 18%, #000 0%, rgba(0, 0, 0, 0.55) 52%, transparent 82%);
  mask-image: radial-gradient(62% 58% at 50% 18%, #000 0%, rgba(0, 0, 0, 0.55) 52%, transparent 82%);
  pointer-events: none;
}

.hero-beam {
  position: absolute;
  top: -28%;
  height: 150%;
  width: 34%;
  background: linear-gradient(180deg, rgba(55, 99, 245, 0) 0%, rgba(55, 99, 245, 0.12) 45%, rgba(55, 99, 245, 0) 100%);
  filter: blur(64px);
  transform: rotate(20deg);
  pointer-events: none;
}
.hero-beam-a {
  left: 8%;
}
.hero-beam-b {
  right: 4%;
  width: 26%;
  background: linear-gradient(180deg, rgba(18, 184, 146, 0) 0%, rgba(18, 184, 146, 0.09) 48%, rgba(18, 184, 146, 0) 100%);
  transform: rotate(-20deg);
}

.hero-flare {
  position: absolute;
  left: 50%;
  bottom: -20%;
  width: 120%;
  height: 70%;
  transform: translateX(-50%);
  background: radial-gradient(45% 50% at 50% 50%, rgba(55, 99, 245, 0.16) 0%, rgba(55, 99, 245, 0) 70%);
  filter: blur(70px);
  pointer-events: none;
}

.hero-title-fire {
  background: linear-gradient(100deg, #2447d6 0%, #6f9bff 100%);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
}
</style>
