import type { Preview } from '@storybook/vue3'
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
    (story) => ({
      components: { story },
      template: '<div class=\'storybook-center-preview\'><story /></div>'
    })
  ]
};

export default preview;
