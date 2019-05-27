<template>
  <vue-perfect-scrollbar class="scroll-area" :settings="settings">
    <transition-group name="fade" tag="ul" class="list-unstyled">
      <message v-for="message in messages" :key="message.id" :message="message"></message>
    </transition-group>
  </vue-perfect-scrollbar>
</template>

<script>
import { mapGetters } from 'vuex';
import VuePerfectScrollbar from 'vue-perfect-scrollbar';

import Message from './components/Message';

export default {
  data() {
    return {
      settings: {},
    };
  },

  components: {
    Message,
    VuePerfectScrollbar,
  },

  computed: {
    ...mapGetters(['messages']),
  },
};
</script>

<style lang="scss">
.fade-enter-active, .fade-leave-active {
  transition: opacity .2s;
}
  
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}

.scroll-area {
  height: 100%;
}
</style>

