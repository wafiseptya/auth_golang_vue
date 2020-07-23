import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { router } from './routes'
import store from './store';
// import VeeValidate from 'vee-validate';

Vue.config.productionTip = false

// Vue.use(VeeValidate);

new Vue({
  router,
  vuetify,
  store,
  render: h => h(App)
}).$mount('#app')
