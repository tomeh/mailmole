/* eslint no-param-reassign: [2, { "props": false }], */

// Watch function return.
export const mailboxBySlug = state =>
  slug => state.mailboxes.find(mailbox => mailbox.slug === slug) || null;

// Watch function return.
export const messageForMailboxBySlug = state =>
  (mailboxSlug, id) => {
    const mailbox = mailboxBySlug(state)(mailboxSlug);
    const messages = mailbox === null ? [] : mailbox.messages;
    return messages.find(message => message.id === Number(id));
  };

export const messages = (state) => {
  const slug = state.route.params.mailbox;
  const mailbox = mailboxBySlug(state)(slug);

  if (mailbox instanceof Object && Object.prototype.hasOwnProperty.call(mailbox, 'messages')) {
    return mailbox.messages;
  }

  return null;
};

export const numberOfMessages = (state) => {
  const m = messages(state);

  if (m instanceof Array) {
    return m.length;
  }

  return 0;
};

export const activeMailbox = (state) => {
  const slug = state.route.params.mailbox;

  return mailboxBySlug(state)(slug) || null;
};

export const activeMessage = (state) => {
  const mailbox = activeMailbox(state);
  const id = Number(state.route.params.message);

  return messageForMailboxBySlug(state)(mailbox.slug, id) || null;
};
