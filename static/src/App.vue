<template>
  <Suspense>
    <template #default>
      <ArticleUpdates />
    </template>
    <template #fallback>
      <div>Loading...</div>
    </template>
  </Suspense>
</template>

<script>
import ArticleUpdates from "./components/ArticleUpdates.vue";



export default {
  components: { ArticleUpdates },

  
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
</style>
