<template>
  <div>
    <PageHero
      eyebrow="接入流程"
      title="五步接入，双方各做什么"
      desc="接入方负责申请、配置和运行，我方负责部署和发放令牌。五步走完，设备开始执行；授权按账号计，3 元/号/月。"
    >
      <div class="flex flex-wrap gap-2">
        <span class="tag">5 步接入</span>
        <span class="tag">接入方 3 步</span>
        <span class="tag">我方 2 步</span>
        <span class="tag">按号授权</span>
      </div>
    </PageHero>

    <div class="wrap">
      <section class="py-12 sm:py-16">
        <div class="flow-board">
          <header class="flow-head">
            <div>
              <p class="eyebrow">流程总览</p>
              <h2 class="mt-2 text-xl font-semibold tracking-tight text-slate-900 sm:text-2xl">从申请到设备运行，一条线走完</h2>
            </div>
            <div class="owner-legend">
              <span class="legend-pill is-client">接入方</span>
              <span class="legend-pill is-ops">我方</span>
            </div>
          </header>

          <div class="flow-track">
            <template v-for="(step, i) in steps" :key="step.no">
              <article class="flow-card" :class="`is-${step.tone}`">
                <span class="flow-ghost" aria-hidden="true">{{ step.no }}</span>
                <div class="flow-top">
                  <span class="flow-no">{{ step.no }}</span>
                  <span class="owner-pill">{{ step.owner }}</span>
                </div>
                <h3 class="flow-title">{{ step.title }}</h3>
                <p class="flow-desc">{{ step.desc }}</p>
                <div class="flow-output">
                  <span>{{ step.outLabel }}</span>
                  <b>{{ step.output }}</b>
                </div>
              </article>
              <div v-if="i < steps.length - 1" class="flow-arrow" aria-hidden="true">
                <span>→</span>
              </div>
            </template>
          </div>

          <div class="flow-foot">
            <span class="foot-dot"></span>
            <span>接入方：申请 → 配置 → 运行</span>
            <span class="foot-sep"></span>
            <span class="foot-dot is-ops"></span>
            <span>我方：部署 → 发放令牌</span>
          </div>
        </div>
      </section>

      <section class="border-t border-slate-200 py-12 sm:py-14">
        <div class="grid gap-4 md:grid-cols-3">
          <article v-for="item in notes" :key="item.no" class="note-card">
            <span class="note-no">{{ item.no }}</span>
            <h3>{{ item.title }}</h3>
            <p>{{ item.desc }}</p>
          </article>
        </div>
      </section>

      <section class="border-t border-slate-200 py-12 sm:py-14">
        <div class="prep-strip">
          <div>
            <p class="eyebrow">开始前</p>
            <h2 class="mt-2 text-lg font-semibold tracking-tight text-slate-900">准备这四项，流程会快很多</h2>
          </div>
          <div class="prep-items">
            <span v-for="item in prep" :key="item" class="prep-item">{{ item }}</span>
          </div>
        </div>
        <p class="mt-4 text-[12px] leading-6 text-slate-400">
          流程为通用说明，管理台菜单名称以实际版本为准；部署方式、端口和授权验签见「技术架构」与「部署接入」。
        </p>
      </section>

      <CtaBand
        title="先接通一台，再批量铺开"
        desc="把设备台数、每台账号数和要跑的模块发给我们，我们按这个流程给出部署顺序、令牌数量和月费。"
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

type Tone = 'client' | 'ops'

