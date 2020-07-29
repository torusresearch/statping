<template>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <a class="navbar-brand" href="#">
            <img
                src="https://app.tor.us/v1.8.1/img/torus-logo-blue.7d2a4a6b.svg"
                height="32"
                width="115"
                class="d-inline-block align-top"
                alt
                loading="lazy"
            />
        </a>
        <span class="brand-status">Status</span>

        <button
            class="navbar-toggler"
            type="button"
            aria-controls="navbarText"
            aria-expanded="false"
            aria-label="Toggle navigation"
            @click="showMenu = !showMenu"
        >
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="navbar-collapse" :class="{collapse: !showMenu}" id="navbarText">
            <ul class="navbar-nav ml-auto">
                <li class="nav-item active">
                    <a href="https://docs.tor.us" class="nav-link" target="_blank">Documentation</a>
                </li>
                <li class="nav-item active">
                    <a class="nav-link" href="https://t.me/TorusLabs" target="_blank">Telegram</a>
                </li>
                <li class="nav-item active">
                    <a class="nav-link" href="https://twitter.com/TorusLabs" target="_blank">Twitter</a>
                </li>
            </ul>
        </div>
    </nav>
</template>

<script>
export default {
    name: "Navbar",
    data: () => ({
        showMenu: false,
    }),
};
</script>

<style scoped>
.navbar {
    background: #ffffff;
    border: 1px solid #f5f5f5;
    box-sizing: border-box;
    box-shadow: 0px 14px 28px rgba(46, 91, 255, 0.06);
    border-radius: 6px;
}
.nav-item a {
    font-style: normal;
    font-weight: bold;
    font-size: 12px;
    line-height: 14px;
    color: #757575;
}
.navbar-brand {
    margin-right: 8px;
}
.brand-status {
    font-style: normal;
    font-weight: 300;
    font-size: 18px;
    line-height: 28px;
    color: rgba(15, 18, 34, 0.6);
    margin-bottom: -15px;
}
.navbar-toggler {
    border: 0;
}
</style>