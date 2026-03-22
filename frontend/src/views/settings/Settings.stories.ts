import type { Meta, StoryObj } from "@storybook/vue3-vite";
import Settings from "./Settings.vue";

const meta: Meta<typeof Settings> = {
  component: Settings
}

export default meta;
type Story = StoryObj<typeof Settings>;

export const Primary: Story = {
  args: {
    executablePath: "this/is/the/executable/path"
  }
}
