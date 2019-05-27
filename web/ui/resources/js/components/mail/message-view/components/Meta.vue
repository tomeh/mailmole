<template>
  <div class="message-meta">
    <span class="message-meta__label"><slot name="label"></slot></span>:
    <span class="message-meta__value"><slot name="value"></slot></span>
  </div>
</template>

<style lang="scss">
  .message-meta {
    margin: 3px 0;
  }

  .message-meta__label {
    font-weight: 700;
  }

  // .message-meta__label {
  // }
</style>
