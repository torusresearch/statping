<template>
  <div class="container index_container">
    <Header />

    <div class="col-12 full-col-12">
      <MessageBlock
        v-for="message in messages"
        :key="message.id"
        :message="message"
      />
    </div>

    <!-- <div class="col-12 full-col-12">
            <div v-for="service in services_no_group" v-bind:key="service.id" class="list-group online_list mb-4">
                <div class="list-group-item list-group-item-action">
                    <router-link class="no-decoration font-3" :to="serviceLink(service)">{{service.name}}</router-link>
                    <span class="badge float-right" :class="{'bg-success': service.online, 'bg-danger': !service.online }">{{service.online ? "ONLINE" : "OFFLINE"}}</span>
                    <GroupServiceFailures :service="service"/>
                    <IncidentsBlock :service="service"/>
                </div>
            </div>
        </div>-->

    <div>
      <Group
        v-for="group in groups"
        :key="group.id"
        :group="group"
      />
    </div>

    <!-- <div class="col-12 full-col-12">
            <div v-for="service in services" :ref="service.id" v-bind:key="service.id">
                <ServiceBlock :service="service" />
            </div>
        </div>-->

    <!-- <div>
            <Incidents :services="services" />
        </div> -->
  </div>
</template>

<script>
const Group = () => import('@/components/Index/Group');
const Header = () => import('@/components/Index/Header');
const MessageBlock = () => import('@/components/Index/MessageBlock');
const ServiceBlock = () => import('@/components/Service/ServiceBlock');
const GroupServiceFailures = () =>
    import('@/components/Index/GroupServiceFailures');
// const IncidentsBlock = () => import("@/components/Index/IncidentsBlock");
const Incidents = () => import('@/components/Index/Incidents');

export default {
    name: 'Index',
    components: {
        // IncidentsBlock,
        // GroupServiceFailures,
        // ServiceBlock,
        MessageBlock,
        Group,
        Header,
        // Incidents,
    },
    data () {
        return {
            logged_in: false,
        };
    },
    computed: {
        loading_text () {
            if (this.$store.getters.groups.length === 0) {
                return 'Loading Groups';
            } else if (this.$store.getters.services.length === 0) {
                return 'Loading Services';
            } else if (this.$store.getters.messages == null) {
                return 'Loading Announcements';
            }

            return '';
        },
        loaded () {
            return this.$store.getters.services.length !== 0;
        },
        core () {
            return this.$store.getters.core;
        },
        messages () {
            return this.$store.getters.messages.filter(
                (m) => this.inRange(m) && m.service === 0
            );
        },
        groups () {
            return this.$store.getters.groupsInOrder;
        },
        // services() {
        //     return this.$store.getters.servicesInOrder;
        // },
        // services_no_group() {
        //     return this.$store.getters.servicesNoGroup;
        // },
    },
    async mounted () {},
    methods: {
        async checkLogin () {
            const token = this.$cookies.get('statping_auth');
            if (!token) {
                this.$store.commit('setLoggedIn', false);
                return;
            }
            try {
                const jwt = await Api.check_token(token);
                this.$store.commit('setAdmin', jwt.admin);
                if (jwt.username) {
                    this.$store.commit('setLoggedIn', true);
                }
            } catch (e) {
                console.error(e);
            }
        },
        inRange (message) {
            return this.isBetween(
                this.now(),
                message.start_on,
                message.start_on === message.end_on
                    ? this.maxDate().toISOString()
                    : message.end_on
            );
        },
    },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.container {
    box-shadow: none !important;
    font-family: Roboto,-apple-system,BlinkMacSystemFont,"Segoe UI","Helvetica Neue",Arial,"Noto Sans",sans-serif,"Apple Color Emoji","Segoe UI Emoji","Segoe UI Symbol","Noto Color Emoji";
}
.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
    opacity: 0;
}
</style>
