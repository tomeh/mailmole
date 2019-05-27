<template>
  <li
    class="message pr-1"
    :class="{ 'show-controls': hovered, hovered: hovered, active: isActive, read: message.read }"
    v-on:mouseover="mouseOver"
    v-on:mouseleave="mouseLeave"
    >

    <router-link
      :class="{ 'text-light': isActive }"
      :to="{ name: 'mailbox_message', params: { message: message.id }}"
      >
      <span
        class="message__date-received" v-text="dateReceived"
        :class="{ 'text-light': isActive }"
        ></span>
      <div
        class="message__name font-weight-bold"
        :class="{ 'text-light': isActive }"
        v-text="message.subject"
        :title="message.subject"></div>
      <div
        class="snippet mt-1"
        :class="{ 'text-light': isActive, 'font-italic': message.snippet === null }"
        v-text="message.snippet || 'Empty'"></div>
    </router-link>
  </li>
</template>

<script>
import moment from 'moment';

export default {
  props: {
    message: { required: true },
  },

  data() {
    return {
      hovered: false,
    };
  },

  methods: {
    mouseOver() {
      this.hovered = true;
    },

    mouseLeave() {
      this.hovered = false;
    },
  },

  computed: {
    dateReceived() {
      const date = moment.unix(this.message.dates.unix);
      const today = moment();

      if (date.isSame(today, 'd')) {
        return date.format('HH:mm');
      }

      if (date.isSame(today.subtract(1, 'd'), 'd')) {
        return 'Yesterday';
      }

      return date.format('DD/MM');
    },

    isActive() {
      const activeMessage = this.$store.getters.activeMessage;
      return !!activeMessage
        && activeMessage.constructor === Object
        && Object.prototype.hasOwnProperty.call(activeMessage, 'id')
        && activeMessage.id === this.message.id;
    },
  },
};
</script>

<style lang="scss">
@import '~@styles/mailmole/variables';

  .message {
    clear: both;
    cursor: pointer;

    &.active {
      background: map-get($theme-colors, 'neutral');
    }

    &:not(.active).hovered {
      background: #D8D8D8
    }

    a {
      color: inherit;
      text-decoration: none;

      display: block;
      font-size: 0.875em;
      padding: .5em .5em .5em .25em;
      border-left: .4em transparent solid;
    }

    &:not(.read) {
      a {
        border-left: .4em map-get($theme-colors, 'primary') solid;
      }
    }

    .message__name {
      margin: 0;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }

    .snippet {
      color: grey;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      clear: both;
    }

    .message__date-received {
      font-size: 0.875em;
      color: grey;
      float: right;
      padding-left: 5px;
    }
  }
</style>
