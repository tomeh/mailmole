<template>
  <div class="dropdown-menu">
    <mailbox-list></mailbox-list>
    <mailbox-name
      :formOpen="addMailboxFormOpen" 
      :errors="addMailboxErrors" 
      :pending="addMailboxPending" 
      @commit="newMailbox" 
      @cancel="cancelAddMailbox"
    ></mailbox-name>
    <div class="dropdown-divider"></div>

    <a class="dropdown-item" @click.stop.prevent="openAddMailboxForm" href="#">Create a new mailbox</a>
    <alert :show="showGeneralErrorAlert" @close="showGeneralErrorAlert = false">An error occured and your request could not be completed</alert>
  </div>
</template>

<script>
import Alert from '@/components/common/modal/Alert';
import MailboxList from '@/components/mail/mailbox-list/MailboxList';
import MailboxName from '@/components/mail/mailbox-list/MailboxName';

export default {
  data() {
    return {
      addMailboxFormOpen: false,
      addMailboxErrors: [],
      addMailboxPending: false,
      addMailboxLastSubmitted: '',
      showGeneralErrorAlert: false,
    };
  },
  components: {
    Alert,
    MailboxList,
    MailboxName,
  },
  methods: {
    addMailbox() {

    },

    openAddMailboxForm() {
      this.addMailboxFormOpen = true;
    },

    closeAddMailboxForm() {
      this.addMailboxFormOpen = false;
      this.addMailboxPending = false;
    },

    cancelAddMailbox() {
      this.closeAddMailboxForm();
    },

    showAddMailboxFormErrors(errors) {
      this.addMailboxErrors = errors;
    },

    showGeneralError() {
      this.closeForm();
      this.showGeneralErrorAlert = true;
    },

    newMailbox(name) {
      if (!name) {
        this.cancel();
        return;
      }

      if (this.addMailboxLastSubmitted === name) {
        return;
      }

      this.addMailboxLastSubmitted = name;
      this.addMailboxPending = true;

      this.$store.dispatch('addMailbox', name)
        .then((response) => {
          const newMailbox = response.data.mailbox;

          this.addMailboxLastSubmitted = '';
          this.closeAddMailboxForm();
          this.$router.push(`/${newMailbox.slug}`);
        })
        .catch((e) => {
          if (e.response.status === 422) {
            const errors = e.response.data.errors;
            if ('name' in errors) {
              this.addMailboxErrors = errors.name;
            }

            this.addMailboxPending = false;

            return;
          }

          this.showGeneralError();
        });
    },
  },
};

</script>

<style lang="scss">
  .nav-link {
    padding: 0.25rem 1.5rem;
  }
</style>
