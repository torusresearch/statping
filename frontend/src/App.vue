<template>
  <div id="app">
    <Navbar v-if="isIndexPage" />
    <router-view :loaded="loaded"/>
      <Footer v-if="$route.path !== '/setup'"/>
  </div>
</template>

<script>
  import Footer from "./components/Index/Footer";
  import Navbar from "./components/Index/Navbar";

  export default {
  name: 'app',
  components: {
    Footer,
    Navbar
  },
  data () {
    return {
      loaded: false,
      version: ""
    }
  },
      computed: {
          core() {
            return this.$store.getters.core
          },
          isIndexPage() {
            return this.$route.name === 'Index'
          }
      },
  async beforeMount() {
    await this.$store.dispatch('loadCore')

    this.$i18n.locale = this.core.language || "en";
    // this.$i18n.locale = "ru";

      if (!this.core.setup) {
        this.$router.push('/setup')
      }
    if (this.$route.path !== '/setup') {
      if (this.$store.state.admin) {
        await this.$store.dispatch('loadAdmin')
      } else {
        await this.$store.dispatch('loadRequired')
      }
      this.loaded = true
    }
  },
  async mounted() {
    if (this.$route.path !== '/setup') {
      if (this.$store.state.admin) {
          this.logged_in = true
            // await this.$store.dispatch('loadAdmin')
        }
    }
  },
}
</script>

<style lang="scss">
    @import "./assets/css/bootstrap.min.css";
    @import "./assets/scss/main";
</style>
