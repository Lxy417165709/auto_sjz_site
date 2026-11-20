/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,ts}'],
  theme: {
    extend: {
      colors: {
        ink: {
          900: '#2d1a11',
          800: '#3a2116',
          700: '#4a2a1b',
          600: '#5a3524',
        },
        brand: {
          50: '#fff5ef',
          100: '#ffe7d8',
          200: '#ffcdb3',
          300: '#ffa980',
          400: '#ff8352',
          500: '#e85c2c',
          600: '#cf451f',
          700: '#a93318',
          800: '#862915',
          900: '#6b2414',
        },
        accent: {
          300: '#ffe08a',
          400: '#ffd15c',
          500: '#f7b52c',
          600: '#a96f08',
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
        card: '0 1px 2px rgba(74, 31, 15, 0.04), 0 12px 32px -16px rgba(74, 31, 15, 0.16)',
      },
      maxWidth: {
        content: '1180px',
      },
    },
  },
  plugins: [],
}
