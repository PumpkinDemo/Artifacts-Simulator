<template>
  <div class="container">
    <TabF v-if="type === 'TabF'" @pressF="pressF"/>
    <UseResin v-else-if="type === 'ConsumeResin'" @showArtifacts="showArtifacts"/>
    <ShowArtifacts v-else-if="type === 'ShowArtifacts'"/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import TabF from "../components/TabF.vue";
import UseResin from "../components/UseResin.vue";
import ShowArtifacts from "../components/ShowArtifacts.vue";

type ComponentType = "TabF" | "ConsumeResin" | "ShowArtifacts" | "EnhanceArtifact";
@Component({
  components: {
    TabF, UseResin, ShowArtifacts
  }
})
export default class Container extends Vue {
  type: ComponentType = "TabF"
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
