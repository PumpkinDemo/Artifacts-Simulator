<template>
  <div>
    <Header />
    <Container />
    <Footer />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Header from "./Header.vue";
import Container from "./Container.vue";
import Footer from "./Footer.vue";

@Component({
  components: { Header, Container, Footer },
})
export default class HomeView extends Vue {}
</script>

<style scoped>
#homeview {
}
</style>
