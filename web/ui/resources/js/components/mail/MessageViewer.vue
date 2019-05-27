<template>
  <div class="h-100">
    <empty-mailbox v-if="$store.getters.numberOfMessages === 0"></empty-mailbox>

    <split-pane
      class="h-100"
      :initialSplit="28"
      :minSplit="15"
      :maxSplit="50"
      v-else-if="$store.getters.activeMailbox !== null"
      @dragging="dragging"
      @finished="dragFinish"
      >
      <div slot="left" class="h-100">
        <message-list></message-list>
      </div>
      <div slot="right" class="p-2 h-100">
        <router-view v-if="$store.getters.activeMessage !== null"></router-view>
        <div v-else class="h-100">
          Message not found
        </div>
      </div>
    </split-pane>

    <div class="padded h-100" v-else>
      Mailbox not found
    </div>
  </div>
</template>

<script>
import SplitPane from '@/components/common/SplitPane';
import MessageList from '@/components/mail/message-list/MessageList';
import EmptyMailbox from '@/components/mail/message-list/Empty';

import eventBus from '@/lib/events';

function checkRoute(vm, to, next) {
  if ([
    'configure_mailbox',
    'configure_users',
  ].indexOf(to.name) !== -1) {
    return next();
  }

  if ('message' in to.params) {
    return next();
  }

  const slug = to.params.mailbox;
  const mailbox = vm.$store.getters.mailboxBySlug(slug);

  if (!mailbox || !mailbox.messages || mailbox.messages.length === 0) {
    return next();
  }

  const message = mailbox.messages[0];
  return next(`/${slug}/${message.id}`);
}

export default {
  components: {
    SplitPane,
    MessageList,
    EmptyMailbox,
  },

  methods: {
    dragFinish() {
      this.resize();
    },

    dragging() {
      this.resize();
    },

    resize() {
      eventBus.$emit('resize');
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
</style>
