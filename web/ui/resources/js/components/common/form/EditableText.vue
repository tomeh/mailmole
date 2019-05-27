<template>
  <div class="editable-text">
    <mm-input
      v-if="editing"
      :autoSelect="true"
      :initial="value"
      class="form-control"
      @commit="commit"
      @cancel="$emit('cancel')"
      ></mm-input>

    <div v-else>
      <span v-text="value"></span>
      <font-awesome-icon :icon="['fas', 'pencil']" class="" v-on:click="$emit('edit')"/>
    </div>
  </div>
</template>

<script>
import Input from './Input';

export default {
  props: {
    value: { required: false, default: '' },
    editing: { required: false, type: Boolean, default: false },
  },

  data() {
    return {
      lastChange: this.value,
    };
  },

  components: {
    'mm-input': Input,
  },

  methods: {
    commit(value) {
      this.$emit(value === this.lastChange ? 'cancel' : 'commit', value);
    },
  },
};
</script>

<style lang="scss">
  .editable-text {
    i {
      padding: 0 0 0 1em;
    }
  }
</style>
