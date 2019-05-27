<template>
  <router-link
    :to="`/${mailbox.slug}`"
    tag="li"
    class="mailbox dropdown-item"
  >

    <div class="mailbox__name" :class="{ 'has-unread': numUnread > 0 }">
      <span v-text="mailbox.name"></span>
      <span v-if="numUnread > 0">({{ numUnread }})</span>
    </div>
  </router-link>
</template>

<script>
export default {
  props: {
    mailbox: { required: true },
  },

  methods: {
    clearMessagesForMailboxClicked() {
      if (this.empty) return;

      this.$store.commit('clearMessagesForMailbox', this.mailbox);
    },
  },

  computed: {
    numUnread() {
      return this.mailbox.messages.reduce((acc, message) => {
        if (message.read) return acc;

        return acc + 1;
      }, 0);
    },

    empty() {
      return !(this.mailbox.messages instanceof Array) || this.mailbox.messages.length < 1;
    },
  },
};
</script>

<style lang="scss">
//   @import '~@styles/mailmole/variables';

  .mailbox {
    clear: both;
    cursor: pointer;

    &:active {
    //   background-color: theme-color("primary");
    //   color: $light;
    }

    &.router-link-active {
    //   background-color: theme-color("primary");
    //   color: $light;
    }

    .mailbox__name {
      margin: 0;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;

      &.has-unread {
        font-weight: bold;
      }
    }
  }
</style>
