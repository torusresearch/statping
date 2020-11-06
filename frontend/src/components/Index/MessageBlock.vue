<template>
    <div class="alert pb-4 pt-3 mt-5 mb-5" role="alert">
        <div class="d-flex">
            <h3 class="mb-3 torus-font-4">{{message.title}}</h3>
            <span class="ml-auto alert_created torus-font-1">{{ niceDateTorus(message.created_at) }}</span>
        </div>
        <span class="mb-3 torus-font-2">{{message.description}}</span>
        <div class="mt-3 alert_dates">
            <div class="torus-font-2">Start: {{niceDateTorus(message.start_on)}}</div>
            <div class="torus-font-2">End: {{niceDateTorus(message.end_on)}}</div>
        </div>
    </div>
</template>

<script>
export default {
    name: "MessageBlock",
    props: {
        message: {
            type: Object,
        },
    },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.alert {
    background: #ffffff;
    border: 1px solid #d3d5e2;
    box-sizing: border-box;
    box-shadow: 0px 14px 28px rgba(0, 0, 0, 0.06);
    border-radius: 6px;
    color: #0f1222;
}
.alert_dates {
    color: #757575;
}
.alert_created {
    color: #7f8fa4;
}
</style>
