<template>
  <div class="showArtifacts">
    <img v-for="(a, i) in artifactList" :src="getArtifactImageUrl(a.set, a.slot)"/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { locale } from "../locale";
import { Language } from "../type";
import { getArtifactImageUrl, allSets, allSlots } from "../constant";

@Component({})
export default class ShowArtifacts extends Vue {
    artifactList: any = [];

    getArtifactImageUrl = getArtifactImageUrl;

    get language(){
        return this.$store.getters.language;
    }
    
    created(){
        this.artifactList = allSets.map((set, index) => {
            return {
                set,
                slot: allSlots[index % 5]
            }
        })
    }
}
</script>

<style scoped lang="less">

.ShowArtifacts {

}
</style>
