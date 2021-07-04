<template>
  <div class="col-12">
    <div class="row">
      <div class="col-md-3 col-sm-12 mb-4 mb-md-0">
        <div
          id="v-pills-tab"
          class="nav flex-column nav-pills"
          role="tablist"
          aria-orientation="vertical"
        >
          <div
            v-if="version_below"
            class="col-12 small text-center mt-0 pt-0 pb-0 mb-3"
          >
            Update {{ github.tag_name }} Available
            <div class="row">
              <div class="col-6">
                <a
                  href="https://github.com/statping/statping/releases/latest"
                  class="btn btn-sm text-success mt-2"
                >
                  Download
                </a>
              </div>
              <div class="col-6">
                <a
                  href="https://github.com/statping/statping/blob/master/CHANGELOG.md"
                  class="btn btn-sm text-dim mt-2"
                >
                  Changelog
                </a>
              </div>
            </div>
          </div>

          <h6 class="text-muted">
            {{ $t('main_settings') }}
          </h6>

          <a
            id="v-pills-home-tab"
            class="nav-link"
            :class="{active: liClass('v-pills-home-tab')}"
            data-toggle="pill"
            href="#v-pills-home"
            role="tab"
            aria-controls="v-pills-home"
            aria-selected="true"
            @click.prevent="changeTab"
          >
            <font-awesome-icon
              icon="cog"
              class="mr-2"
            /> {{ $t('settings') }}
          </a>
          <a
            id="v-pills-style-tab"
            class="nav-link"
            :class="{active: liClass('v-pills-style-tab')}"
            data-toggle="pill"
            href="#v-pills-style"
            role="tab"
            aria-controls="v-pills-style"
            aria-selected="false"
            @click.prevent="changeTab"
          >
            <font-awesome-icon
              icon="image"
              class="mr-2"
            /> {{ $t('theme') }}
          </a>
          <a
            id="v-pills-oauth-tab"
            class="nav-link"
            :class="{active: liClass('v-pills-oauth-tab')}"
            data-toggle="pill"
            href="#v-pills-oauth"
            role="tab"
            aria-controls="v-pills-oauth"
            aria-selected="false"
            @click.prevent="changeTab"
          >
            <font-awesome-icon
              icon="key"
              class="mr-2"
            /> {{ $t('authentication') }}
          </a>
          <a
            id="v-pills-import-tab"
            class="nav-link"
            :class="{active: liClass('v-pills-import-tab')}"
            data-toggle="pill"
            href="#v-pills-import"
            role="tab"
            aria-controls="v-pills-import"
            aria-selected="false"
            @click.prevent="changeTab"
          >
            <font-awesome-icon
              icon="cloud-download-alt"
              class="mr-2"
            /> {{ $t('import') }}
          </a>
          <a
            id="v-pills-configs-tab"
            class="nav-link"
            :class="{active: liClass('v-pills-configs-tab')}"
            data-toggle="pill"
            href="#v-pills-configs"
            role="tab"
            aria-controls="v-pills-configs"
            aria-selected="false"
            @click.prevent="changeTab"
          >
            <font-awesome-icon
              icon="cogs"
              class="mr-2"
            /> {{ $t('configs') }}
          </a>

          <h6 class="mt-4 text-muted">
            {{ $t('notifiers') }}
          </h6>

          <div id="notifiers_tabs">
            <a
              v-for="(notifier, index) in notifiers"
              :id="`v-pills-${notifier.method.toLowerCase()}-tab`"
              :key="`${notifier.method}`"
              class="nav-link text-capitalize"
              :class="{active: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}"
              data-toggle="pill"
              :href="`#v-pills-${notifier.method.toLowerCase()}`"
              role="tab"
              :aria-controls="`v-pills-${notifier.method.toLowerCase()}`"
              aria-selected="false"
              @click.prevent="changeTab"
            >
              <font-awesome-icon
                :icon="iconName(notifier.icon)"
                class="mr-2"
              /> {{ notifier.title }}
              <span
                v-if="notifier.enabled"
                class="badge badge-pill float-right mt-1"
                :class="{'badge-success': !liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), 'badge-light': liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), 'text-dark': liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}"
              >
                ON
              </span>
            </a>
            <a
              :id="`v-pills-notifier-docs-tab`"
              class="nav-link text-capitalize"
              :class="{active: liClass(`v-pills-notifier-docs-tab`)}"
              data-toggle="pill"
              :href="`#v-pills-notifier-docs`"
              role="tab"
              :aria-controls="`v-pills-notifier-docs`"
              aria-selected="false"
              @click.prevent="changeTab"
            >
              <font-awesome-icon
                icon="question"
                class="mr-2"
              /> {{ $t('variables') }}
            </a>
          </div>

          <h6 class="mt-4 mb-3 text-muted">
            Statping {{ $t('links') }}
          </h6>

          <a
            href="https://github.com/torusresearch/statping/wiki"
            class="mb-2 font-2 text-decoration-none text-muted"
          >
            <font-awesome-icon
              icon="question"
              class="mr-3"
            /> {{ $t('settings.docs') }}
          </a>

          <a
            href="https://github.com/torusresearch/statping/wiki/API"
            class="mb-2 font-2 text-decoration-none text-muted"
          >
            <font-awesome-icon
              icon="laptop"
              class="mr-2"
            /> API {{ $t('settings.docs') }}
          </a>

          <a
            href="https://raw.githubusercontent.com/statping/statping/master/CHANGELOG.md"
            class="mb-2 font-2 text-decoration-none text-muted"
          >
            <font-awesome-icon
              icon="book"
              class="mr-3"
            /> {{ $t('changelog') }}
          </a>

          <a
            href="https://github.com/torusresearch/statping"
            class="mb-2 font-2 text-decoration-none text-muted"
          >
            <font-awesome-icon
              icon="code-branch"
              class="mr-3"
            /> {{ $t('settings.repo') }}
          </a>
        </div>
      </div>
      <div class="col-md-9 col-sm-12">
        <div
          id="v-pills-tabContent"
          class="tab-content"
        >
          <div
            id="v-pills-home"
            class="tab-pane fade"
            :class="{active: liClass('v-pills-home-tab'), show: liClass('v-pills-home-tab')}"
            role="tabpanel"
            aria-labelledby="v-pills-home-tab"
          >
            <CoreSettings />

            <div class="card mt-3">
              <div class="card-header">
                API {{ $t('settings') }}
              </div>
              <div class="card-body">
                <div class="form-group row">
                  <label class="col-sm-3 col-form-label">
                    API {{ $t('secret') }}
                  </label>
                  <div class="col-sm-9">
                    <div class="input-group">
                      <input
                        id="api_secret"
                        v-model="core.api_secret"
                        type="text"
                        class="form-control select-input"
                        readonly
                        @focus="$event.target.select()"
                      >
                      <div class="input-group-append copy-btn">
                        <button
                          class="btn btn-outline-secondary"
                          type="button"
                          @click="copy(core.api_secret)"
                        >
                          {{ $t('copy') }}
                        </button>
                      </div>
                    </div>
                    <small class="form-text text-muted">
                      {{ $t('regen_desc') }}
                    </small>
                  </div>
                </div>
              </div>
              <div class="card-footer">
                <button
                  id="regenkeys"
                  class="btn btn-sm btn-danger float-right"
                  @click="renewApiKeys"
                >
                  {{ $t('regen_api') }}
                </button>
              </div>
            </div>
          </div>

          <div
            id="v-pills-style"
            class="tab-pane fade"
            :class="{active: liClass('v-pills-style-tab'), show: liClass('v-pills-style-tab')}"
            role="tabpanel"
            aria-labelledby="v-pills-style-tab"
          >
            <ThemeEditor />
          </div>

          <div
            id="v-pills-oauth"
            class="tab-pane fade"
            :class="{active: liClass('v-pills-oauth-tab'), show: liClass('v-pills-oauth-tab')}"
            role="tabpanel"
            aria-labelledby="v-pills-oauth-tab"
          >
            <OAuth />
          </div>

          <div
            id="v-pills-configs"
            class="tab-pane fade"
            :class="{active: liClass('v-pills-configs-tab'), show: liClass('v-pills-configs-tab')}"
            role="tabpanel"
            aria-labelledby="v-pills-configs-tab"
          >
            <Configs />
          </div>

          <div
            id="v-pills-import"
            class="tab-pane fade"
            :class="{active: liClass('v-pills-import-tab'), show: liClass('v-pills-import-tab')}"
            role="tabpanel"
            aria-labelledby="v-pills-import-tab"
          >
            <Importer />
          </div>

          <div
            :id="`v-pills-notifier-docs-tab`"
            class="tab-pane fade"
            :class="{active: liClass(`v-pills-notifier-docs-tab`), show: liClass(`v-pills-notifier-docs-tab`)}"
            role="tabpanel"
            :aria-labelledby="`v-pills-notifier-docs-tab`"
          >
            <Variables />
          </div>

          <div
            v-for="(notifier, index) in notifiers"
            :id="`v-pills-${notifier.method.toLowerCase()}-tab`"
            :key="`${notifier.method}_${index}`"
            class="tab-pane fade"
            :class="{active: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`), show: liClass(`v-pills-${notifier.method.toLowerCase()}-tab`)}"
            role="tabpanel"
            :aria-labelledby="`v-pills-${notifier.method.toLowerCase()}-tab`"
          >
            <Notifier :notifier="notifier" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Api from '../API';
