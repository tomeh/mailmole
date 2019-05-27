import store from '@/stores/mail';
import router from '@/routes/mail';

window.Echo.private(`message-bus.${window.mailmole.user.id}`)
  .listen('NewMessage', (payload) => {
    const currentMailboxSlug = store.state.route.params.mailbox;
    const routeTo = (currentMailboxSlug === payload.mailbox_slug
        && store.getters.mailboxBySlug(payload.mailbox_slug).messages.length === 0);

    store.commit('addMessage', payload);

    if (routeTo) {
      router.push(`/${payload.mailbox_slug}/${payload.message.id}`);
    }
  });
