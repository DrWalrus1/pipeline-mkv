<script setup lang="ts">
import type { DiscInfo } from '@/domain/disc_info';
import DiscInfoCard from '../components/DiscInfoCard.vue';

const fakeDiscHistoryList: DiscInfo[] = [
  {
    name: "Movie 1",
    language: "English",
    type: "Blu-ray",
    titles: []
  },
  {
    name: "Movie 2",
    language: "English",
    type: "DVD",
    titles: []
  },
  {
    name: "Movie 3",
    language: "Japanese",
    type: "Blu-ray",
    titles: []
  }
]
</script>
<template>
  <div class="disc-info-grid-item">
    <div v-for="(item, index) in fakeDiscHistoryList" class="disc-info-grid-item">
      <DiscInfoCard :discInfo=item url="https://m.media-amazon.com/images/I/91+ShpVWyiL._AC_UF894,1000_QL80_.jpg" />
    </div>
  </div>

</template>

<style>
.disc-info-grid-container {
  display: grid;
  grid-auto-rows: 1fr 1fr 1fr;
  grid-auto-columns: 100px;
}
</style>
