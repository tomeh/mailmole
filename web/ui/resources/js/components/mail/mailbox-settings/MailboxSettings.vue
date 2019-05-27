<template>
  <div class="container-fluid h-100">
    <div class="mailbox-settings h-100 p-3 row">
      <div class="col spark-settings-tabs">

        <aside class="mb-4">
          <router-link class="back-to-mailbox" :to="`/${mailbox.slug}`"> <!-- Named route doesn't show empty -->
            <i class="far fa-long-arrow-left"></i><span class="ml-2">Back to mailbox</span>
          </router-link>
        </aside>

        <aside>
          <h3 class="nav-heading">Settings</h3> 

          <ul class="nav co-nav flex-column mb-4">
            <li class="nav-item">
              <router-link class="nav-link" 
                :class="{ active: $route.name === 'configure_mailbox' }"
                :to="{ name: 'configure_mailbox' }"><i class="fas fa-fw fa-envelope"></i> Mailbox</router-link>
            </li>

            <li class="nav-item">
              <router-link class="nav-link" 
                :class="{ active: $route.name === 'configure_integrations' }"
                :to="{ name: 'configure_integrations' }"><i class="fas fa-fw fa-code"></i> Integrations</router-link>
            </li>

            <li class="nav-item">
              <router-link class="nav-link" 
                :class="{ active: $route.name === 'configure_users' }"
                :to="{ name: 'configure_users' }"><i class="fas fa-fw fa-users"></i> Users</router-link>
            </li>

          </ul>
        </aside>
      </div>

      <div class="col-9">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    mailbox() {
      return this.$store.getters.activeMailbox;
    },
  },
};
</script>

<style lang="scss">
  .mailbox-settings {
    overflow: scroll;
    .mailbox-settings__error {
      color: red;
    }

    .back-to-mailbox {
      text-decoration: none;
    }
  }

  .nav-item {
    i {
      margin-right: 1em;
    }
  }
</style>
