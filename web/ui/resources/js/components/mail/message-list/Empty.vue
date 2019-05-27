<template>
  <div class="empty-mailbox p-3">
    <article class="text-center">
      <h1 class="setting-heading mt-5 mb-5 text-center">
        <i class="far fa-inbox mb-3"></i> <br>
        This mailbox is empty
      </h1>

      <quick-start :mailbox="activeMailbox" class="mb-2"></quick-start>

      <router-link
        class="empty-mailbox_view-settings"
        :to="{ name: 'configure_mailbox', params: { mailbox: activeMailbox.slug }}"
      >
        <font-awesome-icon :icon="['fas', 'cog']" class=""/>
        <span v-text="'View settings'"></span>
      </router-link>
    </article>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import QuickStart from './components/QuickStart';

export default {
  components: {
    QuickStart,
  },

  computed: mapGetters(['activeMailbox']),
};
</script>

<style lang="scss">
@import '~@styles/mailmole/variables';

.empty-mailbox {
  article {
    padding-bottom: 1.4em;

    h1.setting-heading {
      // font-size: em;
    }

    .setting-heading {
      font-weight: 700;
    }

    .empty-mailbox_view-settings {
      color: map-get($theme-colors, 'primary');
    }
  }
}

</style>
