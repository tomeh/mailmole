import MailboxSettings from '@/components/mail/mailbox-settings/MailboxSettings';
import Main from '@/components/mail/mailbox-settings/components/Main';
import Integrations from '@/components/mail/mailbox-settings/components/Integrations';
import Users from '@/components/mail/mailbox-settings/components/Users';

export default {
  path: '/:mailbox/configure',
  component: MailboxSettings,

  children: [
    {
      path: '/:mailbox/configure',
      name: 'configure_mailbox',
      component: Main,
    }, {
      path: '/:mailbox/configure/integrations',
      name: 'configure_integrations',
      component: Integrations,
    }, {
      path: '/:mailbox/configure/users',
      name: 'configure_users',
      component: Users,
    },
  ],
};
