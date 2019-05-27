<template>
<div>
  <p><a href="#">Sending test emails with swaks</a></p>

<pre>
swaks --to user@example.com \
  --server {{ host }}:{{ port }} \
  --auth-user {{ username }} \
  --auth-password {{ password }} \
  --header="Subject: Test email" \
  --body="Mailmole.io test email using swaks\!"
</pre>

</div>
</template>

<script>
import integrationMixin from '@/lib/integration';

export default {
  mixins: [integrationMixin],
};
</script>
