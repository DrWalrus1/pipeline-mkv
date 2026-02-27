import type { Meta, StoryObj } from "@storybook/vue3-vite";
import HomeView from "./HomeView.vue";
import { defineComponent } from "vue";
import type { MakeMkvService } from "@/services/gomakemkv/gomakemkv";
import type { DiscInfo } from "@/domain/disc_info";

class FakeMakeMkvService implements MakeMkvService {
  public GetInsertedDiscInfo(): Promise<string> {
    return new Promise<string>((resolve) => {
      resolve("Hallo")
    })
  }

  public GetRecentDiscInfos(limit: number, offset: number): Promise<DiscInfo[]> {
    const fakeDiscHistoryList: DiscInfo[] = [
      {
        name: "Movie 1",
        language: "English",
        type: "Blu-ray",
        titles: []
      },
      {
        name: "Movie 2",
        language: "English",
        type: "DVD",
        titles: []
      },
      {
        name: "Movie 3",
        language: "Japanese",
        type: "Blu-ray",
        titles: []
      }
    ]
    return new Promise<DiscInfo[]>((resolve) => {
      resolve(fakeDiscHistoryList.slice(0, limit))


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
