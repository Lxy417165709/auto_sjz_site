import type { RouteRecordRaw } from 'vue-router'

export const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'home',
    component: () => import('@/pages/HomePage.vue'),
    meta: { title: 'auto_sjz - 按号收费的 Android 设备管理平台', description: '一个号 3 元/月，设备不按台收费。中台批量管理设备与账号，配置统一、异常通知，授权按账号实例结算。' },
  },
  {
    path: '/features',
    name: 'features',
    component: () => import('@/pages/FeaturesPage.vue'),
    meta: { title: '业务能力 - auto_sjz', description: '按号收费、中台批量管理、监控通知、统一配置下发与到期保留账号，业务视角看 auto_sjz 能解决什么。' },
  },
  {
    path: '/pricing',
    name: 'pricing',
    component: () => import('@/pages/PricingPage.vue'),
    meta: { title: '按号收费 - auto_sjz', description: '3 元/号/月，按账号实例计费，设备台数不计费。一台设备 5 个号即 15 元/月。' },
  },
  {
    path: '/start',
    name: 'start',
    component: () => import('@/pages/StartPage.vue'),
    meta: { title: '接入流程 - auto_sjz', description: '设备如何接入云控、如何添加账号并授权、如何添加模块：三步接入顺序与常见卡点。' },
  },
  {
    path: '/tech',
    name: 'tech',
    component: () => import('@/pages/TechPage.vue'),
    meta: { title: '技术架构 - auto_sjz', description: 'Manager 根节点 + 子节点的树形部署、设备节点归属、授权验签链路、默认端口与技术栈。' },
  },
  {
    path: '/console',
    name: 'console',
    component: () => import('@/pages/ConsolePage.vue'),
    meta: { title: '管理台演示 - auto_sjz', description: '可交互的管理台示意：业务配置（设备、账号、模块）与业务状态（在线、执行、异常通知）。' },
  },
  {
    path: '/deploy',
    name: 'deploy',
    component: () => import('@/pages/DeployPage.vue'),
    meta: { title: '部署接入 - auto_sjz', description: 'Docker 部署 discovery 与 Manager，根节点与子节点分开编排，设备通过 discovery 自动寻址。' },
  },
  {
    path: '/docs',
    name: 'docs',
    component: () => import('@/pages/DocsPage.vue'),
    meta: { title: '文档 - auto_sjz', description: '业务快速开始、系统架构与端口、授权计费、监控通知，以及开发部署文档。' },
  },
  {
    path: '/faq',
    name: 'faq',
    component: () => import('@/pages/FaqPage.vue'),
    meta: { title: '常见问题 - auto_sjz', description: '按号收费、授权到期、批量管理、监控通知与部署相关问题。' },
  },
  {
    path: '/contact',
    name: 'contact',
    component: () => import('@/pages/ContactPage.vue'),
    meta: { title: '联系与咨询 - auto_sjz', description: '咨询报价、接入方案与私有化部署。' },
  },
  { path: '/:pathMatch(.*)*', redirect: '/' },
]
