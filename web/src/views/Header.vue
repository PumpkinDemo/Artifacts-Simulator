<template>
  <header class="header">
    <div class="header-left">
      <el-select
        class="language-select"
        v-model="language"
        :placeholder="locale('Please select a language')"
        @change="handleLanguageChange"
        size="small"
      >
        <el-option
          v-for="item in languageOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        >
        </el-option>
      </el-select>
    </div>

    <div class="header-center">
      <el-select
        class="domain-select"
        v-model="domain"
        :placeholder="locale('Please select a domain', language)"
        @change="handleDomainChange"
        size="small"
      >
        <el-option
          v-for="item in domainOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        >
        </el-option>
      </el-select>
    </div>

    <div class="header-right">
    </div>
  </header>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "vue-property-decorator";
import { Select } from "element-ui";
import { Language, ArtifactDomain } from "../type";
import { locale, makeOptions, languageOptions, availableLanguages } from "../locale";
import { allDomains } from "../constant";

@Component({})
export default class Header extends Vue {
  languageOptions = languageOptions;

  language = this.$store.getters.language;

  domain = "";

  @Watch("$store.getters.domain")
  changeDomain(domain: ArtifactDomain){
    this.domain = domain;
  }

  get domainOptions() {
    return makeOptions(allDomains, this.language);
  }

  get locale(){
      return (a: string) => locale(a, this.language);
  };

  lower(domain: string){
    return domain.toLowerCase().replace(/[\s']/g, '')
  }

  created(){
    let { lang = "EN", domain = "" } = this.$route.query;

    lang = (lang as string).toLowerCase();
    const languages = availableLanguages.map(a => a.toLowerCase());
    if(!languages.includes(lang)){
      lang = "EN";
    }
    lang = lang.toUpperCase();
    this.language = lang;

    domain = (domain as string).toLowerCase();
    const domains = allDomains.map(this.lower);
    if(domains.includes(domain)){
      domain = allDomains.find(a => this.lower(a) === domain) as string;
    }
    this.domain = domain as string;
    this.$store.commit("alterDomain", domain);
  }

  handleLanguageChange(lan: Language) {
    this.$store.commit("alterLanguage", lan);
    this.$router.push({
      query: Object.assign(JSON.parse(JSON.stringify(this.$route.query)), {lang: lan})
    })
  }

  handleDomainChange(domain: string){
    this.$store.commit("alterDomain", domain);
    this.$router.push({
      query: Object.assign(JSON.parse(JSON.stringify(this.$route.query)), {domain: this.lower(domain)})
    })
  }
}
</script>

<style scoped lang="less">
.header {
  display: flex;
  flex-direction: row;
  padding: 20px 0;
  
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
