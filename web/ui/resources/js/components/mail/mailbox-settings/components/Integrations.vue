<template>
  <div class="row card-group">
    <div class="col-12">
      <div class="card w-100">
        <div class="card-header">
          Integrations
        </div>
        <div class="card-body">
          <integrations :mailbox="mailbox"></integrations>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Integrations from '../common/Integrations';

export default {
  components: {
    Integrations,
  },

  computed: {
    mailbox() {
      return this.$store.getters.activeMailbox;
    },
  },
};
</script>
