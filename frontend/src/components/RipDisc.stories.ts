import type { Meta, StoryObj } from '@storybook/vue3'
import RipDisc from './RipDisc.vue'

const meta: Meta<typeof RipDisc> = {
  component: RipDisc,
}

export default meta
type Story = StoryObj<typeof RipDisc>

export const Primary: Story = {
  args: {
    titles: [
      {
        id: 1,
        name: 'Title 1',
      },
      {
        id: 2,
        name: 'Title 2',
      },
      {
        id: 3,
        name: 'Title 3',
      },
    ],
  },
}
