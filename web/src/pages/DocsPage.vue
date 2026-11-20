<template>
  <div>
    <PageHero
      eyebrow="文档"
      title="接入、授权、配置，三件事写清楚"
      desc="不追求大而全，只把上线必须知道的链路、端口与时效规则整理出来；细节以对应版本的部署文档为准。"
    >
      <div class="flex flex-wrap gap-2">
        <span class="tag">快速开始</span>
        <span class="tag">系统架构</span>
        <span class="tag">授权计费</span>
        <span class="tag">配置下发</span>
      </div>
    </PageHero>

    <div class="wrap">
      <div v-if="!docs.length" class="py-16">
        <div v-if="siteState.error" class="rounded-2xl border border-red-200 bg-red-50 p-6 text-sm text-red-700">
          文档加载失败（{{ siteState.error }}）。
          <button class="ml-2 font-medium underline" @click="reload">重新加载</button>
        </div>
        <div v-else class="grid gap-4 sm:grid-cols-2">
          <div v-for="i in 4" :key="i" class="h-28 animate-pulse rounded-2xl bg-slate-100"></div>
        </div>
      </div>

      <div v-else class="grid gap-10 py-14 sm:py-16 lg:grid-cols-[200px_1fr]">
        <aside class="lg:sticky lg:top-24 lg:h-fit">
          <p class="eyebrow">目录</p>
          <nav class="mt-3 space-y-1" aria-label="文档目录">
            <a
              v-for="doc in docs"
              :key="doc.slug"
              :href="`#${doc.slug}`"
              class="block rounded-lg px-3 py-2 text-sm text-slate-600 transition hover:bg-slate-50 hover:text-slate-900"
            >
              {{ doc.title }}
            </a>
          </nav>
          <div class="mt-6 rounded-xl border border-slate-200 bg-slate-50 p-4 text-[12px] leading-5 text-slate-500">
            文档随版本更新。生产环境请以对应发布版本里的部署文档为准。
          </div>
        </aside>

        <div class="space-y-8">
          <article v-for="(doc, index) in docs" :id="doc.slug" :key="doc.slug" class="scroll-mt-24 rounded-3xl border border-slate-200 p-6 sm:p-8">
            <div class="flex flex-wrap items-center gap-3">
              <span class="font-mono text-[11px] text-slate-400">DOC {{ String(index + 1).padStart(2, '0') }}</span>
              <span v-if="doc.updated" class="tag">更新于 {{ doc.updated }}</span>
            </div>
            <h2 class="mt-2 text-xl font-semibold tracking-tight text-slate-900">{{ doc.title }}</h2>
            <p class="mt-2 text-sm leading-6 text-slate-500">{{ doc.summary }}</p>

            <div v-for="section in doc.sections" :key="section.heading" class="mt-6">
              <h3 class="text-sm font-semibold text-slate-900">{{ section.heading }}</h3>
              <div class="prose-doc mt-2">
                <p v-for="p in section.paragraphs" :key="p">{{ p }}</p>
                <ul v-if="section.bullets?.length">
                  <li v-for="b in section.bullets" :key="b">{{ b }}</li>
                </ul>
              </div>
            </div>
          </article>
        </div>
      </div>

      <CtaBand
        title="文档没覆盖到的，直接问"
        desc="接入报错、地址解析异常、授权到期行为，把环境信息发我们，一起排查。"
        primary-label="联系技术支持"
        primary-to="/contact"
        secondary-label="查看常见问题"
        secondary-to="/faq"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import PageHero from '@/components/PageHero.vue'
import CtaBand from '@/components/CtaBand.vue'
import { loadSite, siteState } from '@/lib/site'

const docs = computed(() => siteState.data?.docs ?? [])

function reload() {
  void loadSite()
}
</script>
