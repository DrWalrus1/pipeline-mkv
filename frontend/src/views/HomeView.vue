<script setup lang="ts">
import { computedAsync } from '@vueuse/core'
import { inject } from 'vue'
import { GetMakeMkvService, type MakeMkvService } from '@/services/gomakemkv/gomakemkv'
const makeMkvService = inject<MakeMkvService>("MakeMkvService")
GetMakeMkvService(makeMkvService)


const discInfo = computedAsync(async () => { return await makeMkvService.GetInsertedDiscInfo() }, "Goodbye")

</script>

<template>
  <main>
    <button>Load Disc Info</button>
    <pre id="disc-info-dump">{{ discInfo }}</pre>
  </main>
</template>
