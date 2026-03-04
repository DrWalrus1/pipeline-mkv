<script setup lang="ts">
import type { DiscInfo } from '@/domain/disc_info';
const { discInfo, url } = defineProps<{ discInfo: DiscInfo, url: string }>();
let titleSize = discInfo.name.length > 15 ? "long-title" : "normal-title"

</script>

<template>
  <div class="discInfoCard">
    <img class="discPoster" :src="url" />
    <div class="discInfoCardInfo">
      <h1 :class=titleSize>{{ discInfo.name }}</h1>
      <h5>{{ discInfo.language }}</h5>
      <h5>{{ discInfo.type }}</h5>
    </div>
  </div>

</template>

<style lang="css">
/* Common Styling */
.discInfoCard {
  border-radius: 10px;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
  /*TODO: Sync to theme */
  background-color: white;
  transition: 0.3s;

}

.discInfoCard:hover {
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
  cursor: pointer;
}

.discPoster {
  aspect-ratio: auto;
}

/* Small Mobile */
@media (min-width:20px) {
  .discInfoCard {
    height: 10rem;
    display: inline-flex;
    width: 100%;
  }

  .discPoster {
    height: 10rem;
    border-radius: 10px 0px 0px 10px;
  }

  .discInfoCardInfo {
    padding-left: 5px;
    padding-right: 5px;
  }

  .discInfoCardInfo>h1 {
    margin-top: 5px;
    margin-bottom: 5px;
    margin-right: 5px;
  }

  .normal-title {
    font-size: 24px
  }

  .long-title {
    font-size: 18px
  }

  .discInfoCardInfo>h5 {
    margin-top: 10px;
    margin-bottom: 10px;
  }
}

/* laptops and desktops */
@media (min-width:800px) {
  .discInfoCard {
    display: block;
    width: 275px;
  }

  .discPoster {
    border-radius: 10px 10px 0 0;
    width: 275px;
    height: 440px;
  }

  .discInfoCardInfo {
    padding: 2px 16px;
  }
}
</style>
