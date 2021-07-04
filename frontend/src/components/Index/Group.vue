<template>
  <div class="col-12 full-col-12">
    <h4
      v-if="group.name !== 'Empty Group'"
      class="group_header mb-3 mt-4 font-4"
    >
      {{ group.name }}
    </h4>
    <div class="row">
      <div
        v-for="(service, index) in $store.getters.servicesInGroup(group.id)"
        :key="index"
        class="col-sm-12 col-lg-4 mb-4"
      >
        <div class="torus-service p-4">
          <div class="d-flex">
            <router-link
              class="no-decoration font-1 font-weight-bold"
              :to="serviceLink(service)"
            >
              {{ service.name }}
            </router-link>
            <span
              class="badge py-1 px-3 ml-auto text-capitalize float-right"
              :class="{'bg-success': service.online, 'bg-danger': !service.online }"
            >
              {{ service.online ? $t('online') : $t('offline') }}
            </span>
          </div>
          <GroupServiceFailures :service="service" />
          <!-- <IncidentsBlock :service="service" /> -->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Api from '../../API';
import GroupServiceFailures from './GroupServiceFailures';
// import IncidentsBlock from "./IncidentsBlock";

export default {
    name: 'Group',
    components: {
        GroupServiceFailures,
        // IncidentsBlock,
    },
    props: {
        group: Object,
    },
};
</script>

<style scoped>
.torus-service {
    background: #ffffff;
    border: 1px solid #f2f2f2;
    box-sizing: border-box;
    box-shadow: 0px 14px 28px rgba(46, 91, 255, 0.06);
    border-radius: 6px;
}
.badge {
    background: #2dcc70;
    border-radius: 3px;
}
</style>
