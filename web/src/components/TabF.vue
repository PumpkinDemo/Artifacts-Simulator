<template>
  <div class="tabF">
    <div class="tabF-tip">{{tip}}</div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { locale } from "../locale";

@Component({})
export default class TabF extends Vue {
    get language(){
        return this.$store.getters.language
    }

    get tip(){
        return locale("Press F to get artifacts", this.language);
    }
}
</script>

<style scoped lang="less">
.tabF {
    &-tip{
        min-height: 100px;
        width: 700px
    }
}
</style>
