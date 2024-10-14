<template>
	<view class="content">
		<uni-swiper-dot v-if="imgList.length>0" class="uni-swiper-dot-box" @clickItem=clickItem :info="imgList"
			:current="current" :mode="mode" field="name">
			<swiper @change="change" :current="swiperDotIndex">
				<swiper-item v-for="(item, index) in imgList" :key="index">
					<view>
						<image mode="aspectFill" @click="toggle('center')" :src="item.url">
						</image>
					</view>
				</swiper-item>
			</swiper>
		</uni-swiper-dot>
		<view>
					<uni-popup ref="popup" background-color="rgba(0, 0, 0, 0.7)" @change="change2" :closeOnClickOverlay="true">
					    <view class="popup-content" @touchstart="touchStart" @touchend="touchEnd">
					      <image mode="aspectFit" :src="showImg" class="popup-image"></image>
					    </view>
					  </uni-popup>
				</view>
		<view>
			内容：{{context}}
		</view>
		<view class="page">
			<scroll-view class="scroll-view" scroll-y scroll-with-animation :scroll-top="top">
				<view style="padding: 30rpx 30rpx 240rpx;">
					<view class="message" :class="[item.isTeacher]" v-for="(item,index) in messageList" :key="index">
						<view class="content" v-if="item.type === 2">
							<image :src="item.content" mode="widthFix"></image>
						</view>
						<view class="content" v-else>
							{{ item.content }}
						</view>
					</view>
				</view>
			</scroll-view>
			<view class="tool">
				<input type="text" v-model="content" class="input" />
				<image src="../../static/thumb.png" mode="widthFix" class="thumb" @click="chooseImage"></image>
				<button @click="send" type="primary" size="mini">发送</button>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				type: 'center',
				id: 0,
				current: 0,
				showImg:"",
				content: '',
				top: 0,
				scurrent: 0,
				startX: 0, // 开始触摸点的X坐标
				startY: 0, // 开始触摸点的Y坐标
				mode: 'nav',
				dotsStyles: {},
				swiperDotIndex: 0,
				context: "书法作品回顾，回去多加练习",
				imgList: [{
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
				messageList: [{
						isTeacher: "teacher",
						name: "王老师",
						type: 1, //1文字 2图片 3 视频
						content: "啊实打实的啊实打实阿斯顿",
						timer: "2024-12-12 14:14:14"
					},
					{
						isTeacher: "self",
						name: "艾克无锡",
						type: 1, //1文字 2图片 3 视频
						content: "sadasd aasdasd asd adas ",
						timer: "2024-12-12 14:14:14"
					},
					{
						isTeacher: "self",
						name: "和摄入",
						type: 2, //1文字 2图片 3 视频
						content: "https://preview.qiantucdn.com/58pic/36/84/89/64858PICHz458PIC758PICRf4vT6R_PIC2018.jpg%21w1024_new_3072",
						timer: "2024-12-12 14:14:14"
					}
				]
			}
		},
		onLoad: function(options) {
			// options 是一个包含页面传递参数的对象
			this.id = options.id
			console.log('Received id:', options.id);

			// 你可以根据需要使用这个参数
			// 比如，发送网络请求来获取对应 id 的数据
		},
		methods: {
			change(e) {
				this.current = e.detail.current
			},
			change2(){
				
			},
			clickItem(e) {
			
				this.swiperDotIndex = e
			},
			toggle(type) {
				console.log("1111111111")
				
				console.log(this.current)
				this.showImg=this.imgList[this.current].url
							this.type = type
							// open 方法传入参数 等同在 uni-popup 组件上绑定 type属性
							this.$refs.popup.open(type)
						},
			send() {
				console.log(111111)
				this.messageList.push({
					content: this.content,
					isTeacher: 'self',
					name: "和摄入",
					type: 1, 
					timer: "",
				})
				this.content = ''
				this.scrollToBottom()

			},
			chooseImage() {
				uni.chooseImage({
					// sourceType: 'album',
					success: (res) => {
						this.messageList.push({
							content: res.tempFilePaths[0],
							isTeacher: 'self',
							name: "和摄入",
							type: 2, 
							timer: "",
						})
						this.scrollToBottom()
					
					}
				})
			},
			scrollToBottom() {
				this.top = this.messageList.length * 1000
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
						
			     // 判断是否为水平滑动（避免误触发）
			     if (Math.abs(deltaX) > Math.abs(deltaY)) {
			       if (deltaX < -50) {
			         // 左滑，下一张
					 console.log("左滑")
			         this.nextImage();
			       } else if (deltaX > 50) {
			         // 右滑，上一张
					  console.log("右滑")
			         this.prevImage();
			       }
			     }
			   },
						
			   // 显示下一张图片
			   nextImage() {
			     if (this.scurrent < this.imgList.length - 1) {
			       this.scurrent++;
			     } else {
			       this.scurrent = 0; // 循环回到第一张
			     }
			     this.showImg = this.imgList[this.scurrent].url;
			   },
						
			   // 显示上一张图片
			   prevImage() {
			     if (this.scurrent > 0) {
			       this.scurrent--;
			     } else {
			       this.scurrent = this.imgList.length - 1; // 循环回到最后一张
			     }
			     this.showImg = this.imgList[this.scurrent].url;
			   }
			 
		}
	}
