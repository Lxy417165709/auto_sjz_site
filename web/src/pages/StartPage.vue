<template>
  <div>
    <PageHero
      eyebrow="接入流程"
      title="设备、账号、模块，三步接入"
      desc="接入按顺序做三步：先把设备接入云控，再给设备添加账号并授权，最后把业务模块绑到账号上。全程在 Manager 管理台完成，官网只做说明，不承接开通与下单。"
    >
      <div class="flex flex-wrap gap-2">
        <span class="tag">设备接入</span>
        <span class="tag">添加账号</span>
        <span class="tag">添加模块</span>
        <span class="tag">三步接入</span>
      </div>
    </PageHero>

    <div class="wrap">
      <section class="py-14 sm:py-16">
        <div class="grid gap-4 md:grid-cols-3">
          <article v-for="step in overview" :key="step.no" class="rounded-2xl border border-slate-200 bg-white p-6 shadow-card">
            <div class="flex items-center justify-between">
              <span class="font-mono text-xs text-brand-600">{{ step.no }}</span>
              <span class="tag">{{ step.hint }}</span>
            </div>
            <h2 class="mt-4 text-[15px] font-semibold text-slate-900">{{ step.title }}</h2>
            <p class="mt-2 text-[13px] leading-6 text-slate-500">{{ step.desc }}</p>
          </article>
        </div>
      </section>

      <section class="border-t border-slate-200 py-14 sm:py-16">
        <div class="grid gap-8 lg:grid-cols-[0.85fr_1.15fr] lg:gap-12">
          <div>
            <p class="eyebrow">第一步</p>
            <h2 class="h2 mt-2">设备接入云控</h2>
            <p class="mt-4 text-sm leading-7 text-slate-600">
              设备不写死 Manager 的 IP：启动后先向 discovery 问「Manager 在哪」，再拉取属于自己的配置快照，写入本地后开始执行业务。所以换 Manager 地址、加子节点，都不用逐台改设备。
            </p>
            <div class="mt-6 rounded-2xl border border-brand-100 bg-brand-50/60 p-5">
              <p class="text-xs font-semibold text-brand-700">设备侧只读</p>
              <p class="mt-2 text-[13px] leading-6 text-slate-600">
                设备只读取配置、不提供改配置入口。离线期间也不会改本地配置，恢复后自动对齐管理台的最新快照，避免现场手改和平台不一致。
              </p>
            </div>
          </div>

          <ol class="space-y-3">
            <li v-for="(item, i) in deviceSteps" :key="item.title" class="flex gap-4 rounded-2xl border border-slate-200 bg-white p-5">
              <span class="mt-0.5 flex h-7 w-7 flex-none items-center justify-center rounded-full bg-ink-900 font-mono text-[11px] text-white">{{ i + 1 }}</span>
              <div>
                <p class="text-sm font-medium text-slate-900">{{ item.title }}</p>
                <p class="mt-1 text-[13px] leading-6 text-slate-500">{{ item.desc }}</p>
              </div>
            </li>
          </ol>
        </div>
      </section>

      <section class="border-t border-slate-200 py-14 sm:py-16">
        <div class="grid gap-8 lg:grid-cols-[0.85fr_1.15fr] lg:gap-12">
          <div>
            <p class="eyebrow">第二步</p>
            <h2 class="h2 mt-2">添加账号并授权</h2>
            <p class="mt-4 text-sm leading-7 text-slate-600">
              账号是计费单位，也是执行单位：一个账号对应一份授权，3 元/月。设备只是载体，加设备不额外收费，账号增减才影响月费。
            </p>
            <div class="mt-6 divide-y divide-slate-200 rounded-2xl border border-slate-200 bg-slate-50/70">
              <div v-for="fact in accountFacts" :key="fact.label" class="flex items-baseline justify-between gap-3 px-4 py-3">
                <span class="font-mono text-base font-semibold text-slate-900">{{ fact.value }}</span>
                <span class="text-right text-[12px] leading-5 text-slate-500">{{ fact.label }}</span>
              </div>
            </div>
          </div>

          <ol class="space-y-3">
            <li v-for="(item, i) in accountSteps" :key="item.title" class="flex gap-4 rounded-2xl border border-slate-200 bg-white p-5">
              <span class="mt-0.5 flex h-7 w-7 flex-none items-center justify-center rounded-full bg-ink-900 font-mono text-[11px] text-white">{{ i + 1 }}</span>
              <div>
                <p class="text-sm font-medium text-slate-900">{{ item.title }}</p>
                <p class="mt-1 text-[13px] leading-6 text-slate-500">{{ item.desc }}</p>
              </div>
            </li>
          </ol>
        </div>
      </section>

      <section class="border-t border-slate-200 py-14 sm:py-16">
        <div class="grid gap-8 lg:grid-cols-[0.85fr_1.15fr] lg:gap-12">
          <div>
            <p class="eyebrow">第三步</p>
            <h2 class="h2 mt-2">添加模块并下发</h2>
            <p class="mt-4 text-sm leading-7 text-slate-600">
              模块是跑在账号上的业务能力，决定这个号「做什么」。加模块不加钱：模块只改执行内容，计费仍按账号实例算。
            </p>
            <div class="mt-6 rounded-2xl border border-amber-200/70 bg-amber-50/70 p-5">
              <p class="text-xs font-semibold text-amber-700">模块与授权分开</p>
              <p class="mt-2 text-[13px] leading-6 text-slate-600">
                模块下发不影响授权有效期。授权到期只会停止执行，账号、设备与模块配置都保留，续费后继续跑。
              </p>
            </div>
          </div>

          <div>
            <ol class="space-y-3">
              <li v-for="(item, i) in moduleSteps" :key="item.title" class="flex gap-4 rounded-2xl border border-slate-200 bg-white p-5">
                <span class="mt-0.5 flex h-7 w-7 flex-none items-center justify-center rounded-full bg-ink-900 font-mono text-[11px] text-white">{{ i + 1 }}</span>
                <div>
                  <p class="text-sm font-medium text-slate-900">{{ item.title }}</p>
                  <p class="mt-1 text-[13px] leading-6 text-slate-500">{{ item.desc }}</p>
                </div>
              </li>
            </ol>
            <p class="mt-4 text-[12px] leading-6 text-slate-400">
              模块的导入方式与参数项随 Manager 版本不同，以管理台实际页面为准；配置只存管理台，设备按快照执行，改参数不用逐台改。
            </p>
          </div>
        </div>
      </section>

      <section class="border-t border-slate-200 py-14 sm:py-16">
        <div class="grid gap-4 lg:grid-cols-2">
          <div class="rounded-3xl border border-slate-200 p-7 sm:p-8">
            <p class="eyebrow">接入前准备</p>
            <h2 class="mt-2 text-xl font-semibold tracking-tight text-slate-900">先把这四件事定下来</h2>
            <ul class="mt-5 space-y-4">
              <li v-for="item in prep" :key="item.title" class="flex gap-3">
                <span class="mt-1 flex h-5 w-5 flex-none items-center justify-center rounded-full bg-accent-500/15 text-[11px] text-accent-600">✓</span>
                <div>
                  <p class="text-sm font-medium text-slate-900">{{ item.title }}</p>
                  <p class="mt-0.5 text-[13px] leading-6 text-slate-500">{{ item.desc }}</p>
                </div>
              </li>
            </ul>
          </div>

          <div class="rounded-3xl border border-slate-200 bg-slate-50/70 p-7 sm:p-8">
            <p class="eyebrow">常见卡点</p>
            <h2 class="mt-2 text-xl font-semibold tracking-tight text-slate-900">多半是这四处没对上</h2>
            <dl class="mt-5 space-y-4">
              <div v-for="item in blockers" :key="item.q">
                <dt class="text-sm font-medium text-slate-900">{{ item.q }}</dt>
                <dd class="mt-1 text-[13px] leading-6 text-slate-500">{{ item.a }}</dd>
              </div>
            </dl>
          </div>
        </div>
        <p class="mt-5 text-[12px] leading-6 text-slate-400">
          流程为通用说明，管理台菜单名称以实际版本为准；需要架构细节（discovery、端口、授权验签）可转到技术架构与部署接入。
        </p>
      </section>

      <CtaBand
        title="先跑通一台，再按分组铺开"
        desc="告诉我们设备台数、每台账号数和要跑的模块，我们给出接入顺序、账号实例数与月费；需要私有化部署也可以直接说。"
        primary-label="接入咨询"
        primary-to="/contact"
        secondary-label="查看技术架构"
        secondary-to="/tech"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import CtaBand from '@/components/CtaBand.vue'
