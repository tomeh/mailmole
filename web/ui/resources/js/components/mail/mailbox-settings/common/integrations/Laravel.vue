<template>

<pre>
return [
  'driver'   => 'smtp',
  'host'     => '{{ host }}',
  'port'     => '{{ port }}',
  'username' => '{{ username }}',
  'password' => '{{ password }}',
];
</pre>

</template>

<script>
import integrationMixin from '@/lib/integration';

export default {
  mixins: [integrationMixin],
};
</script>