const steps: { no: string; owner: string; tone: Tone; title: string; desc: string; outLabel: string; output: string }[] = [
  {
    no: '01', owner: '接入方', tone: 'client', title: '申请管理中台',
    desc: '提交设备台数、账号数量和使用场景，确认部署方式、接入顺序与月费。',
    outLabel: '准备', output: '设备 / 账号清单',
  },
  {
    no: '02', owner: '我方', tone: 'ops', title: '部署管理中台',
    desc: '部署 Manager 与 discovery，开通登录地址，交付初始管理员账号。',
    outLabel: '交付', output: '登录地址',
  },
  {
    no: '03', owner: '接入方', tone: 'client', title: '登录配置并申请授权',
    desc: '登录管理台，配置设备、账号、模块；确认配置无误后提交授权申请。',
    outLabel: '提交', output: '授权申请',
  },
  {
    no: '04', owner: '我方', tone: 'ops', title: '发放令牌',
    desc: '按申请账号数量发放授权令牌；令牌决定账号有效期，设备按令牌执行。',
    outLabel: '发放', output: '账号令牌',
  },
  {
    no: '05', owner: '接入方', tone: 'client', title: '设备运行',
    desc: '设备拉取配置与令牌后开始执行；后续加减账号、调整模块仍在管理台完成。',
    outLabel: '结果', output: '设备上线运行',
  },
]

const notes = [
  {
    no: 'A',
    title: '申请一次给全',
    desc: '设备台数、每台账号数、要跑的模块一起给，部署和令牌数量一次算清，少来回。',
  },
  {
    no: 'B',
    title: '部署后先验通',
    desc: '拿到登录地址先确认管理员能登录、设备网络能访问，再进入配置，避免后段卡住。',
  },
  {
    no: 'C',
    title: '令牌按号发',
    desc: '授权按账号计，3 元/号/月；设备台数不计费。加号补令牌，减号下个周期生效。',
  },
]

const prep = ['设备台数与分布', '每台账号数量', '可访问管理台的网络', '配置负责人']
</script>

