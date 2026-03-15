<script setup lang="ts">
const { titleId, titleName } = defineProps<{ titleId: string, titleName: string }>();
</script>

<template>
  <div style="display: block;">
    <input :id=titleId type="checkbox"></input>
    <label for={{titleId}}>{{ titleName }}</label>
  </div>
</template>
