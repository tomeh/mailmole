<template>
  <input
    ref="input"
    type="text"
    v-model="value"
    @blur.stop="blur"
    @keyup.enter.stop="enter"
    @keydown.esc.stop.prevent="esc"
    >
</template>

<script>
export default {
  props: {
    initial: { required: false, default: '' },
    autoSelect: { required: false, type: Boolean, default: false },
    resetOnCancel: { required: false, type: Boolean, default: true },
  },

  data() {
    return {
      value: this.initial,
    };
  },

  mounted() {
    if (this.autoSelect) {
      this.$refs.input.focus();
      this.$refs.input.select();
    }
  },

  methods: {
    blur() {
      if (this.value !== this.initial) {
        return this.$emit('commit', this.value);
      }

      return this.$emit('cancel');
    },
    enter() {
      this.$refs.input.blur();
    },
    esc() {
      if (this.resetOnCancel) {
        this.value = this.initial;
      }

      this.$emit('cancel');
    },
  },
};
</script>

<style lang="scss">
</style>
