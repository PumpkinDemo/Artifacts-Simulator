<template>
  <div class="useResin">
    <div class="useResin-block useResin-block1 pseudobutton" @click="useResin(20)">
      {{ text20Resin }}
      <img src="../assets/FragileResin.png"/>
    </div>
    <div class="useResin-block useResin-block2 pseudobutton" @click="useResin(40)">
      {{ text40Resin }}
      <img src="../assets/CondensedResin.png"/>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "vue-property-decorator";
import { locale } from "../locale";
import { Language } from "../type";
import { artifactFrom } from "../constant";

@Component({})
export default class UseResin extends Vue {
  get language() {
    return this.$store.getters.language;
  }

  get text20Resin() {
    return locale("Use 20 fragile resins", this.language);
  }

  get text40Resin() {
    return locale("Use a condensed resin", this.language);
  }

  useResin(resins: 20 | 40) {
    console.log("Use resin", resins)
    let domain = this.$store.getters.domain

    if(!domain){
      domain = 'Domain of Guyun';
      this.$store.commit('alterDomain', domain);
    }
    
    fetch('/gain', {
      method: 'post',
      body: JSON.stringify({ resin: resins, domain:domain })
    })
    .then(res => res.json())
    .then(res => {
      this.$store.commit('alterGainedArtifacts', res.map(artifactFrom));
      console.log(res)
      this.$emit('showArtifacts')
    })
  }
}
</script>

<style scoped lang="less">
@keyframes twinkle1 {
  0% {
    opacity: 0.5;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0.5;
  }
}

@keyframes twinkle2 {
  0% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
  100% {
    opacity: 1;
  }
}

.useResin {
  padding-top: 400px;
  display: flex;
  &-block {
    display: flex;
    justify-content: center;
    min-height: 30px;
    line-height: 30px;
    font-size: 20px;
    font-weight: bold;
    min-width: 300px;
    background: rgba(255, 255, 255, 0.3);
    margin: auto;
    padding: 15px;
    img {
      height: 30px;
      width: 30px;
    }
  }
  &-block1 {
    margin-right: 50px;
    animation: twinkle1 4s ease-in-out infinite;
  }
  &-block2 {
    animation: twinkle2 4s ease-in-out infinite;
  }
}
</style>