<template>
  <div class="container">
    <TabF/>
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
export default class Container extends Vue {}
</script>

<style scoped>
.container {
  min-height: 600px;
}
</style>
