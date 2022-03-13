<template>
  <div class="container">
    <TabF v-if="type === 'TabF'" @pressF="pressF"/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import TabF from "../components/TabF.vue";

@Component({
  components: {
    TabF
  }
})
export default class Container extends Vue {
  type = "TabF"
  pressF(){
    this.type = "ConsumeResin"
  }
}
</script>

<style scoped>
.container {
  min-height: 600px;
  width: 100%;
}
</style>
