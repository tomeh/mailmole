<template>
  <div class="mailbox-settings__general">
    <dl class="">
      <dt>Host</dt>
      <dd v-text="host || noneConfiguredText"></dd>
      <dt>Port</dt>
      <dd v-text="portsText"></dd>
      <dt>Username</dt>
      <dd v-text="username"></dd>
      <dt>Password</dt>
      <dd v-text="password"></dd>
    </dl>
  </div>
</template>

<script>
import integrationMixin from '@/lib/integration';

export default {
  mixins: [integrationMixin],

  data() {
    return {
      noneConfiguredText: 'None configured',
    };
  },

  props: {
    mailbox: { required: true },
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
.mailbox-settings__general {
  dd {
    margin-bottom: 10px;
  }
}
</style>
