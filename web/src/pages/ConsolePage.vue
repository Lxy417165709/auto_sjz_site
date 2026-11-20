<template>
  <div>
    <PageHero
      eyebrow="管理台演示"
      title="业务配置和业务状态，就长这样"
      desc="配置页按「设备 → 账号 → 模块」组织，状态页按设备聚合账号数据。模块只展示信息、生产、出售、仓库四类，账号与数值均为示意数据。"
    >
      <div class="flex flex-wrap gap-2">
        <span class="tag">设备 → 账号 → 模块</span>
        <span class="tag">按设备聚合</span>
        <span class="tag">3 元 / 号 / 月</span>
        <span class="tag">仅四类模块</span>
        <span class="tag">可交互</span>
        <span class="tag">示意数据</span>
      </div>
    </PageHero>

    <div class="wrap py-10 sm:py-12">
      <div class="demo-shell">
        <div class="demo-topbar">
          <div class="flex min-w-0 items-center gap-3">
            <span class="brand-mark">
              <svg class="h-4 w-4" viewBox="0 0 24 24" fill="none" aria-hidden="true">
                <path d="M12 3.5c.9 2.3-.1 3.7-1 4.7-.8 1-1.4 1.9-1.4 3a2.4 2.4 0 0 0 4.8 0c0-.8-.3-1.5-.9-2.1-.1.8-.5 1.4-1.1 1.8.6-2.4-.1-4.6-1-6.4z" fill="currentColor" />
              </svg>
            </span>
            <div class="min-w-0">
              <p class="truncate text-sm font-semibold text-white">炎火 Manager</p>
              <p class="truncate text-[10px] text-slate-400">演示环境 · 不写入真实管理台</p>
            </div>
          </div>

          <div class="demo-tabs">
            <button
              v-for="item in tabs"
              :key="item.key"
              type="button"
              class="demo-tab"
              :class="{ active: activeTab === item.key }"
              @click="activeTab = item.key"
            >
              {{ item.label }}
            </button>
          </div>

          <div class="topbar-right">
            <label class="mask-toggle">
              <span>脱敏</span>
              <button type="button" role="switch" :aria-checked="masked" class="switch" :class="{ on: masked }" @click="masked = !masked">
                <i></i>
              </button>
            </label>
            <span class="hidden text-[11px] text-slate-400 sm:inline">刷新于 {{ refreshLabel }}</span>
            <button type="button" class="topbar-refresh" @click="refresh">刷新</button>
          </div>
        </div>

        <!-- 业务配置 -->
        <div v-if="activeTab === 'config'" class="accounts-page">
          <aside class="acc-side">
            <div class="side-head">
              <span>设备与账号</span>
              <div class="side-ops">
                <button type="button" @click="expandAll">展开</button>
                <button type="button" @click="collapseAll">收起</button>
              </div>
            </div>

            <div v-for="device in devices" :key="device.id" class="device-group">
              <div class="device-row" :class="{ 'is-off': !device.enabled }" @click="toggleDevice(device)">
                <span class="dev-arrow" :class="{ open: isExpanded(device.id) }">▸</span>
                <span class="dev-dot" :class="deviceDot(device)"></span>
                <span class="dev-name" :title="device.name">{{ device.name }}</span>
                <span v-if="device.approval === 'pending'" class="el-tag tag-warning flex-none">待审核</span>
                <button
                  type="button"
                  class="dev-add"
                  :disabled="device.approval === 'pending'"
                  :title="device.approval === 'pending' ? '设备审核通过后才能添加账号' : '添加账号'"
                  @click.stop="openAddAccount(device)"
                >
                  +账号
                </button>
                <span class="dev-count">{{ device.accounts.length }}</span>
                <button type="button" class="dev-act" :class="{ enable: !device.enabled }" @click.stop="toggleDeviceEnabled(device)">
                  {{ device.enabled ? '禁用' : '开启' }}
                </button>
              </div>

              <div v-if="isExpanded(device.id)" class="acc-list">
                <button
                  v-for="account in device.accounts"
                  :key="account.id"
                  type="button"
                  class="acc-item"
                  :class="{ active: selectedAccountId === account.id, off: !account.enabled }"
                  @click="selectAccount(device, account)"
                >
                  <span class="acc-dot" :class="accountDot(account)"></span>
                  <span class="acc-label">{{ account.type }}：{{ displayAccount(account) }}</span>
                  <span v-if="!account.enabled" class="off-tag">已禁</span>
                  <span v-if="account.license === '已到期'" class="off-tag warn">待授权</span>
                </button>
                <p v-if="!device.accounts.length" class="empty-dev">
                  {{ device.approval === 'pending' ? '等待设备接入并审核' : '暂无账号，点 +账号 添加' }}
                </p>
              </div>
            </div>
          </aside>

          <main class="acc-main">
            <template v-if="selectedAccount">
              <div class="det-head">
                <div class="det-title">
                  <span class="det-label">{{ accountView.type }}：{{ displayAccount(accountView) }}</span>
                  <span class="dev-key" :title="currentDevice.key">{{ currentDevice.key }}</span>
                  <span class="el-tag" :class="accountView.enabled ? 'tag-success' : 'tag-info'">{{ accountView.enabled ? '已启用' : '已停用' }}</span>
                </div>
                <div class="det-ops">
                  <span class="lbl">启用</span>
                  <button type="button" role="switch" :aria-checked="accountView.enabled" class="switch" :class="{ on: accountView.enabled }" @click="toggleAccount(accountView)">
                    <i></i>
                  </button>
                  <button type="button" class="el-btn" @click="openAddModule">+ 添加模块</button>
                  <button type="button" class="el-btn" @click="syncAccount">⇄ 同步配置</button>
                  <button type="button" class="el-btn" @click="batchSync">⇉ 批量同步</button>
                  <button type="button" class="el-btn danger-text" @click="askDeleteAccount">删除账号</button>
                </div>
              </div>

              <div v-if="accountView.license === '已到期'" class="license-alert">
                <div>
                  <p class="font-medium text-amber-800">授权已到期，执行已停止</p>
                  <p class="mt-0.5 text-[12px] leading-5 text-amber-700">账号、设备和模块配置都保留，续费后继续执行，不用重新登录。</p>
                </div>
                <button type="button" class="el-btn" @click="openRenew(accountView)">去续费</button>
              </div>

              <section class="mod-section">
                <div class="mod-section-head">
                  <div>
                    <h4>账号模块</h4>
                    <p>模块决定这个号「做什么」；仓库为设备级模块，同设备账号共用。</p>
                  </div>
                  <div class="flex flex-col items-end gap-1">
                    <span class="mod-count">{{ accountView.modules.length }} 个模块</span>
                    <span class="text-[10px] text-slate-400 md:hidden">表格可左右滑动</span>
                  </div>
                </div>

                <div v-if="accountView.modules.length" class="mod-grid">
                  <section v-for="m in accountView.modules" :key="m.key" class="module-card" :class="{ 'is-off': !m.enabled }">
                    <header class="mc-head">
                      <div class="mc-title">
                        <span class="mc-name">{{ moduleLabel(m.key) }}</span>
                        <span v-if="m.scope === 'device'" class="el-tag tag-warning">设备级</span>
                        <span v-if="m.trigger === 'exclusive'" class="el-tag tag-info">独占</span>
                      </div>
                      <div class="mc-ops">
                        <span class="mc-label">优先级</span>
                        <input v-model.number="m.priority" class="prio-input" type="number" min="0" max="999" @change="normalizePriority(m)" />
                        <select v-model="m.trigger" class="mini-select" aria-label="触发方式">
                          <option value="follow">跟随</option>
                          <option value="auto">自动</option>
                          <option value="exclusive">独占</option>
                        </select>
                        <button type="button" role="switch" :aria-checked="m.enabled" class="switch" :class="{ on: m.enabled }" @click="m.enabled = !m.enabled">
                          <i></i>
                        </button>
                        <button type="button" class="el-btn-text danger" @click="removeModule(m)">删除模块</button>
                      </div>
                    </header>

                    <div class="mc-body">
                      <template v-if="m.key === 'profile'">
                        <div class="pf-top">
                          <div class="pf-cell">
                            <span class="pf-label">角色名</span>
                            <b class="pf-value">{{ masked ? displayChar(accountView.charName) : accountView.charName }}</b>
                          </div>
                          <div class="pf-cell center">
                            <span class="pf-label">段位</span>
                            <b class="pf-value brand">{{ accountView.rank }}</b>
                          </div>
                          <div class="pf-cell right">
                            <span class="pf-label">等级</span>
                            <b class="pf-value">Lv.{{ accountView.level }}</b>
                          </div>
                        </div>
                        <div class="pf-nums">
                          <div class="pf-cell">
                            <span class="pf-label">哈弗币</span>
                            <b class="pf-num success">{{ fmtNum(accountView.havoc) }} M</b>
                          </div>
                          <div class="pf-cell">
                            <span class="pf-label">总资产</span>
                            <b class="pf-num warning">{{ fmtNum(accountView.assets) }} M</b>
                          </div>
                          <div class="pf-cell">
                            <span class="pf-label">仓库占比</span>
                            <b class="pf-num">{{ accountView.ratio }}%</b>
                          </div>
                        </div>
                      </template>

                      <template v-else-if="m.key === 'afk'">
                        <div class="afk-stats">
                          <div class="afk-stat"><span>累计局数</span><b>{{ m.afk.totalGames }}</b></div>
                          <div class="afk-stat"><span>上次结果</span><b class="success">{{ m.afk.lastResult }}</b></div>
                          <div class="afk-stat"><span>上次结束</span><b class="muted">11-20 13:46</b></div>
                        </div>
                        <div class="pass-box">
                          <div class="flex items-center justify-between gap-3">
                            <span class="text-[12px] font-medium text-slate-700">🎫 本周通行证经验</span>
                            <span class="text-[11px] text-slate-400">11-17 起</span>
                          </div>
                          <div class="mt-2 flex items-center justify-between text-[12px]">
                            <b class="text-slate-800">{{ m.afk.passport.exp.toLocaleString() }} / {{ m.afk.passport.max.toLocaleString() }}</b>
                            <span class="font-medium text-emerald-600">{{ Math.round((m.afk.passport.exp / m.afk.passport.max) * 100) }}%</span>
                          </div>
                          <div class="mini-progress mt-1.5">
                            <span :style="{ width: Math.round((m.afk.passport.exp / m.afk.passport.max) * 100) + '%' }"></span>
                          </div>
                        </div>
                        <div class="afk-row">
                          <span class="mc-label">冷却时长（秒，10~3600）</span>
                          <input v-model.number="m.afk.cooldown" class="prio-input !w-[76px]" type="number" min="10" max="3600" />
                        </div>
                        <div class="afk-times">
                          <div class="flex items-center justify-between">
                            <span class="mc-label">执行时段</span>
                            <button type="button" class="el-btn-text" @click="addAfkWindow(m)">+ 添加时段</button>
                          </div>
                          <div v-for="(w, wi) in m.afk.windows" :key="wi" class="afk-win">
                            <input v-model="w.start" type="text" maxlength="5" inputmode="numeric" placeholder="22:00" class="time-input" />
                            <span class="text-slate-400">~</span>
                            <input v-model="w.end" type="text" maxlength="5" inputmode="numeric" placeholder="08:00" class="time-input" />
                            <span class="afk-win-text">{{ windowText(w) }}</span>
                            <button type="button" class="el-btn-text danger" @click="removeAfkWindow(m, wi)">删除</button>
                          </div>
                        </div>
                      </template>

                      <template v-else-if="m.key === 'produce'">
                        <div class="mc-table-scroll"><table class="mc-table">
                          <thead>
                            <tr><th style="width: 112px">特勤部门</th><th>产物</th><th style="width: 126px">当前制造</th><th style="width: 88px">倒计时</th><th style="width: 58px">启用</th><th style="width: 54px"></th></tr>
                          </thead>
                          <tbody>
                            <tr v-for="(item, ii) in m.produce.items" :key="ii">
                              <td>{{ item.dept }}</td>
                              <td>
                                <select v-model="item.prop" class="mini-select !w-full !max-w-[150px]">
                                  <option v-for="opt in produceOptions(item.dept)" :key="opt" :value="opt">{{ opt }}</option>
                                </select>
                              </td>
                              <td><span class="mc-chip" :class="item.status === '制造中' ? 'busy' : 'done'">{{ item.name }}</span></td>
                              <td class="muted">{{ item.left }}</td>
                              <td><button type="button" role="switch" :aria-checked="item.enabled" class="switch !h-[15px] !w-[30px]" :class="{ on: item.enabled }" @click="item.enabled = !item.enabled"><i></i></button></td>
                              <td><button type="button" class="el-btn-text danger" @click="removeProduceItem(m, ii)">移除</button></td>
                            </tr>
                          </tbody>
                        </table></div>
                        <div class="add-row"><button type="button" class="el-btn-text" @click="addProduceItem(m)">+ 添加生产物品</button></div>
                      </template>

                      <template v-else-if="m.key === 'sell'">
                        <div class="mc-table-scroll"><table class="mc-table">
                          <thead><tr><th>物品</th><th style="width: 112px">出售方式</th><th style="width: 100px">最低价</th><th style="width: 58px">启用</th><th style="width: 54px"></th></tr></thead>
                          <tbody>
                            <tr v-for="(item, ii) in m.sell.items" :key="ii">
                              <td>{{ item.prop }}</td>
                              <td>
                                <select v-model="item.mode" class="mini-select !w-full">
                                  <option value="market">集市最低价</option>
                                  <option value="single">单件出售</option>
                                  <option value="all">全部出售</option>
                                </select>
                              </td>
                              <td><input v-model.number="item.price" class="prio-input !w-[74px]" type="number" min="0" /></td>
                              <td><button type="button" role="switch" :aria-checked="item.enabled" class="switch !h-[15px] !w-[30px]" :class="{ on: item.enabled }" @click="item.enabled = !item.enabled"><i></i></button></td>
                              <td><button type="button" class="el-btn-text danger" @click="removeSellItem(m, ii)">移除</button></td>
                            </tr>
                          </tbody>
                        </table></div>
                        <div class="add-row"><button type="button" class="el-btn-text" @click="addSellItem(m)">+ 添加出售物品</button></div>
                      </template>

                      <template v-else-if="m.key === 'buy'">
                        <div class="mc-table-scroll"><table class="mc-table">
                          <thead><tr><th>物品</th><th style="width: 58px">当前</th><th style="width: 64px">目标</th><th style="width: 76px">最高价</th><th style="width: 98px">定时</th><th style="width: 58px">启用</th></tr></thead>
                          <tbody>
                            <tr v-for="(item, ii) in m.buy.items" :key="ii">
                              <td>{{ item.prop }}</td>
                              <td class="muted">{{ item.current }}</td>
                              <td><input v-model.number="item.target" class="prio-input !w-[56px]" type="number" min="0" /></td>
                              <td><input v-model.number="item.maxPrice" class="prio-input !w-[68px]" type="number" min="0" /></td>
                              <td>
                                <select v-model="item.schedule" class="mini-select !w-full">
                                  <option value="每2小时">每 2 小时</option>
                                  <option value="每6小时">每 6 小时</option>
                                  <option value="每天">每天</option>
                                </select>
                              </td>
                              <td><button type="button" role="switch" :aria-checked="item.enabled" class="switch !h-[15px] !w-[30px]" :class="{ on: item.enabled }" @click="item.enabled = !item.enabled"><i></i></button></td>
                            </tr>
                          </tbody>
                        </table></div>
                        <div class="add-row"><button type="button" class="el-btn-text" @click="addBuyItem(m)">+ 添加购买物品</button></div>
                      </template>

                      <template v-else-if="m.key === 'stock'">
                        <div class="mc-table-scroll"><table class="mc-table">
                          <thead><tr><th>物品</th><th style="width: 64px">数量</th><th style="width: 100px">定时</th><th style="width: 78px">倒计时</th><th style="width: 58px">启用</th></tr></thead>
                          <tbody>
                            <tr v-for="(item, ii) in m.stock.items" :key="ii">
                              <td>{{ item.prop }}</td>
                              <td>{{ item.count }}</td>
                              <td class="muted">{{ item.schedule }}</td>
                              <td class="muted">{{ item.left }}</td>
                              <td><button type="button" role="switch" :aria-checked="item.enabled" class="switch !h-[15px] !w-[30px]" :class="{ on: item.enabled }" @click="item.enabled = !item.enabled"><i></i></button></td>
                            </tr>
                          </tbody>
                        </table></div>
                      </template>

                      <template v-else-if="m.key === 'exchange'">
                        <div class="mc-table-scroll"><table class="mc-table">
                          <thead><tr><th>物品</th><th style="width: 112px">渠道</th><th style="width: 82px">倒计时</th><th style="width: 58px">启用</th></tr></thead>
                          <tbody>
                            <tr v-for="(item, ii) in m.exchange.items" :key="ii">
                              <td>{{ item.prop }}</td>
                              <td>
                                <select v-model="item.channel" class="mini-select !w-full"><option value="market">集市</option><option value="ammo">子弹</option></select>
                              </td>
                              <td class="muted">{{ item.left }}</td>
                              <td><button type="button" role="switch" :aria-checked="item.enabled" class="switch !h-[15px] !w-[30px]" :class="{ on: item.enabled }" @click="item.enabled = !item.enabled"><i></i></button></td>
                            </tr>
                          </tbody>
                        </table></div>
                      </template>

                      <template v-else-if="m.key === 'warehouse'">
                        <div class="flex flex-wrap items-center justify-between gap-2 text-[12px]">
                          <span>已用 <b class="text-slate-800">{{ m.warehouse.used }}</b> / {{ m.warehouse.total }}</span>
                          <span>占比 <b :style="{ color: ratioColor(m.warehouse.pct) }">{{ m.warehouse.pct }}%</b></span>
                        </div>
                        <div class="warehouse-bar mt-2"><span :style="{ width: m.warehouse.pct + '%', background: ratioColor(m.warehouse.pct) }"></span></div>
                        <p class="mc-foot">数据更新于 11-20 14:30</p>
                      </template>
                    </div>
                  </section>
                </div>

                <div v-else class="mod-empty">
                  <span class="mod-empty-icon">
                    <svg viewBox="0 0 24 24" fill="none" aria-hidden="true">
                      <path d="M12 5v14M5 12h14" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" />
                    </svg>
                  </span>
                  <p class="mod-empty-title">还没有账号模块</p>
                  <p class="mod-empty-desc">添加常用模块，或从其他账号同步已有配置。模块只改执行内容，不影响计费。</p>
                  <div class="flex flex-wrap justify-center gap-2">
                    <button type="button" class="el-btn el-btn-primary" @click="openAddModule">+ 添加模块</button>
                    <button type="button" class="el-btn" @click="syncFromOther">从其他账号同步</button>
                  </div>
                </div>
              </section>
            </template>

            <div v-else class="acc-empty">
              <p class="text-sm font-medium text-slate-800">该设备还没有账号</p>
              <p class="mt-1 text-[12px] leading-5 text-slate-500">账号是执行和计费单位，一个号一份授权，3 元/月；设备台数不计费。</p>
              <button type="button" class="el-btn el-btn-primary mt-4" @click="openAddAccount(selectedDevice)">+ 添加账号</button>
            </div>
          </main>
        </div>

        <!-- 业务状态 -->
        <div v-else class="status-page">
          <div class="status-banner">
            <div>
              <h2>业务状态</h2>
              <p>按设备聚合 · 演示数据</p>
            </div>
            <div class="banner-actions">
              <span class="hidden text-[11px] text-slate-400 md:inline">刷新于 {{ refreshLabel }}</span>
              <label class="inline-flex items-center gap-2 text-[12px] text-slate-600">
                <span>按设备分组</span>
                <button type="button" role="switch" :aria-checked="grouped" class="switch" :class="{ on: grouped }" @click="grouped = !grouped"><i></i></button>
              </label>
              <button type="button" class="el-btn" @click="showTokens = true">授权令牌 <b class="text-[#e8480f]">{{ availableTokens }}</b></button>
              <button type="button" class="el-btn el-btn-primary" @click="refresh">刷新</button>
            </div>
          </div>
          <p class="mb-2 text-[11px] text-slate-400 md:hidden">表格可左右滑动，查看完整字段 →</p>

          <template v-if="grouped">
            <section v-for="device in groupedDevices" :key="device.id" class="biz-group">
              <div class="group-head" @click="toggleGroup(device.id)">
                <span class="group-chevron" :class="{ open: !isGroupCollapsed(device.id) }">▸</span>
                <span class="device-title">{{ device.key }}</span>
                <span class="text-[12px] text-slate-500">{{ device.accounts.length }} 个账号</span>
                <span class="el-tag" :class="heartbeatClass(device)">{{ heartbeatText(device) }}</span>
                <div class="device-chips">
                  <span>{{ device.meta.os }} · {{ device.meta.model }}</span>
                  <span>{{ device.meta.version }}</span>
                  <span>电量 {{ device.meta.battery }}%{{ device.meta.charging ? ' ⚡' : '' }}</span>
                  <span>温度 {{ device.meta.temp }}℃</span>
                  <span>内存 {{ device.meta.memory }}%</span>
                  <span>磁盘 {{ device.meta.disk }}%</span>
                </div>
                <button type="button" class="el-btn ml-auto flex-none" @click.stop="cloudControl(device)">云控</button>
              </div>

              <div v-if="!isGroupCollapsed(device.id)" class="table-scroll">
                <table class="biz-table grouped-table">
                  <colgroup>
                    <col style="width: 24px" /><col style="width: 58px" /><col style="width: 108px" /><col style="width: 44px" /><col style="width: 70px" /><col style="width: 80px" /><col style="width: 52px" /><col style="width: 64px" /><col style="width: 86px" /><col style="width: 86px" /><col style="width: 84px" /><col style="width: 102px" /><col style="width: 176px" />
                  </colgroup>
                  <thead>
                    <tr>
                      <th></th><th>账号类型</th><th>账号</th><th>序号</th><th>状态</th><th>角色名</th>
                      <th><button type="button" class="th-sort" @click="setSort('level')">等级 <i>{{ sortMark('level') }}</i></button></th>
                      <th>段位</th>
                      <th><button type="button" class="th-sort" @click="setSort('havoc')">哈弗币(M) <i>{{ sortMark('havoc') }}</i></button></th>
                      <th><button type="button" class="th-sort" @click="setSort('assets')">总资产(M) <i>{{ sortMark('assets') }}</i></button></th>
                      <th><button type="button" class="th-sort" @click="setSort('ratio')">仓库占比 <i>{{ sortMark('ratio') }}</i></button></th>
                      <th>更新时间</th><th>账号时效</th>
                    </tr>
                  </thead>
                  <tbody>
                    <template v-for="(account, index) in sortedAccounts(device.accounts)" :key="account.id">
                      <tr class="data-row" :class="{ 'is-off': !account.enabled }" @click="toggleRow(account.id)">
                        <td><span class="row-chevron" :class="{ open: isRowExpanded(account.id) }">▸</span></td>
                        <td><span class="el-tag" :class="account.type === 'QQ' ? 'tag-primary' : 'tag-success'">{{ account.type }}</span></td>
                        <td class="mono">{{ displayAccount(account) }}</td>
                        <td class="muted">{{ String(index + 1).padStart(2, '0') }}</td>
                        <td><span class="status-dot" :class="accountStatusClass(account.status)"></span>{{ account.status }}</td>
                        <td>{{ masked ? displayChar(account.charName) : account.charName }}</td>
                        <td>Lv.{{ account.level }}</td>
                        <td>{{ account.rank }}</td>
                        <td class="num">{{ fmtNum(account.havoc) }}</td>
                        <td class="num">{{ fmtNum(account.assets) }}</td>
                        <td>
                          <span class="ratio-cell">
                            <span class="mini-progress"><span :style="{ width: account.ratio + '%', background: ratioColor(account.ratio) }"></span></span>
                            <span :style="{ color: ratioColor(account.ratio) }">{{ account.ratio }}%</span>
                          </span>
                        </td>
                        <td class="muted">{{ account.updated }}</td>
                        <td>
                          <span class="el-tag" :class="licenseClass(account.license)">{{ account.license }}</span>
                          <span class="ml-1 text-[10px] text-slate-400">{{ shortDate(account.expire) }}</span>
                          <button type="button" class="el-btn-text ml-1" @click.stop="openRenew(account)">续费</button>
                        </td>
                      </tr>
                      <tr v-if="isRowExpanded(account.id)" class="expand-row">
                        <td colspan="13">
                          <div class="expand-inner">
                            <span class="text-[11px] font-semibold text-slate-500">生产物品</span>
                            <span v-for="(p, pi) in account.produce" :key="pi" class="produce-chip" :class="{ done: p.status === '已完成' }">
                              {{ p.name }} · {{ p.status }}<template v-if="p.left !== '—'"> · 剩余 {{ p.left }}</template>
                            </span>
                            <span class="ml-auto text-[10px] text-slate-400">设备 {{ device.key }} · 数据更新于 {{ account.updated }}</span>
                          </div>
                        </td>
                      </tr>
                    </template>
                  </tbody>
                </table>
              </div>
            </section>

            <div class="grand-total">
              <span>全部合计：哈弗币 <b>{{ fmtNum(totalHavoc) }}</b> M · 总资产 <b>{{ fmtNum(totalAssets) }}</b> M</span>
              <span class="text-[11px] text-slate-400">{{ totalAccounts }} 个账号 · {{ groupedDevices.length }} 台设备 · 在线 {{ onlineDevices }}/{{ groupedDevices.length }} · 月费 ¥{{ totalAccounts * 3 }}（3 元/号）</span>
            </div>
          </template>

          <template v-else>
            <section class="biz-group ungrouped">
              <div class="table-scroll">
                <table class="biz-table ungrouped-table">
                  <colgroup>
                    <col style="width: 24px" /><col style="width: 100px" /><col style="width: 52px" /><col style="width: 100px" /><col style="width: 42px" /><col style="width: 64px" /><col style="width: 76px" /><col style="width: 50px" /><col style="width: 60px" /><col style="width: 82px" /><col style="width: 82px" /><col style="width: 78px" /><col style="width: 96px" /><col style="width: 168px" />
                  </colgroup>
                  <thead>
                    <tr>
                      <th></th><th>设备</th><th>账号类型</th><th>账号</th><th>序号</th><th>状态</th><th>角色名</th>
                      <th><button type="button" class="th-sort" @click="setSort('level')">等级 <i>{{ sortMark('level') }}</i></button></th>
                      <th>段位</th>
                      <th><button type="button" class="th-sort" @click="setSort('havoc')">哈弗币(M) <i>{{ sortMark('havoc') }}</i></button></th>
                      <th><button type="button" class="th-sort" @click="setSort('assets')">总资产(M) <i>{{ sortMark('assets') }}</i></button></th>
                      <th><button type="button" class="th-sort" @click="setSort('ratio')">仓库占比 <i>{{ sortMark('ratio') }}</i></button></th>
                      <th>更新时间</th><th>账号时效</th>
                    </tr>
                  </thead>
                  <tbody>
                    <template v-for="account in sortedAccounts(allAccounts)" :key="account.id">
                      <tr class="data-row" :class="{ 'is-off': !account.enabled }" @click="toggleRow(account.id)">
                        <td><span class="row-chevron" :class="{ open: isRowExpanded(account.id) }">▸</span></td>
                        <td class="mono">{{ deviceKeyOf(account) }}</td>
                        <td><span class="el-tag" :class="account.type === 'QQ' ? 'tag-primary' : 'tag-success'">{{ account.type }}</span></td>
                        <td class="mono">{{ displayAccount(account) }}</td>
                        <td class="muted">{{ String(account.seq).padStart(2, '0') }}</td>
                        <td><span class="status-dot" :class="accountStatusClass(account.status)"></span>{{ account.status }}</td>
                        <td>{{ masked ? displayChar(account.charName) : account.charName }}</td>
                        <td>Lv.{{ account.level }}</td>
                        <td>{{ account.rank }}</td>
                        <td class="num">{{ fmtNum(account.havoc) }}</td>
                        <td class="num">{{ fmtNum(account.assets) }}</td>
                        <td>
                          <span class="ratio-cell">
                            <span class="mini-progress"><span :style="{ width: account.ratio + '%', background: ratioColor(account.ratio) }"></span></span>
                            <span :style="{ color: ratioColor(account.ratio) }">{{ account.ratio }}%</span>
                          </span>
                        </td>
                        <td class="muted">{{ account.updated }}</td>
                        <td>
                          <span class="el-tag" :class="licenseClass(account.license)">{{ account.license }}</span>
                          <span class="ml-1 text-[10px] text-slate-400">{{ shortDate(account.expire) }}</span>
                          <button type="button" class="el-btn-text ml-1" @click.stop="openRenew(account)">续费</button>
                        </td>
                      </tr>
                      <tr v-if="isRowExpanded(account.id)" class="expand-row">
                        <td colspan="14">
                          <div class="expand-inner">
                            <span class="text-[11px] font-semibold text-slate-500">生产物品</span>
                            <span v-for="(p, pi) in account.produce" :key="pi" class="produce-chip" :class="{ done: p.status === '已完成' }">
                              {{ p.name }} · {{ p.status }}<template v-if="p.left !== '—'"> · 剩余 {{ p.left }}</template>
                            </span>
                            <span class="ml-auto text-[10px] text-slate-400">设备 {{ deviceKeyOf(account) }} · 数据更新于 {{ account.updated }}</span>
                          </div>
                        </td>
                      </tr>
                    </template>
                  </tbody>
                </table>
              </div>
            </section>

            <div class="grand-total">
              <span>全部合计：哈弗币 <b>{{ fmtNum(totalHavoc) }}</b> M · 总资产 <b>{{ fmtNum(totalAssets) }}</b> M</span>
              <span class="text-[11px] text-slate-400">{{ totalAccounts }} 个账号 · {{ groupedDevices.length }} 台设备 · 在线 {{ onlineDevices }}/{{ groupedDevices.length }} · 月费 ¥{{ totalAccounts * 3 }}（3 元/号）</span>
            </div>
          </template>
        </div>

        <div v-if="toastMsg" class="demo-toast">{{ toastMsg }}</div>
      </div>

      <p class="mt-4 text-[12px] leading-6 text-slate-400">
        界面按 Manager 的业务配置与业务状态两屏还原，账号、角色与数值均为演示数据；点击操作只在本页生效，不会写入真实管理台。
      </p>
    </div>

    <!-- 新增账号 -->
    <div v-if="showAddAccount" class="modal-mask" @click.self="showAddAccount = false">
      <div class="modal-panel">
        <div class="modal-head">
          <div><h3>新增账号</h3><p>账号是执行和计费单位，一个号一份授权。</p></div>
          <button type="button" class="modal-close" @click="showAddAccount = false">×</button>
        </div>
        <div class="modal-body">
          <label class="form-row"><span>设备</span>
            <select v-model="addForm.deviceId" class="form-control" @change="addForm.copyFrom = ''">
              <option v-for="d in devices.filter((x) => x.approval === 'approved')" :key="d.id" :value="d.id">{{ d.name }} · {{ d.key }}</option>
            </select>
          </label>
          <div class="form-row"><span>类型</span>
            <div class="segmented">
              <button v-for="t in ['QQ', '微信']" :key="t" type="button" :class="{ active: addForm.type === t }" @click="addForm.type = t as 'QQ' | '微信'">{{ t }}</button>
            </div>
          </div>
          <label class="form-row"><span>账号</span><input v-model.trim="addForm.number" class="form-control" placeholder="输入 QQ 号或微信号" /></label>
          <label class="form-row"><span>复制配置</span>
            <select v-model="addForm.copyFrom" class="form-control">
              <option value="">不复制，使用默认模块</option>
              <option v-for="a in copyCandidates" :key="a.id" :value="a.id">{{ a.type }}：{{ displayAccount(a) }}（{{ a.modules.length }} 个模块）</option>
            </select>
          </label>
          <div class="form-row"><span>启用</span>
            <button type="button" role="switch" :aria-checked="addForm.enabled" class="switch" :class="{ on: addForm.enabled }" @click="addForm.enabled = !addForm.enabled"><i></i></button>
            <span class="text-[12px] text-slate-500">添加后立即参与执行</span>
          </div>
          <p v-if="addError" class="text-[12px] text-red-600">{{ addError }}</p>
        </div>
        <div class="modal-foot">
          <button type="button" class="el-btn" @click="showAddAccount = false">取消</button>
          <button type="button" class="el-btn el-btn-primary" @click="confirmAddAccount">确认添加</button>
        </div>
      </div>
    </div>

    <!-- 添加模块 -->
    <div v-if="showAddModule" class="modal-mask" @click.self="showAddModule = false">
      <div class="modal-panel">
        <div class="modal-head">
          <div><h3>添加模块</h3><p>模块只改执行内容，不加钱；计费仍按账号实例算。</p></div>
          <button type="button" class="modal-close" @click="showAddModule = false">×</button>
        </div>
        <div class="modal-body !py-2">
          <button v-for="item in availableModules" :key="item.key" type="button" class="module-cand" @click="addModule(item.key)">
            <span class="flex-1 text-left"><b>{{ item.label }}</b><small>{{ item.desc }}</small></span>
            <span class="el-btn-text">添加</span>
          </button>
          <p v-if="!availableModules.length" class="py-6 text-center text-[12px] text-slate-400">当前账号已添加全部模块</p>
        </div>
        <div class="modal-foot">
          <button type="button" class="el-btn" @click="syncFromOther">从其他账号同步</button>
          <button type="button" class="el-btn el-btn-primary" @click="showAddModule = false">完成</button>
        </div>
      </div>
    </div>

    <!-- 删除账号 -->
    <div v-if="deleteTarget" class="modal-mask" @click.self="deleteTarget = null">
      <div class="modal-panel !w-[420px]">
        <div class="modal-head"><div><h3>删除账号</h3><p>账号与模块配置会从管理台移除。</p></div><button type="button" class="modal-close" @click="deleteTarget = null">×</button></div>
        <div class="modal-body"><p class="text-[13px] leading-6 text-slate-600">确认删除 <b>{{ deleteTarget.type }}：{{ displayAccount(deleteTarget) }}</b>？授权与执行会同时停止，此操作不可撤销。</p></div>
        <div class="modal-foot"><button type="button" class="el-btn" @click="deleteTarget = null">取消</button><button type="button" class="el-btn el-btn-danger" @click="confirmDeleteAccount">删除账号</button></div>
      </div>
    </div>

    <!-- 续费 -->
    <div v-if="renewTarget" class="modal-mask" @click.self="renewTarget = null">
      <div class="modal-panel !w-[520px]">
        <div class="modal-head"><div><h3>账号续费</h3><p>续费只延长授权，账号、设备与模块配置原样保留。</p></div><button type="button" class="modal-close" @click="renewTarget = null">×</button></div>
        <div class="modal-body">
          <div class="renew-grid">
            <div><span>账号</span><b>{{ renewTarget.type }}：{{ displayAccount(renewTarget) }}</b></div>
            <div><span>当前状态</span><span class="el-tag" :class="licenseClass(renewTarget.license)">{{ renewTarget.license }}</span></div>
            <div><span>当前到期</span><b>{{ renewTarget.expire }}</b></div>
            <div><span>可用令牌</span><b class="text-[#e8480f]">{{ availableTokens }}</b></div>
          </div>
          <p class="mt-4 text-[12px] font-medium text-slate-700">选择续费时长</p>
          <div class="renew-options">
            <button v-for="opt in renewOptions" :key="opt.days" type="button" class="renew-option" :class="{ active: renewDays === opt.days }" @click="renewDays = opt.days">
              <b>{{ opt.label }}</b><span>消耗 {{ opt.days }} 枚令牌</span>
            </button>
          </div>
          <p class="mt-3 text-[12px] text-slate-500">续费后到期：<b class="text-slate-800">{{ renewPreview }}</b></p>
          <p class="mt-2 text-[11px] leading-5 text-slate-400">时效由 Auth 统一签发，设备按签名租约到期；离线设备同样按绝对时间失效。</p>
        </div>
        <div class="modal-foot"><button type="button" class="el-btn" @click="renewTarget = null">取消</button><button type="button" class="el-btn el-btn-primary" @click="confirmRenew">确认续费</button></div>
      </div>
    </div>

    <!-- 授权令牌 -->
    <div v-if="showTokens" class="modal-mask" @click.self="showTokens = false">
      <div class="modal-panel !w-[480px]">
        <div class="modal-head"><div><h3>授权令牌</h3><p>令牌用于账号授权续期，演示环境不产生实际消耗。</p></div><button type="button" class="modal-close" @click="showTokens = false">×</button></div>
        <div class="modal-body">
          <div class="token-stats">
            <div><b>{{ availableTokens }}</b><span>可用令牌</span></div>
            <div><b>{{ usedTokens }}</b><span>本月已用</span></div>
            <div><b>{{ expiringAccounts.length }}</b><span>7 天内到期</span></div>
          </div>
          <p class="mt-4 text-[12px] font-medium text-slate-700">即将到期</p>
          <div class="mt-2 divide-y divide-slate-100 rounded-lg border border-slate-200">
            <div v-for="a in expiringAccounts" :key="a.id" class="flex items-center justify-between gap-3 px-3 py-2 text-[12px]">
              <span>{{ a.type }}：{{ displayAccount(a) }}</span>
              <span class="text-slate-400">{{ shortDate(a.expire) }}</span>
              <button type="button" class="el-btn-text" @click="showTokens = false; openRenew(a)">续费</button>
            </div>
          </div>
        </div>
        <div class="modal-foot"><button type="button" class="el-btn el-btn-primary" @click="showTokens = false">知道了</button></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue'
