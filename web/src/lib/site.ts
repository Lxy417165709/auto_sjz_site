import { reactive } from 'vue'
import type { SiteData } from './types'

interface SiteState {
  data: SiteData | null
  loading: boolean
  error: string
}

export const siteState = reactive<SiteState>({
  data: null,
  loading: false,
  error: '',
})

let pending: Promise<void> | null = null

/** 站点文案由 Go 服务的 /api/site 一次返回，整站只请求一次。 */
export function loadSite(): Promise<void> {
  if (pending) return pending
  siteState.loading = true
  pending = fetch('/api/site', { headers: { Accept: 'application/json' } })
    .then(async (resp) => {
      if (!resp.ok) throw new Error(`HTTP ${resp.status}`)
      siteState.data = (await resp.json()) as SiteData
      siteState.error = ''
    })
    .catch((err: unknown) => {
      siteState.error = err instanceof Error ? err.message : '内容加载失败'
    })
    .finally(() => {
      siteState.loading = false
      pending = null
    })
  return pending
}

/** 登录与购买统一跳转 Manager；未配置时退回联系页。 */
export function managerEntry(): string {
  const fromMeta = siteState.data?.meta.managerUrl?.trim()
  if (fromMeta) return fromMeta
  const fromContent = siteState.data?.site.cta?.href?.trim()
  if (fromContent) return fromContent
  return '/contact'
}
