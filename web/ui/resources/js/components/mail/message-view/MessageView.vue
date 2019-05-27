<template>
  <article class="message-view">
    <header>
      <h1 v-text="message.subject"></h1>

      <message-meta>
        <span slot="label">From</span>
        <span slot="value">{{ from }}</span>
      </message-meta>

      <message-meta>
        <span slot="label">To</span>
        <span slot="value">
          {{ to }}
          <span v-if="extraRecipients > 0">
            <span class="extra-recipients">(+{{ extraRecipients }})</span>
            <a href="#" @click.prevent="extraRecipientsOpen = !extraRecipientsOpen">
              <i class="fas"
                :class="{ 'fa-caret-down': !extraRecipientsOpen, 'fa-caret-up': extraRecipientsOpen }"
                ></i>
            </a>
          </span>
        </span>
      </message-meta>

      <div v-if="extraRecipientsOpen">
        <message-meta v-if="message.cc">
          <span slot="label">Cc</span>
          <span slot="value">{{ cc }}</span>
        </message-meta>

        <message-meta v-if="message.bcc">
          <span slot="label">Bcc</span>
          <span slot="value">{{ bcc }}</span>
        </message-meta>
      </div>

      <div class="date-received" v-text="dateReceived"></div>
      <actions class="actions" @deleteMessage="deleteMessage"></actions>
      <alert :show="showDeleteAlert" @close="showDeleteAlert = false"><p>There was a problem deleting the message.</p></alert>
    </header>

    <div class="message-body">
      <spinner v-if="loading"></spinner>
      <iframe
        :class="{ 'loading': loading }"
        :style="{ height: iframeHeight + 'px' }"
        ref="iframe"
        frameborder="0"
        scrolling="no"
        ></iframe>
    </div>
  </article>
</template>

<script>
import moment from 'moment';
import eventBus from '@/lib/events';
import Spinner from 'vue-simple-spinner';
import Alert from '@/components/common/modal/Alert';

import Actions from './components/Actions';
import MessageMeta from './components/Meta';

/**
 * Check that the version of the message (html, text) is ok.
 */
function checkRoute(vm, to, next) {
  // Check that this check applies to the route being navigated to.
  if (to.name !== 'mailbox_message') {
    return next();
  }

  const getters = vm.$store.getters;
  const mailboxSlug = to.params.mailbox;
  const messageId = to.params.message;
  const mailbox = getters.mailboxBySlug(mailboxSlug);
  const message = getters.messageForMailboxBySlug(mailboxSlug, messageId);
  if (!message) {
    return next(`/${mailbox.slug}`);
  }

  let version = to.query.view;
  if (!version) {
    // Default to html.
    version = 'html';
  }

  // If the version exists in the message then proceed.
  if ((version in message.src) && message.src[version]) {
    return next();
  }

  // Get the first valid version.
  version = Object.keys(message.src).find(aVersion => message.src[aVersion]);

  if (!version) {
    // Go home and pray.
    return next('/');
  }

  // Construct a url for the valid version.
  return next(`/${mailbox.slug}/${message.id}?view=${version}`);
}

export default {
  props: {
    view: { required: false, type: String, default: 'html' },
  },

  data() {
    return {
      iframeHeight: 100,
      loading: true,
      showDeleteAlert: false,
      extraRecipientsOpen: false,
    };
  },

  components: {
    Actions,
    MessageMeta,
    Spinner,
    Alert,
  },

  mounted() {
    eventBus.$on('resize', () => this.resizeIframe());
    this.load();
  },

  watch: {
    src() {
      this.loading = true;
      this.load();
    },
  },

  methods: {
    load() {
      const contentWindow = this.$refs.iframe.contentWindow;
      const ifDocument = contentWindow.document;

      return axios.get(this.src).then((content) => {
        ifDocument.open();
        ifDocument.write(content.data);
        ifDocument.close();

        this.resizeIframe();
        this.loading = false;
      });
    },

    resizeIframe() {
      const iframe = this.$refs.iframe;

      if (!iframe) {
        return;
      }

      this.iframeHeight = iframe.contentWindow.document.documentElement.offsetHeight;
    },

    buildAddress(params) {
      if (params.name) {
        return `${params.name} [${params.address}]`;
      }

      return params.address;
    },

    deleteMessage() {
      this.$store.dispatch('deleteMessage', {
        mailboxSlug: this.$store.getters.activeMailbox.slug,
        message: this.message,
      }).then(() => {
        this.$router.push(`/${this.$store.getters.activeMailbox.slug}`);
      }).catch(() => {
        this.showDeleteAlert = true;
      });
    },
  },

  computed: {
    dateReceived() {
      return moment.unix(this.message.dates.unix).format('MMMM Do YYYY, h:mm:ss a');
    },

    message() {
      const message = this.$store.getters.activeMessage;
      const mailbox = this.$store.getters.activeMailbox;
      if (!message.read) {
        this.$store.dispatch('markMessageRead', { message, mailboxSlug: mailbox.slug });
      }

      return message;
    },

    extraRecipients() {
      return (this.toNum - 1)
        + this.ccNum
        + this.bccNum;
    },

    from() {
      return this.buildAddress(this.message.raw.from.value[0]);
    },

    to() {
      return this.buildAddress(this.message.raw.to.value[0]);
    },

    toNum() {
      return 'to' in this.message && this.message.raw.to
        ? this.message.raw.to.value.length : 0;
    },

    cc() {
      return this.buildAddress(this.message.raw.cc.value[0]);
    },

    ccNum() {
      return 'cc' in this.message && this.message.raw.cc
        ? this.message.raw.cc.value.length : 0;
    },

    bcc() {
      return this.buildAddress(this.message.raw.bcc.value[0]);
    },

    bccNum() {
      return 'bcc' in this.message && this.message.raw.bcc
        ? this.message.raw.bcc.value.length : 0;
    },

    src() {
      return this.message.src[this.view];
    },
  },

  beforeRouteEnter(to, from, next) {
    return next(vm => checkRoute(vm, to, next));
  },

  beforeRouteUpdate(to, from, next) {
    return checkRoute(this, to, next);
  },
};
</script>

<style lang="scss">
.message-view {
  height: 100%;
  overflow-y: scroll;

  header {
    position: relative;
    padding-bottom: 10px;
    margin-bottom: 10px;
    border-bottom: 1px solid #d6d6d6;

    h1 {
      margin: 0;
      font-size: 1em;
      font-weight: 700;
    }

    .date-received {
        font-size: 0.875em;
    }

    .actions {
      position: absolute;
      bottom: 0;
      right: 0;
    }

    .extra-recipients {
      font-size: 0.8em;
      padding: 0 3px;
    }
  }

  .message-body {
    position: relative;
    flex-grow: 1;

    iframe {
      width: 100%;
      height: 300px;
      border: none;
      position: relative;
      z-index: 1;

      &.loading {
        visibility: hidden;
      }
    }
  }
}
</style>
