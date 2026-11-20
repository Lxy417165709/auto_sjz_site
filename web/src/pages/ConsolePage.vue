<template>
  <div>
    <PageHero
      eyebrow="在线体验"
      title="真实的 Manager 界面，直接在这里点"
      desc="下面不是截图，而是 Manager 前端的 demo 构建：同一套 Vue 组件、同一套业务配置与业务状态页面，接演示数据运行。展开设备、切账号、看模块都能操作，但不会写入真实环境。"
    >
      <div class="flex flex-wrap gap-2">
        <span class="tag">同一份前端代码</span>
        <span class="tag">业务配置 / 业务状态</span>
        <span class="tag">演示数据</span>
        <span class="tag">可交互</span>
      </div>
    </PageHero>

    <div class="console-wrap">
      <div class="console-frame">
        <div class="frame-bar">
          <div class="frame-title">
            <span class="dot-row" aria-hidden="true"><i></i><i></i><i></i></span>
            <span class="frame-name">auto_sjz Manager</span>
            <span class="frame-badge">DEMO</span>
            <span class="frame-hint">演示环境 · 不写真实数据</span>
          </div>
          <a :href="demoHref" target="_blank" rel="noopener" class="frame-open">
            新窗口打开
            <svg class="h-3.5 w-3.5" viewBox="0 0 24 24" fill="none" aria-hidden="true">
              <path d="M14 4h6v6M20 4l-9 9M18 14v5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h5" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" />
            </svg>
          </a>
        </div>

        <iframe
          :src="demoHref"
          class="console-iframe"
          title="auto_sjz Manager 演示环境"
          loading="eager"
        ></iframe>
      </div>

      <div class="mt-4 flex flex-col gap-2 text-[12px] leading-6 text-slate-500 sm:flex-row sm:items-center sm:justify-between">
        <p>界面里左侧「业务配置 / 业务状态」就是 Manager 的两个真实工作区；账号、设备、数值均为演示数据。</p>
        <p class="text-slate-400">业务状态表格可横向滚动，账号时效在最右侧</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import PageHero from '@/components/PageHero.vue'

const demoBase = (import.meta.env.VITE_CONSOLE_DEMO_URL || `${import.meta.env.BASE_URL}console-demo/`).replace(/\/?$/, '/')
const demoHref = `${demoBase}#/accounts`
</script>

<style scoped>
.console-wrap {
  width: 100%;
  max-width: 1480px;
  margin: 0 auto;
  padding: 24px 16px 32px;
}
@media (min-width: 640px) {
  .console-wrap { padding: 28px 24px 36px; }
}

.console-frame {
  overflow: hidden;
  border: 1px solid #dce5f2;
  border-radius: 16px;
  background: #f4f7fb;
  box-shadow: 0 24px 60px -32px rgba(28, 58, 118, 0.35);
}
.frame-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  padding: 10px 14px;
  border-bottom: 1px solid #e3eaf4;
  background: linear-gradient(180deg, #f8fbff, #f1f5fb);
}
.frame-title {
  display: flex;
  min-width: 0;
  align-items: center;
  gap: 10px;
}
.dot-row {
  display: inline-flex;
  gap: 5px;
}
.dot-row i {
  width: 7px;
  height: 7px;
  border-radius: 50%;
  background: #cbd5e1;
}
.frame-name {
  font-size: 13px;
  font-weight: 650;
  color: #1f2937;
  white-space: nowrap;
}
.frame-badge {
  padding: 1px 7px;
  border: 1px solid rgba(47, 107, 255, 0.28);
  border-radius: 999px;
  background: rgba(47, 107, 255, 0.1);
  color: #1d5fd8;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.04em;
}
.frame-hint {
  overflow: hidden;
  color: #8492a6;
  font-size: 11px;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.frame-open {
  display: inline-flex;
  flex: none;
  align-items: center;
  gap: 5px;
  padding: 5px 10px;
  border: 1px solid #d7e2f0;
  border-radius: 8px;
  background: #fff;
  color: #33507a;
  font-size: 11px;
  font-weight: 600;
  transition: border-color 0.15s ease, color 0.15s ease;
}
.frame-open:hover {
  border-color: #b9cdf0;
  color: #1d5fd8;
}
.console-iframe {
  display: block;
  width: 100%;
  height: clamp(700px, 84vh, 1000px);
  border: 0;
  background: #f4f7fb;
}
@media (max-width: 640px) {
  .frame-hint {
    display: none;
  }
  .frame-bar {
    padding: 9px 11px;
  }
  .console-iframe {
    height: 760px;
  }
}
</style>
