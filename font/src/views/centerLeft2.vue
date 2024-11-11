<template>
  <div id="centerRight1">
    <div class="bg-color-black">
      <div class="d-flex pt-2 pl-2">
        <span>
          <icon name="chart-line" class="text-icon"></icon>
        </span>
        <div class="d-flex">
          <span class="fs-xl text mx-2">eis轮播</span>
        </div>
      </div>
      <div class="d-flex jc-center body-box">
        <!-- 轮播组件 -->
        <div class="dv-scr-board">
          <div class="carousel-wrapper">
            <!-- 轮播项 -->
            <div
              class="carousel-item"
              :class="{'current-item': currentIndex === index}"
              v-for="(item, index) in carouselItems"
              :key="index"
            >
              <span v-html="item"></span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      currentIndex: 0, // 当前显示的索引
      carouselItems: [
        "<span class='colorRed'>1</span>",
        "<span class='colorGrass'>2</span>",
        "<span class='colorBlue'>3</span>",
        "<span class='colorYellow'>4</span>",
        "<span class='colorPink'>5</span>",
        "<span class='colorPurple'>6</span>"
      ],
      timer: null // 用于自动切换的定时器
    };
  },
  methods: {
    // 启动自动轮播
    startAutoPlay() {
      this.timer = setInterval(() => {
        this.nextItem();
      }, 2000); // 每2秒切换一次
    },
    // 切换到下一项
    nextItem() {
      this.currentIndex = (this.currentIndex + 1) % this.carouselItems.length;
    }
  },
  mounted() {
    this.startAutoPlay(); // 页面加载时开始自动播放
  },
  beforeDestroy() {
    clearInterval(this.timer); // 组件销毁时清除定时器
  }
};
</script>

<style lang="scss" scoped>
$box-height: 410px;
$box-width: 360px;
$highlight-color: #37a2da;
$bg-dark: #1a1d2e;
$bg-light: #0f1325;

#centerRight1 {
  padding: 16px;
  padding-top: 20px;
  height: $box-height;
  width: $box-width;
  border-radius: 10px;
  background: linear-gradient(135deg, $bg-dark, $bg-light);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);

  .bg-color-black {
    height: $box-height - 30px;
    border-radius: 10px;
    background-color: $bg-dark;
    padding: 10px;
  }

  .header {
    font-size: 18px;
    color: $highlight-color;
    display: flex;
    align-items: center;
    .text {
      color: #c3cbde;
      font-weight: bold;
    }
  }

  .body-box {
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 10px;
    overflow: hidden;
    padding: 10px;
  }

  .dv-scr-board {
    width: 100%;
    height: 100%;
    max-width: 320px;
    max-height: 350px;
    display: flex;
    justify-content: center;
    align-items: center;
    border: 1px solid rgba(255, 255, 255, 0.2);
    border-radius: 8px;
    background-color: rgba(0, 0, 0, 0.5);
    color: #fff;
    font-family: Arial, sans-serif;
    font-size: 16px;

    .carousel-wrapper {
      display: flex;
      flex-direction: column; /* 上下排列 */
      justify-content: center;
      align-items: center;
      transition: transform 0.5s ease-in-out;
    }

    .carousel-item {
      min-width: 100%;
      padding: 10px;
      display: flex;
      justify-content: center;
      align-items: center;
      color: white;
      font-weight: bold;
      opacity: 0.5; /* 默认透明度较低 */
      transition: opacity 0.5s ease-in-out;
    }

    .current-item {
      opacity: 1; /* 当前项完全不透明 */
      font-size: 24px; /* 当前项字体稍大 */
    }
  }
}
</style>
