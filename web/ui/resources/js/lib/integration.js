export default {
  props: {
    mailbox: { required: true, type: Object },
  },

  computed: {
    host() {
      return this.mailbox.settings.host;
    },

    port() {
      return this.mailbox.settings.ports[0];
    },

    username() {
      return this.mailbox.settings.username;
    },

    password() {
      return this.mailbox.settings.password;
    },
  },
};
