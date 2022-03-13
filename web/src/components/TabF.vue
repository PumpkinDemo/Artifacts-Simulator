<template>
  <div class="tabF">
    <div class="tabF-tip">{{tip}}</div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "vue-property-decorator";
import { locale } from "../locale";
import { Language } from "../type";

@Component({})
export default class TabF extends Vue {
    get language(){
        return this.$store.getters.language;
    }

    get tip(){
        return locale("Press F to get artifacts", this.language);
    }

    mounted(){
        document.onkeydown = e => {
            // 兼容
            let e1 = e || event || window.event || arguments.callee.caller.arguments[0]

            if (e1 && e1.keyCode == 70) {
                console.log('Press F');
                this.$emit("pressF");
            }
        }
    }

    beforeDestroy(){
        document.onkeydown = e => {}
    }
}
</script>

<style scoped lang="less">
@keyframes twinkle {
  0% {
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}

.tabF {
    padding-top: 400px;
    &-tip{
        min-height: 30px;
        line-height: 30px;
        font-size: 20px;
        font-weight: bold;
        width: 500px;
        background: rgba(255, 255, 255, .3);
        margin: auto;
        padding: 15px;
        animation: twinkle 2s ease-in-out infinite;
    }
}
</style>
