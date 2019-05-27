<template>
  <nav id="app-nav" class="d-flex align-items-center flex-row justify-content-between pl-2 pr-2 mb-2 text-light bg-neutral navbar-spark">
    <div class="d-flex flex-row align-items-center" v-if="activeMailbox">
      <div id="mailbox-list-dropdown" class="dropdown">
        <button
          title="Menu"
          class="btn nav-btn btn-no-outline btn-sm border-none text-light pt-0 pb-0"
          id="mainMenuDropdownButton"
          data-toggle="dropdown"
          aria-haspopup="true"
          aria-expanded="false"
        >
          <span class="align-middle h-100">
              <span class="mailbox-name float-left pr-2" v-text="$store.getters.activeMailbox.name"></span>
              <div class="float-right">
                <font-awesome-icon :icon="['fas', 'angle-down']" class="align-middle"/>
              </div>
          </span>
        </button>
        <main-menu aria-labelledby="mainMenuDropdownButton"></main-menu>
      </div>

      <div>
        <router-link
          title="Mailbox settings"
          :to="cogLink"
          class="btn nav-btn configure-btn btn-no-outline border-none text-light pt-0 pb-0 ml-2"
          :class="{ configuring: configuring }"
        ><font-awesome-icon :icon="['fas', 'cog']" class="align-middle h-100"/></router-link>
      </div>

      <div v-if="numberOfMessages > 0">
        <a href="#"
          @click.prevent="clearAll"
          title="Clear all messages"
          class="btn nav-btn btn-no-outline border-none text-light pt-0 pb-0 ml-2"
        >
          <font-awesome-icon :icon="['fas', 'ban']" class="align-middle h-100"/>
        </a>
      </div>
    </div>
    <div v-else>
        <button
          title="Menu"
          class="btn nav-btn nav-btn-primary btn-no-outline border-none text-light pt-0 pb-0"
          @click="showCreateMailboxPrompt = true"
          aria-haspopup="true"
        >
        Create mailbox
        </button>

        <mm-prompt
          :show="showCreateMailboxPrompt"
          :pending="showCreateMailboxPrompt && pendingCreateMailboxPrompt"
          :errors="errorsCreateMailboxPrompt"
          placeholder="Mailbox name"
          @cancel="showCreateMailboxPrompt = false; errorsCreateMailboxPrompt = []"
          @submit="submitted"
        >
        </mm-prompt>
    </div>

    <router-link class="navbar-brand d-flex mm-logo mm-logo-light pl-2" to="/">
      <mm-logo></mm-logo>
    </router-link>

    <div class="d-flex flex-row">
      <div class="dropdown">
        <a href='#'
          title="User menu"
          class="pt-0 pb-0"
          id="userMenuDropdownButton"
          data-toggle="dropdown"
          aria-haspopup="true"
          aria-expanded="false"
        >
            <img v-if="avatarUrl" :src="avatarUrl" class="dropdown-toggle-image">
            <font-awesome-icon v-else :icon="['fas', 'user-circle']" size="3x"/>
        </a>
        <user-menu class="dropdown-menu-right" aria-labelledby="userMenuDropdownButton"></user-menu>
      </div>
    </div>
  </nav>
</template>

<script>
import { mapState, mapGetters } from 'vuex';
import Prompt from '@/components/common/modal/Prompt';
import errorsHelper from '@/lib/laravel-errors-helper';
import MainMenu from './MainMenu';
import UserMenu from './UserMenu';
import Logo from './Logo';

export default {
  data() {
    return {
      foo: 'bar',
      showCreateMailboxPrompt: false,
      pendingCreateMailboxPrompt: false,
      errorsCreateMailboxPrompt: [],
    };
  },

  components: {
    MmPrompt: Prompt,
    MmLogo: Logo,
    MainMenu,
    UserMenu,
  },

  methods: {
    submitted(mailboxName) {
      this.pendingCreateMailboxPrompt = true;
      this.$store.dispatch('addMailbox', mailboxName)
        .then((response) => {
          const newMailbox = response.data.mailbox;

          this.showCreateMailboxPrompt = false;
          this.pendingCreateMailboxPrompt = false;
          this.$router.push(`/${newMailbox.slug}`);
        })
        .catch((e) => {
          switch (e.response.status) {
            case 422:
              this.errorsCreateMailboxPrompt = errorsHelper(e.response.data.errors);
              break;
            default:
              this.errorsCreateMailboxPrompt = ['An unknown error occured'];
          }

          this.pendingCreateMailboxPrompt = false;
        });
    },

    clearAll() {
      this.$store.dispatch('clearMessagesForMailbox', this.activeMailbox)
        .then(() => {
          this.$router.push(`/${this.activeMailbox.slug}`);
        });
    },
  },

  computed: {
    ...mapGetters(['activeMailbox', 'numberOfMessages']),
    ...mapState({
      avatarUrl: state => state.user.photo_url,
    }),
    cogLink() {
      if (this.$route.name === 'configure_mailbox') {
        return `/${this.activeMailbox.slug}`;
      }

      return { name: 'configure_mailbox', params: { mailbox: this.activeMailbox.slug } };
    },

    configuring() {
      return this.$route.name === 'configure_mailbox';
    },
  },
};
</script>

<style lang="scss">
@import '~@styles/mailmole/variables';

#app-nav {
  position: relative;

  .navbar-brand {
    position: absolute;
    left: 50%;
    width: 150px;
    margin-left: -75px;
  }

  .btn.nav-btn {
    background-color: map-get($theme-colors, "neutral-light") !important;
    height: 36px;

    &.nav-btn-primary {
        background-color: map-get($theme-colors, "primary") !important;
    }
  }

  .btn.nav-btn.configure-btn.configuring {
    background-color: map-get($theme-colors, "neutral-light-inverse") !important;
    color: map-get($theme-colors, "neutral-light") !important;

    &.nav-btn-primary {
        background-color: map-get($theme-colors, "primary") !important;
    }
  }

  .btn.nav-btn:not(.configuring):hover {
    background-color: map-get($theme-colors, "neutral-light") !important;

    &.nav-btn-primary {
        background-color: map-get($theme-colors, "primary") !important;
    }
  }

  #mainMenuDropdownButton {
    min-width: 200px;
  }

  #userMenuDropdownButton {
      color: white;
  }

  .spark-nav-profile-photo {
    height: 30px;
    width: 30px;
  }
}
</style>

<style lang="scss">
#mailbox-list-dropdown .dropdown-menu {
  min-width: 16rem;

}

.navbar-brand {
  svg {
  }
}
</style>

