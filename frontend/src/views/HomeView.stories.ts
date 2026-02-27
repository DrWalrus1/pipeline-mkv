import type { Meta, StoryObj } from "@storybook/vue3-vite";
import HomeView from "./HomeView.vue";
import { defineComponent } from "vue";
import type { MakeMkvService } from "@/services/gomakemkv/gomakemkv";

class FakeMakeMkvService implements MakeMkvService {
  public GetDiscInfo(): Promise<string> {
    return new Promise<string>((resolve) => {
      resolve("Hallo")
    })
  }

}
let fakeInstance = new FakeMakeMkvService()

const meta: Meta<typeof HomeView> = {
  component: HomeView,
  decorators: [
    () => defineComponent({
      template: '<story />',
      provide: {
        MakeMkvService: fakeInstance
      }
    })
  ]
}

export default meta;
type Story = StoryObj<typeof HomeView>;

export const Primary: Story = {}
