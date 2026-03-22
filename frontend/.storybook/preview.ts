import type { Preview } from '@storybook/vue3-vite'
import '../src/assets/style.css'
import './storybook.css'

const preview: Preview = {
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },

    a11y: {
      // 'todo' - show a11y violations in the test UI only
      // 'error' - fail CI on a11y violations
      // 'off' - skip a11y checks entirely
      test: 'todo',
    },
  },
  decorators: [
    (_, context) => {
      switch (context.parameters.type) {
        case 'component':
          return { template: "<div class='storybook-center-preview'><story /></div>" }
        default:
          return { template: '<story />' }
      }
    },
  ],
}

export default preview
