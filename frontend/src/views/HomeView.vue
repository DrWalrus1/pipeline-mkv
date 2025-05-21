<script setup lang="ts">
import { useWebSocket } from '@vueuse/core'
import { watch } from 'vue'
function LoadDiscInfo() {
  const socket = useWebSocket('ws:///api/info', {
    heartbeat: {
      message: 'ping',
      interval: 1000,
      pongTimeout: 1000
    }
  });
  watch(socket.data, (newData) => {
    console.log(newData)
  })

}
</script>

<template>
  <main>
    <button @click="LoadDiscInfo">Load Disc Info</button>
    <pre>
      <code id="disc-info-dump">
        {
        "Hello": "World"
        }
      </code>
    </pre>
  </main>
</template>
