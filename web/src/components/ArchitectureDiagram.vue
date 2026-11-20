<template>
  <section class="mt-4">
    <div class="flex flex-wrap items-center gap-3 text-xs text-slate-500">
      <span class="eyebrow">系统关系</span>
      <span class="hidden h-px flex-1 bg-slate-200 sm:block"></span>
      <span>设备不写死 IP，地址变更只改 discovery 登记</span>
    </div>

    <div class="mt-5 grid gap-4 lg:grid-cols-5">
      <div
        v-for="(step, index) in steps"
        :key="step.title"
        class="relative rounded-2xl border border-slate-200 bg-white p-4 shadow-card"
      >
        <div class="flex items-center justify-between">
          <span class="font-mono text-[11px] text-slate-400">0{{ index + 1 }}</span>
          <span v-if="step.port" class="tag font-mono">{{ step.port }}</span>
        </div>
        <p class="mt-2 text-sm font-semibold text-slate-900">{{ step.title }}</p>
        <p class="mt-1 text-xs leading-5 text-slate-500">{{ step.desc }}</p>
      </div>
    </div>

    <div class="mt-3 flex flex-wrap items-center gap-x-6 gap-y-2 rounded-2xl border border-dashed border-slate-300 bg-slate-50/70 px-4 py-3 text-xs text-slate-600">
      <span class="font-medium text-slate-800">授权侧</span>
      <span>Auth 签发 license lease → Manager 下发 → 设备本地验签 → 按绝对到期时间执行</span>
    </div>
  </section>
</template>

<script setup lang="ts">
const steps = [
  { title: '设备启动', desc: '读取内置的 discovery 地址，准备注册。', port: '' },
  { title: 'discovery 解析', desc: '返回 Manager 对外注册的地址，设备不写死 IP。', port: '5003' },
  { title: '拉取快照', desc: '向 Manager 注册并获取配置快照，子节点示例端口 5102、5202。', port: '5001' },
  { title: '写入本地 JSON', desc: '快照落盘，离线期间只读，不接受本地改配置。', port: '' },
  { title: '执行业务', desc: '按快照与授权租约执行，到期按同一时间点失效。', port: '' },
]
</script>
