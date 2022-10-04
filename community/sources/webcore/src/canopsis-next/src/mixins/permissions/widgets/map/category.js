import { USERS_PERMISSIONS } from '@/constants';

import { authMixin } from '@/mixins/auth';

export const permissionsWidgetsMapCategory = {
  mixins: [authMixin],
  computed: {
    hasAccessToCategory() {
      return this.checkAccess(USERS_PERMISSIONS.business.map.actions.category);
    },
  },
};
