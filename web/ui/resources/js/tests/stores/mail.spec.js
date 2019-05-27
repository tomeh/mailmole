/* eslint-env mocha */

import expect from 'expect';
import * as getters from '../../stores/mail/getters';
import * as mutations from '../../stores/mail/mutations';

// Fake application state.
let state;

beforeEach(() => {
  state = {
    mailboxes: [{
      id: 1,
      name: 'mailbox 1',
      messages: [{
        id: 1,
        subject: 'mailbox 1 message 1',
      }, {
        id: 2,
        subject: 'mailbox 1 message 2',
      }],
    }, {
      id: 2,
      name: 'mailbox 2',
      messages: [{
        id: 3,
        subject: 'mailbox 2 message 3',
      }, {
        id: 4,
        subject: 'mailbox 2 message 4',
      }],
    }, {
      id: 3,
      name: 'mailbox 3',
    }],
  };
});

describe('The mail store getters', () => {
  it('should find a mailbox by id.', () => {
    expect(getters.mailboxBySlug(state)(2).name).toEqual('mailbox 2');
  });

  it('should find a message for a mailbox both by id.', () => {
    expect(getters.messageForMailboxBySlug(state)(2, 4).subject).toEqual('mailbox 2 message 4');
  });

  it('should get all messages for current mailbox.', () => {
    state.route = {
      params: {
        mailbox: 1,
      },
    };

    expect(getters.messages(state)).toBeInstanceOf(Array);
    expect(getters.messages(state)[0].id).toEqual(1);
  });

  it('should get the number of messages.', () => {
    state.route = {
      params: {
        mailbox: 1,
      },
    };

    expect(getters.numberOfMessages(state)).toEqual(2);
  });

  it('should get the current active mailbox.', () => {
    state.route = {
      params: {
        mailbox: 1,
      },
    };

    expect(getters.activeMailbox(state).name).toEqual('mailbox 1');
  });

  it('should get the current active message.', () => {
    state.route = {
      params: {
        mailbox: 1,
        message: 2,
      },
    };

    expect(getters.activeMessage(state).subject).toEqual('mailbox 1 message 2');
  });
});

describe('The mail store mutations', () => {
  it('should set mailboxes on state.', () => {
    const oldState = {};

    mutations.setMailboxes(oldState, state.mailboxes);

    expect(oldState).toHaveProperty('mailboxes');
    expect(oldState.mailboxes).toBeInstanceOf(Array);
    expect(oldState.mailboxes).toHaveLength(3);
  });

  it('should delete a mailbox.', () => {
    mutations.deleteMailbox(state, state.mailboxes[1]);

    expect(state.mailboxes).toHaveLength(2);
  });

  it('should delete a message.', () => {
    mutations.deleteMessage(state, {
      mailboxSlug: state.mailboxes[1].slug,
      message: state.mailboxes[1].messages[0],
    });

    expect(state.mailboxes[1].messages).toHaveLength(1);
    expect(state.mailboxes[1].messages[0].id).toEqual(4);
  });

  it('should add a mailbox.', () => {
    mutations.addMailbox(state, {
      slug: 'slug',
    });

    expect(state.mailboxes).toHaveLength(4);
    expect(state.mailboxes[3].slug).toEqual('slug');
  });

  xit('should add a message.', () => {
    mutations.addMessage(state, {
      mailbox_slug: 1,
      message: {
        id: 100,
      },
    });

    expect(state.mailboxes[0].messages).toHaveLength(3);
    expect(state.mailboxes[0].messages[2].id).toEqual(100);
  });

  it('should reorder mailboxes.');
  it('should update a mailbox.');
  it('should mark a message as read.');
});
