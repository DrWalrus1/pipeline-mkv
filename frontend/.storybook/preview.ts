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
  },
  decorators: [
    (_, context) => {
      switch (context.parameters.type) {
        case 'component':
          return { template: '<div class=\'storybook-center-preview\'><story /></div>' }
        default:
          return { template: '<story />' }
      }
    }
  ]
};

export default preview;
