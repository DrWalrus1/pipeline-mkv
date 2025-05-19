<script setup lang="ts">
import { ref, defineProps } from 'vue';
const props = defineProps<{
  disc: {
    id: number;
    name: string;
    titles: Array<{ id: number; name: string }>;
  };
}>();
</script>

<template>

</template>
