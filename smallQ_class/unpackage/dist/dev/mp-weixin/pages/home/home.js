"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  data() {
    return {
      coverList: [
        {
          url: "https://editor-img.888ban.com/gif_complete_img/2021-07-22/small_ac13fea4de624388803bca40ed36807f.png?auth_key=2285856000-0-0-736166abe8ebbf358a21073531e7075c"
        },
        {
          url: "https://preview.qiantucdn.com/58pic/36/84/89/64858PICHz458PIC758PICRf4vT6R_PIC2018.jpg%21w1024_new_3072"
        },
        {
          url: "https://img95.699pic.com/desgin_photo/40162/0663_list.jpg%21/fw/431/clip/0x300a0a0"
        }
      ],
      notice: "新校区，新起点，大家一起冲！",
      classList: [
        {
          id: 1,
          name: "托班",
          url: "https://nimg.ws.126.net/?url=http%3A%2F%2Fdingyue.ws.126.net%2F2024%2F0716%2F61aab088j00sgpro2000vd20046005jg0046005j.jpg&thumbnail=660x2147483647&quality=80&type=jpg"
        },
        {
          id: 2,
          name: "美术",
          url: "https://yikaow.com/images/tr2.jpg"
        },
        {
          id: 3,
          name: "陶艺",
          url: "https://wx4.sinaimg.cn/thumb150/a821da45ly1fksfoz2sfhj20qo0zktfn.jpg"
        },
        {
          id: 4,
          name: "阅读",
          url: "http://t15.baidu.com/it/u=1624261450,1270904063&fm=224&app=112&f=JPEG?w=155&h=155"
        }
      ],
      imageList: [{
        url: ""
      }],
      class2List: [
        {
          id: 1,
          name: "陶艺A-01",
          content: "动手能力超强的一个班",
          cover: "https://img1.baidu.com/it/u=1884410828,871274811&fm=253&fmt=auto&app=138&f=JPEG?w=750&h=500"
        },
        {
          id: 2,
          name: "陶艺A-02",
          content: "动手能力超强的一个班",
          cover: "https://img1.baidu.com/it/u=1884410828,871274811&fm=253&fmt=auto&app=138&f=JPEG?w=750&h=500"
        },
        {
          id: 3,
          name: "陶艺B-03",
          content: "动手能力超强的一个班",
          cover: "https://img1.baidu.com/it/u=1884410828,871274811&fm=253&fmt=auto&app=138&f=JPEG?w=750&h=500"
        },
        {
          id: 4,
          name: "陶艺C-01",
          content: "动手能力超强的一个班",
          cover: "https://img1.baidu.com/it/u=1884410828,871274811&fm=253&fmt=auto&app=138&f=JPEG?w=750&h=500"
        }
      ],
      background: ["color1", "color2", "color3"],
      indicatorDots: true,
      autoplay: true,
      interval: 3e3,
      duration: 1e3
    };
  },
  methods: {}
};
if (!Array) {
  const _easycom_uni_notice_bar2 = common_vendor.resolveComponent("uni-notice-bar");
  const _easycom_uni_grid_item2 = common_vendor.resolveComponent("uni-grid-item");
  const _easycom_uni_grid2 = common_vendor.resolveComponent("uni-grid");
  const _easycom_uni_card2 = common_vendor.resolveComponent("uni-card");
  const _easycom_uni_section2 = common_vendor.resolveComponent("uni-section");
  (_easycom_uni_notice_bar2 + _easycom_uni_grid_item2 + _easycom_uni_grid2 + _easycom_uni_card2 + _easycom_uni_section2)();
}
const _easycom_uni_notice_bar = () => "../../uni_modules/uni-notice-bar/components/uni-notice-bar/uni-notice-bar.js";
const _easycom_uni_grid_item = () => "../../uni_modules/uni-grid/components/uni-grid-item/uni-grid-item.js";
const _easycom_uni_grid = () => "../../uni_modules/uni-grid/components/uni-grid/uni-grid.js";
const _easycom_uni_card = () => "../../uni_modules/uni-card/components/uni-card/uni-card.js";
const _easycom_uni_section = () => "../../uni_modules/uni-section/components/uni-section/uni-section.js";
if (!Math) {
  (_easycom_uni_notice_bar + _easycom_uni_grid_item + _easycom_uni_grid + _easycom_uni_card + _easycom_uni_section)();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: common_vendor.f($data.coverList, (item, index, i0) => {
      return {
        a: item.url,
        b: index
      };
    }),
    b: _ctx.heightFix,
    c: $data.indicatorDots,
    d: $data.autoplay,
    e: $data.interval,
    f: $data.duration,
    g: common_vendor.p({
      ["show-icon"]: true,
      scrollable: true,
      text: $data.notice
    }),
    h: common_vendor.f($data.classList, (item, index, i0) => {
      return {
        a: item.url,
        b: common_vendor.t(item.name),
        c: index,
        d: "bd3213e4-2-" + i0 + ",bd3213e4-1",
        e: common_vendor.p({
          index
        })
      };
    }),
    i: common_vendor.p({
      column: 4,
      highlight: true
    }),
    j: common_vendor.f($data.class2List, (item, index, i0) => {
      return {
        a: common_vendor.t(item.name),
        b: item.cover,
        c: common_vendor.t(item.content),
        d: index,
        e: "bd3213e4-4-" + i0 + ",bd3213e4-3",
        f: common_vendor.p({
          thumbnail: _ctx.avatar,
          index
        })
      };
    }),
    k: common_vendor.p({
      title: "班级列表",
      type: "line"
    })
  };
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render]]);
wx.createPage(MiniProgramPage);
