import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

const vuetify = createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'dark',
    themes: {
      light: {
        colors: {
          primary: '#1565C0',
          secondary: '#42A5F5',
          accent: '#FF6F00',
          error: '#D32F2F',
          warning: '#FFA000',
          info: '#0288D1',
          success: '#388E3C',
        },
      },
      dark: {
        colors: {
          primary: '#42A5F5',
          secondary: '#1565C0',
          accent: '#FFB300',
          error: '#EF5350',
          warning: '#FFB74D',
          info: '#29B6F6',
          success: '#66BB6A',
        },
      },
    },
  },
})

export default vuetify