<style scoped>
.flow-board {
  overflow: hidden;
  border: 1px solid #dce5f2;
  border-radius: 22px;
  background:
    radial-gradient(120% 100% at 0% 0%, rgba(47, 107, 255, 0.06), transparent 42%),
    radial-gradient(100% 100% at 100% 0%, rgba(183, 121, 31, 0.06), transparent 38%),
    #fff;
  box-shadow: 0 24px 60px -42px rgba(28, 58, 118, 0.45);
}
.flow-head {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  gap: 20px;
  padding: 26px 26px 20px;
  border-bottom: 1px solid #e8eef7;
}
.owner-legend {
  display: flex;
  flex: none;
  align-items: center;
  gap: 8px;
}
.legend-pill,
.owner-pill {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border-radius: 999px;
  font-size: 11px;
  font-weight: 600;
  line-height: 1;
  white-space: nowrap;
}
.legend-pill {
  padding: 7px 12px;
}
.legend-pill.is-client {
  border: 1px solid rgba(47, 107, 255, 0.24);
  background: rgba(47, 107, 255, 0.09);
  color: #1d5fd8;
}
.legend-pill.is-ops {
  border: 1px solid rgba(183, 121, 31, 0.26);
  background: rgba(183, 121, 31, 0.1);
  color: #9a6417;
}
.flow-track {
  display: flex;
  align-items: stretch;
  padding: 24px 26px 26px;
}
.flow-card {
  position: relative;
  flex: 1 1 0;
  min-width: 0;
  overflow: hidden;
  border: 1px solid #e6ebf3;
  border-top-width: 3px;
  border-radius: 16px;
  background: #fff;
  padding: 16px 15px 14px;
  box-shadow: 0 12px 28px -24px rgba(28, 58, 118, 0.5);
}
.flow-card.is-client {
  border-top-color: #2f6bff;
  background: linear-gradient(180deg, rgba(47, 107, 255, 0.045), #fff 34%);
}
.flow-card.is-ops {
  border-top-color: #b7791f;
  background: linear-gradient(180deg, rgba(183, 121, 31, 0.05), #fff 34%);
}
.flow-ghost {
  position: absolute;
  top: 4px;
  right: 10px;
  color: rgba(15, 23, 42, 0.055);
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 42px;
  font-weight: 700;
  line-height: 1;
  pointer-events: none;
}
.flow-top {
  position: relative;
  z-index: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 6px;
}
.flow-no {
  color: #8a98ab;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.06em;
}
.owner-pill {
  padding: 5px 8px;
}
.flow-card.is-client .owner-pill {
  border: 1px solid rgba(47, 107, 255, 0.24);
  background: rgba(47, 107, 255, 0.09);
  color: #1d5fd8;
}
.flow-card.is-ops .owner-pill {
  border: 1px solid rgba(183, 121, 31, 0.26);
  background: rgba(183, 121, 31, 0.1);
  color: #9a6417;
}
.flow-title {
  margin-top: 18px;
  color: #172033;
  font-size: 15px;
  font-weight: 650;
  line-height: 1.45;
  letter-spacing: -0.01em;
}
.flow-desc {
  margin-top: 7px;
  color: #64748b;
  font-size: 12px;
  line-height: 1.75;
}
.flow-output {
  display: flex;
  align-items: baseline;
  gap: 8px;
  margin-top: 14px;
  border-top: 1px dashed #e5eaf2;
  padding-top: 10px;
  font-size: 11px;
  line-height: 1.4;
}
.flow-output span {
  flex: none;
  color: #97a3b4;
}
.flow-output b {
  overflow: hidden;
  color: #33415a;
  font-weight: 600;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.flow-arrow {
  display: flex;
  flex: 0 0 30px;
  align-items: center;
  justify-content: center;
  color: #a9b6c8;
  font-size: 17px;
}
.flow-foot {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 8px;
  border-top: 1px solid #e8eef7;
  padding: 14px 26px;
  color: #64748b;
  font-size: 11px;
}
.foot-dot {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #2f6bff;
}
.foot-dot.is-ops {
  margin-left: 8px;
  background: #b7791f;
}
.foot-sep {
  width: 1px;
  height: 12px;
  margin: 0 8px;
  background: #dbe3ef;
}
.note-card {
  border: 1px solid #e3e9f3;
  border-radius: 16px;
  background: #fff;
  padding: 18px 18px 16px;
}
.note-no {
  display: inline-flex;
  width: 22px;
  height: 22px;
  align-items: center;
  justify-content: center;
  border-radius: 7px;
  background: #f1f5fb;
  color: #5d6d84;
  font-family: ui-monospace, SFMono-Regular, Menlo, monospace;
  font-size: 11px;
  font-weight: 700;
}
.note-card h3 {
  margin-top: 10px;
  color: #172033;
  font-size: 14px;
  font-weight: 650;
}
.note-card p {
  margin-top: 6px;
  color: #64748b;
  font-size: 12px;
  line-height: 1.75;
}
.prep-strip {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 24px;
  border: 1px solid #e3e9f3;
  border-radius: 18px;
  background: #f8fafc;
  padding: 20px 22px;
}
.prep-items {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  justify-content: flex-end;
}
.prep-item {
  border: 1px solid #dbe4f0;
  border-radius: 999px;
  background: #fff;
  padding: 7px 11px;
  color: #52637a;
  font-size: 12px;
  white-space: nowrap;
}
@media (max-width: 1023px) {
  .flow-track {
    flex-direction: column;
    padding: 20px;
  }
  .flow-card {
    flex: none;
  }
  .flow-arrow {
    flex: 0 0 28px;
    transform: rotate(90deg);
  }
  .flow-ghost {
    font-size: 52px;
  }
  .prep-strip {
    flex-direction: column;
    align-items: flex-start;
    gap: 14px;
  }
  .prep-items {
    justify-content: flex-start;
  }
}
@media (max-width: 639px) {
  .flow-head {
    flex-direction: column;
    align-items: flex-start;
    padding: 20px 18px 16px;
  }
  .flow-track {
    padding: 16px 18px;
  }
  .flow-foot {
    padding: 12px 18px;
  }
  .foot-sep {
    display: none;
  }
  .foot-dot.is-ops {
    margin-left: 0;
  }
}
</style>
