<template>
  <div class="flex min-h-screen flex-col bg-white">
    <SiteHeader />
    <main class="flex-1">
      <router-view />
    </main>
    <SiteFooter />
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import SiteHeader from '@/components/SiteHeader.vue'
import SiteFooter from '@/components/SiteFooter.vue'
import { loadSite } from '@/lib/site'

onMounted(() => {
  void loadSite()
})
</script>
