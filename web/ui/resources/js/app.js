import Vue from 'vue';
import { sync } from 'vuex-router-sync';

import '@/bootstrap';
import '@/fontawesome';
import store from '@/stores/mail';
import router from '@/routes/mail';
import Nav from '@/components/mail/nav/Nav';
import eventBus from '@/lib/events';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

import '@/listeners/mail';

// Sync vue-router into the vuex store.
sync(store, router);

window.onresize = () => eventBus.$emit('resize');

Vue.component('font-awesome-icon', FontAwesomeIcon);

new Vue({ // eslint-disable-line no-new
  el: '#app',
  store,
  router,
  components: {
    'mm-navbar': Nav,
  },
});
