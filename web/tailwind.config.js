/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,ts}'],
  theme: {
    extend: {
      colors: {
        ink: {
          900: '#070d18',
          800: '#0b1220',
          700: '#111c31',
          600: '#1a2740',
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
        card: '0 1px 2px rgba(15, 23, 42, 0.04), 0 12px 32px -16px rgba(15, 23, 42, 0.18)',
      },
      maxWidth: {
        content: '1180px',
      },
    },
  },
  plugins: [],
}
