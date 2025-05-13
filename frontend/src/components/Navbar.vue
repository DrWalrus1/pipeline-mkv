<script setup lang="ts">
import { defineComponent, ref } from 'vue';

const searchQuery = ref('');

const handleNewClick = () => {
  // Handle new button click logic here
  console.log('New button clicked');
  console.log('Search Query:', searchQuery.value);
};

const handleSettingsClick = () => {
  // Handle settings button click logic here
  console.log('Settings button clicked');
};

</script>

<template>
  <nav class="navbar">
    <div class="logo">PIPELINEMKV</div>
    <div class="input-container">
      <input type="text" placeholder="5 Movies in progress..." class="input-text" />
      <button class="new-button">New</button>
    </div>
    <button class="settings-button">Settings</button>
  </nav>
</template>

<style scoped>
.navbar {
  background-color: #374151;
  padding: 1rem;
  border-radius: 0.5rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  display: flex;
  align-items: center;
  gap: 1rem;
}

.logo {
  color: #ffffff;
  font-weight: bold;
  font-size: 1.25rem;
}

.input-container {
  flex-grow: 1;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.input-text {
  background-color: #ffffff;
  color: #374151;
  padding: 0.5rem 1rem;
  border-radius: 0.375rem;
  outline: none;
  width: 200px;
  border: none;
}

.input-text:focus {
  outline: 2px solid #3b82f6;
}

.new-button {
  background-color: #2563eb;
  color: #ffffff;
  padding: 0.5rem 1rem;
  border-radius: 0.375rem;
  cursor: pointer;
  border: none;
}

.new-button:hover {
  background-color: #1e40af;
}

.new-button:focus {
  outline: 2px solid #3b82f6;
}

.settings-button {
  background-color: #4b5563;
  color: #ffffff;
  padding: 0.5rem 1rem;
  border-radius: 0.375rem;
  cursor: pointer;
  border: none;
}

.settings-button:hover {
  background-color: #374151;
}

.settings-button:focus {
  outline: 2px solid #6b7280;
}
</style>
