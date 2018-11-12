<template>
  <div class="bg"></div>
</template>

<script>
  import QbHeader from '@/views/header'
  // import API from '@/api'
  export default {
    data () {
      return {
      }
    },
    components: {
      QbHeader
    },
    created () {
    },
    methods: {
    }
  }
</script>

<style>
  .home-wapper {
    min-width: 1330px;
    margin: 0 auto;
    padding: 0 18%;
  }
  .bg {
    min-width: 1440px;
    height: 1000px;
    margin-top: -20px;
    background-image: url(../../assets/img/bg.jpeg);
    background-size: 100% 100%;
    background-repeat: no-repeat;
  }
</style>
