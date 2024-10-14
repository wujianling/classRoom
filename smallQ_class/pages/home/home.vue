<template>
	<view>
		<view class="uni-margin-wrap">
			<swiper class="swiper" circular :indicator-dots="indicatorDots" :autoplay="autoplay" :interval="interval"
				:duration="duration">
				<swiper-item v-for="(item,index) in coverList" :key="index">
					<view class="swiper-item uni-bg-red">
						<image style="" :mode="heightFix" :src="item.url">
						</image>
					</view>
				</swiper-item>

			</swiper>
		</view>
		<view>
			<uni-notice-bar show-icon scrollable :text="notice" />
		</view>
		<view>
			<swiper class="swiper2">
				<swiper-item >
					<uni-grid :column="4" :highlight="true">
						<uni-grid-item v-for="(item, index) in classList" :index="index" :key="index">
							<view class="grid-item-box">
								<image :src="item.url" class="image" mode="aspectFill" />
								<text class="text">{{ item.name }}</text>
							</view>
						</uni-grid-item>
					</uni-grid>
				</swiper-item>
			</swiper>
		</view>
		<view>
			<uni-section title="班级列表" type="line">
				<uni-card    :thumbnail="avatar" v-for="(item, index) in class2List" :index="index" :key="index">
					<template v-slot:title>
						{{item.name}}
					</template>
					<image style="width: 100%; height: 360rpx;" :src="item.cover"></image>
					<text
						class="uni-body uni-mt-5">{{item.content}}</text>
					
				</uni-card>
			</uni-section>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				coverList: [{
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
				classList: [{
						id: 1,
						name: "托班",
						url: "https://nimg.ws.126.net/?url=http%3A%2F%2Fdingyue.ws.126.net%2F2024%2F0716%2F61aab088j00sgpro2000vd20046005jg0046005j.jpg&thumbnail=660x2147483647&quality=80&type=jpg",
					},
					{

						id: 2,
						name: "美术",
						url: "https://yikaow.com/images/tr2.jpg",
					},
					{
						id: 3,
						name: "陶艺",
						url: "https://wx4.sinaimg.cn/thumb150/a821da45ly1fksfoz2sfhj20qo0zktfn.jpg",
					}, {
						id: 4,
						name: "阅读",
						url: "http://t15.baidu.com/it/u=1624261450,1270904063&fm=224&app=112&f=JPEG?w=155&h=155",
					}
				],
				imageList: [{
					url: ""
				}],
				class2List:[
					{
						id:1,
						name:"陶艺A-01",
						content:"动手能力超强的一个班",
						cover:"https://img1.baidu.com/it/u=1884410828,871274811&fm=253&fmt=auto&app=138&f=JPEG?w=750&h=500"
					},
					{
						id:2,
						name:"陶艺A-02",
						content:"动手能力超强的一个班",
						cover:"https://img1.baidu.com/it/u=1884410828,871274811&fm=253&fmt=auto&app=138&f=JPEG?w=750&h=500"
					},{
						id:3,
						name:"陶艺B-03",
						content:"动手能力超强的一个班",
						cover:"https://img1.baidu.com/it/u=1884410828,871274811&fm=253&fmt=auto&app=138&f=JPEG?w=750&h=500"
					},
					{
						id:4,
						name:"陶艺C-01",
						content:"动手能力超强的一个班",
						cover:"https://img1.baidu.com/it/u=1884410828,871274811&fm=253&fmt=auto&app=138&f=JPEG?w=750&h=500"
					}
				],
				background: ['color1', 'color2', 'color3'],
				indicatorDots: true,
				autoplay: true,
				interval: 3000,
				duration: 1000
			}
		},
		methods: {

		}
	}
</script>

<style lang="scss">
	.uni-margin-wrap {
		width: 690rpx;
		width: 100%;
	}

	.swiper {
		height: 300rpx;
	}
	.swiper2{
		height: 200rpx;
	}
	.swiper-item {
		display: block;
		height: 300rpx;
		line-height: 300rpx;
		text-align: center;
	}

	.swiper-list {
		margin-top: 40rpx;
		margin-bottom: 0;
	}

	.uni-common-mt {
		margin-top: 60rpx;
		position: relative;
	}

	.image {
		width: 100rpx;
		height: 100rpx;
	}

	.text {
		font-size: 14px;
		margin-top: 5px;
	}

	.info {
		position: absolute;
		right: 20rpx;
	}

	.grid-item-box {
		flex: 1;
		// position: relative;
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 5px 0;
	}

	.uni-padding-wrap {
		width: 550rpx;
		padding: 0 100rpx;
	}
</style>