import PageHero from '@/components/PageHero.vue'

type TabKey = 'config' | 'status'
type Trigger = 'follow' | 'auto' | 'exclusive'
type ModuleKey = 'profile' | 'afk' | 'produce' | 'sell' | 'buy' | 'stock' | 'exchange' | 'warehouse'
type AccountStatus = '运行中' | '挂机中' | '空闲' | '异常' | '已停用'
type LicenseStatus = '有效' | '即将到期' | '已到期'

interface ModuleCfg {
  key: ModuleKey
  enabled: boolean
  priority: number
  trigger: Trigger
  scope?: 'device'
  [key: string]: any
}

interface ProduceItem { name: string; status: string; left: string }

interface Account {
  id: string
  deviceId: string
  type: 'QQ' | '微信'
  account: string
  seq: number
  enabled: boolean
  status: AccountStatus
  charName: string
  level: number
  rank: string
  havoc: number
  assets: number
  ratio: number
  updated: string
  license: LicenseStatus
  expire: string
  modules: ModuleCfg[]
  produce: ProduceItem[]
}

interface Device {
  id: string
  key: string
  name: string
  group: string
  enabled: boolean
  online: boolean
  heartbeatSec: number
  approval: 'approved' | 'pending'
  meta: { os: string; model: string; version: string; battery: number; charging: boolean; temp: number; memory: number; disk: number }
  accounts: Account[]
}

