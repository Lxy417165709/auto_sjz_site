import type { RouteLocationNormalized } from 'vue-router'

function upsertMeta(attr: 'name' | 'property', key: string, value: string) {
  let el = document.head.querySelector<HTMLMetaElement>(`meta[${attr}="${key}"]`)
  if (!el) {
    el = document.createElement('meta')
    el.setAttribute(attr, key)
    document.head.appendChild(el)
  }
  el.setAttribute('content', value)
}

/** 每次路由切换更新标题、描述与 canonical，确保单页应用也能被正常索引。 */
export function applySeo(to: RouteLocationNormalized) {
  const title = (to.meta.title as string) || '炎火'
  const description = (to.meta.description as string) || ''

  document.title = title
  if (description) upsertMeta('name', 'description', description)

  const canonical = `${window.location.origin}${to.path === '/' ? '/' : to.path.replace(/\/$/, '')}`
  let link = document.head.querySelector<HTMLLinkElement>('link[rel="canonical"]')
  if (!link) {
    link = document.createElement('link')
    link.rel = 'canonical'
    document.head.appendChild(link)
  }
  link.href = canonical

  upsertMeta('property', 'og:title', title)
  if (description) upsertMeta('property', 'og:description', description)
  upsertMeta('property', 'og:url', canonical)
}
