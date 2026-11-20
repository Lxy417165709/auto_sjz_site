<template>
  <section class="mt-2 rounded-3xl border border-slate-200 bg-white p-5 shadow-card sm:p-8">
    <div class="flex flex-wrap items-center justify-between gap-3">
      <span class="eyebrow">部署拓扑</span>
      <div class="flex flex-wrap items-center gap-4 text-[11px] text-slate-400">
        <span class="inline-flex items-center gap-1.5"><span class="h-px w-4 bg-slate-300"></span>隶属 / 注册</span>
        <span class="inline-flex items-center gap-1.5"><span class="h-px w-4 border-t border-dashed border-slate-300"></span>旁挂服务</span>
      </div>
    </div>

    <div class="mt-4 flex flex-wrap items-center gap-x-3 gap-y-2 rounded-xl border border-dashed border-slate-300 bg-slate-50/70 px-4 py-2.5 text-[12px] text-slate-600">
      <span class="tag">旁挂服务</span>
      <span><span class="font-medium text-slate-900">Auth</span> 签发授权租约</span>
      <span class="text-slate-300">·</span>
      <span><span class="font-medium text-slate-900">discovery</span> 解析 Manager 地址 <span class="font-mono text-slate-500">5003</span></span>
      <span class="text-[11px] text-slate-400">不承载配置与业务数据</span>
    </div>

    <div class="mt-6 flex justify-center">
      <div class="w-full max-w-md rounded-2xl bg-ink-900 px-5 py-4 text-white">
        <div class="flex items-start justify-between gap-3">
          <div>
            <p class="text-[11px] font-semibold uppercase tracking-[0.18em] text-accent-400">根节点</p>
            <h3 class="mt-1 text-sm font-semibold">Manager 根节点</h3>
          </div>
          <span class="rounded-full border border-white/15 px-2 py-0.5 font-mono text-[11px] text-slate-300">5001</span>
        </div>
        <p class="mt-2 text-[12px] leading-5 text-slate-400">配置唯一权威源：节点管理、设备注册、快照下发。</p>
      </div>
    </div>

    <div class="mx-auto hidden h-6 w-px bg-slate-300 sm:block"></div>

    <div class="ml-4 mt-4 grid gap-3 border-l border-slate-200 pl-4 sm:ml-0 sm:mt-0 sm:grid-cols-2 sm:gap-0 sm:border-0 sm:pl-0">
      <div v-for="(node, i) in nodes" :key="node.code" class="relative sm:px-3 sm:pt-8">
        <span class="hidden h-8 w-px bg-slate-300 sm:absolute sm:left-1/2 sm:top-0 sm:block"></span>
        <span
          class="hidden h-px bg-slate-300 sm:absolute sm:top-0 sm:block"
          :class="i === 0 ? 'sm:left-1/2 sm:right-0' : 'sm:left-0 sm:right-1/2'"
        ></span>

        <div class="rounded-2xl border border-slate-200 bg-slate-50/50 p-4">
          <div class="flex items-start justify-between gap-3">
            <div>
              <p class="text-[11px] font-semibold text-brand-600">子节点 {{ node.code }}</p>
              <h3 class="mt-0.5 text-sm font-semibold text-slate-900">Manager 子节点</h3>
            </div>
            <span class="tag font-mono">{{ node.port }}</span>
          </div>
          <p class="mt-2 text-[12px] leading-5 text-slate-500">{{ node.group }}</p>

          <div class="mt-3 flex flex-wrap gap-1.5">
            <span
              v-for="device in node.devices"
              :key="device"
              class="inline-flex items-center gap-1.5 rounded-lg border border-slate-200 bg-white px-2 py-1 text-[11px] text-slate-600"
            >
              <span class="h-1.5 w-1.5 rounded-full bg-accent-500"></span>{{ device }}
            </span>
            <span class="inline-flex items-center rounded-lg border border-dashed border-slate-200 px-2 py-1 text-[11px] text-slate-400">按需扩展</span>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-5 flex flex-wrap items-center gap-x-2 gap-y-1 border-t border-slate-100 pt-4 text-[12px] text-slate-500">
      <span class="font-medium text-slate-700">设备接入路径</span>
      <span>设备启动</span>
      <span class="text-slate-300">→</span>
      <span>discovery 解析地址</span>
      <span class="text-slate-300">→</span>
      <span>向所属 Manager 节点注册</span>
      <span class="text-slate-300">→</span>
      <span>拉取快照并执行</span>
    </div>
  </section>
</template>

<script setup lang="ts">
const nodes = [
  {
    code: 'A',
    port: '5102',
    group: '门店 / 批次 A 的设备与账号，独立编排部署。',
    devices: ['设备 01', '设备 02', '设备 03'],
  },
  {
    code: 'B',
    port: '5202',
    group: '门店 / 批次 B 的设备与账号，按规模继续加节点。',
    devices: ['设备 04', '设备 05', '设备 06'],
  },
]
</script>
