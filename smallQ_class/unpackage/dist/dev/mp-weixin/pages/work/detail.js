"use strict";
const common_vendor = require("../../common/vendor.js");
const common_assets = require("../../common/assets.js");
const _sfc_main = {
  data() {
    return {
      type: "center",
      id: 0,
      current: 0,
      showImg: "",
      content: "",
      top: 0,
      scurrent: 0,
      startX: 0,
      // 开始触摸点的X坐标
      startY: 0,
      // 开始触摸点的Y坐标
      mode: "nav",
      dotsStyles: {},
      swiperDotIndex: 0,
      context: "书法作品回顾，回去多加练习",
      imgList: [
        {
          url: "https://preview.qiantucdn.com/58pic/36/84/89/64858PICHz458PIC758PICRf4vT6R_PIC2018.jpg%21w1024_new_3072",
          name: "发送到地方"
        },
        {
          url: "https://img95.699pic.com/desgin_photo/40162/0663_list.jpg%21/fw/431/clip/0x300a0a0",
          name: "地方v从vv从从v"
        },
        {
          url: "https://preview.qiantucdn.com/58pic/36/84/89/64858PICHz458PIC758PICRf4vT6R_PIC2018.jpg%21w1024_new_3072",
          name: "8u拮抗剂拮抗剂"
        },
        {
          url: "https://img95.699pic.com/desgin_photo/40162/0663_list.jpg%21/fw/431/clip/0x300a0a0",
          name: "就看美女"
        },
        {
          url: "https://preview.qiantucdn.com/58pic/36/84/89/64858PICHz458PIC758PICRf4vT6R_PIC2018.jpg%21w1024_new_3072",
          name: "刚发噶得分"
        }
      ],
      messageList: [
        {
          isTeacher: "teacher",
          name: "王老师",
          type: 1,
          //1文字 2图片 3 视频
          content: "啊实打实的啊实打实阿斯顿",
          timer: "2024-12-12 14:14:14"
        },
        {
          isTeacher: "self",
          name: "艾克无锡",
          type: 1,
          //1文字 2图片 3 视频
          content: "sadasd aasdasd asd adas ",
          timer: "2024-12-12 14:14:14"
        },
        {
          isTeacher: "self",
          name: "和摄入",
          type: 2,
          //1文字 2图片 3 视频
          content: "https://preview.qiantucdn.com/58pic/36/84/89/64858PICHz458PIC758PICRf4vT6R_PIC2018.jpg%21w1024_new_3072",
          timer: "2024-12-12 14:14:14"
        }
      ]
    };
  },
  onLoad: function(options) {
    this.id = options.id;
    console.log("Received id:", options.id);
  },
  methods: {
    change(e) {
      this.current = e.detail.current;
    },
    change2() {
    },
    clickItem(e) {
      this.swiperDotIndex = e;
    },
    toggle(type) {
      console.log("1111111111");
      console.log(this.current);
      this.showImg = this.imgList[this.current].url;
      this.type = type;
      this.$refs.popup.open(type);
    },
    send() {
      console.log(111111);
      this.messageList.push({
        content: this.content,
        isTeacher: "self",
        name: "和摄入",
        type: 1,
        timer: ""
      });
      this.content = "";
      this.scrollToBottom();
    },
    chooseImage() {
      common_vendor.index.chooseImage({
        // sourceType: 'album',
        success: (res) => {
          this.messageList.push({
            content: res.tempFilePaths[0],
            isTeacher: "self",
            name: "和摄入",
            type: 2,
            timer: ""
          });
          this.scrollToBottom();
        }
      });
    },
    scrollToBottom() {
      this.top = this.messageList.length * 1e3;
    },
    touchStart(e) {
      this.startX = e.touches[0].pageX;
      this.startY = e.touches[0].pageY;
    },
    // 处理滑动事件
    touchEnd(e) {
      let endX = e.changedTouches[0].pageX;
      let endY = e.changedTouches[0].pageY;
      let deltaX = endX - this.startX;
      let deltaY = endY - this.startY;
      if (Math.abs(deltaX) > Math.abs(deltaY)) {
        if (deltaX < -50) {
          console.log("左滑");
          this.nextImage();
        } else if (deltaX > 50) {
          console.log("右滑");
          this.prevImage();
        }
      }
    },
    // 显示下一张图片
    nextImage() {
      if (this.scurrent < this.imgList.length - 1) {
        this.scurrent++;
      } else {
        this.scurrent = 0;
      }
      this.showImg = this.imgList[this.scurrent].url;
    },
    // 显示上一张图片
    prevImage() {
      if (this.scurrent > 0) {
        this.scurrent--;
      } else {
        this.scurrent = this.imgList.length - 1;
      }
      this.showImg = this.imgList[this.scurrent].url;
    }
  }
};
if (!Array) {
  const _easycom_uni_swiper_dot2 = common_vendor.resolveComponent("uni-swiper-dot");
  const _easycom_uni_popup2 = common_vendor.resolveComponent("uni-popup");
  (_easycom_uni_swiper_dot2 + _easycom_uni_popup2)();
}
const _easycom_uni_swiper_dot = () => "../../uni_modules/uni-swiper-dot/components/uni-swiper-dot/uni-swiper-dot.js";
const _easycom_uni_popup = () => "../../uni_modules/uni-popup/components/uni-popup/uni-popup.js";
if (!Math) {
  (_easycom_uni_swiper_dot + _easycom_uni_popup)();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: $data.imgList.length > 0
  }, $data.imgList.length > 0 ? {
    b: common_vendor.f($data.imgList, (item, index, i0) => {
      return {
        a: common_vendor.o(($event) => $options.toggle("center"), index),
        b: item.url,
        c: index
      };
    }),
    c: common_vendor.o((...args) => $options.change && $options.change(...args)),
    d: $data.swiperDotIndex,
    e: common_vendor.o($options.clickItem),
    f: common_vendor.p({
      info: $data.imgList,
      current: $data.current,
      mode: $data.mode,
      field: "name"
    })
  } : {}, {
    g: $data.showImg,
    h: common_vendor.o((...args) => $options.touchStart && $options.touchStart(...args)),
    i: common_vendor.o((...args) => $options.touchEnd && $options.touchEnd(...args)),
    j: common_vendor.sr("popup", "5233752c-1"),
    k: common_vendor.o($options.change2),
    l: common_vendor.p({
      ["background-color"]: "rgba(0, 0, 0, 0.7)",
      closeOnClickOverlay: true
    }),
    m: common_vendor.t($data.context),
    n: common_vendor.f($data.messageList, (item, index, i0) => {
      return common_vendor.e({
        a: item.type === 2
      }, item.type === 2 ? {
        b: item.content
      } : {
        c: common_vendor.t(item.content)
      }, {
        d: common_vendor.n(item.isTeacher),
        e: index
      });
    }),
    o: $data.top,
    p: $data.content,
    q: common_vendor.o(($event) => $data.content = $event.detail.value),
    r: common_assets._imports_0,
    s: common_vendor.o((...args) => $options.chooseImage && $options.chooseImage(...args)),
    t: common_vendor.o((...args) => $options.send && $options.send(...args))
  });
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-5233752c"]]);
wx.createPage(MiniProgramPage);
