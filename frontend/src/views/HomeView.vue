<script setup lang="ts">
import { computedAsync } from '@vueuse/core'
import { inject } from 'vue'
// import { GetMakeMkvService, type MakeMkvService } from '@/services/gomakemkv/gomakemkv'
import Navbar from '@/components/Navbar.vue'
import DiscInfoGrid from '@/layouts/DiscInfoGrid.vue'
// const makeMkvService = inject<MakeMkvService>("MakeMkvService")
// GetMakeMkvService(makeMkvService)
//
//
// const discInfo = computedAsync(async () => { return await makeMkvService.GetInsertedDiscInfo() }, "Goodbye")

</script>

<template>
  <Navbar :inProgressMovieCount=5 />
  <DiscInfoGrid style="padding-top: 10px;" />
</template>
