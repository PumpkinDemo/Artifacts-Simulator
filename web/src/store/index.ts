import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    language: "EN",
    domain: "",
    gainedArtifacts: []
  },
  getters: {
    language: state => state.language,
    domain: state => state.domain,
    gainedArtifacts: state => state.gainedArtifacts
  },
  mutations: {
    alterLanguage: (state, language) => {
      console.log("ALTER LANGUAGE TO", language)
      state.language = language;
    },
    alterDomain: (state, domain) => {
      console.log("ALTER DOMAIN TO", domain)
      state.domain = domain;
    },
    alterGainedArtifacts: (state, artifacts) => {
      console.log(`GAIN ${artifacts.length} ARTIFACTS`)
      state.gainedArtifacts = artifacts;
    }
  },
  actions: {},
  modules: {},
});
