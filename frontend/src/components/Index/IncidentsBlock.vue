<template>
    <div class="row">
        <div v-for="incident in incidents" class="col-12" :key="incident.id">
            <span v-if="!incidentsView" class="braker mt-1 mb-3"></span>
            <div class="torus-font-4">
                <div class="font-weight-bold">Incident: {{incident.title}}</div>
                <span class="font-2 float-right">{{niceDate(incident.created_at)}}</span>
            </div>
            <span class="torus-incident_info torus-font-4" v-html="incident.description"></span>

            <UpdatesBlock :incident="incident" />
        </div>
    </div>
</template>

<script>
import Api from "../../API";
import UpdatesBlock from "@/components/Index/UpdatesBlock";

export default {
    name: "IncidentsBlock",
    components: { UpdatesBlock },
    props: {
        service: {
            type: Object,
            required: false,
        },
        incidentsView: {
            type: Boolean,
            required: false,
        },
        incidentList: {
            type: Array,
            required: false,
        },
    },
    data() {
        return {
            incidents: null,
        };
    },
    mounted() {
        if (this.incidentsView) {
            this.incidents = this.incidentList;
        } else {
            this.getIncidents();
        }
    },
    methods: {
        badgeClass(val) {
            switch (val.toLowerCase()) {
                case "resolved":
                    return "badge-success";
                case "update":
                    return "badge-info";
                case "investigating":
                    return "badge-danger";
            }
        },
        async getIncidents() {
            this.incidents = await Api.incidents_service(this.service.id);
        },
        async incident_updates(incident) {
            await Api.incident_updates(incident).then((d) => {
                return d;
            });
            return o;
        },
    },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.torus-incident_info {
    color: #757575;
}
</style>
