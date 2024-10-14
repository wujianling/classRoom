<template>
	<uni-list>
		<uni-list :border="true">
			<uni-list-chat v-for="(item,index) in messageList" :key="index" :to="`/pages/work/detail?id=`+item.id"   :title="item.className" :avatar="item.cover" :note="item.msg" :time="item.lastTime" :show-bage="item.unReadCount>0?true:false" :badge-text="item.unReadCount"></uni-list-chat>
		</uni-list>
	</uni-list>

</template>

<script>
	export default {
		data() {
			return {
				messageList:[
					{
						id:1,
						cover:"http://t15.baidu.com/it/u=3639854910,1717044149&fm=224&app=112&f=JPEG?w=200&h=200",
						msg:"张老师发来一份作业",
						lastTime:"12分钟前",
						unReadCount:8,
						className:"书画A-01",
						
					},
					{
						id:2,
						cover:"http://t15.baidu.com/it/u=3639854910,1717044149&fm=224&app=112&f=JPEG?w=200&h=200",
						msg:"张老师回复了消息",
						lastTime:"2小时前",
						unReadCount:3,
						className:"书画A-02",
						
					},
					{
						id:3,
						cover:"http://t15.baidu.com/it/u=3639854910,1717044149&fm=224&app=112&f=JPEG?w=200&h=200",
						msg:"我发送了一条消息",
						lastTime:"3天前",
						unReadCount:0,
						className:"书画A-32",
						
					}
				]
			}
		},
		
		methods: {
			goToDetail(id){
				console.log(1)
				uni.navigateTo({
					url:"/pages/work/detail?id="+id,
					
				})
			}
		}
	}
</script>

<style>

</style>
