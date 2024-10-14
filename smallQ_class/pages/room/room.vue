<template>
		<view class="container">
			<view v-for="(item,index) in list" :key="index" >
				<view class="example">
						<view class="room_line">
							<view class="room_line_title">科目：</view>
							<view class="room_line_value">{{item.typer}}</view>
						</view>
						<view class="room_line">
							<view class="room_line_title">教室：</view>
							<view class="room_line_value">{{item.roomName}}</view>
						</view>
						<view class="room_line">
							<view class="room_line_title">时间：</view>
							<view class="room_line_value">{{item.time}}</view>
						</view>
						<view class="room_line">
							<view class="room_line_title">时长：</view>
							<view class="room_line_value">{{item.duration}}(分钟)</view>
						</view>
						<view class="room_line">
							<view class="room_line_title">老师：</view>
							<view class="room_line_value">{{item.tearcher.name}}</view>
						</view>
						<view class="room_line">
							<view class="room_line_title">备注：</view>
							<view class="room_line_value">{{item.content}}</view>
						</view>
						<view class="room_line" style="margin-bottom: 120rpx;margin-top: 120rpx;">
							<button style="width: 600rpx;border-radius: 8%;" disabled="disabled">请假</button>
						</view>
						
					</view>
					
				</view>
	
		</view>
</template>

<script>
	export default {
		data() {
			return {
				list: [{
					time: '6月12日 14:30',
					roomName: '书法001-x',
					duration: 120,
					typer: "书法",
					tearcher:{
						name:"老王",
						id:1
					},
					content:"没有任何备注"
				},
				{
					time: '',
					roomName: '书法001-x',
					duration: 120,
					typer: "书法",
					tearcher:{
						name:"老王",
						id:1
					}
				},
				{
					time: '',
					roomName: '书法001-x',
					duration: 120,
					typer: "书法",
					tearcher:{
						name:"老王",
						id:1
					}
				}
				]
			}
		},
		methods: {

		}
	}
</script>

<style lang="scss">
	.example {
		padding: 15px;
		margin-left: 80rpx;
		margin-top: 80rpx;
		border-radius: 5%;
		background-color: #33ccff;
		width: 80%;
		height: 100%;
	}
	.room_line{
		
		margin-top: 40rpx;
		display: flex;
		align-items: center;
	}
	.room_line_title{
		color: #000;
		margin-left: 60rpx;
		font-size: 26px;
	}
	.room_line_value{
		color: #fff;
		font-size: 20px;
	}
	.segmented-control {
		margin-bottom: 15px;
	}

	.button-group {
		margin-top: 15px;
		display: flex;
		justify-content: space-around;
	}

	.form-item {
		display: flex;
		align-items: center;
	}

	.button {
		display: flex;
		align-items: center;
		height: 35px;
		margin-left: 10px;
	}
</style>