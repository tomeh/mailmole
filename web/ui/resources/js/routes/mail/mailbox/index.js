import MessageViewer from '@/components/mail/MessageViewer';
import MessageView from '@/components/mail/message-view/MessageView';
import store from '@/stores/mail';

export default {
  path: '/:mailbox',
  component: MessageViewer,
  name: 'mailbox',

  beforeEnter: (to, from, next) => {
    const slug = to.params.mailbox;
    const mailbox = store.getters.mailboxBySlug(slug);

    if (!mailbox) {
      return next('/');
    }

    if (to.params.message) {
      return next();
    }

    if (!mailbox || !mailbox.messages || mailbox.messages.length === 0) {
      return next();
    }

    const message = mailbox.messages[0];

    return next(`/${slug}/${message.id}`);
  },

  children: [
    {
      path: '/:mailbox',
      component: MessageView,
      props: (route) => ({ view: route.query.view }), // eslint-disable-line
      children: [
        {
          path: ':message',
          name: 'mailbox_message',
          component: MessageView,
        },
      ],
    },
  ],
};
