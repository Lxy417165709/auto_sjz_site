<template>
  <div>
    <PageHero
      eyebrow="业务能力"
      title="少花冤枉钱，少做重复操作"
      desc="按号收费、中台批量管理、监控通知——业务上最常被问的三件事放在最前面。架构、端口和部署细节都挪到技术区，不干扰判断。"
    >
      <div class="flex flex-wrap gap-2">
        <span class="tag">3 元 / 号 / 月</span>
        <span class="tag">设备不计费</span>
        <span class="tag">批量管理</span>
        <span class="tag">监控通知</span>
      </div>
    </PageHero>

    <div class="wrap">
      <section class="grid gap-4 py-14 sm:grid-cols-3 sm:py-16">
        <div v-for="stat in stats" :key="stat.label" class="rounded-2xl border border-slate-200 bg-white p-6 shadow-card">
          <p class="font-mono text-2xl font-semibold text-slate-900">{{ stat.value }}</p>
          <p class="mt-1 text-sm font-medium text-slate-800">{{ stat.label }}</p>
          <p class="mt-2 text-[13px] leading-6 text-slate-500">{{ stat.desc }}</p>
        </div>
      </section>

      <section>
        <p class="eyebrow">能力清单</p>
        <h2 class="h2 mt-2">六件事，覆盖日常管理</h2>
        <div class="mt-8 grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          <article v-for="item in capabilities" :key="item.title" class="card flex flex-col">
            <span class="tag w-fit">{{ item.tag }}</span>
            <h3 class="mt-3 text-[15px] font-semibold text-slate-900">{{ item.title }}</h3>
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

      <section class="mt-14 rounded-3xl border border-[#e7e5e2] bg-[#f7f6f3] px-5 py-10 sm:px-9">
        <p class="eyebrow">收费方式对比</p>
        <h2 class="h2 mt-2">按设备，还是按号</h2>
        <div class="mt-6 overflow-x-auto">
          <table class="w-full min-w-[680px] border-separate border-spacing-y-2 text-left text-sm">
            <thead>
              <tr class="text-xs uppercase tracking-[0.12em] text-slate-500">
                <th class="w-[24%] px-4 py-2 font-medium">对比项</th>
                <th class="w-[32%] px-4 py-2 font-medium">传统按设备</th>
                <th class="w-[44%] px-4 py-2 font-medium">按号方案</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="row in compareRows" :key="row.item" class="bg-white text-slate-700">
                <td class="rounded-l-xl px-4 py-3 font-medium text-slate-900">{{ row.item }}</td>
                <td class="px-4 py-3 text-slate-500">{{ row.device }}</td>
                <td class="rounded-r-xl px-4 py-3 text-brand-700">{{ row.account }}</td>
              </tr>
            </tbody>
          </table>
        </div>
        <p class="mt-4 text-[12px] leading-5 text-slate-500">按号计费以账号实例为单位，3 元/号/月；设备台数不参与计费。</p>
      </section>

      <section class="grid gap-4 py-14 sm:py-16 lg:grid-cols-3">
        <div v-for="scene in scenes" :key="scene.title" class="rounded-2xl border border-slate-200 p-6">
          <h3 class="text-sm font-semibold text-slate-900">{{ scene.title }}</h3>
          <p class="mt-2 text-[13px] leading-6 text-slate-500">{{ scene.desc }}</p>
        </div>
      </section>

      <CtaBand
        title="先算清账号，再谈部署"
        desc="告诉我们设备台数与每台账号数，我们按 3 元/号/月 给出月费，再讨论接入方式。"
        primary-label="算一下月费"
        primary-to="/pricing"
        secondary-label="看看管理台"
        secondary-to="/console"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import PageHero from '@/components/PageHero.vue'
import CtaBand from '@/components/CtaBand.vue'

const stats = [
  { value: '¥3', label: '单个账号月费', desc: '按账号实例结算，设备台数不额外收费。' },
  { value: '1 个中台', label: '设备与账号集中管理', desc: '批量下发配置、批量查看状态，不用逐台登录。' },
  { value: '持续监控', label: '在线、异常、到期', desc: '关键状态集中呈现，异常按配置通知。' },
]

const capabilities = [
  {
    tag: '收费',
    title: '按号收费，成本跟着账号走',
    desc: '一个号 3 元/月，一台设备挂 5 个号就是 15 元/月；加设备不加钱，账号增减直接反映到账单。',
    points: ['设备台数不计费', '账号按实例结算', '月费可预估、可对账'],
  },
  {
    tag: '管理',
    title: '管理中台，批量管理',
    desc: '设备、分组、账号与配置集中在一个后台，批量变更一次下发，异常设备集中处理。',
    points: ['按分组批量下发', '集中查看设备状态', '账号实例统一维护'],
  },
  {
    tag: '监控',
    title: '监控通知，异常不漏',
    desc: '设备在线/离线、执行异常、授权到期集中监控，触发后按配置的通知渠道提醒。',
    points: ['在线状态可见', '执行异常提醒', '到期提前通知'],
  },
  {
    tag: '配置',
    title: '统一配置，不用逐台改',
    desc: '配置只在中台维护，设备拉取快照执行；离线期间只读，恢复后自动对齐最新配置。',
    points: ['单一配置源', '变更批量生效', '离线不误改'],
  },
  {
    tag: '授权',
    title: '到期只停执行，不删号不登出',
    desc: '授权按绝对到期时间失效，账号与设备配置原样保留，续费后继续执行，不用重建环境。',
    points: ['账号不删除', '不强制退出登录', '离线同样按点失效'],
  },
  {
    tag: '部署',
    title: '独立部署，数据自己掌握',
    desc: 'Docker 私有化部署，Manager 与 discovery 可以分开，账号与业务数据留在自己的机器上。',
    points: ['单机即可起步', '按规模加子节点', '不依赖厂商云'],
  },
]

const compareRows = [
  { item: '计费单位', device: '按设备台数', account: '按账号实例（3 元/号/月）' },
  { item: '设备增减', device: '每加一台就加一份费用', account: '加设备不额外收费' },
  { item: '账号增减', device: '一台设备挂多少号都一样', account: '加一个号算一个，减一个省一个' },
  { item: '闲置成本', device: '设备闲着也要付整套', account: '只为实际账号付费' },
  { item: '账单核对', device: '按台数对，难对应到账号', account: '账号数 × 3 元，账目直接可对' },
]

const scenes = [
  { title: '多账号运营', desc: '一台设备多个号，按号结算更贴近实际成本结构。' },
  { title: '设备逐步扩容', desc: '先小规模跑通，再按账号数扩，预算跟着业务走。' },
  { title: '长期稳定运行', desc: '到期不删号、不登出，续费即恢复，减少重复配置。' },
]
</script>
