<template>
  <div class="integrations">
    <div class="form-group w-50">
      <select name="integration" class="form-control" v-model="currentView">
        <option :key="integration" v-for="(label, integration) in integrations" :value="integration" v-text="label"></option>
      </select>
    </div>
    <div class="code-viewer">
      <component :is="currentView" :mailbox="mailbox"></component>
    </div>
  </div>
</template>

<script>

export default {

  props: {
    mailbox: { required: true, type: Object },
  },

};

</script>

<script>
import swaks from './integrations/Swaks';
import laravel from './integrations/Laravel';
import laravelEnv from './integrations/LaravelEnv';
import telnet from './integrations/Telnet';

export default {

  props: {
    mailbox: { required: true, type: Object },
  },

  data() {
    return {
      integrations: {
        swaks: 'Swaks',
        laravel: 'Laravel',
        laravelEnv: 'Laravel (.env)',
        telnet: 'Telnet',
      },
      currentView: 'swaks',
    };
  },

  components: {
    swaks,
    laravel,
    laravelEnv,
    telnet,
  },

};

</script>

<style lang="scss">
.integrations {
  select {
    margin-bottom: 1em;
  }
}
</style>