const tabs: { key: TabKey; label: string }[] = [
  { key: 'config', label: '业务配置' },
  { key: 'status', label: '业务状态' },
]

const activeTab = ref<TabKey>('config')
const masked = ref(false)
const refreshedAt = ref(new Date())
const toastMsg = ref('')
let toastTimer: ReturnType<typeof setTimeout> | undefined

const refreshLabel = computed(() => {
  const t = refreshedAt.value
  return `${String(t.getHours()).padStart(2, '0')}:${String(t.getMinutes()).padStart(2, '0')}:${String(t.getSeconds()).padStart(2, '0')}`
})

function notify(message: string) {
  toastMsg.value = message
  if (toastTimer) clearTimeout(toastTimer)
  toastTimer = setTimeout(() => (toastMsg.value = ''), 2600)
}

function refresh() {
  refreshedAt.value = new Date()
  notify('已刷新演示数据')
}

const moduleMeta: Record<ModuleKey, { label: string; desc: string; priority: number }> = {
  profile: { label: '信息', desc: '角色名、等级、段位与资产采集', priority: 10 },
  afk: { label: '挂机', desc: '自动对局、冷却与执行时段', priority: 120 },
  produce: { label: '生产', desc: '特勤部门产物与制造队列', priority: 80 },
  sell: { label: '出售', desc: '集市上架与最低价控制', priority: 60 },
  buy: { label: '购买', desc: '按目标库存补货，限制最高价', priority: 40 },
  stock: { label: '库存', desc: '定时盘点仓库物品数量', priority: 30 },
  exchange: { label: '兑换', desc: '按渠道把物品兑换成资源', priority: 50 },
  warehouse: { label: '仓库', desc: '设备级仓库容量与占用', priority: 200 },
}

