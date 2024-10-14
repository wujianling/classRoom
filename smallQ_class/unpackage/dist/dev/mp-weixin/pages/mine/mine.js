"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  data() {
    return {
      name: "爱无悔",
      avatar: "http://img.duoziwang.com/2021/02/1621261310479632.jpg",
      studentInfo: {
        name: "赵小月"
      }
    };
  },
  methods: {}
};
if (!Array) {
  const _easycom_uni_icons2 = common_vendor.resolveComponent("uni-icons");
  _easycom_uni_icons2();
}
const _easycom_uni_icons = () => "../../uni_modules/uni-icons/components/uni-icons/uni-icons.js";
if (!Math) {
  _easycom_uni_icons();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: !$data.avatar
  }, !$data.avatar ? {} : {}, {
    b: $data.avatar
  }, $data.avatar ? {
    c: common_vendor.o((...args) => _ctx.handleToAvatar && _ctx.handleToAvatar(...args)),
    d: $data.avatar
  } : {}, {
    e: !$data.name
  }, !$data.name ? {
    f: common_vendor.o((...args) => _ctx.handleToLogin && _ctx.handleToLogin(...args))
  } : {}, {
    g: $data.name
  }, $data.name ? {
    h: common_vendor.t($data.name)
  } : {}, {
    i: $data.studentInfo.name
  }, $data.studentInfo.name ? {
    j: common_vendor.t($data.studentInfo.name),
    k: common_vendor.p({
      type: "search",
      size: "20"
    }),
    l: common_vendor.o((...args) => _ctx.handleToInfo && _ctx.handleToInfo(...args))
  } : {}, {
    m: common_vendor.o((...args) => _ctx.handleJiaoLiuQun && _ctx.handleJiaoLiuQun(...args)),
    n: common_vendor.o((...args) => _ctx.handleJiaoLiuQun && _ctx.handleJiaoLiuQun(...args)),
    o: common_vendor.o((...args) => _ctx.handleBuilding && _ctx.handleBuilding(...args)),
    p: common_vendor.o((...args) => _ctx.handleToEditInfo && _ctx.handleToEditInfo(...args)),
    q: common_vendor.o((...args) => _ctx.handleAbout && _ctx.handleAbout(...args)),
    r: common_vendor.o((...args) => _ctx.handleToSetting && _ctx.handleToSetting(...args))
  });
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render]]);
wx.createPage(MiniProgramPage);