import PageHero from '@/components/PageHero.vue'

const overview = [
  {
    no: '01',
    hint: '设备',
    title: '接入云控',
    desc: '设备装客户端、填 discovery 地址，启动后自动注册上线，出现在管理台的设备列表里。',
  },
  {
    no: '02',
    hint: '账号',
    title: '加号并授权',
    desc: '在设备或分组下新增账号实例，一个号一份授权，3 元/月；设备台数不计费。',
  },
  {
    no: '03',
    hint: '模块',
    title: '绑模块跑业务',
    desc: '选择或导入业务模块，配好参数并绑定账号范围，先灰度验证，再批量下发。',
  },
]

const deviceSteps = [
  { title: '准备服务地址', desc: '先部署好 Manager 与 discovery，拿到 discovery 地址；私有化部署方式见「部署接入」。' },
  { title: '建设备分组', desc: '在管理台按门店、批次或用途建分组，后面的设备和账号都按分组管理。' },
  { title: '设备装客户端', desc: '在设备上安装 auto_sjz 客户端，填入 discovery 地址；不需要逐台写 Manager 的 IP。' },
  { title: '上线并确认状态', desc: '设备启动后自动注册，管理台设备列表出现该设备，状态变为在线即接入成功。' },
  { title: '归组并拉取配置', desc: '把设备归到分组，设备拉取配置快照写入本地；后续配置变更都在管理台完成。' },
]

