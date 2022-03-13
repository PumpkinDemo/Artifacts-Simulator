import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    language: "EN",
    domain: ""
  },
  getters: {
    language: state => state.language,
    domain: state => state.domain
  },
  mutations: {
    alterLanguage: (state, language) => {
      console.log("ALTER LANGUAGE TO", language)
      state.language = language;
    },
    alterDomain: (state, domain) => {
      console.log("ALTER DOMAIN TO", domain)
      state.domain = domain;
    }
  },
  actions: {},
  modules: {},
});
