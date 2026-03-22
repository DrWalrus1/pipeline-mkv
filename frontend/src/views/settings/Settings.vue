<script setup lang="ts">
import type { Settings } from './settings';

const settings = defineProps<Settings>()
</script>

<template>
  <label for="executablePathInput">Executable Path: </label>
  <input type="text" name="executablePath" id="executablePathInput" :value=settings.executablePath>
  <br>
  <label for="registrationKeyInput">MakeMkv Registration Key: </label>
  <input type="text" name="registrationPath" id="registrationKeyInput" :value=settings.registrationKey>
  <br>
  <label for="metadataServiceTokenInput">The Movie Database API Key: </label>
  <input type="text" name="metadataServiceToken" id="metadataServiceTokenInput" :value=settings.metadataServiceToken>
</template>
