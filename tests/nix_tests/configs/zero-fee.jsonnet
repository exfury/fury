local config = import 'default.jsonnet';

config {
  'highbury_710-1'+: {
    genesis+: {
      app_state+: {
        feemarket+: {
          params+: {
            min_gas_price: '0',
            no_base_fee: true,
            base_fee: '0',            
          },
        },
      },
    },
  },
}
