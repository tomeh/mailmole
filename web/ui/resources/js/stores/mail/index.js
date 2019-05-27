import Vue from 'vue';
import Vuex from 'vuex';
import * as actions from './actions';
import * as getters from './getters';
import * as mutations from './mutations';
import { initialMailboxes, user } from './setup';

Vue.use(Vuex);

const state = {
  mailboxes: [],
  user,
};

const store = new Vuex.Store({
  state,
  getters,
  actions,
  mutations,
});

store.commit('setMailboxes', initialMailboxes());

export default store;
