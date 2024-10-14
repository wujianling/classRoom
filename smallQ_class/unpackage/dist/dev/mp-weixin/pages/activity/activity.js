"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  data() {
    return {
      activityList: [
        {
          id: 1,
          cover: "https://p6.itc.cn/q_70/images03/20220415/5aa7b2be438940b28c77167113849c19.png",
          content: "欢乐教育，第三期火热报名了！"
        },
        {
          id: 2,
          cover: "https://img1.baidu.com/it/u=3572012950,3980820241&fm=253&fmt=auto&app=138&f=JPEG?w=1028&h=500",
          content: "开大会了！哈哈哈哈哈。"
        },
        {
          id: 3,
          cover: "https://p6.itc.cn/q_70/images03/20220415/5aa7b2be438940b28c77167113849c19.png",
          content: "欢乐教育，第三期火热报名了！"
        },
        {
          id: 4,
          cover: "https://img1.baidu.com/it/u=3572012950,3980820241&fm=253&fmt=auto&app=138&f=JPEG?w=1028&h=500",
          content: "开大会了！哈哈哈哈哈。"
        },
        {
          id: 5,
          cover: "https://p6.itc.cn/q_70/images03/20220415/5aa7b2be438940b28c77167113849c19.png",
          content: "欢乐教育，第三期火热报名了！"
        },
        {
          id: 6,
          cover: "https://img1.baidu.com/it/u=3572012950,3980820241&fm=253&fmt=auto&app=138&f=JPEG?w=1028&h=500",
          content: "开大会了！哈哈哈哈哈。"
        }
      ]
    };
  },
  methods: {
    onClick(id) {
      console.log("id");
      console.log(id);
    }
  }
};
if (!Array) {
  const _easycom_uni_card2 = common_vendor.resolveComponent("uni-card");
  _easycom_uni_card2();
}
const _easycom_uni_card = () => "../../uni_modules/uni-card/components/uni-card/uni-card.js";
if (!Math) {
  _easycom_uni_card();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: common_vendor.f($data.activityList, (item, index, i0) => {
      return {
        a: item.cover,
        b: common_vendor.t(item.content),
        c: index,
        d: common_vendor.o(($event) => $options.onClick(item.id), index),
        e: "8c5c05e4-0-" + i0,
        f: common_vendor.p({
          cover: item.cover,
          index
        })
      };
    })
  };
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render]]);
wx.createPage(MiniProgramPage);
