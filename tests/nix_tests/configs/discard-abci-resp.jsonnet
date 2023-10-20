local config = import 'default.jsonnet';

config {
  'highbury_710-1'+: {
    config+: {
      storage: {
        discard_abci_responses: true,
      },
    },
  },
}