const semver = require('semver');

const CoreSettings = () => import(/* webpackChunkName: "dashboard" */ '@/forms/CoreSettings');
// const FormIntegration = () => import(/* webpackChunkName: "dashboard" */ '@/forms/Integration');
const Notifier = () => import(/* webpackChunkName: "dashboard" */ '@/forms/Notifier');
const OAuth = () => import(/* webpackChunkName: "dashboard" */ '@/forms/OAuth');
const ThemeEditor = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/ThemeEditor');
const Importer = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/Importer');
const Variables = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/Variables');
const Configs = () => import(/* webpackChunkName: "dashboard" */ '@/components/Dashboard/Configs');

export default {
    name: 'Settings',
    components: {
        Configs,
        Importer,
        Variables,
        OAuth,
        ThemeEditor,
        // FormIntegration,
        Notifier,
        CoreSettings
    },
    data () {
        return {
            tab: 'v-pills-home-tab',
            github: null,
        };
    },
    computed: {
        core () {
            return this.$store.getters.core;
        },
        notifiers () {
            return this.$store.getters.notifiers;
        },
        version_below () {
            if (!this.github || !this.core.version) {
                return false;
            }
            return semver.gt(semver.coerce(this.github.tag_name), semver.coerce(this.core.version));
        }
    },
    mounted () {

    },
    created () {
        this.update();
    },
    methods: {
        async update () {
            await this.getGithub();
        },
        async getGithub () {
            try {
                this.github = await Api.github_release();
            } catch (e) {
                console.error(e);
            }
        },
        changeTab (e) {
            this.tab = e.target.id;
        },
        liClass (id) {
            return this.tab === id;
        },
        async renew () {
            await Api.renewApiKeys();
            const core = await Api.core();
            this.$store.commit('setCore', core);
            this.core = core;
            await this.logout();
        },
        async renewApiKeys () {
            const modal = {
                visible: true,
                title: 'Reset API Key',
                body: `Are you sure you want to reset the API keys? You will be logged out.`,
                btnColor: 'btn-danger',
                btnText: 'Reset',
                func: () => this.renew(),
            };
            this.$store.commit('setModal', modal);
        },
        async logout () {
            await Api.logout();
            this.$store.commit('setHasAllData', false);
            this.$store.commit('setToken', null);
            this.$store.commit('setAdmin', false);
            this.$store.commit('setUser', false);
            // this.$cookies.remove("statping_auth")
            await this.$router.push('/logout');
        }
    }
};
</script>
