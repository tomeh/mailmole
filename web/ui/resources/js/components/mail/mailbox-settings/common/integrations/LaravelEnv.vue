<template>

<pre>
MAIL_DRIVER=smtp
MAIL_HOST={{ host }}
MAIL_PORT={{ port }}
MAIL_USERNAME={{ username }}
MAIL_PASSWORD={{ password }}
</pre>

</template>

<script>
import integrationMixin from '@/lib/integration';

export default {
  mixins: [integrationMixin],
};
</script>