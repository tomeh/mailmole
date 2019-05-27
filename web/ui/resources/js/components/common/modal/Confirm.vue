<template>
  <modal :show="show" @close="cancel" v-on:keyup.enter="confirm" v-on:keyup.esc="cancel">
    <template slot="header"></template>
    <slot></slot>
    <template slot="footer">
      <div class="float-left">
        <spinner v-if="(pending && spinnerWhenPending)" class="spinner"></spinner>
      </div>
      <div class="float-right">
        <button @click="cancel" class="btn btn-default btn-sm">Cancel</button>
        <button @click="confirm" class="btn btn-primary btn-sm" :disabled="pending">OK</button>
      </div>
      <div class="clearfix"></div>
    </template>
  </modal>
</template>

<script>
import Spinner from 'vue-simple-spinner';
import Modal from './Modal';

export default {
  props: {
    show: { type: Boolean },
    pending: { type: Boolean, default: false },
    spinnerWhenPending: { type: Boolean, default: true },
  },

  components: {
    Modal,
    Spinner,
  },

  mounted() {
    document.addEventListener('keydown', (e) => {
      if (this.show && e.keyCode === 13) {
        this.confirm();
      }

      if (this.show && e.keyCode === 27) {
        this.cancel();
      }
    });
  },

  methods: {
    confirm() {
      this.$emit('confirm');
    },

    cancel() {
      this.$emit('cancel');
    },
  },
};
</script>

<style lang="scss">

</style>
