<template>
  <div class="actions pb-2 pr-2">
    <div class="button-group" v-if="attachmentsNum > 0">
      <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
        <i class="far fa-paperclip"></i> {{ attachmentsNum }}
      </button>
      <div class="dropdown-menu">
        <attachment class="dropdown-item" v-for="attachment in attachments" :message="message" :attachment="attachment" :key="attachment.filename"></attachment>
      </div>
    </div>

    <div class="btn-group versions" role="group" aria-label="Basic example">
      <router-link
        tag="button"
        type="button"
        class="btn btn-secondary message-version-link py-1 px-2"
        :to="route('html')"
        :disabled="disabled('html')"
        :title="disabled('html') ? 'No Html version' : 'Html'"
        >
          <font-awesome-icon :icon="['fas', 'code']"/>
        </router-link>

      <router-link
        tag="button"
        type="button"
        class="btn btn-secondary message-version-link py-1 px-2"
        :to="route('text')"
        :disabled="disabled('text')"
        :title="disabled('text') ? 'No Text version' : 'Text'"
        >
          <font-awesome-icon :icon="['fas', 'font']"/>
        </router-link>

      <router-link
        tag="button"
        type="button"
        class="btn btn-secondary message-version-link py-1 px-2"
        :to="route('raw')"
        title="Raw"
        >
          <font-awesome-icon :icon="['fas', 'terminal']"/>
        </router-link>
    </div>

    <button class="btn btn-default btn-sm" @click="$emit('deleteMessage')">
      <i class="fas fa-trash-alt"></i>
      <span class="sr-only">Delete message</span>
    </button>
  </div>
</template>

<script>
import Attachment from './Attachment';

export default {
  components: {
    Attachment,
  },

  methods: {
    viewVersion() {
      switch (this.$router.currentRoute.query.view) {
        case 'text':
          return 'Text';

        case 'raw':
          return 'Raw';

        default:
          return 'Html';
      }
    },

    route(version) {
      const mailbox = this.$store.getters.activeMailbox;
      const route = {
        name: 'mailbox_message',
        params: {
          mailbox: mailbox.slug,
          message: this.message.id,
        },
      };

      if (version !== 'html') {
        route.query = {
          view: version,
        };
      }

      return route;
    },

    disabled(version) {
      const message = this.$store.getters.activeMessage;

      if (!message.src[version]) {
        return true;
      }

      return false;
    },
  },

  computed: {
    attachmentsNum() {
      return this.attachments.length;
    },

    attachments() {
      return this.message.attachments;
    },

    message() {
      return this.$store.getters.activeMessage;
    },
  },
};
</script>

<style lang="scss">
.actions {

  > * {
    margin-left: 4px;
  }

  .btn-group.versions {
    button {
      font-size: .875em;
    }
  }
}
</style>
