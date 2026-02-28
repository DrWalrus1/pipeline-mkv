<script setup lang="ts">
import type { DiscInfo } from '@/domain/disc_info';
import DiscInfoCard from '../components/DiscInfoCard.vue';

const urls = [
  "https://m.media-amazon.com/images/I/91+ShpVWyiL._AC_UF894,1000_QL80_.jpg",
  "https://posters.movieposterdb.com/21_08/2020/11032374/l_11032374_5dc62e2c.jpg",
  "https://posters.movieposterdb.com/25_05/2025/32820897/l_gekijo-ban-kimetsu-no-yaiba-mugen-jo-hen-movie-poster_c68cb10f.jpg"
]
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
  <div class="disc-info-grid-container">
    <DiscInfoCard v-for="(item, index) in fakeDiscHistoryList" class="disc-info-grid-item" :discInfo=item
      :url=urls[index] />
  </div>

</template>

<style>
.disc-info-grid-container {
  display: grid;
  grid-template-rows: 1fr 1fr 1fr;
  grid-template-columns: 1fr 1fr 1fr;
  justify-items: center;
  align-items: center;
}

/* .disc-info-grid-item { */
/**/
/* } */
</style>
