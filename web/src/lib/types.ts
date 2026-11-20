export interface LinkItem {
  label: string
  href: string
}

export interface NavItem extends LinkItem {}

export interface Contact {
  email: string
  wechat: string
  working: string
}

export interface FooterColumn {
  title: string
  links: LinkItem[]
}

export interface Site {
  name: string
  tagline: string
  description: string
  nav: NavItem[]
  cta: LinkItem
  secondaryCta: LinkItem
  contact: Contact
  footer: FooterColumn[]
  icp: string
  copyright: string
}

export interface FeatureItem {
  title: string
  desc: string
}

export interface PricingCalculator {
  defaultDevices: number
  defaultAccountsPerDevice: number
  hint: string
}

export interface Pricing {
  headline: string
  summary: string
  unitLabel: string
  unitPriceCny: number | null
  currency: string
  billingRules: string[]
  calculator: PricingCalculator
  includes: FeatureItem[]
  notes: string[]
}

export interface FAQItem {
  id: string
  group: string
  question: string
  answer: string
}

export interface DocStepItem {
  text: string
  href?: string
  linkLabel?: string
}

export interface DocStep {
  title: string
  items: DocStepItem[]
}

export interface DocSection {
  heading: string
  paragraphs: string[]
  bullets: string[]
  steps?: DocStep[]
}

export interface Doc {
  slug: string
  audience?: 'business' | 'developer'
  title: string
  summary: string
  updated: string
  sections: DocSection[]
}

export interface SiteData {
  site: Site
  pricing: Pricing
  faq: FAQItem[]
  docs: Doc[]
  meta: {
    generatedAt: string
    managerUrl: string
  }
}
