import type { Meta, StoryObj } from "@storybook/vue3-vite";
import NewMovie from "./NewMovie.vue";

const meta: Meta<typeof NewMovie> = {
  component: NewMovie
}

export default meta;
type Story = StoryObj<typeof NewMovie>;

export const Primary: Story = {}
