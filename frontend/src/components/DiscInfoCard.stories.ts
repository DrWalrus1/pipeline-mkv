import type { Meta, StoryObj } from "@storybook/vue3";
import DiscInfoCard from "./DiscInfoCard.vue";
import type { DiscInfo } from "@/domain/disc_info";

let x: DiscInfo = {
  name: "Demon Slayer",
  language: "English",
  type: "Blu-Ray Disc",
  titles: []
}

const meta: Meta<typeof DiscInfoCard> = {
  component: DiscInfoCard,
}

export default meta;
type Story = StoryObj<typeof DiscInfoCard>;

export const Primary: Story = {
  args: {
    discInfo: x,
    url: "https://m.media-amazon.com/images/I/91+ShpVWyiL._AC_UF894,1000_QL80_.jpg"
  }
}
