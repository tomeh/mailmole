<template>
  <div class="list-controls">
    <div class="list-controls__control-container" v-if="formOpen">

      <mm-input
        class="w-100 mt-2"
        :autoSelect="true"
        :disabled="pending" 
        :class="{ pending }" 
        @commit="onCommit" 
        @cancel="onCancel"
        ></mm-input>

      <spinner v-if="pending" class="spinner" size="small"></spinner>

      <small class="list-controls__error" v-if="errors.length" v-text="errors[0]"></small>
    </div>
  </div>
</template>

<script>
import Spinner from 'vue-simple-spinner';
import Input from '@/components/common/form/Input';

export default {

  props: {
    formOpen: { required: true },
    errors: { type: Array, default: [] },
    pending: { type: Boolean, default: false },
  },

  components: {
    'mm-input': Input,
    Spinner,
  },

  methods: {
    onCommit(name) {
      this.$emit('commit', name);
    },

    onCancel() {
      this.$emit('cancel');
    },
  },

};
</script>

<style lang="scss">
  .list-controls {
    clear: both;
    padding: 0 .5em;

    .list-controls__error {
      color: red;
    }

    .btn.btn-link {
      padding: 0;
    }

    .list-controls__control-container {
      position: relative;
      .spinner {
        position: absolute;
        right: 10px;
        top: calc(50% - 8px);
      }
    }
  }
</style>


