export default {
  mode: "spa",
  /*
   ** Headers of the page
   */
  head: {
    titleTemplate: "%s - Catalog APP",
    title: process.env.npm_package_name || "",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      {
        hid: "description",
        name: "description",
        content: process.env.npm_package_description || ""
      }
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }]
  },
  /*
   ** Customize the progress-bar color
   */
  loading: { color: "#fff" },
  /*
   ** Global CSS
   */
  css: [],
  /*
   ** Plugins to load before mounting the App
   */
  plugins: [],
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [
    /**
     * Doc: https://github.com/nuxt-community/nuxt-tailwindcss
     */
    "@nuxtjs/tailwindcss"
  ],
  /**
   * Doc: https://github.com/Developmint/nuxt-purgecss#options
   */
  purgeCSS: {
    extractors: () => [
      {
        extractor(content) {
          // Custom extractor supporting additions for TailwindUI.
          return content.match(/[\w-/.:]+(?<!:)/g);
        },
        extensions: ["html", "vue", "js"]
      }
    ]
  },
  /*
   ** Nuxt.js modules
   */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    "@nuxtjs/axios",
    // Doc: https://github.com/nuxt-community/dotenv-module
    "@nuxtjs/dotenv"
  ],
  /*
   ** Axios module configuration
   ** See https://axios.nuxtjs.org/options
   */
  axios: {},
  /*
   ** Build configuration
   */
  build: {
    /*
     ** You can extend webpack config here
     */
    extend(config, ctx) {}
  },
  env: {
    BASE_URL: process.env.BASE_URL,
    API_URL: process.env.API_URL
  }
};
