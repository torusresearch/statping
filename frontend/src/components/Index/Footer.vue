<template>
  <footer>
    <div
      v-if="!core.footer"
      class="footer text-center mb-4 p-2"
    >
      <a
        href="https://github.com/torusresearch/statping"
        target="_blank"
      >
        Statping {{ core.version }} made with <FontAwesomeIcon
          icon="heart"
          class="text-danger"
        />
      </a> |
      <router-link :to="$store.state.admin ? '/dashboard' : '/login'">
        {{ $t('top_nav.dashboard') }}
      </router-link>
    </div>
    <div
      v-else
      class="footer text-center mb-4 p-2"
      v-html="core.footer"
    />
  </footer>
</template>

<script>

export default {
    name: 'Footer',
    computed: {
        core () {
            return this.$store.getters.core;
        },
        commit () {
            return this.$store.getters.core.commit.slice(0,8);
        },
        admin () {
            return this.$store.getters.admin;
        },
    }
};
</script>

<style scoped>
.hlight {
  color: #f6cbcb;
}
</style>
