<template>
  <div>
    <PageHero
      eyebrow="技术架构"
      title="设备、discovery、Manager、Auth，各管一段"
      desc="技术细节集中在这里：链路、模块职责、授权验签、端口与部署形态。业务价值与计费规则在「业务能力」「按号收费」里，两边不混着讲。"
    >
      <div class="flex flex-wrap gap-2">
        <span class="tag">Go + Vue</span>
        <span class="tag">Docker</span>
        <span class="tag">discovery</span>
        <span class="tag">license lease</span>
      </div>
    </PageHero>

    <div class="wrap">
      <section class="py-14 sm:py-16">
        <div class="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
          <div>
            <p class="eyebrow">部署拓扑</p>
            <h2 class="h2 mt-2">一棵 Manager 树，设备挂在节点下</h2>
          </div>
          <router-link to="/deploy" class="text-sm font-medium text-brand-600 hover:text-brand-700">看部署接入 →</router-link>
        </div>
        <p class="lead mt-4 max-w-3xl">
          根节点是配置唯一权威源，负责节点管理与快照下发；子节点按门店或设备规模拆分，设备只与所属节点注册、拉快照。设备不写死 IP，启动后先用 discovery 解析地址。
        </p>
        <ManagerTree class="mt-6" />
      </section>

      <section class="rounded-3xl border border-slate-200 bg-slate-50/60 px-5 py-10 sm:px-9">
        <p class="eyebrow">模块职责</p>
        <h2 class="h2 mt-2">谁负责什么，排障时先定位环节</h2>
        <div class="mt-6 overflow-x-auto">
          <table class="w-full min-w-[760px] border-separate border-spacing-y-2 text-left text-sm">
            <thead>
              <tr class="text-xs uppercase tracking-[0.12em] text-slate-500">
                <th class="px-4 py-2 font-medium">模块</th>
                <th class="px-4 py-2 font-medium">负责</th>
                <th class="px-4 py-2 font-medium">不做</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="row in modules" :key="row.name" class="bg-white text-slate-700">
                <td class="rounded-l-xl px-4 py-3 font-medium text-slate-900">{{ row.name }}</td>
                <td class="px-4 py-3 text-slate-600">{{ row.owns }}</td>
                <td class="rounded-r-xl px-4 py-3 text-slate-500">{{ row.avoids }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>

      <section class="grid gap-4 py-14 sm:py-16 lg:grid-cols-[1fr_1fr]">
        <div class="rounded-3xl border border-slate-200 p-7">
          <p class="eyebrow">授权链路</p>
          <h2 class="mt-2 text-xl font-semibold text-slate-900">Auth 签发，设备本地验签</h2>
          <ol class="mt-5 space-y-4">
            <li v-for="(step, i) in licenseFlow" :key="step.title" class="flex gap-3">
              <span class="mt-0.5 flex h-6 w-6 flex-none items-center justify-center rounded-full bg-brand-50 font-mono text-[11px] text-brand-600">{{ i + 1 }}</span>
              <div>
                <p class="text-sm font-medium text-slate-900">{{ step.title }}</p>
                <p class="mt-1 text-[13px] leading-6 text-slate-500">{{ step.desc }}</p>
              </div>
            </li>
          </ol>
        </div>

        <div class="rounded-3xl border border-slate-200 bg-ink-900 p-7 text-slate-200">
          <p class="text-[11px] font-semibold uppercase tracking-[0.22em] text-accent-400">端口</p>
          <h2 class="mt-2 text-xl font-semibold text-white">默认端口规划</h2>
          <dl class="mt-5 space-y-3 font-mono text-sm">
            <div v-for="row in ports" :key="row.port" class="flex items-center justify-between border-b border-white/10 pb-3">
              <dt class="text-slate-400">{{ row.component }}</dt>
              <dd class="text-white">{{ row.port }}</dd>
            </div>
          </dl>
          <p class="mt-4 text-[12px] leading-5 text-slate-400">
            Manager 与 discovery 不要求同机，只要 Manager 能连 discovery、设备能连 Manager 注册出去的地址。
          </p>
        </div>
      </section>

      <section>
        <p class="eyebrow">技术栈</p>
        <h2 class="h2 mt-2">用的都是能自己维护的东西</h2>
        <div class="mt-8 grid gap-4 sm:grid-cols-2 lg:grid-cols-5">
          <div v-for="stack in stacks" :key="stack.name" class="rounded-2xl border border-slate-200 bg-white p-5">
            <p class="text-[11px] font-semibold uppercase tracking-[0.16em] text-brand-600">{{ stack.scope }}</p>
            <h3 class="mt-1.5 text-sm font-semibold text-slate-900">{{ stack.name }}</h3>
            <p class="mt-2 text-[12px] leading-5 text-slate-500">{{ stack.desc }}</p>
          </div>
        </div>
      </section>

      <section class="mt-12 rounded-3xl border border-slate-200 bg-slate-50/70 p-7 sm:p-9">
        <div class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <p class="eyebrow">继续往下看</p>
            <h2 class="mt-2 text-xl font-semibold text-slate-900">部署步骤、端口与配置下发细节</h2>
            <p class="mt-2 max-w-2xl text-[13px] leading-6 text-slate-500">
              具体编排、网络要求、上线检查项在部署接入；授权与快照的完整说明在文档。
            </p>
          </div>
          <div class="flex flex-wrap gap-3">
            <router-link to="/deploy" class="btn-primary">部署接入</router-link>
            <router-link to="/docs" class="btn-secondary">打开文档</router-link>
          </div>
        </div>
      </section>

      <CtaBand
        title="要私有化部署或接口对接，直接聊"
        desc="可以把现有设备规模、网络形态和部署要求发给我们，技术同学直接对接。"
        primary-label="联系技术支持"
        primary-to="/contact"
        secondary-label="了解按号收费"
        secondary-to="/pricing"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import PageHero from '@/components/PageHero.vue'
import ManagerTree from '@/components/ManagerTree.vue'
import CtaBand from '@/components/CtaBand.vue'

const modules = [
  { name: '设备端', owns: '执行业务、缓存快照、本地验签', avoids: '不维护全局配置、不改写他人配置' },
  { name: 'Manager', owns: '配置权威源、设备注册、快照下发、节点管理', avoids: '不签发授权时效' },
  { name: 'discovery', owns: 'Manager 地址登记与解析', avoids: '不承载配置与业务数据' },
  { name: 'Auth', owns: '订单、token、账号实例授权与租约签发', avoids: '不管设备执行与配置内容' },
  { name: '监控', owns: '设备在线、执行异常、授权到期的采集与告警触发', avoids: '不介入业务执行与授权判定' },
]

const licenseFlow = [
  { title: 'Auth 签发租约', desc: '按账号实例签发，写入绝对到期时间，作为唯一时效依据。' },
  { title: 'Manager 下发', desc: '设备注册后随快照拿到属于自己的租约，不需要自己拼装。' },
  { title: '设备本地验签', desc: '设备校验签名后缓存，离线也按同一份租约执行。' },
  { title: '按绝对时间失效', desc: '在线、离线、关机都用同一个到期点，不因离线顺延。' },
]

const ports = [
  { component: 'Manager 根节点', port: '5001' },
  { component: 'Manager 子节点', port: '5102 / 5202' },
  { component: 'discovery', port: '5003' },
  { component: '官网（本仓库）', port: '8080' },
]

const stacks = [
  { scope: '设备端', name: 'Python / AirScript', desc: '设备侧执行脚本，快照落 JSON，本地验签。' },
  { scope: '管理后台', name: 'Go + Vue', desc: 'Manager 后端与管理界面，配置与设备集中管理。' },
  { scope: '部署', name: 'Docker Compose', desc: '根节点、子节点分开编排，按规模拆分。' },
  { scope: '监控', name: 'metrics_infra', desc: '采集设备在线与执行指标，为通知提供数据源。' },
  { scope: '官网', name: 'Go + Vue', desc: '本仓库：静态产物 + 内容 API，独立部署。' },
]
</script>
