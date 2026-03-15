<script setup lang="ts">
import TitleSelection from '@/components/TitleSelection.vue';
/*
  TODO:
  1. Get hardware/disc drive info
  2. Show list of drives
  3. Pick drive -> sends request to backend to load disc info
  4. show rip options
  5. click rip button
  */
</script>

<template>
  <!-- extract into selection component -->
  <label for="disc-selection">Select Disc:</label>
  <select id="disc-selection">
    <option value="disc:0">Disc 0</option>
  </select>

  <!--Disc info -->
  <div id="disc-info">
    <!-- show poster maybe -->
    <h1>Demon Slayer</h1>
    <h2>Mugen Train Arc</h2>
    <h3>Titles</h3>
    <div>
      <!-- on click, show title metadata -->
      <TitleSelection titleId="title-1" titleName="Title 1" />
      <TitleSelection titleId="title-2" titleName="Title 2" />
      <TitleSelection titleId="title-3" titleName="Title 3" />
    </div>
  </div>

  <!-- Rip Button -->
  <button>Rip Disc</button>
</template>
