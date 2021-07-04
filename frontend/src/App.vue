<template>
  <div id="app">
    <Navbar v-if="isIndexPage" />
    <router-view
      v-if="loaded"
      :loaded="loaded"
    />
    <div
      v-else
      class="loader"
    >
      <template>
        <VueSimpleSpinner size="big" />
      </template>
    </div>
    <!-- <Footer v-if="$route.path !== '/setup'"/> -->
  </div>
</template>

<script>
// import Footer from "./components/Index/Footer";
import Navbar from './components/Index/Navbar';
import VueSimpleSpinner from 'vue-simple-spinner';

export default {
    name: 'App',
    components: {
    // Footer,
        Navbar,
        VueSimpleSpinner
    },
    data () {
        return {
            loaded: false,
            version: ''
        };
    },
    computed: {
        core () {
            return this.$store.getters.core;
        },
        isIndexPage () {
            return this.$route.name === 'Index';
        }
    },
    async beforeMount () {
        await this.$store.dispatch('loadCore');

        this.$i18n.locale = this.core.language || 'en';
        // this.$i18n.locale = "ru";

        if (!this.core.setup) {
            this.$router.push('/setup');
        }
        if (this.$route.path !== '/setup') {
            if (this.$store.state.admin) {
                await this.$store.dispatch('loadAdmin');
            } else {
                await this.$store.dispatch('loadRequired');
            }
            this.loaded = true;
        }
        this.loaded = true;
    },
    async mounted () {
        if (this.$route.path !== '/setup') {
            if (this.$store.state.admin) {
                this.logged_in = true;
            // await this.$store.dispatch('loadAdmin')
            }
        }
    },
};
</script>

<style lang="scss">
    @import "./assets/css/bootstrap.min.css";
    @import "./assets/scss/index";
</style>

<style scoped>

.loader {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color:#fcfcfc;
}

</style>
