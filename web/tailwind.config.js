/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,ts}'],
  theme: {
    extend: {
      colors: {
        ink: {
          950: '#08080a',
          900: '#0c0c0f',
          800: '#131317',
          700: '#1a1a1f',
          600: '#26262c',
        },
        brand: {
          50: '#fff4ed',
          100: '#ffe6d5',
          200: '#ffc9a8',
          300: '#ffa270',
          400: '#ff7a3d',
          500: '#ff5a1f',
          600: '#e8480f',
          700: '#c23708',
          800: '#9a2b06',
          900: '#7a2205',
        },
        accent: {
          300: '#ffe08a',
          400: '#ffc857',
          500: '#f5a623',
          600: '#c77c00',
        },
      },
      fontFamily: {
        sans: [
          'Inter',
          'ui-sans-serif',
          'system-ui',
          '-apple-system',
          'Segoe UI',
          'PingFang SC',
          'Hiragino Sans GB',
          'Microsoft YaHei',
          'sans-serif',
        ],
        mono: ['ui-monospace', 'SFMono-Regular', 'Menlo', 'Consolas', 'monospace'],
      },
      boxShadow: {
        card: '0 1px 2px rgba(10, 10, 12, 0.04), 0 16px 40px -24px rgba(10, 10, 12, 0.16)',
      },
      maxWidth: {
        content: '1200px',
      },
    },
  },
  plugins: [],
}
