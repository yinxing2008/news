//厦门大学计算机专业 | 前华为工程师
//专注《零基础学编程系列》 http://lblbc.cn/blog
//包含：Java | 安卓 | 前端 | Flutter | iOS | 小程序 | 鸿蒙
//公众号：蓝不蓝编程
<template>
	<view class="page-wrapper">
		<view class="search_wrapper" @tap="gotoSearch">
			<view class="search_row">
				<image class="search-img" src="/static/images/search.png"></image>
				<view class="search-text">点击搜索</view>
			</view>
		</view>

		<scroll-view class="tab-row" :scroll-x="true">
			<view :class="'tab  ' + (index == current ? 'tab-selected' : '')" @tap="tabSelect" :data-id="index"
				v-for="(item, index) in categories" :key="item.k">
				{{ item.name }}
			</view>
		</scroll-view>

		<view class="row" :data-id="item.id" v-for="(item, index) in apps" :key="item.k" @click="gotoDetail(item.link)">
			<div class="col">
				<div class="news-title">{{ item.title }}</div>
				<div class="news-comment">
					 {{ item.author }} {{ item.commentCount }}评论
				</div>
			</div>
			<image class="img" :src="item.imgUrl" ></image>
		</view>
	</view>
</template>
<script>
	import {
		queryCategories,
		queryNewsByCategory
	} from '@/common/api'
	export default {
		data() {
			return {
				categories: [],
				apps: [],
				current: 0
			}
		},
		mounted() {
			this.queryCategories();
		},
		methods: {
			tabSelect: function(e) {
				this.current = e.currentTarget.dataset.id;
				this.queryNewsByCategory();
			},
			queryCategories() {
				queryCategories().then(resp => {
					this.categories = resp.data
					this.queryNewsByCategory();
				})
			},
			queryNewsByCategory() {
				queryNewsByCategory(this.getCurrentCategoryId()).then(resp => {
					this.apps = resp.data
				})
			},

			getCurrentCategoryId() {
				return this.categories[this.current].id;
			},
			gotoSearch() {
				uni.navigateTo({
					url: '/pages/search'
				});
			},
			gotoDetail(url) {
				location.href = url;
			}
		}
	}
</script>
<style>
	.page-wrapper {}

	.search_wrapper {
		background-color: #418df9;
		padding-top: 10px;
		padding-bottom: 10px;
	}

	.search_row {
		display: flex;
		flex-direction: row;
		background-color: white;
		border-radius: 10rem;
		margin-left: 1rem;
		margin-right: 1rem;
		padding-top: 5px;
		padding-bottom: 5px;
		align-items: center;
	}

	.search-img {
		height: 20px;
		width: 20px;
		margin-left: 10px;
	}

	.search-text {
		margin-left: 10px;
	}

	.tab-row {
		padding-top: 15px;
		width: 100%;
		display: flex;
		align-items: center;
		box-sizing: border-box;
		white-space: nowrap;
		background-color: white;
	}

	.tab {
		width: 80px;
		margin-left: 10px;
		margin-right: 10px;
		display: inline-block;
		padding-bottom: 10px;
		text-align: center;
	}

	.tab-selected {
		border-bottom: 2px solid;
		color: #418df9;
	}

	.row {
	  display: flex;
	  margin-left: 10px;
	  margin-right: 10px;
	  margin-top: 20px;
	}
	.col {
	  font-size: 15px;
	  flex: 1;
	}
	.news-title {
	  color: #333333;
	}
	.news-comment {
	  color: #666666;
	}
	.img {
	  margin-left: 5px;
	  width: 100px;
	  height: 60px;
	}
</style>