function makeModule(key: ModuleKey): ModuleCfg {
  const base: ModuleCfg = { key, enabled: true, priority: moduleMeta[key].priority, trigger: 'follow' }
  if (key === 'warehouse') base.scope = 'device'
  if (key === 'afk') {
    base.afk = {
      totalGames: 36,
      lastResult: '撤离成功',
      cooldown: 60,
      passport: { exp: 12800, max: 16000 },
      windows: [{ start: '22:00', end: '08:00' }],
    }
  }
  if (key === 'produce') {
    base.produce = {
      items: [
        { dept: '工作台', prop: '紫罗兰', name: '紫罗兰', status: '制造中', left: '02:41', enabled: true },
        { dept: '制药台', prop: '医疗包', name: '医疗包', status: '制造中', left: '00:58', enabled: true },
        { dept: '技术中心', prop: '战术手电', name: '空闲', status: '空闲', left: '—', enabled: true },
      ],
    }
  }
  if (key === 'sell') {
    base.sell = {
      items: [
        { prop: '紫罗兰', mode: 'market', price: 1800, enabled: true },
        { prop: '海盗弯刀', mode: 'market', price: 6200, enabled: true },
      ],
    }
  }
  if (key === 'buy') {
    base.buy = {
      items: [
        { prop: '标准弹药', current: 12, target: 40, maxPrice: 820, schedule: '每2小时', enabled: true },
        { prop: '高级护甲', current: 3, target: 10, maxPrice: 4200, schedule: '每6小时', enabled: true },
      ],
    }
  }
  if (key === 'stock') {
    base.stock = { items: [{ prop: '金条', count: 18, schedule: '每6小时', left: '03:12', enabled: true }] }
  }
  if (key === 'exchange') {
    base.exchange = { items: [{ prop: '赛季补给箱', channel: 'market', left: '01:20', enabled: true }] }
  }
  if (key === 'warehouse') {
    base.warehouse = { used: 248, total: 400, pct: 62 }
  }
  return base
}

function makeAccount(base: Omit<Account, 'modules' | 'produce'>, moduleKeys: ModuleKey[]): Account {
  return {
    ...base,
    modules: moduleKeys.map(makeModule),
    produce: [
      { name: '紫罗兰', status: '制造中', left: '02:41' },
      { name: '标准弹药', status: '排队中', left: '—' },
      { name: '医疗包', status: '已完成', left: '×12' },
    ],
  }
}

