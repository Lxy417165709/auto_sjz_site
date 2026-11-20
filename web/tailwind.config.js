/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,ts}'],
  theme: {
    extend: {
      colors: {
        ink: {
          950: '#050b16',
          900: '#0b1a33',
          800: '#122647',
          700: '#1b3259',
          600: '#26436f',
        },
        brand: {
          50: '#eef3ff',
          100: '#dbe5ff',
          200: '#bccfff',
          300: '#8fb0ff',
          400: '#5c86ff',
          500: '#3763f5',
          600: '#2447d6',
          700: '#1d38a8',
          800: '#1b3185',
          900: '#1a2d6b',
        },
        accent: {
          300: '#7fe8d2',
          400: '#34d6b2',
          500: '#12b892',
          600: '#0d9477',
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
        card: '0 1px 2px rgba(15, 23, 42, 0.04), 0 16px 40px -24px rgba(15, 23, 42, 0.18)',
      },
      maxWidth: {
        content: '1200px',
      },
    },
  },
  plugins: [],
}
