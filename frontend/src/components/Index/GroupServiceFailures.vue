<template>
  <div>
    <div class="d-flex mt-2">
      <div
        v-for="(d, index) in failureData"
        :key="index"
        class="flex-fill service_day"
        :class="{'mini_error': d.amount > 0, 'mini_success': d.amount === 0}"
      >
        <!-- <span
          v-if="d.amount != 0"
          class="small"
        >
          {{ d.amount }}
        </span> -->
      </div>
    </div>
    <div class="row">
      <div class="col-6 text-left service_note text-muted">
        30 Days Ago
      </div>
      <div class="col-6 text-right service_note text-muted">
        Today
      </div>
    </div>
    <div class="row mt-1">
      <div
        class="col-8 text-left service_note_2"
        :class="{'text-muted': service.online, 'text-danger': !service.online}"
      >
        {{ service_txt }}
      </div>
    </div>
    <div class="daily-failures small text-right text-dim">
      {{ hover_text }}
    </div>
  </div>
</template>

<script>
import Api from '../../API';

export default {
    name: 'GroupServiceFailures',
    components: {

    },
    props: {
        service: {
            type: Object,
            required: true
        }
    },
    data () {
        return {
            failureData: [],
            hover_text: '',
            loaded: false,
            visible: false,
        };
    },
    computed: {
        service_txt () {
            return this.smallText(this.service);
        }
    },
    mounted () {
        this.lastDaysFailures();
    },
    methods: {
        visibleChart (isVisible, entry) {
            if (isVisible && !this.visible) {
                this.visible = true;
                this.lastDaysFailures().then(() =>  this.loaded = true);
            }
        },
        mouseout () {
            this.hover_text = '';
        },
        mouseover (e) {
            let txt = `${e.amount} Failures`;
            if (e.amount === 0) {
                txt = `No Issues`;
            }
            this.hover_text = `${e.date.toLocaleDateString()} - ${txt}`;
        },
        async lastDaysFailures () {
            const start = this.nowSubtract(86400 * 30);
            const data = await Api.service_failures_data(this.service.id, this.toUnix(start), this.toUnix(this.startToday()), '24h');
            data.forEach((d) => {
                const date = this.parseISO(d.timeframe);
                this.failureData.push({ month: 1, day: date.getDate(), amount: d.amount });
            });
        }
    }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.service_day {
  height: 50px;
}
.service_note {
  font-size: 0.625rem;
}
.service_note_2 {
  font-size: 0.75rem;
}
.mini_success {
  background-color: #48d337;
}
.mini_error{
  background-color: #ffca46;
}
</style>
