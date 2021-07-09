<template>
  <div>
    <h1
      id="title"
      class="col-12 header-title mb-3 font-4"
    >
      Status Page
    </h1>
    <div class="col-12 mb-5 header-desc font-1">
      <p class="mb-4">
        This is Torus' service status notification page where you can get all the real-time information on health status on all the services and systems. For any service interruption, we will post it here with transparency so you can follow along.
      </p>
      <p>
        Please don't hesitate to reach out to us at our <a
          href="#"
          onclick="$crisp.push(['do', 'chat:open'])"
        >
          helpdesk
        </a> if you experience any technical issue. We will do our best to get back to you as soon as we can.
      </p>
    </div>
  </div>
</template>

<script>
export default {
    name: 'Header',
    computed: {
        core () {
            return this.$store.getters.core;
        }
    }
};
</script>
