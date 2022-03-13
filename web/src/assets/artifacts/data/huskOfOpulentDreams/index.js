import goblet from "./goblet.png";
import plume from "./plume.png";
import flower from "./flower.png";
import circlet from "./circlet.png";
import sands from "./sands.png";

export default {
  flower: {
    chs: "荣花之期",
    url: flower,
  },
  plume: {
    chs: "华馆之羽",
    url: plume,
  },
  sands: {
    chs: "众生之谣",
    url: sands,
  },
  goblet: {
    chs: "梦醒之瓢",
    url: goblet,
  },
  circlet: {
    chs: "形骸之笠",
    url: circlet,
  },
  chs: "华馆梦醒形骸记",
  eng: "huskOfOpulentDreams",
  minStar: 4,
  maxStar: 5,
  effectText: {
    chs: {
      2: "防御力提高30%。",
      4: "装备此圣遗物套装的角色在以下情况下，将获得「问答」效果：在场上用岩元素攻击命中敌人后获得一层，每0.3秒至多触发一次；在队伍后台中，每3秒获得一层。问答至多叠加4层，每层能提供6%防御力与6%岩元素伤害加成。每6秒，若未获得问答效果，将损失一层。",
    },
  },
};
