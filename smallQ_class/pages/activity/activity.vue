<template>
	<view class="container">
		<uni-card :cover="item.cover" v-for="(item, index) in activityList" :index="index" :key="index"
			@click="onClick(item.id)">
			<image slot="cover" style="width: 100%" :src="item.cover"></image>
			<text class="uni-body">{{item.content}}</text>
			<view slot="actions" class="card-actions">
				<!-- <view class="card-actions-item" @click="actionsClick('分享')">
					<uni-icons type="pengyouquan" size="18" color="#999"></uni-icons>
					<text class="card-actions-item-text">分享</text>
				</view>
				<view class="card-actions-item" @click="actionsClick('点赞')">
					<uni-icons type="heart" size="18" color="#999"></uni-icons>
					<text class="card-actions-item-text">点赞</text>	
				</view>
				<view class="card-actions-item" @click="actionsClick('评论')">
					<uni-icons type="chatbubble" size="18" color="#999"></uni-icons>
					<text class="card-actions-item-text">评论</text>
				</view> -->
			</view>
		</uni-card>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				activityList: [{
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
			}
		},
		methods: {
			onClick(id) {
				console.log("id")
				console.log(id)
			}
		}
	}
</script>

<style lang="scss">
	.container {
		overflow: hidden;
	}

	.custom-cover {
		flex: 1;
		flex-direction: row;
		position: relative;
	}

	.cover-content {
		position: absolute;
		bottom: 0;
		left: 0;
		right: 0;
		height: 40px;
		background-color: rgba($color: #000000, $alpha: 0.4);
		display: flex;
		flex-direction: row;
		align-items: center;
		padding-left: 15px;
		font-size: 14px;
		color: #fff;
	}

	.card-actions {
		display: flex;
		flex-direction: row;
		justify-content: space-around;
		align-items: center;
		height: 45px;
		border-top: 1px #eee solid;
	}

	.card-actions-item {
		display: flex;
		flex-direction: row;
		align-items: center;
	}

	.card-actions-item-text {
		font-size: 12px;
		color: #666;
		margin-left: 5px;
	}

	.cover-image {
		flex: 1;
		height: 150px;
	}

	.no-border {
		border-width: 0;
	}
</style>