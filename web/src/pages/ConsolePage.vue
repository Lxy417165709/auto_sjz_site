<template>
  <div>
    <PageHero
      eyebrow="管理台演示"
      title="配置与状态，都在中台里"
      desc="下面是可以点的示意界面：左边切设备，右边看账号与配置；状态页可以筛选设备、查看账号运行情况和异常通知。数据为演示数据，实际操作以 Manager 为准。"
    >
      <div class="flex flex-wrap gap-2">
        <span class="tag">业务配置</span>
        <span class="tag">业务状态</span>
        <span class="tag">可交互</span>
        <span class="tag">示意数据</span>
      </div>
    </PageHero>

    <div class="wrap py-10 sm:py-12">
      <div class="overflow-hidden rounded-3xl border border-slate-200 bg-slate-50 shadow-card">
        <div class="flex flex-wrap items-center justify-between gap-3 bg-ink-900 px-4 py-3 text-white sm:px-6">
          <div class="flex flex-none items-center gap-3">
            <span class="flex h-8 w-8 items-center justify-center rounded-lg bg-brand-600 text-white">
              <svg class="h-4 w-4" viewBox="0 0 24 24" fill="none" aria-hidden="true">
                <path d="M3 12h4l2-5 4 10 2-5h6" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" />
              </svg>
            </span>
            <div>
              <p class="text-sm font-semibold">auto_sjz Manager</p>
              <p class="text-[10px] text-slate-400">演示环境 · 示意数据</p>
            </div>
          </div>

          <div class="order-3 flex w-full justify-center rounded-full border border-white/15 bg-white/5 p-1 sm:order-none sm:w-auto">
            <button
              v-for="item in tabs"
              :key="item.key"
              type="button"
              class="rounded-full px-3.5 py-1.5 text-[12px] transition"
              :class="tab === item.key ? 'bg-white text-slate-900' : 'text-slate-300 hover:text-white'"
              @click="tab = item.key"
            >
              {{ item.label }}
            </button>
          </div>

          <div class="flex items-center gap-2 text-[11px] text-slate-400">
            <span class="hidden sm:inline">刷新于 {{ lastRefreshLabel }}</span>
            <button type="button" class="rounded-full border border-white/15 px-3 py-1 text-slate-300 transition hover:text-white" @click="refresh">
              刷新
            </button>
          </div>
        </div>

        <div v-if="tab === 'config'" class="grid lg:grid-cols-[250px_1fr]">
          <aside class="border-b border-slate-200 bg-white p-4 lg:border-b-0 lg:border-r">
            <div class="flex items-center justify-between">
              <p class="text-[11px] font-semibold uppercase tracking-[0.14em] text-slate-500">设备与分组</p>
              <span class="text-[11px] text-slate-400">{{ devices.length }} 台</span>
            </div>
            <input v-model.trim="deviceQuery" class="field mt-3 !py-2 text-[13px]" placeholder="搜索设备或分组" />

            <div class="mt-3 space-y-1.5">
              <div v-for="group in filteredGroups" :key="group.name">
                <button
                  type="button"
                  class="flex w-full items-center justify-between rounded-lg px-2 py-1.5 text-left text-[12px] font-medium text-slate-600 hover:bg-slate-50"
                  @click="toggleGroup(group.name)"
                >
                  <span class="flex items-center gap-1.5">
                    <span class="text-slate-400">{{ collapsed.includes(group.name) ? '▸' : '▾' }}</span>
                    {{ group.name }}
                  </span>
                  <span class="text-[11px] text-slate-400">{{ group.devices.length }}</span>
                </button>

                <div v-if="!collapsed.includes(group.name)" class="mt-1 space-y-1">
                  <button
                    v-for="device in group.devices"
                    :key="device.id"
                    type="button"
                    class="flex w-full items-center justify-between gap-2 rounded-xl px-2.5 py-2 text-left transition"
                    :class="selectedDeviceId === device.id ? 'bg-brand-50 text-brand-800' : 'text-slate-600 hover:bg-slate-50'"
                    @click="selectedDeviceId = device.id"
                  >
                    <span class="flex min-w-0 items-center gap-2">
                      <span class="h-1.5 w-1.5 flex-none rounded-full" :class="deviceDot(device.status)"></span>
                      <span class="truncate text-[13px]">{{ device.name }}</span>
                    </span>
                    <span class="flex-none text-[11px] text-slate-400">{{ device.accounts.length }} 号</span>
                  </button>
                </div>
              </div>
              <p v-if="!filteredGroups.length" class="px-2 py-6 text-center text-[12px] text-slate-400">没有匹配的设备</p>
            </div>
          </aside>

          <section class="min-w-0 bg-white p-4 sm:p-6">
            <div class="flex flex-wrap items-start justify-between gap-3">
              <div>
                <div class="flex flex-wrap items-center gap-2">
                  <h2 class="text-lg font-semibold text-slate-900">{{ selectedDevice.name }}</h2>
                  <span class="rounded-full px-2 py-0.5 text-[11px]" :class="devicePill(selectedDevice.status)">{{ deviceLabel(selectedDevice.status) }}</span>
                  <span class="tag">{{ selectedDevice.group }}</span>
                </div>
                <p class="mt-1 text-[12px] text-slate-500">最近心跳 {{ selectedDevice.heartbeat }} · 今日执行 {{ selectedDevice.executions }} 次</p>
              </div>
              <button type="button" class="btn-primary !px-4 !py-2 text-[13px]" @click="addAccount">+ 添加账号</button>
            </div>

            <div class="mt-5 grid gap-4 xl:grid-cols-[1.15fr_0.85fr]">
              <div class="rounded-2xl border border-slate-200">
                <div class="flex flex-wrap items-center justify-between gap-2 border-b border-slate-100 px-4 py-3">
                  <div>
                    <p class="text-sm font-semibold text-slate-900">账号实例</p>
                    <p class="mt-0.5 text-[11px] text-slate-500">
                      {{ selectedDevice.accounts.length }} 个号 · {{ selectedDevice.accounts.length * 3 }} 元/月
                    </p>
                  </div>
                  <span class="text-[11px] text-slate-400">一个号一份授权</span>
                </div>

                <div class="divide-y divide-slate-100">
                  <div v-for="account in selectedDevice.accounts" :key="account.id" class="flex flex-wrap items-center gap-3 px-4 py-3">
                    <span class="flex h-8 w-8 flex-none items-center justify-center rounded-full bg-slate-100 text-[11px] font-medium text-slate-600">
                      {{ account.platform.slice(0, 1) }}
                    </span>
                    <div class="min-w-[96px] flex-1">
                      <p class="truncate text-[13px] font-medium text-slate-900">{{ account.name }}</p>
                      <p class="text-[11px] text-slate-400">{{ account.platform }}</p>
                    </div>

                    <div class="flex flex-none flex-wrap items-center gap-2">
                      <select v-model="account.module" class="field !w-[92px] !px-2 !py-1 text-[11px]" aria-label="账号模块">
                        <option v-for="module in moduleOptions" :key="module" :value="module">{{ module }}</option>
                      </select>

                      <span class="rounded-full px-2 py-0.5 text-[10px]" :class="accountPill(account.status)">{{ accountStatusLabel(account.status) }}</span>

                      <button
                        type="button"
                        role="switch"
                        :aria-checked="account.status === 'running'"
                        :disabled="account.status === 'pending'"
                        class="relative h-5 w-9 flex-none rounded-full transition disabled:cursor-not-allowed disabled:opacity-40"
                        :class="account.status === 'running' ? 'bg-accent-500' : 'bg-slate-200'"
                        @click="toggleAccount(account)"
                      >
                        <span class="absolute top-0.5 h-4 w-4 rounded-full bg-white shadow transition" :class="account.status === 'running' ? 'left-[18px]' : 'left-0.5'"></span>
                      </button>

                      <button type="button" class="text-[11px] text-slate-400 transition hover:text-red-600" @click="removeAccount(account)">删除</button>
                    </div>
                  </div>
                </div>

                <p v-if="selectedDevice.accounts.some((a) => a.status === 'pending')" class="border-t border-slate-100 px-4 py-2.5 text-[11px] text-amber-700">
                  有账号待授权：授权由 Auth 签发后才开始执行，演示环境不实际扣费。
                </p>
              </div>

              <div class="rounded-2xl border border-slate-200 p-4">
                <div class="flex items-center justify-between">
                  <p class="text-sm font-semibold text-slate-900">设备配置</p>
                  <span class="text-[11px] text-slate-400">只存 Manager</span>
                </div>

                <label class="mt-4 block">
                  <span class="text-[11px] font-medium text-slate-600">默认模块</span>
                  <select v-model="config.module" class="field mt-1.5 !py-2 text-[13px]">
                    <option v-for="module in moduleOptions" :key="module" :value="module">{{ module }}</option>
                  </select>
                </label>

                <div class="mt-3 grid grid-cols-2 gap-3">
                  <label class="block">
                    <span class="text-[11px] font-medium text-slate-600">执行开始</span>
                    <input v-model="config.start" type="time" class="field mt-1.5 !px-2.5 !py-2 text-[13px]" />
                  </label>
                  <label class="block">
                    <span class="text-[11px] font-medium text-slate-600">执行结束</span>
                    <input v-model="config.end" type="time" class="field mt-1.5 !px-2.5 !py-2 text-[13px]" />
                  </label>
                </div>

                <div class="mt-3 flex items-center justify-between rounded-xl border border-slate-200 px-3 py-2.5">
                  <div>
                    <p class="text-[12px] font-medium text-slate-800">失败重试</p>
                    <p class="text-[10px] text-slate-400">执行失败后自动重试次数</p>
                  </div>
                  <div class="flex items-center gap-2">
                    <button type="button" class="h-7 w-7 rounded-full border border-slate-200 text-slate-500 disabled:opacity-40" :disabled="config.retry <= 0" @click="config.retry--">−</button>
                    <span class="w-5 text-center font-mono text-sm text-slate-900">{{ config.retry }}</span>
                    <button type="button" class="h-7 w-7 rounded-full border border-slate-200 text-slate-500 disabled:opacity-40" :disabled="config.retry >= 5" @click="config.retry++">+</button>
                  </div>
                </div>

                <div class="mt-3 space-y-2">
                  <button type="button" class="flex w-full items-center justify-between rounded-xl border border-slate-200 px-3 py-2.5" @click="config.auto = !config.auto">
                    <span class="text-[12px] font-medium text-slate-800">自动执行</span>
                    <span class="relative h-5 w-9 rounded-full transition" :class="config.auto ? 'bg-brand-600' : 'bg-slate-200'">
                      <span class="absolute top-0.5 h-4 w-4 rounded-full bg-white shadow transition" :class="config.auto ? 'left-[18px]' : 'left-0.5'"></span>
                    </span>
                  </button>
                  <button type="button" class="flex w-full items-center justify-between rounded-xl border border-slate-200 px-3 py-2.5" @click="config.notify = !config.notify">
                    <span class="text-[12px] font-medium text-slate-800">异常通知</span>
                    <span class="relative h-5 w-9 rounded-full transition" :class="config.notify ? 'bg-brand-600' : 'bg-slate-200'">
                      <span class="absolute top-0.5 h-4 w-4 rounded-full bg-white shadow transition" :class="config.notify ? 'left-[18px]' : 'left-0.5'"></span>
                    </span>
                  </button>
                </div>

                <div v-if="dirty" class="mt-4 rounded-xl border border-amber-200 bg-amber-50 px-3 py-2.5">
                  <p class="text-[11px] text-amber-800">有未保存修改，设备下次拉取快照后生效。</p>
                  <div class="mt-2 flex gap-2">
                    <button type="button" class="rounded-full bg-amber-600 px-3 py-1 text-[11px] font-medium text-white" @click="saveConfig">保存并下发</button>
                    <button type="button" class="rounded-full border border-amber-300 px-3 py-1 text-[11px] text-amber-800" @click="resetConfig">放弃</button>
                  </div>
                </div>
                <p v-else class="mt-4 text-[11px] text-slate-400">上次下发 2026-11-20 09:12</p>
              </div>
            </div>
          </section>
        </div>

        <div v-else class="bg-white p-4 sm:p-6">
          <div class="grid grid-cols-2 gap-3 xl:grid-cols-4">
            <div v-for="kpi in kpis" :key="kpi.label" class="rounded-2xl border border-slate-200 px-4 py-3">
              <p class="text-[11px] text-slate-500">{{ kpi.label }}</p>
              <p class="mt-1 font-mono text-xl font-semibold" :class="kpi.tone">{{ kpi.value }}</p>
              <p class="mt-0.5 text-[10px] text-slate-400">{{ kpi.hint }}</p>
            </div>
          </div>

          <div v-if="!incidentDismissed" class="mt-4 flex flex-wrap items-center justify-between gap-3 rounded-2xl border border-red-200 bg-red-50 px-4 py-3">
            <div class="flex items-start gap-2">
              <span class="mt-0.5 h-2 w-2 flex-none rounded-full bg-red-500"></span>
              <div>
                <p class="text-[12px] font-medium text-red-800">设备 07 连续 3 次执行失败，已暂停自动执行</p>
                <p class="text-[11px] text-red-600">建议检查账号登录状态与模块参数。</p>
              </div>
            </div>
            <div class="flex gap-2">
              <button type="button" class="rounded-full bg-red-600 px-3 py-1 text-[11px] text-white" @click="focusIncident">查看设备</button>
              <button type="button" class="rounded-full border border-red-300 px-3 py-1 text-[11px] text-red-700" @click="incidentDismissed = true">忽略</button>
            </div>
          </div>

          <div class="mt-5 grid gap-5 xl:grid-cols-[1fr_300px]">
            <div class="min-w-0">
              <div class="flex flex-wrap items-center justify-between gap-3">
                <div class="flex flex-wrap gap-1.5">
                  <button
                    v-for="filter in filters"
                    :key="filter.key"
                    type="button"
                    class="rounded-full px-3 py-1.5 text-[12px] transition"
                    :class="statusFilter === filter.key ? 'bg-brand-600 text-white' : 'border border-slate-200 text-slate-600 hover:border-brand-200'"
                    @click="statusFilter = filter.key"
                  >
                    {{ filter.label }}
                  </button>
                </div>
                <input v-model.trim="statusQuery" class="field !w-44 !py-1.5 text-[12px]" placeholder="搜索设备 / 分组" />
              </div>

              <div class="mt-3 hidden overflow-x-auto rounded-2xl border border-slate-200 sm:block">
                <table class="w-full min-w-[640px] text-left text-[12px]">
                  <thead class="bg-slate-50 text-[11px] text-slate-500">
                    <tr>
                      <th class="px-4 py-2.5 font-medium">设备</th>
                      <th class="px-4 py-2.5 font-medium">分组</th>
                      <th class="px-4 py-2.5 font-medium">状态</th>
                      <th class="px-4 py-2.5 font-medium">账号</th>
                      <th class="px-4 py-2.5 font-medium">最近心跳</th>
                      <th class="px-4 py-2.5 font-medium">今日执行</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-slate-100">
                    <tr
                      v-for="device in statusDevices"
                      :key="device.id"
                      class="cursor-pointer transition hover:bg-slate-50"
                      :class="selectedStatusId === device.id ? 'bg-brand-50/60' : ''"
                      @click="selectedStatusId = device.id"
                    >
                      <td class="px-4 py-2.5 font-medium text-slate-900">{{ device.name }}</td>
                      <td class="px-4 py-2.5 text-slate-500">{{ device.group }}</td>
                      <td class="px-4 py-2.5">
                        <span class="rounded-full px-2 py-0.5 text-[10px]" :class="devicePill(device.status)">{{ deviceLabel(device.status) }}</span>
                      </td>
                      <td class="px-4 py-2.5 text-slate-600">{{ runningOf(device) }} / {{ device.accounts.length }}</td>
                      <td class="px-4 py-2.5 text-slate-500">{{ device.heartbeat }}</td>
                      <td class="px-4 py-2.5 font-mono text-slate-600">{{ device.executions }}</td>
                    </tr>
                  </tbody>
                </table>
              </div>
              <div class="mt-3 space-y-2 sm:hidden">
                <button
                  v-for="device in statusDevices"
                  :key="device.id"
                  type="button"
                  class="w-full rounded-xl border px-3 py-2.5 text-left transition"
                  :class="selectedStatusId === device.id ? 'border-brand-200 bg-brand-50/60' : 'border-slate-200 bg-white'"
                  @click="selectedStatusId = device.id"
                >
                  <div class="flex items-center justify-between gap-2">
                    <span class="text-[13px] font-medium text-slate-900">{{ device.name }}</span>
                    <span class="rounded-full px-2 py-0.5 text-[10px]" :class="devicePill(device.status)">{{ deviceLabel(device.status) }}</span>
                  </div>
                  <p class="mt-1 text-[11px] text-slate-500">{{ device.group }} · 心跳 {{ device.heartbeat }}</p>
                  <div class="mt-1.5 flex items-center gap-3 text-[11px] text-slate-500">
                    <span>账号 {{ runningOf(device) }} / {{ device.accounts.length }}</span>
                    <span>今日执行 {{ device.executions }}</span>
                  </div>
                </button>
              </div>

              <p v-if="!statusDevices.length" class="py-6 text-center text-[12px] text-slate-400">没有匹配的设备</p>

              <div class="mt-4 rounded-2xl border border-slate-200 p-4">
                <div class="flex flex-wrap items-center justify-between gap-3">
                  <div>
                    <p class="text-sm font-semibold text-slate-900">{{ selectedStatusDevice.name }} · 账号运行情况</p>
                    <p class="mt-0.5 text-[11px] text-slate-500">
                      {{ selectedStatusDevice.group }} · {{ deviceLabel(selectedStatusDevice.status) }} · 心跳 {{ selectedStatusDevice.heartbeat }}
                    </p>
                  </div>
                  <button type="button" class="rounded-full border border-slate-300 px-3 py-1 text-[11px] text-slate-600 hover:border-brand-300" @click="pullSnapshot">
                    重新拉取快照
                  </button>
                </div>
                <div class="mt-3 grid gap-2 sm:grid-cols-2">
                  <div v-for="account in selectedStatusDevice.accounts" :key="account.id" class="flex items-center justify-between gap-2 rounded-xl bg-slate-50 px-3 py-2">
                    <div class="min-w-0">
                      <p class="truncate text-[12px] text-slate-800">{{ account.name }}</p>
                      <p class="text-[10px] text-slate-400">{{ account.module }}</p>
                    </div>
                    <span class="rounded-full px-2 py-0.5 text-[10px]" :class="accountPill(account.status)">{{ accountStatusLabel(account.status) }}</span>
                  </div>
                </div>
              </div>
            </div>

            <aside class="rounded-2xl border border-slate-200 p-4">
              <div class="flex items-center justify-between">
                <p class="text-sm font-semibold text-slate-900">最近通知</p>
                <button v-if="notifications.length" type="button" class="text-[11px] text-slate-400 hover:text-brand-600" @click="notifications = []">全部已读</button>
              </div>
              <ul class="mt-3 space-y-2.5">
                <li v-for="note in notifications" :key="note.id" class="rounded-xl border border-slate-100 bg-slate-50/70 px-3 py-2">
                  <div class="flex items-center justify-between gap-2">
                    <span class="rounded-full px-1.5 py-0.5 text-[9px]" :class="notePill(note.type)">{{ note.type }}</span>
                    <span class="text-[10px] text-slate-400">{{ note.time }}</span>
                  </div>
                  <p class="mt-1.5 text-[11px] leading-5 text-slate-600">{{ note.text }}</p>
                </li>
              </ul>
              <p v-if="!notifications.length" class="mt-4 text-[11px] text-slate-400">没有待处理通知，状态正常。</p>
              <p class="mt-4 border-t border-slate-100 pt-3 text-[10px] leading-5 text-slate-400">
                通知渠道在部署时配置，演示页不连接真实消息通道。
              </p>
            </aside>
          </div>
        </div>
      </div>

      <p class="mt-4 text-[12px] leading-6 text-slate-400">
        以上为交互示意，设备、账号与执行数据均为演示数据；实际界面、菜单与字段以当前 Manager 版本为准。
      </p>
    </div>

    <div class="pointer-events-none fixed inset-x-0 bottom-6 z-50 flex justify-center px-4">
      <p v-if="toast" class="pointer-events-auto rounded-full bg-ink-900 px-4 py-2 text-[12px] text-white shadow-2xl">{{ toast }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue'
import PageHero from '@/components/PageHero.vue'

type DeviceStatus = 'online' | 'offline' | 'error'
type AccountStatus = 'running' | 'stopped' | 'pending' | 'error'
type TabKey = 'config' | 'status'

interface Account {
  id: string
  name: string
  platform: string
  module: string
  status: AccountStatus
}

interface Device {
  id: string
  name: string
  group: string
  status: DeviceStatus
  accounts: Account[]
  heartbeat: string
  executions: number
}

interface DeviceConfig {
  module: string
  start: string
  end: string
  retry: number
  auto: boolean
  notify: boolean
}

const moduleOptions = ['日常巡检', '内容发布', '消息回复', '数据采集']
const tabs: { key: TabKey; label: string }[] = [
  { key: 'config', label: '业务配置' },
  { key: 'status', label: '业务状态' },
]
const filters = [
  { key: 'all', label: '全部' },
  { key: 'online', label: '在线' },
  { key: 'offline', label: '离线' },
  { key: 'error', label: '异常' },
]

function buildAccounts(deviceIndex: number, count: number, deviceStatus: DeviceStatus): Account[] {
  return Array.from({ length: count }, (_, i) => {
    const n = i + 1
    let status: AccountStatus = 'running'
    if (n === count) status = 'stopped'
    if (deviceStatus === 'offline') status = 'stopped'
    if (deviceStatus === 'error') status = n === 1 ? 'error' : 'stopped'
    if (deviceIndex === 1 && n === 2) status = 'pending'
    return {
      id: `d${deviceIndex}-a${n}`,
      name: `账号 ${String(n).padStart(2, '0')}`,
      platform: n % 3 === 0 ? '视频号' : '抖音',
      module: moduleOptions[(deviceIndex + n) % moduleOptions.length],
      status,
    }
  })
}

function buildDevice(index: number, group: string, status: DeviceStatus, executions: number, count = 5): Device {
  return {
    id: `d${index}`,
    name: `设备 ${String(index).padStart(2, '0')}`,
    group,
    status,
    accounts: buildAccounts(index, count, status),
    heartbeat: status === 'online' ? '刚刚' : status === 'error' ? '2 分钟前' : '18 分钟前',
    executions,
  }
}

const devices = reactive<Device[]>([
  buildDevice(1, '门店 A', 'online', 186),
  buildDevice(2, '门店 A', 'online', 172),
  buildDevice(3, '门店 A', 'online', 168),
  buildDevice(4, '门店 A', 'online', 154),
  buildDevice(5, '门店 B', 'online', 148),
  buildDevice(6, '门店 B', 'online', 132),
  buildDevice(7, '门店 B', 'error', 96),
  buildDevice(8, '门店 B', 'online', 128),
  buildDevice(9, '门店 C', 'online', 121),
  buildDevice(10, '门店 C', 'online', 118),
  buildDevice(11, '门店 C', 'offline', 0),
  buildDevice(12, '测试组', 'offline', 43),
])

const tab = ref<TabKey>('config')
const deviceQuery = ref('')
const statusQuery = ref('')
const statusFilter = ref('all')
const collapsed = ref<string[]>([])
const selectedDeviceId = ref('d1')
const selectedStatusId = ref('d7')
const incidentDismissed = ref(false)
const lastRefresh = ref(new Date())

const groups = computed(() => {
  const map = new Map<string, Device[]>()
  for (const device of devices) {
    const list = map.get(device.group) ?? []
    list.push(device)
    map.set(device.group, list)
  }
  return Array.from(map, ([name, list]) => ({ name, devices: list }))
})

const filteredGroups = computed(() => {
  const query = deviceQuery.value.toLowerCase()
  if (!query) return groups.value
  return groups.value
    .map((group) => ({
      name: group.name,
      devices: group.devices.filter((device) => device.name.toLowerCase().includes(query) || group.name.toLowerCase().includes(query)),
    }))
    .filter((group) => group.devices.length)
})

const selectedDevice = computed(() => devices.find((device) => device.id === selectedDeviceId.value) ?? devices[0])
const selectedStatusDevice = computed(() => devices.find((device) => device.id === selectedStatusId.value) ?? devices[0])

const configs = reactive<Record<string, DeviceConfig>>({})
const savedConfigs = reactive<Record<string, DeviceConfig>>({})

for (const device of devices) {
  const config: DeviceConfig = { module: '日常巡检', start: '09:00', end: '22:00', retry: 2, auto: true, notify: true }
  configs[device.id] = { ...config }
  savedConfigs[device.id] = { ...config }
}

const config = computed(() => configs[selectedDevice.value.id])
const dirty = computed(() => JSON.stringify(configs[selectedDevice.value.id]) !== JSON.stringify(savedConfigs[selectedDevice.value.id]))

const toast = ref('')
let toastTimer: ReturnType<typeof setTimeout> | undefined

function showToast(text: string) {
  toast.value = text
  if (toastTimer) clearTimeout(toastTimer)
  toastTimer = setTimeout(() => {
    toast.value = ''
  }, 2400)
}

function toggleGroup(name: string) {
  if (collapsed.value.includes(name)) {
    collapsed.value = collapsed.value.filter((item) => item !== name)
  } else {
    collapsed.value = [...collapsed.value, name]
  }
}

function addAccount() {
  const device = selectedDevice.value
  const n = device.accounts.length + 1
  device.accounts.push({
    id: `${device.id}-a${n}-${Date.now()}`,
    name: `账号 ${String(n).padStart(2, '0')}`,
    platform: '抖音',
    module: '日常巡检',
    status: 'pending',
  })
  showToast(`已添加账号，待授权后开始执行（3 元/号/月）`)
}

function removeAccount(account: Account) {
  const device = selectedDevice.value
  const index = device.accounts.findIndex((item) => item.id === account.id)
  if (index === -1) return
  device.accounts.splice(index, 1)
  showToast(`已删除 ${account.name}，下个计费周期减少 3 元`)
}

function toggleAccount(account: Account) {
  if (account.status === 'pending') {
    showToast('待授权账号需要先在 Auth 完成授权')
    return
  }
  account.status = account.status === 'running' ? 'stopped' : 'running'
  showToast(account.status === 'running' ? `${account.name} 已恢复执行` : `${account.name} 已暂停执行`)
}

function saveConfig() {
  const device = selectedDevice.value
  savedConfigs[device.id] = { ...configs[device.id] }
  showToast('配置已保存，设备下次拉取快照后生效')
}

function resetConfig() {
  const device = selectedDevice.value
  configs[device.id] = { ...savedConfigs[device.id] }
  showToast('已放弃未保存的修改')
}

const statusDevices = computed(() => {
  const query = statusQuery.value.toLowerCase()
  return devices.filter((device) => {
    const matchFilter = statusFilter.value === 'all' || device.status === statusFilter.value
    const matchQuery = !query || device.name.toLowerCase().includes(query) || device.group.toLowerCase().includes(query)
    return matchFilter && matchQuery
  })
})

function runningOf(device: Device) {
  if (device.status !== 'online') return 0
  return device.accounts.filter((account) => account.status === 'running').length
}

const kpis = computed(() => [
  {
    label: '在线设备',
    value: `${devices.filter((device) => device.status === 'online').length} / ${devices.length}`,
    hint: '离线设备不计费但仍保留配置',
    tone: 'text-slate-900',
  },
  {
    label: '运行账号',
    value: `${devices.reduce((sum, device) => sum + runningOf(device), 0)} / ${devices.reduce((sum, device) => sum + device.accounts.length, 0)}`,
    hint: `按号计费，本月 ${devices.reduce((sum, device) => sum + device.accounts.length, 0) * 3} 元`,
    tone: 'text-brand-700',
  },
  {
    label: '今日执行',
    value: devices.reduce((sum, device) => sum + device.executions, 0).toLocaleString('zh-CN'),
    hint: '设备上报的执行次数合计',
    tone: 'text-slate-900',
  },
  {
    label: '待处理',
    value: `${devices.filter((device) => device.status !== 'online').length} 项`,
    hint: '离线与执行异常待确认',
    tone: 'text-red-600',
  },
])

function focusIncident() {
  statusFilter.value = 'error'
  statusQuery.value = ''
  selectedStatusId.value = 'd7'
  tab.value = 'status'
  showToast('已筛选异常设备，定位到设备 07')
}

function pullSnapshot() {
  showToast(`${selectedStatusDevice.value.name} 已发起快照拉取，等待设备响应`)
}

const notifications = ref([
  { id: 'n1', type: '授权', text: '门店 A / 账号 03 授权将在 3 天后到期。', time: '10:24' },
  { id: 'n2', type: '异常', text: '设备 07 连续 3 次执行失败，已暂停自动执行。', time: '09:48' },
  { id: 'n3', type: '设备', text: '设备 09 离线超过 10 分钟，最后一次心跳 09:31。', time: '09:31' },
])

const lastRefreshLabel = computed(() =>
  lastRefresh.value.toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit', second: '2-digit', hour12: false }),
)

function refresh() {
  lastRefresh.value = new Date()
  for (const device of devices) {
    if (device.status === 'online') device.heartbeat = '刚刚'
  }
  showToast('状态已刷新')
}

function deviceLabel(status: DeviceStatus) {
  return status === 'online' ? '在线' : status === 'offline' ? '离线' : '异常'
}

function accountStatusLabel(status: AccountStatus) {
  return { running: '运行中', stopped: '已停止', pending: '待授权', error: '异常' }[status]
}

function devicePill(status: DeviceStatus) {
  return status === 'online'
    ? 'bg-accent-500/15 text-accent-600'
    : status === 'error'
      ? 'bg-red-100 text-red-700'
      : 'bg-slate-100 text-slate-500'
}

function accountPill(status: AccountStatus) {
  return status === 'running'
    ? 'bg-accent-500/15 text-accent-600'
    : status === 'error'
      ? 'bg-red-100 text-red-700'
      : status === 'pending'
        ? 'bg-amber-100 text-amber-700'
        : 'bg-slate-100 text-slate-500'
}

function deviceDot(status: DeviceStatus) {
  return status === 'online' ? 'bg-accent-500' : status === 'error' ? 'bg-red-500' : 'bg-slate-300'
}

function notePill(type: string) {
  return type === '异常'
    ? 'bg-red-100 text-red-700'
    : type === '授权'
      ? 'bg-amber-100 text-amber-700'
      : 'bg-slate-200 text-slate-600'
}
</script>
