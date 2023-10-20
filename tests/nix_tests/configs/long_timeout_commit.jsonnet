local default = import 'default.jsonnet';

default {
  'highbury_710-1'+: {
    config+: {
      consensus+: {
        timeout_commit: '5s',
      },
    },
  },
}
