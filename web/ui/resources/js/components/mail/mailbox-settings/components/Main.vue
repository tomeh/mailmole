<template>
  <div class="row card-group">
    <div class="col-12">
      <div class="card w-100 mb-3">
        <div class="card-header">
          Mailbox Name
        </div>
        <form class="card-body" @submit.stop>
          <div class="form-group">
            <editable-text
              id="mailboxName"
              :value="mailbox.name"
              :editing="editing"
              @edit="edit"
              @commit="onCommit"
              @cancel="cancel"
              ></editable-text>
            <div class="mailbox-settings__error" v-if="errors.name.length" v-text="errors.name[0]"></div>
          </div>
        </form>
      </div>

      <div class="card w-100 mb-3">
        <div class="card-header">
          SMTP Settings
        </div>
        <div class="card-body">
          <general :mailbox="mailbox"></general>
        </div>
      </div>

      <div class="card w-100 mb-3">
        <div class="card-header">
          Danger
        </div>
        <div class="card-body">
          <button class="btn btn-danger" @click="deleteMailbox">Delete mailbox!</button>
          <confirm
            class="confirm-danger"
            :show="showDeleteConfirm"
            :pending="pendingDeleteConfirm && showDeleteConfirm"
            @confirm="onConfirmDelete"
            @cancel="showDeleteConfirm = false">
            Are you sure that you want to delete this mailbox?
          </confirm>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Confirm from '@/components/common/modal/Confirm';
import EditableText from '@/components/common/form/EditableText';
import General from '../common/General';

const initialErrors = {
  name: [],
};

export default {
  data() {
    return {
      editing: false,
      errors: initialErrors,
      showDeleteConfirm: false,
      pendingDeleteConfirm: false,
    };
  },

  components: {
    Confirm,
    EditableText,
    General,
  },

  computed: {
    mailbox() {
      return this.$store.getters.activeMailbox;
    },
  },

  methods: {
    edit() {
      this.editing = true;
    },

    cancel() {
      this.editing = false;
    },

    onCommit(name) {
      const mailbox = Object.assign({}, this.mailbox);
      mailbox.name = name;
      this.errors = initialErrors;

      this.$store
        .dispatch('renameMailbox', {
          id: mailbox.slug,
          mailbox,
        })
        .then(() => {
          this.editing = false;
        })
        .catch((e) => {
          if (e.response.status === 422) {
            this.errors = e.response.data.errors;
            this.pending = false;

            return;
          }

          this.showGeneralError();
        });
    },

    onConfirmDelete() {
      this.pendingDeleteConfirm = true;
      this.$store
        .dispatch('deleteMailbox', this.mailbox)
        .then(() => {
          this.$router.push('/');
          this.pendingDeleteConfirm = false;
        })
        .catch(() => {
          this.pendingDeleteConfirm = false;
        });
    },

    deleteMailbox() {
      this.showDeleteConfirm = true;
    },
  },
};
</script>

<style lang="scss">
  @import '~@styles/mailmole/variables';

  #mailboxName {
    font-weight: 700;
  }

  .confirm-danger {
   .btn.btn-primary {
      background-color: $red !important;
    }
  }
</style>
