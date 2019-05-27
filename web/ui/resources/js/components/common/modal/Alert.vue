<template>
  <modal :show="show" @close="close">
    <template slot="header"></template>
    <slot></slot>
    <template slot="footer">
      <div class="float-right">
        <button @click="close" class="btn btn-primary btn-sm">Close</button>
      </div>
      <div class="clearfix"></div>
    </template>
  </modal>
</template>

<script>
import Modal from './Modal';

export default {
  props: {
    show: { type: Boolean },
  },

  components: {
    Modal,
  },

  methods: {
    close() {
      this.$emit('close');
    },
  },
};
</script>

<style lang="scss">

</style>