const devices = reactive<Device[]>([
  {
    id: 'd1', key: 'a1b2c3d4e5f6', name: '门店A · 主机01', group: '门店 A', enabled: true, online: true, heartbeatSec: 8, approval: 'approved',
    meta: { os: 'Android 14', model: '小米 14', version: 'v1.0.0(3f8a21c)', battery: 86, charging: true, temp: 31, memory: 62, disk: 71 },
    accounts: [
      makeAccount({ id: 'a1', deviceId: 'd1', type: 'QQ', account: '124578963', seq: 1, enabled: true, status: '运行中', charName: '林深', level: 62, rank: '钻石Ⅱ', havoc: 428.6, assets: 1286.3, ratio: 62, updated: '11-20 14:32', license: '有效', expire: '2026-12-05 18:00' }, ['profile', 'produce', 'sell', 'warehouse']),
      makeAccount({ id: 'a2', deviceId: 'd1', type: '微信', account: 'wxid_lxy8821', seq: 2, enabled: true, status: '运行中', charName: '陈默', level: 58, rank: '铂金Ⅰ', havoc: 312.4, assets: 942.8, ratio: 48, updated: '11-20 14:28', license: '即将到期', expire: '2026-11-23 12:00' }, ['profile', 'produce', 'sell']),
      makeAccount({ id: 'a3', deviceId: 'd1', type: 'QQ', account: '369258147', seq: 3, enabled: false, status: '已停用', charName: '周野', level: 41, rank: '黄金Ⅲ', havoc: 88.2, assets: 265.9, ratio: 31, updated: '11-19 22:10', license: '已到期', expire: '2026-11-18 09:00' }, ['profile']),
    ],
  },
  {
    id: 'd2', key: 'b7c8d9e0f1a2', name: '门店A · 主机02', group: '门店 A', enabled: true, online: true, heartbeatSec: 34, approval: 'approved',
    meta: { os: 'Android 13', model: 'Redmi K70', version: 'v1.0.0(3f8a21c)', battery: 72, charging: false, temp: 34, memory: 71, disk: 64 },
    accounts: [
      makeAccount({ id: 'b1', deviceId: 'd2', type: 'QQ', account: '741852963', seq: 1, enabled: true, status: '空闲', charName: '苏晚', level: 66, rank: '钻石Ⅰ', havoc: 501.2, assets: 1533.6, ratio: 70, updated: '11-20 14:31', license: '即将到期', expire: '2026-11-22 20:00' }, ['profile', 'produce', 'sell', 'warehouse']),
      makeAccount({ id: 'b2', deviceId: 'd2', type: '微信', account: 'wxid_sjz_009', seq: 2, enabled: true, status: '运行中', charName: '陆行', level: 55, rank: '铂金Ⅱ', havoc: 288.3, assets: 864.2, ratio: 45, updated: '11-20 14:20', license: '有效', expire: '2026-12-11 10:00' }, ['profile', 'produce']),
      makeAccount({ id: 'b3', deviceId: 'd2', type: 'QQ', account: '159357486', seq: 3, enabled: true, status: '异常', charName: '江屿', level: 49, rank: '铂金Ⅲ', havoc: 196.7, assets: 602.5, ratio: 52, updated: '11-20 13:58', license: '有效', expire: '2026-12-18 15:00' }, ['profile', 'sell']),
    ],
  },
  {
    id: 'd3', key: 'c3d4e5f6a7b8', name: '门店B · 主机03', group: '门店 B', enabled: true, online: false, heartbeatSec: 999, approval: 'pending',
    meta: { os: 'Android 14', model: '一加 12', version: '待接入', battery: 0, charging: false, temp: 0, memory: 0, disk: 0 },
    accounts: [],
  },
  {
    id: 'd4', key: 'd4e5f6a7b8c9', name: '门店B · 主机04', group: '门店 B', enabled: true, online: true, heartbeatSec: 12, approval: 'approved',
    meta: { os: 'Android 14', model: '小米 14 Pro', version: 'v1.0.0(9b2ce71)', battery: 94, charging: true, temp: 29, memory: 54, disk: 58 },
    accounts: [
      makeAccount({ id: 'c1', deviceId: 'd4', type: 'QQ', account: '852741963', seq: 1, enabled: true, status: '运行中', charName: '顾北', level: 60, rank: '钻石Ⅲ', havoc: 402.1, assets: 1198.4, ratio: 58, updated: '11-20 14:33', license: '有效', expire: '2026-12-02 09:00' }, ['profile', 'produce', 'sell', 'warehouse']),
      makeAccount({ id: 'c2', deviceId: 'd4', type: '微信', account: 'wxid_blue77', seq: 2, enabled: true, status: '运行中', charName: '沈知', level: 57, rank: '铂金Ⅰ', havoc: 276.9, assets: 812.7, ratio: 41, updated: '11-20 14:27', license: '有效', expire: '2026-12-09 17:00' }, ['profile', 'produce']),
      makeAccount({ id: 'c3', deviceId: 'd4', type: 'QQ', account: '963852741', seq: 3, enabled: false, status: '已停用', charName: '韩川', level: 38, rank: '黄金Ⅱ', havoc: 61.5, assets: 188.3, ratio: 27, updated: '11-19 18:45', license: '已到期', expire: '2026-11-19 10:00' }, ['profile']),
    ],
  },
  {
    id: 'd5', key: 'e5f6a7b8c9d0', name: '测试组 · 主机05', group: '测试组', enabled: true, online: false, heartbeatSec: 128, approval: 'approved',
    meta: { os: 'Android 13', model: '真我 GT5', version: 'v1.0.0(3f8a21c)', battery: 41, charging: false, temp: 36, memory: 78, disk: 83 },
    accounts: [
      makeAccount({ id: 'e1', deviceId: 'd5', type: 'QQ', account: '321654987', seq: 1, enabled: true, status: '运行中', charName: '许星河', level: 64, rank: '钻石Ⅱ', havoc: 455.8, assets: 1362.9, ratio: 66, updated: '11-20 14:05', license: '有效', expire: '2026-12-15 12:00' }, ['profile', 'produce', 'sell', 'warehouse']),
      makeAccount({ id: 'e2', deviceId: 'd5', type: '微信', account: 'wxid_test01', seq: 2, enabled: true, status: '空闲', charName: '叶青', level: 52, rank: '铂金Ⅲ', havoc: 231.4, assets: 704.8, ratio: 39, updated: '11-20 14:02', license: '即将到期', expire: '2026-11-24 08:00' }, ['profile', 'produce']),
      makeAccount({ id: 'e3', deviceId: 'd5', type: 'QQ', account: '654987321', seq: 3, enabled: true, status: '运行中', charName: '温言', level: 59, rank: '钻石Ⅳ', havoc: 366.2, assets: 1089.5, ratio: 54, updated: '11-20 14:12', license: '有效', expire: '2026-12-20 16:00' }, ['profile', 'produce', 'sell']),
    ],
  },
])

const allAccounts = computed(() => devices.flatMap((d) => d.accounts))
const groupedDevices = computed(() => devices.filter((d) => d.accounts.length > 0))
const onlineDevices = computed(() => groupedDevices.value.filter((d) => d.online).length)
const totalAccounts = computed(() => allAccounts.value.length)
const totalHavoc = computed(() => allAccounts.value.reduce((sum, a) => sum + a.havoc, 0))
const totalAssets = computed(() => allAccounts.value.reduce((sum, a) => sum + a.assets, 0))

const selectedDeviceId = ref('d1')
const selectedAccountId = ref('a1')
const expandedDevices = ref<string[]>(['d1'])

const selectAccount = (device: Device, account: Account) => {
  selectedDeviceId.value = device.id
  selectedAccountId.value = account.id
}

const selectedAccount = computed<Account | null>(() => {
  for (const device of devices) {
    const found = device.accounts.find((a) => a.id === selectedAccountId.value)
    if (found) return found
  }
  return null
})

const emptyAccount: Account = {
  id: '', deviceId: '', type: 'QQ', account: '', seq: 0, enabled: false, status: '已停用', charName: '', level: 0, rank: '', havoc: 0, assets: 0, ratio: 0, updated: '', license: '已到期', expire: '', modules: [], produce: [],
}
const accountView = computed(() => selectedAccount.value ?? emptyAccount)
const selectedDevice = computed(() => devices.find((d) => d.id === selectedAccount.value?.deviceId) ?? devices[0])
const currentDevice = computed(() => selectedDevice.value)

const isExpanded = (id: string) => expandedDevices.value.includes(id)
function toggleDevice(device: Device) {
  expandedDevices.value = isExpanded(device.id) ? expandedDevices.value.filter((id) => id !== device.id) : [...expandedDevices.value, device.id]
}
function expandAll() { expandedDevices.value = devices.map((d) => d.id) }
function collapseAll() { expandedDevices.value = [] }
function toggleDeviceEnabled(device: Device) {
  device.enabled = !device.enabled
  notify(`${device.name} 已${device.enabled ? '开启' : '禁用'}（演示）`)
}

function displayRaw(value: string) {
  if (!masked.value) return value
  if (value.length <= 5) return value
  return `${value.slice(0, 3)}****${value.slice(-3)}`
}
const displayAccount = (account: Account) => displayRaw(account.account)
const displayChar = (name: string) => (name.length > 1 ? `${name.slice(0, 1)}**` : name)

function deviceDot(device: Device) {
  if (device.approval === 'pending') return 'pending'
  if (!device.online) return 'off'
  return 'on'
}
function accountDot(account: Account) {
  if (!account.enabled || account.license === '已到期') return 'off'
  if (account.license === '即将到期') return 'soon'
  return 'on'
}

function toggleAccount(account: Account) {
  account.enabled = !account.enabled
  account.status = account.enabled ? '空闲' : '已停用'
  notify(`${displayAccount(account)} 已${account.enabled ? '启用' : '停用'}（演示）`)
}

function syncAccount() { notify('已从设备最近一次配置快照同步（演示）') }
function batchSync() { notify(`已按设备分组批量同步 ${currentDevice.value.accounts.length} 个账号（演示）`) }

const showAddAccount = ref(false)
const addError = ref('')
const addForm = reactive({ deviceId: 'd1', type: 'QQ' as 'QQ' | '微信', number: '', copyFrom: '', enabled: true })
const copyCandidates = computed(() => devices.find((d) => d.id === addForm.deviceId)?.accounts ?? [])

function openAddAccount(device: Device) {
  if (device.approval === 'pending') {
    notify('设备还需审核通过，才能添加账号')
    return
  }
  addForm.deviceId = device.id
  addForm.type = 'QQ'
  addForm.number = ''
  addForm.copyFrom = ''
  addForm.enabled = true
  addError.value = ''
  showAddAccount.value = true
}

function fmtDateTime(date: Date) {
  const y = date.getFullYear()
  const m = String(date.getMonth() + 1).padStart(2, '0')
  const d = String(date.getDate()).padStart(2, '0')
  const hh = String(date.getHours()).padStart(2, '0')
  const mm = String(date.getMinutes()).padStart(2, '0')
  return `${y}-${m}-${d} ${hh}:${mm}`
}

