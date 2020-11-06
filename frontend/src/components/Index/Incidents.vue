<template>
    <div class="col-12 full-col-12">
        <h4 class="group_header mb-3 mt-4 font-4">Past Incidents</h4>
        <div class="row">
            <div class="col-12 mb-4" v-for="(incidentDate, index) in incidentDates" :key="index">
                <div class="torus-incident">
                    <div
                        class="torus-incident_date torus-font-5 font-weight-bold"
                    >{{ incidentDate }}</div>
                    <div class="dropdown-divider"></div>
                    <IncidentsBlock v-if="incidents[incidentDate]" :incidents-view="true" :incident-list="incidents[incidentDate]" />
                    <div v-else class="torus-incident_info torus-font-4">No incidents repoorted today</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import Api from "../../API";
import IncidentsBlock from "./IncidentsBlock";
import { subDays } from 'date-fns';

export default {
    name: "Incidents",
    components: { IncidentsBlock },
    props: {
        services: {
            type: Array,
            default: [],
        },
    },
    data: () => ({
        incidents: [],
        incidentDates: [],
    }),
    methods: {
        async getIncidents(services) {
            const promises = [];
            const incidents = await Promise.all(
                services.map(
                    async (service) => await Api.incidents_service(service.id)
                )
            ).then((incidents) => {
                return [].concat(...incidents);
            });

            const finalIncidents = incidents.reduce((accumulator, incident) => {
                const dateKey = this.dateFormat(incident.updated_at);
                if (!accumulator[dateKey]) accumulator[dateKey] = [];
                accumulator[dateKey].push(incident);
                return accumulator;
            }, []);
            this.incidents = finalIncidents;

            const currentDate = new Date()
            this.incidentDates = []
            for(let step = 0; step < 5; step++) {
                this.incidentDates.push(this.format(currentDate, "yyyy-MM-dd"))
                currentDate.setDate(currentDate.getDate() - 1)
            }
        },
        dateFormat(date) {
            return this.format(this.parseISO(date), "yyyy-MM-dd");
        },
    },
    watch: {
        services(newValue) {
            this.getIncidents(newValue);
        },
    },
    mounted() {
        this.getIncidents(this.services);
    },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.torus-incident_date {
    color: #0f1222;
}
.torus-incident_info {
    color: #757575;
}
</style>
