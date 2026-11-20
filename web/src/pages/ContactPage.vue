<template>
  <div>
    <PageHero
      eyebrow="联系与咨询"
      title="说说设备规模，我们给出接入方案"
      desc="需要报价、接入方案或私有化部署，填一下基本信息即可。官网不保存账号与订单，登录和购买仍在 Manager 完成。"
    />

    <div class="wrap grid gap-8 py-14 sm:py-16 lg:grid-cols-[1.1fr_0.9fr]">
      <form class="rounded-3xl border border-slate-200 p-6 shadow-card sm:p-8" @submit.prevent="submit">
        <h2 class="text-lg font-semibold text-slate-900">提交接入咨询</h2>
        <p class="mt-1 text-[13px] text-slate-500">带 <span class="text-brand-600">*</span> 为必填，邮箱或电话留一个即可。</p>

        <div class="mt-6 grid gap-4 sm:grid-cols-2">
          <label class="block">
            <span class="text-xs font-medium text-slate-600">称呼 *</span>
            <input v-model.trim="form.name" class="field mt-1.5" placeholder="怎么称呼你" autocomplete="name" />
          </label>
          <label class="block">
            <span class="text-xs font-medium text-slate-600">公司 / 团队</span>
            <input v-model.trim="form.company" class="field mt-1.5" placeholder="选填" autocomplete="organization" />
          </label>
          <label class="block">
            <span class="text-xs font-medium text-slate-600">邮箱</span>
            <input v-model.trim="form.email" class="field mt-1.5" type="email" placeholder="name@company.com" autocomplete="email" />
          </label>
          <label class="block">
            <span class="text-xs font-medium text-slate-600">电话 / 微信</span>
            <input v-model.trim="form.phone" class="field mt-1.5" placeholder="选填" autocomplete="tel" />
          </label>
          <label class="block">
            <span class="text-xs font-medium text-slate-600">预计设备台数</span>
            <input v-model.number="form.devices" class="field mt-1.5" type="number" min="0" max="100000" placeholder="例如 20" />
          </label>
          <label class="block">
            <span class="text-xs font-medium text-slate-600">每台账号数</span>
            <input v-model.number="form.accounts" class="field mt-1.5" type="number" min="0" max="1000" placeholder="例如 5" />
          </label>
        </div>

        <label class="mt-4 block">
          <span class="text-xs font-medium text-slate-600">需求说明</span>
          <textarea
            v-model.trim="form.message"
            class="field mt-1.5 min-h-[110px] resize-y"
            placeholder="例如：已有 20 台设备，每台 5 个账号，想了解部署方式和授权数量"
          ></textarea>
        </label>

        <div class="hidden" aria-hidden="true">
          <label>网站<input v-model="form.website" tabindex="-1" autocomplete="off" /></label>
        </div>

        <p v-if="errorMsg" class="mt-4 rounded-xl border border-red-200 bg-red-50 px-4 py-3 text-[13px] text-red-700">{{ errorMsg }}</p>
        <p v-if="okMsg" class="mt-4 rounded-xl border border-accent-400/40 bg-accent-500/10 px-4 py-3 text-[13px] text-accent-600">{{ okMsg }}</p>

        <div class="mt-6 flex flex-wrap items-center gap-4">
          <button type="submit" class="btn-primary" :disabled="submitting">
            {{ submitting ? '提交中…' : '提交咨询' }}
          </button>
          <span class="text-[11px] text-slate-400">提交后我们会用你留下的方式联系，不会用于其他用途。</span>
        </div>
      </form>

      <aside class="space-y-4">
        <div class="rounded-3xl border border-[#f0dfd3] bg-[#fffaf4] p-6">
          <p class="eyebrow">联系方式</p>
          <dl class="mt-4 space-y-4 text-sm">
            <div>
              <dt class="text-xs text-slate-500">邮箱</dt>
              <dd class="mt-1 font-medium text-slate-900">
                <a v-if="contact?.email" :href="`mailto:${contact.email}`" class="underline decoration-slate-300 underline-offset-4">
                  {{ contact.email }}
                </a>
                <span v-else class="text-slate-400">待配置</span>
              </dd>
            </div>
            <div v-if="contact?.wechat">
              <dt class="text-xs text-slate-500">微信</dt>
              <dd class="mt-1 font-medium text-slate-900">{{ contact.wechat }}</dd>
            </div>
            <div>
              <dt class="text-xs text-slate-500">响应时间</dt>
              <dd class="mt-1 text-slate-700">{{ contact?.working || '工作日 9:00 - 18:00（UTC+8）' }}</dd>
            </div>
          </dl>
        </div>

        <div class="rounded-3xl border border-slate-200 p-6">
          <p class="eyebrow">提交前可以先准备</p>
          <ul class="mt-3 space-y-2 text-[13px] leading-6 text-slate-600">
            <li>设备大致台数与每台账号数</li>
            <li>计划单机部署还是分开部署</li>
            <li>设备所在网络是否能访问公网或指定内网地址</li>
          </ul>
          <p class="mt-4 text-[12px] leading-5 text-slate-400">
            官网不接入 Manager 登录与订单系统，正式开通与续费在 Manager 内完成。
          </p>
        </div>
      </aside>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue'
import PageHero from '@/components/PageHero.vue'
import { siteState } from '@/lib/site'

const contact = computed(() => siteState.data?.site.contact ?? null)

const form = reactive({
  name: '',
  company: '',
  email: '',
  phone: '',
  devices: 0,
  accounts: 0,
  message: '',
  website: '',
})

const submitting = ref(false)
const okMsg = ref('')
const errorMsg = ref('')

async function submit() {
  okMsg.value = ''
  errorMsg.value = ''

  if (!form.name) {
    errorMsg.value = '请填写称呼'
    return
  }
  if (!form.email && !form.phone) {
    errorMsg.value = '请至少留下邮箱或电话'
    return
  }

  submitting.value = true
  try {
    const resp = await fetch('/api/lead', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        name: form.name,
        company: form.company,
        email: form.email,
        phone: form.phone,
        devices: Number(form.devices) || 0,
        message: [form.message, form.accounts ? `每台账号数：${form.accounts}` : ''].filter(Boolean).join('；'),
        source: 'site-contact',
        website: form.website,
      }),
    })
    if (resp.status === 429) {
      errorMsg.value = '提交过于频繁，请稍后再试'
      return
    }
    if (!resp.ok) {
      const data = await resp.json().catch(() => null)
      errorMsg.value = data?.error?.message || '提交失败，请直接邮件联系我们'
      return
    }
    okMsg.value = '已收到，我们会尽快联系你。'
    form.name = ''
    form.company = ''
    form.email = ''
    form.phone = ''
    form.devices = 0
    form.accounts = 0
    form.message = ''
  } catch {
    errorMsg.value = '网络异常，提交失败，请直接邮件联系我们'
  } finally {
    submitting.value = false
  }
}
</script>
