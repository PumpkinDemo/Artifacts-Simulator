<template>
  <div class="showArtifacts" @click="showArtifacts">
    <div class="artifact" v-for="(a, i) in artifactList" :key="i">
        <transition name="sudden">
            <div class="artifact-item star5" v-if="artifactType === 'artifacts' || i === focusedIndex" @click.stop="showDetail(i)">
                <el-tooltip effect="light" placement="top-start" :disabled="artifactType === 'detail'">
                    <template #content>
                        <h4>{{ locale(a.name) }}</h4>
                        <p style="text-indent: 2ch">{{ locale(a.set) + ' ' + locale(a.slot) }}</p>
                        <p style="font-weight: bold">{{ getStatContent(a.mainStat) }}</p>
                        <ul style="margin: 0; padding: 0 2ch">
                            <li v-for="stat in a.subStats" :key="stat">
                                <p>{{ getStatContent(stat) }}</p>
                            </li>
                        </ul>
                    </template>
                    <img :src="getArtifactImageUrl(a.set, a.slot)"/>
                </el-tooltip>
            </div>
        </transition>
        <transition name="fade">
            <div class="artifact-detail" v-if="artifactType === 'detail' && i === focusedIndex" @click.stop="levelUp" >
                <h4>{{ locale(a.name) }}</h4>
                <p style="text-indent: 2ch">{{ locale(a.set) + ' ' + locale(a.slot) }}</p>
                <p style="font-weight: bold">{{ getStatContent(a.mainStat) }}</p>
                <ul style="margin: 0; padding: 0 2ch">
                    <li v-for="stat in a.subStats" :key="stat">
                        <p>{{ getStatContent(stat) }}</p>
                    </li>
                </ul>
            </div>
        </transition>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { locale } from "../locale";
import { Language, Artifact, ArtifactStatNonnullable } from "../type";
import { artifactJSON, getArtifactImageUrl, allSets, allSlots, isPercentageStat } from "../constant";

@Component({})
export default class ShowArtifacts extends Vue {
    artifactType: "artifacts" | "detail"  = "artifacts";

    getArtifactImageUrl = getArtifactImageUrl;

    focusedIndex = 0;

    get language(){
        return this.$store.getters.language;
    }

    get locale(){
        return (a: string) => locale(a, this.language);
    };

    get artifactList(){
        return this.$store.getters.gainedArtifacts
    }

    get getStatContent(){
        return (stat: ArtifactStatNonnullable) => {
            const statType = locale(stat.statType, this.language);
            const isPercentage = isPercentageStat(stat.statType);
            const value = isPercentage? Math.round(stat.value * 10) / 10: Math.round(stat.value)
            return statType + ': ' + value + (isPercentage? '%': '')
        }
    }

    getRandomArtifactImage(count: number){
        let star = Math.round(count)
        if (star <= 0)
            star = 1
        else if (star > 5)
            star = 5
        const sets = allSets[star];
        const indexSet = Math.floor(Math.random() * sets.length);
        const indexSlot = Math.floor(Math.random() * 5);
        
        return getArtifactImageUrl(sets[indexSet], allSlots[indexSlot])
    }

    showDetail(i: number){
        console.log("SHOW DETAIL", i)
        this.artifactType = "detail";
        this.focusedIndex = i;
    }

    showArtifacts(){
        console.log("SHOW ALL ARTIFACTS")
        this.artifactType = "artifacts";
    }

    levelUp(){
        const embryo = artifactJSON(this.artifactList[this.focusedIndex])
        const dogfoods = new Array(6).fill({stars:5, lv:8})
        console.log(dogfoods)
        fetch('/enhance', {
            method: 'POST',
            body: JSON.stringify({
                dogFoods: dogfoods,
                embryo: embryo
            })
        })
        .then(res => res.json())
        .then(res => {
            console.log(res)
        })
    }
}
</script>

<style scoped lang="less">
.sudden-enter-active {
  animation: bounce-in .3s;
}
.sudden-leave-active {
  animation: bounce-in .3s reverse;
}
@keyframes bounce-in {
  0%, 95% {
    transform: scale(0);
  }
  100% {
    transform: scale(1);
  }
}

.fade-enter-active {
  animation: fade .3s;
}
.fade-leave-active {
  animation: fade .3s reverse;
}
@keyframes fade{
  0% {
    opacity: 0;
    width: 0;
  }
  100% {
    opacity: 1;
    width: 400px;
  }
}

.showArtifacts {
    padding: 4ch;
    display: flex;
    justify-content: center;
    .artifact{
        display: flex;
        flex-wrap: no-wrap;
        &-item{
            height: 100px;
            width: 100px;
            margin: 20px;
            background-color: rgba(255, 255, 255, 0.5);  
            > img {
                height: 100px;
                width: 100px;
            }      
        }
        &-detail {
            width: 400px;
            text-align: left;
            display: block;
            padding: 2ch;
            color: white;
            font-size: large;
            background-color: rgba(0, 0, 0, .5);
        }
    }
}

.star5{
    background-color: rgba(213, 168, 96, .7) !important;
}

.star4{
    background-color: rgba(113, 101, 149, .7) !important;
}

.star3{
    background-color: rgba(121, 161, 190, .7) !important;
}

.star2{
    background-color: rgba(117, 159, 135, .7) !important;
}

.star1{
    background-color: rgba(151, 151, 151, .7) !important;
}
</style>
