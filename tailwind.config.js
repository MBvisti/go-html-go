module.exports = {
  purge: {
    mode: 'all',
    preserveHtmlElements: false,
    content: [
      './pkg/views/pages/*.html',
      './pkg/views/layout/*.html',
      './pkg/views/static/js/**/*.js',
    ]
  },
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [],
}
