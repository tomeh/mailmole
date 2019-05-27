/* eslint no-param-reassign: [2, { "props": false }], */

import Vue from 'vue';
import { mailboxBySlug } from './getters';

export const setMailboxes = (state, mailboxes) => {
  Vue.set(state, 'mailboxes', mailboxes);
};

export const deleteMailbox = (state, mailbox) => {
  const i = state.mailboxes.indexOf(mailbox);

  if (i > -1) {
    state.mailboxes.splice(i, 1);
  }
};

export const deleteMessage = (state, payload) => {
  const mailboxIndex = state.mailboxes.findIndex(mailbox => mailbox.slug === payload.mailboxSlug);
  const mailbox = mailboxBySlug(state)(payload.mailboxSlug);
  const i = mailbox.messages.indexOf(payload.message);

  if (i > -1) {
    state.mailboxes[mailboxIndex].messages.splice(i, 1);
  }
};

export const addMailbox = (state, mailbox) => {
  state.mailboxes.push(mailbox);
};

export const addMessage = (state, payload) => {
  const i = state.mailboxes.findIndex(el =>
    el.slug === payload.mailbox_slug,
  );

  if (i === -1) {
    throw Error(`Cannot find mailbox with slug ${payload.mailboxSlug}`);
  }

  const mailbox = state.mailboxes[i];

  const messages = mailbox.messages;
  messages.push(payload.message);
  messages.sort((a, b) => b.dates.unix - a.dates.unix);

  for (let overlength = messages.length - mailbox.max_messages; overlength > 0; overlength -= 1) {
    messages.pop();
  }
};

export const reorderMailboxes = (state, order) => {
  state.mailboxes.sort((a, b) => order.indexOf(a.slug) - order.indexOf(b.slug));
};

export const updateMailbox = (state, payload) => {
  const i = state.mailboxes.findIndex(el =>
    el.slug === payload.slug,
  );

  if (i === -1) {
    throw Error(`Cannot find mailbox with slug ${payload.slug}`);
  }

  Vue.set(state.mailboxes, i, payload.mailbox);
};

export const clearMessagesForMailbox = (state, mailbox) => {
  const i = state.mailboxes.indexOf(mailbox);

  Vue.set(state.mailboxes[i], 'messages', []);
};

export const markMessageRead = (state, payload) => {
  const mailboxIndex = state.mailboxes.findIndex(mailbox => mailbox.slug === payload.mailboxSlug);
  const mailbox = mailboxBySlug(state)(payload.mailboxSlug);
  const i = mailbox.messages.indexOf(payload.message);

  if (i > -1) {
    Vue.set(state.mailboxes[mailboxIndex].messages[i], 'read', true);
  }
};
