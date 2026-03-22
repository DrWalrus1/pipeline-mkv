<script setup lang="ts">
import { reactive } from 'vue'
import type { Settings } from './settings'

const props = defineProps<Settings>()
const emit = defineEmits<{
  (e: "submit", payload: Settings): void
}>()

const localSettings = reactive<Settings>({ ...props })
</script>

<template>
  <label for="executablePathInput">Executable Path: </label>
  <input type="text" id="executablePathInput" v-model="localSettings.executablePath">
  <br>
  <label for="registrationKeyInput">MakeMkv Registration Key: </label>
  <input type="text" id="registrationKeyInput" v-model="localSettings.registrationKey">
  <br>
  <label for="metadataServiceTokenInput">The Movie Database API Key: </label>
  <input type="text" id="metadataServiceTokenInput" v-model="localSettings.metadataServiceToken">
  <br>
  <button @click="emit('submit', { ...localSettings })">Update Settings</button>
</template>
