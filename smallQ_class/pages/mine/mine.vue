<template>
	 <view class="mine-container" >
	    <!--顶部个人信息栏-->
	    <view class="header-section">
	      <view class="flex padding justify-between">
	        <view class="flex align-center">
	          <view v-if="!avatar" class="cu-avatar xl round bg-white">
	            <view class="iconfont icon-people text-gray icon"></view>
	          </view>
	          <image v-if="avatar" @click="handleToAvatar" :src="avatar" class="mine_cover" style="border-radius: 15%;" mode="aspectFill">
	          </image>
	          <view v-if="!name" @click="handleToLogin" class="login-tip">
	            点击登录
	          </view>
	          <view v-if="name"  class="user-info">
	            <view class="u_title">
	              {{name}}
	            </view>
	          </view>
	        </view>
	        <view v-if="studentInfo.name" @click="handleToInfo"  class="flex align-center">
	          <text>学员: {{studentInfo.name}}</text>
	          <uni-icons type="search" size="20" style="margin-left: 10rpx;"></uni-icons>
	        </view>
	      </view>
	    </view>
	
	    <view class="content-section">
	      <view class="mine-actions grid col-3 text-center">
	        <view class="action-item" @click="handleJiaoLiuQun">
	          <view class="iconfont icon-friendfill text-pink icon"></view>
	          <button open-type="contact" type="primary" size="mini" plain>联系客服</button>
	        </view>
			<view class="action-item" @click="handleJiaoLiuQun">
			  <view class="iconfont icon-friendfill text-pink icon"></view>
			  <text class="text">课包查询</text>
			</view>
	        <view class="action-item" @click="handleBuilding">
	          <view class="iconfont icon-service text-blue icon"></view>
	          <text class="text">问题反馈</text>
	        </view>
	      </view>
	
	      <view class="menu-list">
	        <view class="list-cell list-cell-arrow" @click="handleToEditInfo">
	          <view class="menu-item-box">
	            <view class="iconfont icon-user menu-icon"></view>
	            <view>编辑资料</view>
	          </view>
	        </view>
	        <view class="list-cell list-cell-arrow" @click="handleAbout">
	          <view class="menu-item-box">
	            <view class="iconfont icon-aixin menu-icon"></view>
	            <view>关于我们</view>
	          </view>
	        </view>
	        <view class="list-cell list-cell-arrow" @click="handleToSetting">
	          <view class="menu-item-box">
	            <view class="iconfont icon-setting menu-icon"></view>
	            <view>应用设置</view>
	          </view>
	        </view>
	      </view>
	
	    </view>
	  </view>
</template>

<script>
	export default {
		data() {
			return {
				name:"爱无悔",
				avatar:"http://img.duoziwang.com/2021/02/1621261310479632.jpg",
				studentInfo:{
					name:"赵小月"
				}
			}
		},
		methods: {
			
		}
	}
</script>

<style lang="scss">
 .mine-container {
    width: 100%;
    height: 100%;


    .header-section {
      padding: 15px 15px 45px 15px;
      background-color: #3c96f3;
      color: white;

      .login-tip {
        font-size: 18px;
        margin-left: 10px;
      }
.mine_cover{
	width: 80px;
	height: 80px;
	border-radius:10;
}
      .cu-avatar {
        border: 2px solid #eaeaea;

        .icon {
          font-size: 40px;
        }
      }

      .user-info {
        margin-left: 15px;

        .u_title {
          font-size: 18px;
          line-height: 30px;
        }
      }
    }

    .content-section {
      position: relative;
      top: -50px;

      .mine-actions {
        margin: 15px 15px;
        padding: 20px 0px;
        border-radius: 8px;
        background-color: white;

        .action-item {
          .icon {
            font-size: 28px;
          }

          .text {
            display: block;
            font-size: 13px;
            margin: 8px 0px;
          }
        }
      }
    }
  }
</style>
