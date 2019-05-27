<template>
  <draggable v-model='mailboxes' :element="'ul'" class="list-unstyled mb-0">
    <mailbox 
      v-for="mailbox in mailboxes" 
      :key="mailbox.slug" 
      :mailbox="mailbox" 
      ></mailbox>
  </draggable>
</template>

<script>
import Draggable from 'vuedraggable';
// import { mapState } from 'vuex';

import Mailbox from './components/Mailbox';

export default {
  components: {
    Draggable,
    Mailbox,
  },

  computed: {
    // ...mapState(['mailboxes']),

    mailboxes: {
      get() {
        return this.$store.state.mailboxes;
      },
      set(mailboxes) {
        this.$store.dispatch('reorderMailboxes', mailboxes.map(m => m.slug));
      },
    },
  },
};
</script>
