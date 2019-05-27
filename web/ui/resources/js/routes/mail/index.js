import Vue from 'vue';
import Router from 'vue-router';
import NoMailboxes from '@/components/mail/NoMailboxes';
import store from '@/stores/mail';
import mailbox from './mailbox';
import configureMailbox from './mailbox/configure';

Vue.use(Router);

const router = new Router({
  routes: [
    {
      path: '/',
      component: NoMailboxes,

      beforeEnter: (to, from, next) => {
        if (store.state.mailboxes.length === 0) {
          return next();
        }

        const firstMailbox = store.state.mailboxes[0];

        if (!firstMailbox) {
          return next('/start');
        }

        if (!firstMailbox || !firstMailbox.messages || firstMailbox.messages.length === 0) {
          return next(`/${firstMailbox.slug}`);
        }

        const message = firstMailbox.messages[0];

        return next(`/${firstMailbox.slug}/${message.id}`);
      },
    },
    configureMailbox,
    mailbox,
  ],
});

export default router;
