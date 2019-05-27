<template>
  <div ref="root" class="split-pane"
  @mousemove="dragMove"
  @mouseup="dragEnd"
  @mouseleave="dragEnd"
  :class="{ dragging: dragging }">
    <div class="left" :class="{ closed: closed }" :style="{ width: leftSplit + '%' }">
      <slot name="left"></slot>
      <div class="cover" :class="{ on: dragging }"></div>
      <div class="dragger" @mousedown.stop="dragStart"></div>
    </div>
    <div class="right" :style="{ width: rightSplit + '%' }">
      <slot name="right"></slot>
      <div class="cover" :class="{ on: dragging }"></div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    initialSplit: { default: 50, type: Number },
    maxSplit: { default: 100, Type: Number },
    minSplit: { default: 0, type: Number },
    closed: { default: false, type: Boolean },
  },

  data() {
    return {
      split: this.initialSplit,
      dragging: false,
    };
  },

  mounted() {
    this.$refs.root
      .addEventListener('selectstart', (e) => {
        if (this.dragging) {
          e.preventDefault();
        }
      });
  },

  methods: {
    dragStart(e) {
      this.dragging = true;
      this.startX = e.pageX;
      this.startSplit = this.split;
    },

    dragMove(e) {
      if (this.dragging) {
        const dx = e.pageX - this.startX;
        const totalWidth = this.$el.offsetWidth;
        const split = this.startSplit + ((dx / totalWidth) * 100);
        if (this.minSplit > split || split > this.maxSplit) {
          return false;
        }

        this.$emit('dragging');

        this.split = split;

        return true;
      }

      return false;
    },

    dragEnd() {
      this.dragging = false;
      this.$emit('finished', this.split);
    },
  },

  computed: {
    leftSplit() {
      return this.closed ? 0 : this.split;
    },

    rightSplit() {
      return this.closed ? 100 : (100 - this.split);
    },
  },
};
</script>

<style lang="scss">
  .split-pane {
    display: flex;
    height: 100%;


    &.dragging {
      cursor: ew-resize;
    }

    .left, .right {
      position: relative;
      .cover {
        position: absolute;
        width: 99%;
        height: 100%;
        top: 0;
        left: 3px;
        display: block;
        z-index: 10;

        &:not(.on) {
          display: none;
        }
      }
    }

    .left {
      &:not(.closed) {
        border-right: 1px solid #d6d6d6;
      }
    }

    .dragger {
      position: absolute;
      z-index: 99;
      top: 0;
      bottom: 0;
      right: -5px;
      width: 10px;
      cursor: ew-resize;
    }
  }
</style>