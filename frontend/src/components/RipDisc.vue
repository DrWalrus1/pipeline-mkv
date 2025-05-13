<script lang="ts" setup>
import { ref, defineProps } from 'vue';
const props = defineProps<{
  titles: Array<{ id: number; name: string }>;
}>();

</script>

<template>
  <select>
    <option value="disc:1">Disc 1</option>
    <option value="disc:2">Disc 2</option>
    <option value="disc:3">Disc 3</option>
  </select>
  <br>
  <div v-for="title in props.titles" :key="title.id">
    <input type="checkbox" :name="'title-' + title.id" :value="title.id">{{ title.name }}</input>
    <br>
  </div>
  <Button>Rip Disc</Button>
</template>
