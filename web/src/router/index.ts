import type { RouteRecordRaw } from 'vue-router'

export const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'home',
    component: () => import('@/pages/HomePage.vue'),
    meta: { title: 'auto_sjz - Android 设备自动化管理平台', description: '统一管理 Android 设备、账号与自动化任务，Manager 统一下发配置，授权按账号实例计量。' },
  },
  {
    path: '/features',
    name: 'features',
    component: () => import('@/pages/FeaturesPage.vue'),
    meta: { title: '功能 - auto_sjz', description: '设备管理、配置下发、自动寻址、授权验签与多节点扩展，覆盖 Android 设备自动化全链路。' },
  },
  {
    path: '/pricing',
    name: 'pricing',
    component: () => import('@/pages/PricingPage.vue'),
    meta: { title: '按号收费 - auto_sjz', description: '按账号实例计费，一台设备 5 个账号即 5 份授权，设备台数不计费。' },
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
    meta: { title: '文档 - auto_sjz', description: '快速开始、系统架构、授权计费与配置下发说明。' },
  },
  {
    path: '/faq',
    name: 'faq',
    component: () => import('@/pages/FaqPage.vue'),
    meta: { title: '常见问题 - auto_sjz', description: '按号收费、授权到期、discovery 接入与部署相关问题。' },
  },
  {
    path: '/contact',
    name: 'contact',
    component: () => import('@/pages/ContactPage.vue'),
    meta: { title: '联系与试用 - auto_sjz', description: '预约演示、咨询报价与私有化部署方案。' },
  },
  { path: '/:pathMatch(.*)*', redirect: '/' },
]
