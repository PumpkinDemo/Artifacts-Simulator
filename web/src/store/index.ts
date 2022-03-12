import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    language: "EN",
  },
  getters: {
    language: (state) => state.language,
  },
  mutations: {
    alterLanguage: (state, language) => {
      state.language = language;
    },
  },
  actions: {},
  modules: {},
});
