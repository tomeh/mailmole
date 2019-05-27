<template>
  <a :href="`/message/${message.id}/attachment/${attachment.id}/download`">
    <i class="far pr-2" :class="iconClass"></i><span v-text="attachment.original_file_name"></span>
  </a>
</template>

<script>
export default {
  props: {
    attachment: { type: Object, required: true },
  },

  computed: {
    message() {
      return this.$store.getters.activeMessage;
    },

    iconClass() {
      switch (this.attachment.contentType) {
        case 'application/pdf':
          return 'fa-file-pdf';

        case 'image/jpeg':
        case 'image/png':
          return 'fa-file-image';

        default:
          return 'fa-file';
      }
    },
  },
};
</script>

