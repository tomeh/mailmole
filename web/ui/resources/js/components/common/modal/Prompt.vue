<template>
  <modal :show="show" @close="cancel" @opened="opened">
    <template slot="header"></template>

    <div class="form-group">
      <label for="promptInput"><slot></slot></label>
      <input ref="input" v-model="value" name="value" type="text" class="form-control" id="promptInput" :placeholder="placeholder">
      <slot name="input-help" class="form-text text-muted"></slot>

      <ul class="list-unstyled mt-1 mb-0" v-if="errors.length > 0">
        <li class="text-danger" v-for="(error, key) in errors" :key="key">
          <small v-text="error"></small>
        </li>
      </ul>
    </div>

    <template slot="footer">
      <div class="float-left">
        <spinner v-if="(pending && spinnerWhenPending)" class="spinner"></spinner>
      </div>

      <div class="float-right">
        <button @click="cancel" class="btn btn-default btn-sm">Cancel</button>
        <button @click="submit" class="btn btn-primary btn-sm" :disabled="pending || value.length === 0">OK</button>
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
    initialValue: { default: '' },
    placeholder: { default: '' },
    errors: { type: Array, default: () => [] },
  },

  data() {
    return {
      value: this.initialValue,
    };
  },

  mounted() {
    document.addEventListener('keydown', (e) => {
      if (this.show && e.keyCode === 13) {
        this.submit();
      }

      if (this.show && e.keyCode === 27) {
        this.cancel();
      }
    });
  },

  components: {
    Modal,
    Spinner,
  },

  methods: {
    submit() {
      this.$emit('submit', this.value);
    },

    opened() {
      $(this.$refs.input).focus();
      this.$emit('opened');
    },

    cancel() {
      this.value = this.initialValue;
      this.$emit('cancel');
    },
  },
};
</script>

<style lang="scss">

</style>
