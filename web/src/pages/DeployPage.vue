<template>
  <div>
    <PageHero
      eyebrow="部署接入"
      title="Docker 起步，一台机器跑通 discovery 与 Manager"
      desc="根节点与子节点分开编排，Manager 和 discovery 不要求同机。只要 Manager 能连 discovery、设备能连 Manager 注册出去的地址，链路就成立。"
    >
      <div class="flex flex-wrap gap-2">
        <span class="tag font-mono">discovery :5003</span>
        <span class="tag font-mono">Manager Root :5001</span>
        <span class="tag font-mono">子节点 :5102 / :5202</span>
      </div>
    </PageHero>

    <div class="wrap">
      <section class="py-14 sm:py-16">
        <p class="eyebrow">部署顺序</p>
        <h2 class="h2 mt-2">四步完成一次接入</h2>
        <div class="mt-8 space-y-4">
          <article v-for="(step, index) in steps" :key="step.title" class="relative rounded-2xl border border-slate-200 bg-white p-6 shadow-card sm:pl-20">
            <span class="absolute left-6 top-6 hidden h-9 w-9 items-center justify-center rounded-full bg-ink-900 font-mono text-[12px] text-white sm:flex">
              {{ String(index + 1).padStart(2, '0') }}
            </span>
            <h3 class="text-sm font-semibold text-slate-900">{{ step.title }}</h3>
            <p class="mt-2 text-[13px] leading-6 text-slate-500">{{ step.desc }}</p>
            <div class="mt-3 flex flex-wrap gap-2">
              <span v-for="tag in step.tags" :key="tag" class="tag font-mono">{{ tag }}</span>
            </div>
          </article>
        </div>
      </section>

      <section class="grid gap-4 lg:grid-cols-[1fr_1fr]">
        <div class="rounded-3xl border border-slate-200 bg-slate-50/60 p-7">
          <p class="eyebrow">端口规划</p>
          <h2 class="mt-2 text-xl font-semibold text-slate-900">默认端口一览</h2>
          <table class="mt-4 w-full text-left text-sm">
            <thead>
              <tr class="text-xs uppercase tracking-[0.12em] text-slate-500">
                <th class="py-2 font-medium">组件</th>
                <th class="py-2 font-medium">端口</th>
                <th class="py-2 font-medium">说明</th>
            </tr>
            </thead>
            <tbody>
              <tr v-for="row in ports" :key="row.port" class="border-t border-slate-200">
                <td class="py-3 font-medium text-slate-800">{{ row.component }}</td>
                <td class="py-3 font-mono text-brand-700">{{ row.port }}</td>
                <td class="py-3 text-[13px] text-slate-500">{{ row.note }}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <div class="rounded-3xl border border-slate-200 p-7">
          <p class="eyebrow">网络要求</p>
          <h2 class="mt-2 text-xl font-semibold text-slate-900">三条链路必须通</h2>
          <ul class="mt-5 space-y-4">
            <li v-for="item in network" :key="item.title" class="flex gap-3">
              <span class="mt-1 flex h-5 w-5 flex-none items-center justify-center rounded-full bg-accent-500/10 text-accent-600">
                <svg class="h-3 w-3" viewBox="0 0 12 12" fill="none" aria-hidden="true">
                  <path d="M2 6.5l2.5 2.5 5.5-6" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" />
                </svg>
              </span>
              <div>
                <p class="text-sm font-medium text-slate-900">{{ item.title }}</p>
                <p class="mt-1 text-[13px] leading-6 text-slate-500">{{ item.desc }}</p>
              </div>
            </li>
          </ul>
        </div>
      </section>

      <section class="mt-4 grid gap-4 lg:grid-cols-3">
        <div v-for="mode in modes" :key="mode.title" class="rounded-2xl border border-slate-200 p-6">
          <span class="tag">{{ mode.badge }}</span>
          <h3 class="mt-3 text-sm font-semibold text-slate-900">{{ mode.title }}</h3>
          <p class="mt-2 text-[13px] leading-6 text-slate-500">{{ mode.desc }}</p>
        </div>
      </section>

      <section class="mt-10 rounded-3xl border border-slate-200 bg-ink-900 p-7 text-slate-300 sm:p-9">
        <p class="text-[11px] font-semibold uppercase tracking-[0.22em] text-accent-400">上线前检查</p>
        <h2 class="mt-2 text-xl font-semibold text-white">设备注册不上，先看这三处</h2>
        <div class="mt-5 grid gap-4 sm:grid-cols-3">
          <div v-for="check in checks" :key="check.title">
            <p class="text-sm font-medium text-white">{{ check.title }}</p>
            <p class="mt-1 text-[13px] leading-6 text-slate-400">{{ check.desc }}</p>
          </div>
        </div>
      </section>

      <CtaBand
        title="官网与业务系统分开部署，互不影响"
        desc="官网是独立的 Go + Vue 静态站点，不接 Manager 数据库与订单；升级官网不会影响设备接入与业务执行。"
        primary-label="查看部署文档"
        primary-to="/docs"
        secondary-label="联系技术支持"
        secondary-to="/contact"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import PageHero from '@/components/PageHero.vue'
import CtaBand from '@/components/CtaBand.vue'

const steps = [
  { title: '启动 discovery', desc: '先起地址解析服务，确认监听正常，再继续后续步骤。', tags: ['docker compose', ':5003'] },
  { title: '启动 Manager 根节点', desc: '根节点负责配置、注册与快照下发，默认 5001；确认管理后台可访问。', tags: ['root', ':5001'] },
  { title: '按规模加子节点', desc: '设备变多时增加子节点，各自独立编排，示例端口 5102、5202。', tags: ['child', ':5102', ':5202'] },
  { title: '设备接入', desc: '设备只配置 discovery 地址，由 discovery 返回 Manager 地址，再注册拉快照。', tags: ['device', 'auto discover'] },
]

const ports = [
  { component: 'Manager 根节点', port: '5001', note: '配置、注册、快照下发入口' },
  { component: 'Manager 子节点', port: '5102 / 5202', note: '按节点规划递增，独立编排' },
  { component: 'discovery', port: '5003', note: '只做 Manager 地址登记与解析' },
  { component: '官网（本仓库）', port: '8080', note: '独立部署，不参与业务链路' },
]

const network = [
  { title: 'Manager → discovery', desc: 'Manager 启动时向 discovery 登记对外地址，网络必须放通。' },
  { title: '设备 → discovery', desc: '设备只保存 discovery 地址，变更 Manager 机房时无需逐台修改。' },
  { title: '设备 → Manager 注册地址', desc: '注册与快照下发走这条链路，需保证设备侧可达登记出去的地址。' },
]

const modes = [
  { badge: '起步', title: '单机部署', desc: '一台机器同时跑 discovery 与 Manager 根节点，适合试点与小规模设备。' },
  { badge: '拆分', title: 'Manager 与 discovery 分机', desc: '两者网络互通即可，便于按安全域或机房拆分。' },
  { badge: '扩展', title: '根节点 + 多子节点', desc: '设备按子节点分流，配置与账号仍在 Manager 统一管理。' },
]

const checks = [
  { title: 'discovery 解析结果', desc: '设备拿到的地址是否为设备侧可达的地址，而不是容器内网地址。' },
  { title: '端口与安全组', desc: 'Manager 注册端口、discovery 端口是否对设备网段放行。' },
  { title: '快照是否落盘', desc: '设备本地 JSON 是否写入成功，失败时是否保留上一份快照。' },
]
</script>
