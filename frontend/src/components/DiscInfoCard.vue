<script setup lang="ts">
import type { DiscInfo } from '@/domain/disc_info';
import { defineProps } from 'vue';
const { discInfo, url } = defineProps<{ discInfo: DiscInfo, url: string }>();
</script>

<template>
  <div class="discInfoCard">
    <img class="discPoster" :src="url" />
    <div class="discInfoCardInfo">
      <h1>{{ discInfo.name }}</h1>
      <h5>{{ discInfo.language }}</h5>
      <h5>{{ discInfo.type }}</h5>
    </div>
  </div>

</template>

<style lang="css">
.discInfoCard {
  border-radius: 10px;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
  transition: 0.3s;
  width: 275px;
}

.discInfoCard:hover {
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
  cursor: pointer;
}

.discPoster {
  aspect-ratio: auto;
  border-radius: 10px 10px 0 0;
  width: 275px;
}

.discInfoCardInfo {
  padding: 2px 16px;
}
</style>
