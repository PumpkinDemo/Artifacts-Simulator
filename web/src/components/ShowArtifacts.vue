<template>
  <div class="showArtifacts" @click="showArtifacts">
    <div class="artifacts">
        <div class="artifact" v-for="(a, i) in artifactList">
            <transition name="sudden">
                <div class="artifact-item star5" v-if="artifactType === 'artifacts' || i === focusedIndex" @click.stop="showDetail(i)">
                    <el-tooltip effect="light" placement="top-start" :disabled="artifactType === 'detail'">
                        <template #content>
                            <h4>{{ locale(a.name) }}</h4>
                            <p style="text-indent: 2ch">{{ locale(a.set) + ', ' + locale(a.slot) }}</p>
                            <p style="font-weight: bold">{{ getStatContent(a.mainStat) }}</p>
                            <ul style="margin: 0; padding: 0 2ch">
                                <li v-for="stat in a.subStats">
                                    <p>{{ getStatContent(stat) }}</p>
                                </li>
                            </ul>
                        </template>
                        <img :src="getArtifactImageUrl(a.set, a.slot)"/>
                    </el-tooltip>
                </div>
            </transition>
            <transition name="fade">
                <div class="artifact-detail" v-if="artifactType === 'detail' && i === focusedIndex">
                    <p style="display: inline; background-color: white; color: rgba(0, 0, 0, 0.5); padding: 0 1ch;">
                        {{` LV ${a.level} EXP ${a.exp} `}}
                    </p>
                    <h4>{{ locale(a.name) }}</h4>
                    <p style="text-indent: 2ch">
                        {{ locale(a.set) + ', ' + locale(a.slot) }}
                    </p>
                    <p style="font-weight: bold">{{ getStatContent(a.mainStat) }}</p>
                    <ul style="margin: 0; padding: 0 1ch">
                        <li v-for="stat in a.subStats">
                            <p>{{ getStatContent(stat) }}</p>
                        </li>
                    </ul>
                </div>
            </transition>
        </div>
    </div>
    <transition name="sudden">
        <div class="artifact-enhancing" v-if="artifactType === 'detail'">
            <div class="artifact-enhancing-button pseudobutton" v-if="detailType === 'toEnhance'" @click.stop="enterEnhanceInterface">
                {{ locale("Enhance this artifact") }}
            </div>
            <div class="artifact-enhancing-bar" v-else>
                <div class="enhance-selecter">
                    <el-select
                        class="dogfood-select"
                        v-model="dogfoodStr"
                        size="small"
                    >
                        <el-option
                            v-for="item in dogfoodOptions"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        >
                        </el-option>
                    </el-select>
                    <el-button round @click.stop="quickFill" size="small"> {{ locale("Quick select") }} </el-button>
                </div>
                <div class="enhance-dogfood">
                    <div v-for="(a, i) in dogfood" class="dogfood">
                        <img :class="a ? 'star'+a.stars: ''" :src="dogfoodUrls[i]"/>
                    </div>
                </div>
                <div class="enhance-button pseudobutton" @click.stop="levelUp">
                    <i class="el-icon-loading" v-if="loading"></i>
                    {{ locale("Enhance") }}
                </div>
            </div>
        </div>
    </transition>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { Select, Button } from "element-ui";
import { locale } from "../locale";
import { Language, Artifact, ArtifactStatNonnullable } from "../type";
import { artifactFrom, artifactJSON, getArtifactImageUrl, allSets, allSlots, isPercentageStat } from "../constant";

@Component({})
export default class ShowArtifacts extends Vue {
    artifactType: "artifacts" | "detail"  = "artifacts";

    detailType: "enhancing" | "toEnhance"  = "toEnhance";

    dogfood: any[] = [];

    dogfoodStr = "6*1";

    dogfoodUrls: string[] = [];

    get dogfoodOptions(){
        const locale = this.locale;
        return [
            { value: "6*1", label: locale("Six 1-star artifacts") },
            { value: "6*2", label: locale("Six 2-star artifacts") },
            { value: "6*3", label: locale("Six 3-star artifacts") },
            { value: "6*4", label: locale("Six 4-star artifacts") },
            { value: "6*5.0", label: locale("Six 5-star lv0 artifacts") },
            { value: "6*5.4", label: locale("Six 5-star lv4 artifacts") },
            { value: "6*5.8", label: locale("Six 5-star lv8 artifacts") },
            { value: "6*5.12", label: locale("Six 5-star lv12 artifacts") }
        ]
    }

    getArtifactImageUrl = getArtifactImageUrl;

    focusedIndex = 0;

    loading = false;

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

    created(){
        this.dogfood = Array(6).fill({stars: 1, lv: 0});
        this.dogfoodUrls = this.dogfood.map(a => this.getRandomArtifactImage(a.stars));
    }
    
    getRandomArtifactImage(count: number){
        let star = Math.round(count)
        if (star <= 0)
            star = 1
        else if (star > 5)
            star = 5
        const sets = allSets[star];
        const indexSet = Math.floor(Math.random() * sets.length);
        const set = sets[indexSet];
        const mereCircletSets = [ "Prayers for Illumination", "Prayers for Destiny", "Prayers for Wisdom", "Prayers to Springtime" ]
        const slot = mereCircletSets.includes(set)? "CIRCLET": allSlots[Math.floor(Math.random() * 5)];
        
        return getArtifactImageUrl(set, slot)
    }