const accountFacts = [
  { value: '3 元', label: '一个号一个月，设备不叠加' },
  { value: '1 : 1', label: '一个账号一份授权' },
  { value: '不删号', label: '到期只停执行，续费恢复' },
]

const accountSteps = [
  { title: '选设备或分组', desc: '在设备详情或分组下点新增账号，先确定这个号挂在哪台设备上。' },
  { title: '填账号信息', desc: '按统一规则填账号名称、平台与备注，后续对账、排障都靠这套命名区分。' },
  { title: '提交并授权', desc: '提交后由 Auth 签发该账号的授权租约，按号计算；设备台数不参与计费。' },
  { title: '首次启动', desc: '下发配置并在设备上启动该账号，管理台查看运行状态，异常会进入通知。' },
]

const moduleSteps = [
  { title: '选择或导入模块', desc: '从模块库选择现成模块，或按 Manager 支持的方式导入自有模块包。' },
  { title: '配置参数', desc: '设置执行频率、时间窗、失败重试与目标账号范围，先按默认值跑通再调优。' },
  { title: '绑定账号范围', desc: '绑到单个账号、一批账号或整个分组，避免一台一台重复配置。' },
  { title: '灰度验证', desc: '先放 1–2 台设备执行，看执行记录与异常，再决定是否全量。' },
  { title: '批量下发并观察', desc: '全量下发后看执行成功率与通知，需要时改配置重新下发，不用上设备操作。' },
]

const prep = [
  { title: '设备与账号规划', desc: '准备接入多少台设备、每台挂几个号，用来估授权数量与月费。' },
  { title: '网络可达', desc: '设备能访问 discovery 与 Manager；跨网段或门店网络要先确认出口。' },
  { title: '部署方式', desc: 'Docker 单机起步，或按规模拆根节点与子节点，私有化部署在自己机器上。' },
  { title: '配置负责人', desc: '配置只在管理台改，现场不手改；明确谁负责分组、加号和模块参数。' },
]

const blockers = [
  { q: '设备一直离线', a: '先看 discovery 地址是否可达、设备网络出口是否放行、客户端是否已启动。' },
  { q: '账号加完跑不起来', a: '确认该账号已授权；授权按号签发，到期只停执行，续费后恢复。' },
  { q: '模块下发没反应', a: '设备要等下一次拉取快照才生效，确认设备在线且已归到对应分组。' },
  { q: '账单和预期不一致', a: '按账号数核对，不看设备台数；删号后从下一个计费周期起减少。' },
]
</script>
