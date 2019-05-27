import axios from 'axios';

export default class {
  static createMailbox(name) {
    return axios.post('/mailbox', {
      name,
    });
  }

  static updateMailboxBySlug(slug, mailbox) {
    return axios.put(`/mailbox/${slug}`, mailbox);
  }

  static deleteMailboxBySlug(slug) {
    return axios.delete(`/mailbox/${slug}`);
  }

  static clearMailboxBySlug(slug) {
    return axios.post(`/mailbox/${slug}/clear`);
  }

  static deleteMessageById(id) {
    return axios.delete(`/message/${id}`);
  }

  static reorderMailboxes(order) {
    return axios.post('/mailbox/reorder', {
      order,
    });
  }

  static markMessageReadById(id) {
    return axios.patch(`/message/${id}/mark-read`);
  }
}