    showDetail(i: number){
        console.log("SHOW DETAIL", i)
        this.artifactType = "detail";
        this.focusedIndex = i;
    }

    showArtifacts(){
        if ( this.artifactType === 'artifacts' ){
            this.$emit('getOtherArtifacts')
        } else {
            console.log("SHOW ALL ARTIFACTS")
            this.artifactType = "artifacts";
            this.detailType = "toEnhance";
        }
    }

    enterEnhanceInterface(){
        this.detailType = "enhancing";
    }

    quickFill(){
        const [stars, lv = 0] = this.dogfoodStr.split('*')[1].split('.');
        this.dogfood = Array(6).fill({
            stars: +stars, 
            lv: +lv
        });
        this.dogfoodUrls = this.dogfood.map(a => this.getRandomArtifactImage(a.stars))
        console.log(stars, lv, this.dogfood)
    }

    levelUp(){
        this.loading = true;
        const embryo = artifactJSON(this.artifactList[this.focusedIndex])
        fetch('/enhance', {
            method: 'POST',
            body: JSON.stringify({
                dogFoods: this.dogfood,
                embryo: embryo
            })
        })
        .then(res => res.json())
        .then(res => {
            this.loading = false;
            console.log(res, artifactFrom(res))
            this.updateArtifact(artifactFrom(res));
        })
    }

    updateArtifact(a: Artifact){
        const { locale } = this;
        type s = ArtifactStatNonnullable;
        const embryo = this.artifactList[this.focusedIndex];
        console.log(embryo, a)
        if(embryo.level < a.level){
            const isPercentage = isPercentageStat((a.mainStat as s).statType);
            const round = (value: number) => isPercentage? Math.round(value * 10) / 10: Math.round(value);

            let html = `
                <h4> LV&nbsp;${embryo.level}&nbsp;→&nbsp;LV&nbsp;${a.level}  </h4>
                <h5> ${locale((a.mainStat as s).statType)}:&nbsp;${round(embryo.mainStat.value)}&nbsp;→&nbsp;${round((a.mainStat as s).value)}  </h5>
            `;

            let subStats = (a.subStats as s[]).map((stat) => {
                const originalStat = embryo.subStats.find((s: s) => stat.statType === s.statType);
                let originalValue = originalStat? originalStat.value: 0;
                const isPercentage = isPercentageStat(stat.statType);
                const round = (value: number) => isPercentage? Math.round(value * 10) / 10: Math.round(value);
                const res = {
                    statType: locale(stat.statType),
                    from: round(originalValue) + (isPercentage? '%': ''),
                    to: round(stat.value) + (isPercentage? '%': '')
                }
                console.log(isPercentage, originalValue, round(originalValue), stat.value, round(stat.value), res)
                return res;
            }).filter(a => a.from !== a.to);

            if(subStats.length){
                html += subStats.map(stat => {
                    return `<p> ${stat.statType}:&nbsp;${stat.from}&nbsp;→&nbsp;${stat.to}  </p>`
                }
                ).join('');
            }

            this.$alert(html, locale('Artifact level up'), {
                dangerouslyUseHTMLString: true
            });
        }

        this.artifactList[this.focusedIndex] = a;
    }
}
</script>

<style scoped lang="less">
.sudden-enter-active {
  animation: bounce-in .3s;
}

.sudden-leave-active {
  animation: bounce-in .1s reverse;
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

@keyframes fade {
  0% {
    opacity: 0;
    width: 0;
  }
  100% {
    opacity: 1;
    width: 400px;
  }
}

@keyframes twinkle {
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

.showArtifacts {
    padding: 2ch;
    display: flex;
    flex-direction: column;

    .artifacts{
        display: flex;
        justify-content: center;

        .artifact {
            display: flex;
            flex-wrap: no-wrap;

            &-item {
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
                p, h4 {
                    margin: 1ch 0;
                }
            }
        }
    }

    .artifact-enhancing {
        margin: 10px;
        &-button {
            padding: 15px;
            background: rgba(255, 255, 255, 0.3);
            margin: auto;
            min-height: 30px;
            line-height: 30px;
            font-size: 20px;
            font-weight: bold;
            animation: twinkle 4s ease-in-out infinite;  
            width: 250px; 
        }

        &-bar {
            .enhance-selecter{
                .dogfood-select{
                    width: 200px;
                    margin-right: 10px;
                }
            }
            .enhance-dogfood{
                display: flex;
                justify-content: center;
                .dogfood {
                    margin: 10px;
                    width: 70px;
                    height: 70px;
                    img {
                        width: 70px;
                        height: 70px;
                    }
                }
            }
            .enhance-button{
                border-radius: 15px;
                padding: 15px;
                background: rgba(0, 0, 0, .7);
                color: white;
                margin: auto;
                min-height: 30px;
                line-height: 30px;
                font-size: 20px;
                font-weight: bold;
                width: 250px; 
            }
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
