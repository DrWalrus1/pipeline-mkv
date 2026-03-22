import { createRouter, createWebHistory } from "vue-router";
import HomeView from "./views/HomeView.vue";
import NewMovie from "./views/NewMovie.vue";
import SettingsView from "./views/settings/SettingsView.vue";
import type { Settings } from "./views/settings/settings";

const routes = [
  { path: "/", component: HomeView },
  { path: "/new", component: NewMovie },
  {
    path: "/settings",
    component: SettingsView,
    props: {
      executablePath: "Hello",
      registrationKey: "registrationKey",
      metadataServiceToken: "metadataToken",
      onSubmit: (settings: Settings) => alert(`Hello ${settings.executablePath}, ${settings.metadataServiceToken}, ${settings.registrationKey}`)
    }
  }
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})