</script>

<style lang="scss" scoped>
	.scroll-view {
		/* #ifdef H5 */
		height: calc(100vh - 44px);
		/* #endif */
		/* #ifndef H5 */
		height: 100vh;
		/* #endif */
		background: #eee;
		box-sizing: border-box;
	}

	.message {
		display: flex;
		align-items: flex-start;
		margin-bottom: 30rpx;

		.avatar {
			width: 80rpx;
			height: 80rpx;
			border-radius: 10rpx;
			margin-right: 30rpx;
		}

		.content {
			min-height: 80rpx;
			max-width: 60vw;
			box-sizing: border-box;
			font-size: 28rpx;
			line-height: 1.3;
			padding: 20rpx;
			border-radius: 10rpx;
			background: #fff;

			image {
				width: 200rpx;
			}
		}

		&.self {
			justify-content: flex-end;

			.avatar {
				margin: 0 0 0 30rpx;
			}

			.content {
				position: relative;

				&::after {
					position: absolute;
					content: '';
					width: 0;
					height: 0;
					border: 16rpx solid transparent;
					border-left: 16rpx solid #fff;
					right: -28rpx;
					top: 24rpx;
				}
			}
		}

		&.teacher {
			.content {
				position: relative;

				&::after {
					position: absolute;
					content: '';
					width: 0;
					height: 0;
					border: 16rpx solid transparent;
					border-right: 16rpx solid #fff;
					left: -28rpx;
					top: 24rpx;
				}
			}
		}
	}

	.tool {
		position: fixed;
		width: 100%;
		min-height: 120rpx;
		left: 0;
		bottom: 0;
		background: #fff;
		display: flex;
		align-items: flex-start;
		box-sizing: border-box;
		padding: 20rpx 24rpx 20rpx 40rpx;
		padding-bottom: calc(20rpx + constant(safe-area-inset-bottom)/2) !important;
		padding-bottom: calc(20rpx + env(safe-area-inset-bottom)/2) !important;

		.input {
			background: #eee;
			border-radius: 10rpx;
			height: 70rpx;
			margin-right: 30rpx;
			flex: 1;
			padding: 0 20rpx;
			box-sizing: border-box;
			font-size: 28rpx;
		}

		.thumb {
			width: 64rpx;
		}
		
	}
	.popup-content {
	  width: 100vw; /* 全屏宽度 */
	  height: 100vh; /* 全屏高度 */
	  display: flex;
	  justify-content: center;
	  align-items: center;
	}
	
	.popup-image {
	  width: 100%; /* 保证图片按比例缩放 */
	  max-height: 100%;
	}
</style>