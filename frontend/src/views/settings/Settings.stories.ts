import type { Meta, StoryObj } from "@storybook/vue3-vite";
import SettingsView from "./SettingsView.vue";
import type { Settings } from "./settings";

const meta: Meta<typeof SettingsView> = {
  component: SettingsView
}

export default meta;
type Story = StoryObj<typeof SettingsView>;

let settings: Settings = {
  executablePath: "this/is/the/executable/path",
  metadataServiceToken: "MetadataServiceToken",
  registrationKey: "RegistrationKey"
}

export const Primary: Story = {
  args: {
    ...settings,
    onSubmit: (settings: Settings) => alert(`Hello ${settings.executablePath}, ${settings.metadataServiceToken}, ${settings.registrationKey}`)
  }
}
