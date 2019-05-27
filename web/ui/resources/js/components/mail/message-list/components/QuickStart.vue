<template>
  <div class="quickstart">
    <div class="quickstart__setting">
      <span class="quickstart__label">Host</span>
      <span class="value" id="quickstart__setting-host" v-text="mailbox.settings.host || noneConfiguredText"></span>
      <a href="#" class="clipboard" data-clipboard-target="#quickstart__setting-host"><font-awesome-icon :icon="['fas', 'clipboard']" /></a>
    </div>

    <div class="quickstart__setting">
      <span class="quickstart__label">Port</span>
      <span class="value" id="quickstart__setting-port" v-text="portsText"></span>
      <a href="#" class="clipboard" data-clipboard-target="#quickstart__setting-port"><font-awesome-icon :icon="['fas', 'clipboard']" /></a>
    </div>

    <div class="quickstart__setting">
      <span class="quickstart__label">Username</span>
      <span class="value" id="quickstart__setting-username" v-text="mailbox.settings.username"></span>
      <a href="#" class="clipboard" data-clipboard-target="#quickstart__setting-username"><font-awesome-icon :icon="['fas', 'clipboard']" /></a>
    </div>

    <div class="quickstart__setting">
      <span class="quickstart__label">Password</span>
      <span class="value" id="quickstart__setting-password" v-text="mailbox.settings.password"></span>
      <a href="#" class="clipboard" data-clipboard-target="#quickstart__setting-password"><font-awesome-icon :icon="['fas', 'clipboard']" /></a>
    </div>

    <div class="quickstart__setting">
      <span class="quickstart__label">TLS</span>
      <span class="value">Optional</span>
    </div>
  </div>
</template>

<script>
import ClipboardJS from 'clipboard';

new ClipboardJS('.clipboard'); // eslint-disable-line no-new

export default {

  data() {
    return {
      noneConfiguredText: 'None configured',
    };
  },

  props: {
    mailbox: { required: true, type: Object },
  },

  computed: {
    portsText() {
      return this.mailbox.settings.ports
        ? this.mailbox.settings.ports.join(' / ')
        : this.noneConfiguredText;
    },
  },

};
</script>

<style lang="scss">
@import '~@styles/mailmole/variables';

.quickstart {

  .quickstart__setting {
    padding: 5px 0 10px;


    a.clipboard {
      color: map-get($theme-colors, "primary") !important;
      padding-left: 4px;
      opacity: 0;
      transition: opacity .25s ease-in-out;
      -moz-transition: opacity .25s ease-in-out;
      -webkit-transition: opacity .25s ease-in-out;
    }

    &:hover {
      a.clipboard {
        opacity: 0.25;

        &:hover {
          opacity: 1;
        }
      }
    }
  }

  .quickstart__label {
    font-weight: 700;
    &::after {
      content: ":";
    }
  }
}
</style>
