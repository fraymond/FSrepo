<template>
  <section>
    <!-- <span class="info1">此处有大图</span> -->
    <!-- <h1>文件风暴</h1>
    <h4>全球第一个去中心化的分布式存储平台</h4>
    <input type="button" value="立即体验"> -->
    <canvas id='myCanvas' width='' height=''></canvas>
  </section>
</template>
<script>
  // import QbHeader from '@/views/header'
  // import API from '@/api'
  export default {
    data () {
      return {
        isAgree: false,
        captchaPath: '',
        dianArr: [],
        Canvas: document.getElementById('myCanvas'),
        can: this.Canvas.getContext('2d')
      }
    },
    // components: {
    //   QbHeader
    // },
    created () {
      this.init()
    },
    computed: {
      // canvas 宽
      CanvasWidth: function () {
        // `this` 指向 vm 实例
        return document.documentElement.clientWidth
      },
      // canvas 高
      CanvasHeight: function () {
        // `this` 指向 vm 实例
        return document.documentElement.clientHeight
      }
    },
    methods: {
      init () {
        this.drawBGC()
      },
      drawBGC () {
        this.can.clearRect(0, 0, this.CanvasWidth, this.CanvasHeight)
        this.can.beginPath()
        this.can.fillStyle = 'black'
        this.can.fillRect(0, 0, this.CanvasWidth, this.CanvasHeight)
        for (var i = 0; i < 10; i++) {
        //   var dian = new Dian()
        //   this.dianArr.push(dian)
        }
        for (var j = 0; j < this.dianArr.length; j++) {
          var aa = this.dianArr[j]
          aa.this.move()
          aa.this.drawDian()
          if (aa.x > this.CanvasWidth || aa.x < 0 || aa.y > this.CanvasHeight || aa.y < 0) {
            this.dianArr.splice(i, 1)
          }
        }
        this.can.closePath()
        window.requestAnimationFrame(this.drawBGC())
      },
      randNum (min, max) {
        return parseInt(Math.random() * (max - min + 1) + min)
      },
      Dian () {
        this.x = this.randNum(this.CanvasWidth / 2 - 200, (this.CanvasWidth / 2 + 200))
        this.y = this.randNum(this.CanvasHeight / 2 - 200, (this.CanvasHeight / 2 + 200))
        this.w = this.randNum(1, 5)
        this.h = this.randNum(1, 5)
        this.speedX = 0
        this.speedY = 0
        this.bluX = this.randNum(0, 1)   // Y正负
        this.bluY = this.randNum(0, 1)   // X正负
        this.blu = this.randNum(1, 2)    // 将所有点分为两类
        this.blu1 = this.randNum(1, 2)   // 正方向点XY正负
      },
      move () {
        if (this.blu === 1) {
          if (this.blu1 === 1) {
            if (this.bluX === 1) {
              this.speedX += Math.random() / 250
              this.x += this.speedX
            } else {
              this.speedX += Math.random() / 250
              this.x -= this.speedX
            }
            if (this.bluY === 1) {
              this.speedY += Math.random() / 50
              this.y -= this.speedY
            } else {
              this.speedY += Math.random() / 50
              this.y += this.speedY
            }
          } else {
            if (this.bluX === 1) {
              this.speedX += Math.random() / 25
              this.x += this.speedX
            } else {
              this.speedX += Math.random() / 25
              this.x -= this.speedX
            }
            if (this.bluY === 1) {
              this.speedY += Math.random() / 400
              this.y -= this.speedY
            } else {
              this.speedY += Math.random() / 400
              this.y += this.speedY
            }
          }
        } else {
          this.speedY += Math.random() / 50
          this.speedX += Math.random() / 25
          if (this.bluX === 1) {
            this.x += this.speedX
          } else {
            this.x -= this.speedX
          }
          if (this.bluY === 1) {
            this.y -= this.speedY
          } else {
            this.y += this.speedY
          }
        }
      },
      drawDian () {
        this.can.fillStyle = 'white'
        this.can.fillRect(this.x, this.y, this.w, this.h)
      },
      // 去登录页面
      goLogin () {
        this.$router.push({ path: '/login' })
      }
    }
  }
</script>
<style>
  /* .home-wapper {
    min-width: 1330px;
    margin: 0 auto;
    margin-top: -20px;
    padding: 0 18%;
  }
  .home-wapper .info1 {
    width: 1200px;
    height: 900px;
    display: block;
    line-height: 900px;
    text-align: center;
  } */
  * {
    margin: 0;
    padding: 0;
  }
  html {
    overflow: hidden;
  }
</style>

