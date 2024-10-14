"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  data() {
    return {
      list: [
        {
          time: "6月12日 14:30",
          roomName: "书法001-x",
          duration: 120,
          typer: "书法",
          tearcher: {
            name: "老王",
            id: 1
          },
          content: "没有任何备注"
        },
        {
          time: "",
          roomName: "书法001-x",
          duration: 120,
          typer: "书法",
          tearcher: {
            name: "老王",
            id: 1
          }
        },
        {
          time: "",
          roomName: "书法001-x",
          duration: 120,
          typer: "书法",
          tearcher: {
            name: "老王",
            id: 1
          }
        }
      ]
    };
  },
  methods: {}
};
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: common_vendor.f($data.list, (item, index, i0) => {
      return {
        a: common_vendor.t(item.typer),
        b: common_vendor.t(item.roomName),
        c: common_vendor.t(item.time),
        d: common_vendor.t(item.duration),
        e: common_vendor.t(item.tearcher.name),
        f: common_vendor.t(item.content),
        g: index
      };
    })
  };
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render]]);
wx.createPage(MiniProgramPage);
