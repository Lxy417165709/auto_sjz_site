<template>
  <div>
    <PageHero
      eyebrow="功能"
      title="从设备接入到授权执行，全链路由一套系统负责"
      desc="设备端只执行，Manager 只下发，discovery 只解析地址，Auth 只签发时效。职责边界清晰，排障时一眼定位到环节。"
    >
      <div class="flex flex-wrap gap-2">
        <span class="tag">设备注册</span>
        <span class="tag">快照下发</span>
        <span class="tag">账号实例</span>
        <span class="tag">授权租约</span>
        <span class="tag">多节点</span>
      </div>
    </PageHero>

    <div class="wrap">
      <section class="py-14 sm:py-16">
        <p class="eyebrow">功能矩阵</p>
        <h2 class="h2 mt-2">六件事，覆盖设备自动化日常</h2>
        <div class="mt-8 grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          <article v-for="item in features" :key="item.title" class="card flex flex-col">
            <p class="text-[11px] font-semibold uppercase tracking-[0.16em] text-brand-600">{{ item.module }}</p>
            <h3 class="mt-2 text-[15px] font-semibold text-slate-900">{{ item.title }}</h3>
            <p class="mt-2 text-[13px] leading-6 text-slate-500">{{ item.desc }}</p>
            <ul class="mt-4 space-y-1.5 border-t border-slate-100 pt-4 text-[12px] text-slate-500">
              <li v-for="point in item.points" :key="point" class="flex gap-2">
                <span class="mt-1.5 h-1 w-1 flex-none rounded-full bg-brand-400"></span>
                <span>{{ point }}</span>
              </li>
            </ul>
          </article>
        </div>
      </section>

      <section class="rounded-3xl border border-slate-200 bg-slate-50/60 px-5 py-10 sm:px-9">
        <p class="eyebrow">职责边界</p>
        <h2 class="h2 mt-2">四个模块，各自只做一件事</h2>
        <div class="mt-6 overflow-x-auto">
          <table class="w-full min-w-[720px] border-separate border-spacing-y-2 text-left text-sm">
            <thead>
              <tr class="text-xs uppercase tracking-[0.12em] text-slate-500">
                <th class="px-4 py-2 font-medium">模块</th>
                <th class="px-4 py-2 font-medium">负责</th>
                <th class="px-4 py-2 font-medium">不做</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="row in modules" :key="row.name" class="rounded-xl bg-white text-slate-700">
                <td class="rounded-l-xl px-4 py-3 font-medium text-slate-900">{{ row.name }}</td>
                <td class="px-4 py-3 text-slate-600">{{ row.owns }}</td>
                <td class="rounded-r-xl px-4 py-3 text-slate-500">{{ row.avoids }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>

      <section class="grid gap-4 py-14 sm:py-16 lg:grid-cols-3">
        <div v-for="scene in scenes" :key="scene.title" class="rounded-2xl border border-slate-200 p-6">
          <h3 class="text-sm font-semibold text-slate-900">{{ scene.title }}</h3>
          <p class="mt-2 text-[13px] leading-6 text-slate-500">{{ scene.desc }}</p>
        </div>
      </section>

      <CtaBand
        title="把设备接入链路跑通，再谈规模"
        desc="一套 discovery + Manager 就能支撑起步阶段的设备接入；设备变多时，按子节点拆分即可。"
        primary-label="查看部署接入"
        primary-to="/deploy"
        secondary-label="了解按号收费"
        secondary-to="/pricing"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import PageHero from '@/components/PageHero.vue'
import CtaBand from '@/components/CtaBand.vue'

const features = [
  {
    module: '设备端',
    title: '设备注册与快照拉取',
    desc: '设备启动后自动注册，拉取属于自己的配置快照并落盘，业务按快照执行。',
    points: ['注册失败自动重试', '快照写入本地 JSON', '离线期间只读'],
  },
  {
    module: '设备端',
    title: '账号实例运行',
    desc: '一台设备可承载多个账号，每个账号对应一份授权实例，执行时互不干扰。',
    points: ['账号与实例一一对应', '增减账号即时生效', '执行日志按实例归档'],
  },
  {
    module: 'Manager',
    title: '统一配置与分组',
    desc: '按设备、分组维护配置，变更后由设备主动拉取，避免逐台登录改配置。',
    points: ['单一权威源', '变更审计留痕', '根节点 + 子节点'],
  },
  {
    module: 'discovery',
    title: '地址解析与切换',
    desc: '设备只认 discovery 地址，Manager 地址变更只改登记，不动设备。',
    points: ['不写死 IP', '迁移机房无需刷机', '故障时可按分组切流'],
  },
  {
    module: 'Auth',
    title: '授权租约与验签',
    desc: '按账号实例签发租约，设备本地验签，按绝对到期时间执行。',
    points: ['签名不可伪造', '到期不删号不登出', '离线同一时间失效'],
  },
  {
    module: '部署',
    title: 'Docker 独立编排',
    desc: '根节点与子节点 compose 分开，Manager 与 discovery 可跨机部署。',
    points: ['单机即可起步', '按节点横向扩展', '升级不影响设备执行'],
  },
]

const modules = [
  { name: '设备端', owns: '执行业务、缓存快照、本地验签', avoids: '不维护全局配置、不改写他人配置' },
  { name: 'Manager', owns: '配置权威源、设备注册、快照下发、节点管理', avoids: '不签发授权时效' },
  { name: 'discovery', owns: 'Manager 地址登记与解析', avoids: '不承载配置与业务数据' },
  { name: 'Auth', owns: '订单、token、账号实例授权与租约签发', avoids: '不管设备执行与配置内容' },
]

const scenes = [
  { title: '多设备统一变更', desc: '一次改动下发到全部设备，避免逐台操作导致配置不一致。' },
  { title: '账号规模化管理', desc: '按账号实例授权与计费，账号增减直接反映在授权数量上。' },
  { title: '跨机房部署', desc: 'Manager 与 discovery 分开部署，设备只依赖 discovery 与 Manager 可达地址。' },
]
</script>
