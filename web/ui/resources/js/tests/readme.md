Example vue component test

```js
/* eslint-env mocha */

import { mount, createLocalVue } from 'vue-test-utils';
import expect from 'expect';
import Vuex from 'vuex';
import MailViewer from '../MailViewer';

const localVue = createLocalVue();
localVue.use(Vuex);

describe('test', () => {
  let $store;

  beforeEach(() => {
    $store = new Vuex.Store({
      state: {},
    });
  });

  it('works', () => {
    const wrapper = mount(MailViewer, {
      mocks: {
        $store,
      },
    });

    console.log(wrapper.find('div'));
  });
});
```