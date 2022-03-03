<template>
  <layout-base></layout-base>
</template>

<script>
// import ArticleUpdates from "./components/ArticleUpdates.vue";
import LayoutBase from "./components/LayoutBase.vue";



export default {
  components: { LayoutBase },

  
};
</script>

<style>
  /* @import url('https://fonts.googleapis.com/css2?family=Merriweather&display=swap'); */
  @import url('https://fonts.googleapis.com/css2?family=Source+Sans+Pro:wght@300;400&display=swap');
#app {
  font-family: 'Source Sans Pro', sans-serif;
  -webkit-font-smoothing: antialiased;
  font-weight: 400;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
  margin-top: 60px;
}
@media (max-width: 720px) {
  #app {
    margin-top: 0.8rem;
  }
  
}
</style>
