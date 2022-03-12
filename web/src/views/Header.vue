<template>
  <header class="header">
    <div class="header-left">
      <el-select
        class="language-select"
        v-model="language"
        :placeholder="locale('Please select a language', language)"
      >
        <el-option
          v-for="item in languageOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
          @change="handleLanguageChange"
        >
        </el-option>
      </el-select>
    </div>

    <div class="header-center">
      <el-select
        class="domain-select"
        v-model="domain"
        :placeholder="locale('Please select a domain', language)"
      >
        <el-option
          v-for="item in domainOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
          @change="handleLanguageChange"
        >
        </el-option>
      </el-select>
    </div>

    <div class="header-right">
    </div>
  </header>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { Select } from "element-ui";
import { Language } from "../type";
import { locale, makeOptions, languageOptions } from "../locale";
import { allDomains } from "../constant";

@Component({})
export default class Header extends Vue {
  languageOptions = languageOptions;

  get domainOptions() {
    return makeOptions(allDomains, this.language);
  }

  locale = locale;

  store = this.$store;

  language = this.store.getters.language;

  domain = "";

  handleLanguageChange(lan: Language) {
    this.store.commit("alterLanguage", lan);
  }
}
</script>

<style scoped lang="less">
.header {
  display: flex;
  flex-direction: row;
  
  &-left{
    flex: 2;
  }

  &-center{
    flex: 10;
  }

  &-right{
    flex: 2;
  }

  .language-select {
    max-width: 150px;
    font-size: 10px;
    margin: 0 20px;
  }

  .domain-select {
    min-width: 500px;
    margin: 0 20px;
  }
}
</style>
