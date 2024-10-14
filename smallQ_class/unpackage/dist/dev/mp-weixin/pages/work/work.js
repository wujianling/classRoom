"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  data() {
    return {
      messageList: [
        {
          id: 1,
          cover: "http://t15.baidu.com/it/u=3639854910,1717044149&fm=224&app=112&f=JPEG?w=200&h=200",
          msg: "张老师发来一份作业",
          lastTime: "12分钟前",
          unReadCount: 8,
          className: "书画A-01"
        },
        {
          id: 2,
          cover: "http://t15.baidu.com/it/u=3639854910,1717044149&fm=224&app=112&f=JPEG?w=200&h=200",
          msg: "张老师回复了消息",
          lastTime: "2小时前",
          unReadCount: 3,
          className: "书画A-02"
        },
        {
          id: 3,
          cover: "http://t15.baidu.com/it/u=3639854910,1717044149&fm=224&app=112&f=JPEG?w=200&h=200",
          msg: "我发送了一条消息",
          lastTime: "3天前",
          unReadCount: 0,
          className: "书画A-32"
        }
      ]
    };
  },
  methods: {
    goToDetail(id) {
      console.log(1);
      common_vendor.index.navigateTo({
        url: "/pages/work/detail?id=" + id
      });
    }
  }
};
if (!Array) {
  const _easycom_uni_list_chat2 = common_vendor.resolveComponent("uni-list-chat");
  const _easycom_uni_list2 = common_vendor.resolveComponent("uni-list");
  (_easycom_uni_list_chat2 + _easycom_uni_list2)();
}
const _easycom_uni_list_chat = () => "../../uni_modules/uni-list/components/uni-list-chat/uni-list-chat.js";
const _easycom_uni_list = () => "../../uni_modules/uni-list/components/uni-list/uni-list.js";
if (!Math) {
  (_easycom_uni_list_chat + _easycom_uni_list)();
}
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: common_vendor.f($data.messageList, (item, index, i0) => {
      return {
        a: index,
        b: "5c52e54e-2-" + i0 + ",5c52e54e-1",
        c: common_vendor.p({
          to: `/pages/work/detail?id=` + item.id,
          title: item.className,
          avatar: item.cover,
          note: item.msg,
          time: item.lastTime,
          ["show-bage"]: item.unReadCount > 0 ? true : false,
          ["badge-text"]: item.unReadCount
        })
      };
    }),
    b: common_vendor.p({
      border: true
    })
  };
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render]]);
wx.createPage(MiniProgramPage);
