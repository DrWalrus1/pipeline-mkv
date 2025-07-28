<script setup lang="ts">
import { useWebSocket } from '@vueuse/core'
import { ref, watch, computed } from 'vue'
import { API_URL } from '@/config'

const discInfo = ref<any>(null)

function LoadDiscInfo() {
  const socket = useWebSocket(`ws://${API_URL}/api/info`, {
    heartbeat: {
      message: 'ping',
      interval: 1000,
      pongTimeout: 1000
    }
  });
  watch(socket.data, (newData) => {
    discInfo.value = JSON.parse(newData)
  })
}

const formattedDiscInfo = computed(() =>
  discInfo.value
    ? JSON.stringify(discInfo.value, null, 2)
    : '{\n  "Hello": "World"\n}'
)
</script>

<template>
  <main>
    <button @click="LoadDiscInfo">Load Disc Info</button>
    <pre id="disc-info-dump">{{ formattedDiscInfo }}</pre>
  </main>
</template>