function confirmAddAccount() {
  const number = addForm.number.trim()
  if (!number) { addError.value = '请输入账号'; return }
  if (allAccounts.value.some((a) => a.account === number)) { addError.value = '该账号已存在，请勿重复添加'; return }
  const device = devices.find((d) => d.id === addForm.deviceId)
  if (!device) { addError.value = '设备不存在'; return }
  const source = addForm.copyFrom ? allAccounts.value.find((a) => a.id === addForm.copyFrom) : null
  const modules = source ? (JSON.parse(JSON.stringify(source.modules)) as ModuleCfg[]) : [makeModule('profile'), makeModule('produce'), makeModule('sell'), makeModule('warehouse')]
  const expire = new Date(Date.now() + 30 * 86400000)
  const account: Account = {
    id: `new-${Date.now()}`, deviceId: device.id, type: addForm.type, account: number, seq: device.accounts.length + 1,
    enabled: addForm.enabled, status: addForm.enabled ? '空闲' : '已停用', charName: '未采集', level: 0, rank: '未定级',
    havoc: 0, assets: 0, ratio: 0, updated: fmtDateTime(new Date()).slice(5, 16), license: '有效', expire: fmtDateTime(expire),
    modules, produce: [],
  }
  device.accounts.push(account)
  if (!isExpanded(device.id)) expandedDevices.value.push(device.id)
  selectedDeviceId.value = device.id
  selectedAccountId.value = account.id
  showAddAccount.value = false
  notify(`已添加 ${account.type}：${displayAccount(account)}（演示）`)
}

const deleteTarget = ref<Account | null>(null)
function askDeleteAccount() { deleteTarget.value = selectedAccount.value }
function confirmDeleteAccount() {
  const account = deleteTarget.value
  if (!account) return
  const device = devices.find((d) => d.id === account.deviceId)
  if (device) device.accounts = device.accounts.filter((a) => a.id !== account.id)
  if (selectedAccountId.value === account.id) {
    const next = device?.accounts[0] ?? allAccounts.value[0]
    selectedAccountId.value = next?.id ?? ''
    if (next) selectedDeviceId.value = next.deviceId
  }
  deleteTarget.value = null
  notify('账号已删除（演示）')
}

const showAddModule = ref(false)
const addModuleFor = ref<Account | null>(null)
const moduleCatalog: { key: ModuleKey; label: string; desc: string }[] = [
  { key: 'profile', label: '信息', desc: '采集角色名、等级、段位与资产' },
  { key: 'produce', label: '生产', desc: '按特勤部门排产与制造' },
  { key: 'sell', label: '出售', desc: '集市上架与最低价控制' },
  { key: 'warehouse', label: '仓库', desc: '设备级仓库容量与占用' },
]
const availableModules = computed(() => {
  const target = addModuleFor.value
  if (!target) return moduleCatalog
  return moduleCatalog.filter((item) => !target.modules.some((m) => m.key === item.key))
})
function openAddModule() {
  addModuleFor.value = selectedAccount.value
  showAddModule.value = true
}
function addModule(key: ModuleKey) {
  if (!addModuleFor.value) return
  addModuleFor.value.modules.push(makeModule(key))
  notify(`已添加「${moduleLabel(key)}」模块（演示）`)
}
function removeModule(module: ModuleCfg) {
  if (!selectedAccount.value) return
  selectedAccount.value.modules = selectedAccount.value.modules.filter((m) => m.key !== module.key)
}
function syncFromOther() {
  const target = addModuleFor.value ?? selectedAccount.value
  if (!target) return
  const source = allAccounts.value.find((a) => a.id !== target.id && a.deviceId === target.deviceId && a.modules.length)
  if (!source) { notify('同设备下没有可复制的账号配置'); return }
  target.modules = JSON.parse(JSON.stringify(source.modules)) as ModuleCfg[]
  showAddModule.value = false
  notify(`已从 ${displayAccount(source)} 同步模块配置（演示）`)
}

const moduleLabel = (key: ModuleKey) => moduleMeta[key].label
function normalizePriority(module: ModuleCfg) {
  const value = Number(module.priority)
  module.priority = Number.isFinite(value) ? Math.min(999, Math.max(0, Math.round(value))) : 0
}
function addAfkWindow(module: ModuleCfg) { module.afk.windows.push({ start: '09:00', end: '12:00' }) }
function removeAfkWindow(module: ModuleCfg, index: number) { module.afk.windows.splice(index, 1) }
function windowText(window: { start: string; end: string }) {
  const [sh, sm] = window.start.split(':').map(Number)
  const [eh, em] = window.end.split(':').map(Number)
  const minutes = (eh * 60 + em - (sh * 60 + sm) + 1440) % 1440
  return `${Math.floor(minutes / 60)} 小时 ${minutes % 60} 分`
}
function produceOptions(dept: string) {
  if (dept === '制药台') return ['医疗包', '止痛药', '肾上腺素']
  if (dept === '技术中心') return ['战术手电', '激光指示器', '防弹插板']
  return ['紫罗兰', '标准弹药', '高级护甲']
}
function addProduceItem(module: ModuleCfg) {
  module.produce.items.push({ dept: '工作台', prop: '紫罗兰', name: '空闲', status: '空闲', left: '—', enabled: true })
}
function removeProduceItem(module: ModuleCfg, index: number) { module.produce.items.splice(index, 1) }
function addSellItem(module: ModuleCfg) { module.sell.items.push({ prop: '新物品', mode: 'market', price: 0, enabled: true }) }
function removeSellItem(module: ModuleCfg, index: number) { module.sell.items.splice(index, 1) }
function addBuyItem(module: ModuleCfg) { module.buy.items.push({ prop: '新物品', current: 0, target: 10, maxPrice: 0, schedule: '每2小时', enabled: true }) }

const grouped = ref(true)
const collapsedGroups = ref<string[]>([])
const isGroupCollapsed = (id: string) => collapsedGroups.value.includes(id)
function toggleGroup(id: string) {
  collapsedGroups.value = isGroupCollapsed(id) ? collapsedGroups.value.filter((x) => x !== id) : [...collapsedGroups.value, id]
}

const sortKey = ref<'level' | 'havoc' | 'assets' | 'ratio'>('havoc')
const sortDir = ref<1 | -1>(-1)
function setSort(key: 'level' | 'havoc' | 'assets' | 'ratio') {
  if (sortKey.value === key) sortDir.value = sortDir.value === 1 ? -1 : 1
  else { sortKey.value = key; sortDir.value = -1 }
}
function sortMark(key: string) {
  if (sortKey.value !== key) return ''
  return sortDir.value === -1 ? '↓' : '↑'
}
function sortedAccounts(list: Account[]) {
  return [...list].sort((a, b) => (a[sortKey.value] - b[sortKey.value]) * sortDir.value)
}

function heartbeatText(device: Device) {
  if (!device.online) return '离线 · 2 分钟'
  if (device.heartbeatSec <= 30) return `心跳 ${device.heartbeatSec}s`
  if (device.heartbeatSec <= 60) return `心跳 ${device.heartbeatSec}s`
  return `心跳 ${Math.floor(device.heartbeatSec / 60)}m`
}
function heartbeatClass(device: Device) {
  if (!device.online) return 'tag-danger'
  if (device.heartbeatSec <= 30) return 'tag-success'
  if (device.heartbeatSec <= 60) return 'tag-warning'
  return 'tag-danger'
}
function deviceKeyOf(account: Account) {
  return devices.find((d) => d.id === account.deviceId)?.key ?? '—'
}
function cloudControl(device: Device) {
  notify(`已向 ${device.key} 发送云控指令（演示）`)
}

function accountStatusClass(status: AccountStatus) {
  if (status === '运行中') return 'run'
  if (status === '挂机中') return 'hang'
  if (status === '异常') return 'error'
  if (status === '空闲') return 'idle'
  return 'off'
}
function licenseClass(license: LicenseStatus) {
  if (license === '有效') return 'tag-success'
  if (license === '即将到期') return 'tag-warning'
  return 'tag-danger'
}
function ratioColor(value: number) {
  if (value >= 75) return '#f56c6c'
  if (value >= 50) return '#e6a23c'
  return '#67c23a'
}

const expandedRows = ref<string[]>([])
const isRowExpanded = (id: string) => expandedRows.value.includes(id)
function toggleRow(id: string) {
  expandedRows.value = isRowExpanded(id) ? expandedRows.value.filter((x) => x !== id) : [...expandedRows.value, id]
}

const availableTokens = ref(59)
const usedTokens = ref(31)
const expiringAccounts = computed(() => allAccounts.value.filter((a) => a.license === '即将到期'))
const showTokens = ref(false)

const renewTarget = ref<Account | null>(null)
const renewDays = ref(30)
const renewOptions = [
  { days: 7, label: '7 天' },
  { days: 30, label: '30 天' },
  { days: 90, label: '90 天' },
]
function openRenew(account: Account) {
  renewTarget.value = account
  renewDays.value = 30
}
function parseExpire(value: string) {
  const ts = Date.parse(value.replace(' ', 'T'))
  return Number.isNaN(ts) ? Date.now() : ts
}
const renewPreview = computed(() => {
  const target = renewTarget.value
  if (!target) return ''
  const base = Math.max(Date.now(), parseExpire(target.expire))
  return fmtDateTime(new Date(base + renewDays.value * 86400000))
})
function confirmRenew() {
  const account = renewTarget.value
  if (!account) return
  if (availableTokens.value < renewDays.value) {
    notify('可用令牌不足，无法完成续费（演示）')
    return
  }
  const base = Math.max(Date.now(), parseExpire(account.expire))
  account.expire = fmtDateTime(new Date(base + renewDays.value * 86400000))
  account.license = '有效'
  availableTokens.value -= renewDays.value
  usedTokens.value += renewDays.value
  renewTarget.value = null
  notify(`已为 ${displayAccount(account)} 续费 ${renewDays.value} 天（演示）`)
}

function fmtNum(value: number) {
  return value.toLocaleString('zh-CN', { minimumFractionDigits: 1, maximumFractionDigits: 1 })
}
function shortDate(value: string) {
  return value.length >= 16 ? value.slice(5, 16) : value
}
</script>

