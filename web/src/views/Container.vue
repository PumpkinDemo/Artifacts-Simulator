<template>
  <div class="container">
    <TabF v-if="type === 'TabF'" @pressF="pressF"/>
    <UseResin v-else-if="type === 'ConsumeResin'" @showArtifacts="showArtifacts"/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import TabF from "../components/TabF.vue";
import UseResin from "../components/UseResin.vue"

type ContanerType = "TabF" | "ConsumeResin" | "ShowArtifacts" | "EnhanceArtifact";
@Component({
  components: {
    TabF, UseResin
  }
})
export default class Container extends Vue {
  type: ContanerType = "TabF"
  pressF(){
    this.type = "ConsumeResin"
  }
  showArtifacts(){
    this.type = "ShowArtifacts"
  }
}
</script>

<style scoped>
.container {
  min-height: 600px;
  width: 100%;
}
</style>
