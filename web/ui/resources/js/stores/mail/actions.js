/* eslint no-param-reassign: [2, { "props": false }], */

import mailmole from '@/api';

export const addMailbox = ({ commit }, name) => {
  const p = mailmole.createMailbox(name);

  p.then((response) => {
    // Commit the change to the store on success.
    commit('addMailbox', response.data.mailbox);
  });

  return p;
};


export const renameMailbox = ({ commit }, payload) => {
  const slug = payload.mailbox.slug;
  const params = { slug, mailbox: payload.mailbox };

  // Send update to server.
  const p = mailmole.updateMailboxBySlug(slug, params.mailbox);

  p.then((response) => {
    // Update store with updated mailbox on success.
    commit('updateMailbox', { slug, mailbox: response.data.mailbox });
  });

  return p;
};

export const reorderMailboxes = ({ commit }, order) => {
  commit('reorderMailboxes', order);

  const p = mailmole.reorderMailboxes(order);

  return p;
};

export const deleteMailbox = ({ commit }, mailbox) => {
  const p = mailmole.deleteMailboxBySlug(mailbox.slug);

  p.then(() => {
    commit('deleteMailbox', mailbox);
  });

  return p;
};

export const clearMessagesForMailbox = ({ commit }, mailbox) => {
  const p = mailmole.clearMailboxBySlug(mailbox.slug);

  p.then(() => {
    commit('clearMessagesForMailbox', mailbox);
  });

  return p;
};

export const deleteMessage = ({ commit }, payload) => {
  const p = mailmole.deleteMessageById(payload.message.id);

  p.then(() => {
    commit('deleteMessage', payload);
  });

  return p;
};

export const markMessageRead = ({ commit }, payload) => {
  const p = mailmole.markMessageReadById(payload.message.id);

  p.then(() => {
    commit('markMessageRead', payload);
  });

  return p;
};