<style scoped>
.demo-shell {
  overflow: hidden;
  border-radius: 18px;
  border: 1px solid #efd9cc;
  background: #fafaf9;
  box-shadow: 0 1px 2px rgba(15, 23, 42, 0.04), 0 18px 40px -24px rgba(15, 23, 42, 0.24);
}
.demo-topbar {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  background: #0a0a0c;
  padding: 12px 16px;
}
.brand-mark {
  display: inline-flex;
  height: 32px;
  width: 32px;
  flex: none;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  background: #e8480f;
  color: #fff;
}
.demo-tabs {
  order: 3;
  display: flex;
  width: 100%;
  justify-content: center;
  border-radius: 999px;
  border: 1px solid rgba(255, 255, 255, 0.14);
  background: rgba(255, 255, 255, 0.05);
  padding: 3px;
}
@media (min-width: 640px) {
  .demo-tabs { order: 0; width: auto; }
}
.demo-tab {
  border-radius: 999px;
  padding: 6px 15px;
  font-size: 12px;
  color: #cbd5e1;
  transition: all 0.15s ease;
}
.demo-tab.active { background: #fff; color: #0f172a; }
.topbar-right { display: flex; flex-wrap: wrap; align-items: center; gap: 10px; }
.mask-toggle { display: inline-flex; align-items: center; gap: 6px; font-size: 11px; color: #94a3b8; }
.topbar-refresh {
  border-radius: 999px;
  border: 1px solid rgba(255, 255, 255, 0.16);
  padding: 4px 12px;
  font-size: 11px;
  color: #cbd5e1;
}
.topbar-refresh:hover { color: #fff; border-color: rgba(255, 255, 255, 0.32); }

.accounts-page {
  display: grid;
  grid-template-columns: 252px minmax(0, 1fr);
  gap: 14px;
  align-items: start;
  padding: 14px;
}
@media (max-width: 1023px) {
  .accounts-page { grid-template-columns: minmax(0, 1fr); }
  .acc-side { min-height: 0 !important; max-height: 360px !important; }
}
.acc-side {
  min-height: 680px;
  max-height: 780px;
  overflow-y: auto;
  border: 1px solid #e4e7ed;
  border-radius: 10px;
  background: #fff;
  padding: 10px;
  scrollbar-width: thin;
}
.side-head { display: flex; align-items: center; justify-content: space-between; padding: 2px 4px 8px; font-size: 13px; font-weight: 600; color: #303133; }
.side-ops { display: flex; gap: 2px; }
.side-ops button { padding: 0 4px; font-size: 11px; font-weight: 500; color: #e8480f; }
.device-group { border-radius: 8px; margin-bottom: 3px; }
.device-row { display: flex; align-items: center; gap: 4px; padding: 7px 6px; cursor: pointer; border-radius: 6px; font-size: 12px; }
.device-row:hover { background: #fafaf9; }
.device-row.is-off { opacity: 0.62; }
.dev-arrow { font-size: 12px; color: #909399; transition: transform 0.15s ease; }
.dev-arrow.open { transform: rotate(90deg); }
.dev-dot { height: 6px; width: 6px; flex: none; border-radius: 50%; background: #c0c4cc; }
.dev-dot.on { background: #35b37e; box-shadow: 0 0 0 2px rgba(53, 179, 126, 0.16); }
.dev-dot.off { background: #e25555; box-shadow: 0 0 0 2px rgba(226, 85, 85, 0.14); }
.dev-dot.pending { background: #e6a23c; box-shadow: 0 0 0 2px rgba(230, 162, 60, 0.16); }
.dev-name { min-width: 0; flex: 1 1 auto; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; font-weight: 500; color: #303133; }
.dev-add { flex: none; margin-left: -2px; padding: 0 2px; font-size: 10px; color: #e8480f; }
.dev-add:disabled { color: #c0c4cc; cursor: not-allowed; }
.dev-count { flex: 0 0 18px; text-align: center; font-size: 10px; color: #909399; }
.dev-act { flex: none; margin-left: auto; font-size: 10px; color: #909399; opacity: 0.8; }
.dev-act:hover { color: #f56c6c; opacity: 1; }
.dev-act.enable { color: #67c23a; }
.acc-list { display: flex; flex-direction: column; gap: 2px; padding: 2px 0 4px 18px; }
.acc-item { display: flex; align-items: center; gap: 6px; padding: 6px 8px; border-radius: 6px; font-size: 12px; color: #606266; text-align: left; }
.acc-item:hover { background: #fafaf9; }
.acc-item.active { background: #fff4ed; color: #e8480f; font-weight: 600; }
.acc-item.off { opacity: 0.6; }
.acc-dot { height: 7px; width: 7px; flex: none; border-radius: 50%; background: #d3dbe7; }
.acc-dot.on { background: #4c9a5f; box-shadow: 0 0 0 2px rgba(76, 154, 95, 0.18); }
.acc-dot.soon { background: #f2b73c; box-shadow: 0 0 0 2px rgba(242, 183, 60, 0.16); }
.acc-dot.off { background: #c0c4cc; }
.acc-label { min-width: 0; flex: 1; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; }
.off-tag { flex: none; font-size: 10px; color: #909399; }
.off-tag.warn { color: #e6a23c; }
.empty-dev { padding: 4px 6px 8px; font-size: 11px; line-height: 1.6; color: #a8abb2; }

.acc-main { display: flex; min-width: 0; flex-direction: column; gap: 12px; }
.det-head { display: flex; flex-wrap: wrap; align-items: center; justify-content: space-between; gap: 10px; border: 1px solid #e4e7ed; border-radius: 10px; background: #fff; padding: 11px 14px; }
.det-title { display: flex; min-width: 0; flex-wrap: wrap; align-items: center; gap: 8px; }
.det-label { font-size: 15px; font-weight: 600; color: #303133; }
.dev-key { display: inline-flex; height: 24px; width: 112px; align-items: center; justify-content: center; overflow: hidden; border-radius: 4px; border: 1px solid #ffd3bf; background: #fff4ed; padding: 0 6px; font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace; font-size: 10px; color: #e8480f; text-overflow: ellipsis; white-space: nowrap; }
.det-ops { display: flex; flex-wrap: wrap; align-items: center; justify-content: flex-end; gap: 6px; margin-left: auto; }
.lbl { font-size: 11px; color: #909399; }
.license-alert { display: flex; flex-wrap: wrap; align-items: center; justify-content: space-between; gap: 10px; border: 1px solid #faecd8; border-radius: 10px; background: #fdf6ec; padding: 10px 14px; }

.mod-section { border: 1px solid #e4e7ed; border-radius: 10px; background: #fff; padding: 13px 14px 15px; }
.mod-section-head { display: flex; align-items: flex-start; justify-content: space-between; gap: 10px; margin-bottom: 10px; }
.mod-section-head h4 { margin: 0; font-size: 13px; font-weight: 600; color: #303133; }
.mod-section-head p { margin: 3px 0 0; font-size: 11px; line-height: 1.6; color: #909399; }
.mod-count { flex: none; font-size: 11px; color: #909399; }
.mod-grid { display: grid; grid-template-columns: minmax(0, 1fr); gap: 10px; }
.module-card { border: 1px solid #e4e7ed; border-radius: 10px; background: #fff; overflow: hidden; transition: opacity 0.15s ease; }
.module-card.is-off { opacity: 0.72; }
.mc-head { display: flex; flex-wrap: wrap; align-items: center; justify-content: space-between; gap: 8px; border-bottom: 1px solid #faf2eb; background: #fafafa; padding: 8px 12px; }
.mc-title { display: flex; flex-wrap: wrap; align-items: center; gap: 5px; }
.mc-name { font-size: 13px; font-weight: 600; color: #303133; }
.mc-ops { display: flex; flex-wrap: wrap; align-items: center; gap: 6px; }
.mc-label { font-size: 11px; color: #909399; }
.mc-body { padding: 12px; }
.mc-foot { margin: 10px 0 0; font-size: 10px; color: #a8abb2; }

.pf-top { display: grid; grid-template-columns: 1fr 1fr 1fr; gap: 8px; border-bottom: 1px solid #faf2eb; padding-bottom: 10px; }
.pf-cell { display: flex; flex-direction: column; gap: 2px; }
.pf-cell.center { align-items: center; }
.pf-cell.right { align-items: flex-end; }
.pf-label { font-size: 10px; color: #909399; }
.pf-value { font-size: 14px; font-weight: 600; color: #303133; }
.pf-value.brand { color: #e8480f; }
.pf-nums { display: grid; grid-template-columns: 1fr 1fr 1fr; gap: 8px; padding-top: 10px; }
.pf-num { font-size: 14px; font-weight: 700; color: #303133; font-variant-numeric: tabular-nums; }
.pf-num.success { color: #67c23a; }
.pf-num.warning { color: #e6a23c; }

.afk-stats { display: grid; grid-template-columns: repeat(3, 1fr); gap: 8px; }
.afk-stat { display: flex; flex-direction: column; gap: 2px; border-radius: 6px; background: #fafafa; padding: 7px 9px; }
.afk-stat span { font-size: 10px; color: #909399; }
.afk-stat b { font-size: 13px; color: #303133; }
.afk-stat b.success { color: #67c23a; }
.afk-stat b.muted { font-size: 11px; font-weight: 500; color: #909399; }
.pass-box { margin-top: 10px; border: 1px solid #f0eeea; border-radius: 8px; background: #fcfcfd; padding: 9px 11px; }
.afk-row { display: flex; align-items: center; justify-content: space-between; gap: 8px; margin-top: 10px; }
.afk-times { margin-top: 10px; border-top: 1px solid #faf2eb; padding-top: 9px; }
.afk-win { display: flex; flex-wrap: wrap; align-items: center; gap: 6px; margin-top: 6px; font-size: 11px; }
.afk-win-text { color: #909399; }

.mc-table-scroll { overflow-x: auto; }
.mc-table { width: 100%; min-width: 520px; border-collapse: collapse; font-size: 12px; table-layout: fixed; }
.mc-table th { border-bottom: 1px solid #f0eeea; padding: 6px 6px; text-align: left; font-size: 11px; font-weight: 500; color: #909399; white-space: nowrap; }
.mc-table td { border-bottom: 1px solid #faf2eb; padding: 6px; color: #606266; }
.mc-chip { display: inline-flex; max-width: 100%; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; border-radius: 4px; padding: 1px 6px; font-size: 10px; }
.mc-chip.busy { background: #fff4ed; color: #e8480f; }
.mc-chip.done { background: #f4f4f5; color: #909399; }
.add-row { margin-top: 7px; }
.mod-empty { display: flex; min-height: 220px; flex-direction: column; align-items: center; justify-content: center; gap: 7px; border: 1px dashed #efd9cc; border-radius: 12px; background: #fafafa; padding: 26px 20px; text-align: center; }
.mod-empty-icon { display: flex; height: 48px; width: 48px; align-items: center; justify-content: center; border-radius: 14px; border: 1px solid #ffd3bf; background: #fff4ed; color: #e8480f; }
.mod-empty-icon svg { height: 24px; width: 24px; }
.mod-empty-title { font-size: 14px; font-weight: 600; color: #303133; }
.mod-empty-desc { max-width: 380px; font-size: 12px; line-height: 1.7; color: #909399; }
.acc-empty { display: flex; min-height: 280px; flex-direction: column; align-items: center; justify-content: center; border: 1px solid #e4e7ed; border-radius: 10px; background: #fff; padding: 30px; text-align: center; }

.status-page { padding: 14px; }
.status-banner { display: flex; flex-wrap: wrap; align-items: center; justify-content: space-between; gap: 10px; margin-bottom: 12px; padding: 2px 2px 0; }
.status-banner h2 { margin: 0; font-size: 18px; font-weight: 600; color: #303133; }
.status-banner p { margin: 3px 0 0; font-size: 12px; color: #909399; }
.banner-actions { display: flex; flex-wrap: wrap; align-items: center; gap: 10px; }
.biz-group { margin-bottom: 12px; overflow: hidden; border: 1px solid #e4e7ed; border-radius: 10px; background: #fff; }
.group-head { display: flex; flex-wrap: wrap; align-items: center; gap: 8px; border-bottom: 1px solid #f0eeea; background: #fafafa; padding: 9px 12px; cursor: pointer; }
.group-head:hover { background: #fafaf9; }
.group-chevron { font-size: 12px; color: #909399; transition: transform 0.15s ease; }
.group-chevron.open { transform: rotate(90deg); }
.device-title { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace; font-size: 12px; font-weight: 600; color: #303133; }
.device-chips { display: flex; min-width: 0; flex-wrap: wrap; align-items: center; gap: 4px 8px; }
.device-chips span { border-radius: 3px; background: #fff; padding: 2px 6px; font-size: 10px; line-height: 1.5; color: #909399; border: 1px solid #f0eeea; }
.table-scroll { overflow-x: auto; }
.grouped-table { min-width: 1034px; }
.ungrouped-table { min-width: 1090px; }
.biz-table { width: 100%; border-collapse: collapse; table-layout: fixed; font-size: 12px; }
.biz-table th { border-bottom: 1px solid #f0eeea; background: #fff; padding: 8px 7px; text-align: left; font-size: 11px; font-weight: 500; color: #909399; white-space: nowrap; }
.biz-table td { overflow: hidden; border-bottom: 1px solid #faf2eb; padding: 8px 7px; color: #606266; text-overflow: ellipsis; white-space: nowrap; }
.data-row { cursor: pointer; }
.data-row:hover td { background: #f7f6f3; }
.data-row.is-off { opacity: 0.6; }
.data-row .mono { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace; font-size: 11px; }
.data-row .muted { color: #909399; }
.data-row .num { font-variant-numeric: tabular-nums; }
.row-chevron { display: inline-block; color: #c0c4cc; font-size: 11px; transition: transform 0.15s ease; }
.row-chevron.open { transform: rotate(90deg); }
.status-dot { display: inline-block; height: 6px; width: 6px; margin-right: 4px; border-radius: 50%; background: #c0c4cc; vertical-align: 1px; }
.status-dot.run { background: #67c23a; }
.status-dot.hang { background: #e8480f; }
.status-dot.idle { background: #909399; }
.status-dot.error { background: #f56c6c; }
.status-dot.off { background: #c0c4cc; }
.ratio-cell { display: flex; align-items: center; gap: 5px; }
.mini-progress { display: inline-block; height: 6px; width: 48px; overflow: hidden; border-radius: 4px; background: #f0eeea; }
.mini-progress span { display: block; height: 100%; border-radius: 4px; background: #e8480f; }
.warehouse-bar { height: 8px; overflow: hidden; border-radius: 4px; background: #f0eeea; }
.warehouse-bar span { display: block; height: 100%; border-radius: 4px; background: #e8480f; transition: width 0.2s ease; }
.expand-row td { border-bottom: 1px solid #f0eeea; background: #f7f6f3; padding: 10px 14px; white-space: normal; }
.expand-inner { display: flex; flex-wrap: wrap; align-items: center; gap: 6px; }
.produce-chip { border: 1px solid #f0eeea; border-radius: 4px; background: #fff; padding: 2px 7px; font-size: 10px; color: #606266; }
.produce-chip.done { color: #909399; background: #f4f4f5; }
.th-sort { display: inline-flex; align-items: center; gap: 3px; font-size: 11px; color: inherit; }
.th-sort:hover { color: #e8480f; }
.th-sort i { font-style: normal; font-size: 9px; color: #e8480f; }
.grand-total { display: flex; flex-wrap: wrap; align-items: center; justify-content: space-between; gap: 8px; border: 1px solid #e4e7ed; border-radius: 10px; background: #fff; padding: 10px 14px; font-size: 12px; color: #606266; }
.grand-total b { color: #303133; font-variant-numeric: tabular-nums; }

.el-tag { display: inline-flex; flex: none; align-items: center; border: 1px solid; border-radius: 4px; padding: 0 5px; font-size: 10px; line-height: 18px; white-space: nowrap; }
.tag-primary { border-color: #ffd3bf; background: #fff4ed; color: #e8480f; }
.tag-success { border-color: #e1f3d8; background: #f0f9eb; color: #67c23a; }
.tag-warning { border-color: #faecd8; background: #fdf6ec; color: #e6a23c; }
.tag-danger { border-color: #fde2e2; background: #fef0f0; color: #f56c6c; }
.tag-info { border-color: #e9e9eb; background: #f4f4f5; color: #909399; }

.switch { position: relative; display: inline-flex; height: 16px; width: 32px; flex: none; align-items: center; border-radius: 999px; background: #efd9cc; transition: background 0.15s ease; }
.switch.on { background: #e8480f; }
.switch i { position: absolute; left: 2px; top: 2px; height: 12px; width: 12px; border-radius: 50%; background: #fff; transition: left 0.15s ease; }
.switch.on i { left: 18px; }
.prio-input, .time-input { height: 26px; border: 1px solid #efd9cc; border-radius: 4px; background: #fff; padding: 0 5px; font-size: 11px; color: #606266; outline: none; }
.prio-input { width: 54px; text-align: center; }
.prio-input:focus, .time-input:focus, .mini-select:focus, .form-control:focus { border-color: #e8480f; }
.time-input { width: 88px; }
.mini-select { height: 26px; max-width: 100%; border: 1px solid #efd9cc; border-radius: 4px; background: #fff; padding: 0 4px; font-size: 11px; color: #606266; outline: none; }

.el-btn { display: inline-flex; height: 28px; flex: none; align-items: center; justify-content: center; gap: 4px; border: 1px solid #efd9cc; border-radius: 4px; background: #fff; padding: 0 10px; font-size: 11px; color: #606266; transition: all 0.15s ease; }
.el-btn:hover { border-color: #ffb99c; background: #fff4ed; color: #e8480f; }
.el-btn-primary { border-color: #e8480f; background: #e8480f; color: #fff; }
.el-btn-primary:hover { border-color: #ff7a3d; background: #ff7a3d; color: #fff; }
.el-btn-danger { border-color: #f56c6c; background: #f56c6c; color: #fff; }
.el-btn-danger:hover { border-color: #f78989; background: #f78989; color: #fff; }
.el-btn-text { flex: none; padding: 0 4px; font-size: 11px; color: #e8480f; }
.el-btn-text:hover { color: #ff7a3d; }
.el-btn-text.danger, .danger-text { color: #f56c6c; }
.el-btn-text.danger:hover, .danger-text:hover { color: #f78989; }

.demo-toast { position: fixed; bottom: 28px; left: 50%; z-index: 80; transform: translateX(-50%); border-radius: 8px; background: rgba(15, 23, 42, 0.92); padding: 9px 18px; font-size: 12px; color: #fff; box-shadow: 0 12px 30px -12px rgba(0, 0, 0, 0.4); }

.modal-mask { position: fixed; inset: 0; z-index: 70; display: flex; align-items: center; justify-content: center; background: rgba(0, 0, 0, 0.45); padding: 16px; }
.modal-panel { width: 560px; max-width: 100%; max-height: 88vh; overflow-y: auto; border-radius: 12px; background: #fff; box-shadow: 0 24px 60px -20px rgba(0, 0, 0, 0.4); }
.modal-head { display: flex; align-items: flex-start; justify-content: space-between; gap: 12px; border-bottom: 1px solid #f0eeea; padding: 16px 18px 13px; }
.modal-head h3 { margin: 0; font-size: 16px; font-weight: 600; color: #303133; }
.modal-head p { margin: 3px 0 0; font-size: 11px; line-height: 1.5; color: #909399; }
.modal-close { flex: none; font-size: 18px; line-height: 1; color: #909399; }
.modal-close:hover { color: #e8480f; }
.modal-body { padding: 14px 18px; }
.modal-foot { display: flex; flex-wrap: wrap; justify-content: flex-end; gap: 8px; border-top: 1px solid #f0eeea; padding: 12px 18px; }
.form-row { display: flex; align-items: center; gap: 10px; margin-bottom: 11px; font-size: 12px; color: #606266; }
.form-row > span:first-child { width: 70px; flex: none; color: #909399; }
.form-control { width: 100%; height: 32px; border: 1px solid #efd9cc; border-radius: 4px; background: #fff; padding: 0 9px; font-size: 12px; color: #303133; outline: none; }
.segmented { display: inline-flex; border: 1px solid #efd9cc; border-radius: 4px; overflow: hidden; }
.segmented button { padding: 5px 14px; font-size: 12px; color: #606266; background: #fff; }
.segmented button.active { background: #e8480f; color: #fff; }
.module-cand { display: flex; width: 100%; align-items: center; gap: 10px; border-bottom: 1px solid #faf2eb; padding: 9px 2px; text-align: left; }
.module-cand:hover { background: #f7f6f3; }
.module-cand b { display: block; font-size: 13px; color: #303133; }
.module-cand small { display: block; margin-top: 2px; font-size: 11px; color: #909399; }
.renew-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 10px 14px; }
.renew-grid > div { display: flex; flex-direction: column; align-items: flex-start; gap: 3px; border-radius: 6px; background: #fafafa; padding: 8px 10px; }
.renew-grid span:first-child { font-size: 10px; color: #909399; }
.renew-grid b { font-size: 12px; color: #303133; }
.renew-options { display: grid; grid-template-columns: repeat(3, 1fr); gap: 8px; margin-top: 8px; }
.renew-option { display: flex; flex-direction: column; align-items: flex-start; gap: 2px; border: 1px solid #efd9cc; border-radius: 6px; padding: 9px 11px; text-align: left; }
.renew-option b { font-size: 13px; color: #303133; }
.renew-option span { font-size: 10px; color: #909399; }
.renew-option.active { border-color: #e8480f; background: #fff4ed; }
.renew-option.active b { color: #e8480f; }
.token-stats { display: grid; grid-template-columns: repeat(3, 1fr); gap: 8px; }
.token-stats > div { display: flex; flex-direction: column; align-items: center; gap: 2px; border-radius: 8px; background: #fafafa; padding: 10px; }
.token-stats b { font-size: 20px; color: #303133; font-variant-numeric: tabular-nums; }
.token-stats span { font-size: 10px; color: #909399; }

@media (max-width: 640px) {
  .status-page, .accounts-page { padding: 10px; }
  .mc-head { align-items: flex-start; }
  .mc-ops { width: 100%; justify-content: flex-start; gap: 5px; }
  .mc-ops .prio-input { width: 48px; }
  .mc-ops .mini-select { width: 58px; }
  .mc-ops .el-btn-text.danger { margin-left: auto; }
  .demo-topbar { padding: 10px 12px; }
  .status-banner h2 { font-size: 16px; }
  .device-chips span:nth-child(n + 4) { display: none; }
  .grand-total { font-size: 11px; }
  .renew-grid { grid-template-columns: 1fr; }
  .renew-options { grid-template-columns: 1fr; }
}
</style>
