<template>
  <!-- <el-row justify="center" class="tag-row">
    <el-tag class="ml-2" type="success">{{ feed.updated }} </el-tag>
    <el-tag class="ml-2" type="success">{{ feed.language }} </el-tag>

    <el-button
      class="ml-2"
      size="small"
      type="primary"
      plain
      @click="open(feed.link)"
      >Google News Feed</el-button
    >
  </el-row> -->
  <el-scrollbar height="84vh">
    <news-feed-list :articles="articles"></news-feed-list>
  </el-scrollbar>
</template>

<script>
async function fetchArticles() {
  // const news = await fetch("http://localhost:5050/articles");
  const news = await fetch("/articles");
  let newsJson = await news.json();
  const data = newsJson.message;

  return data;
}

import {  ElScrollbar } from "element-plus";
import NewsFeedList from "./NewsFeedList.vue";

import MomentTZ from "moment-timezone";

function formatDate(rawDate) {
  let date = MomentTZ.utc(rawDate);
  return date.format("DD.MM.YYYY HH:mm");
}

export default {
  components: { NewsFeedList, ElScrollbar },

  async setup() {
    let data = await fetchArticles();
    let feed = data.feed;
    let articles = data.Articles;

    feed.updated = formatDate(feed.updated);
    console.log(articles);

    return { feed, articles };
  },

  props: {
    isMobile: Boolean,
  },

  methods: {
    open(url) {
      window.open(url);
    },
  },
};
</script>

<style lang="scss">
.ml-2 {
  margin-left: 1rem;
}
.tag-row {
  margin-bottom: 1rem;
}

.article-card {
  font-size: 1.2rem;
  /* line-height: 1.25rem; */
  word-wrap: normal;
  margin-bottom: 2.6rem;

  .article-info {
    font-size: 0.85rem;
    margin-top: 0.5rem;

    .el-link {
      font-size: inherit;
    }
    .article-date {
      color: #41b883;
      font-family: "Courier New", Courier, monospace;
    }
  }
}
@media (min-width: 1500px) {
  .article-card {
    width: 50%;
    max-width: 40vw;
  }
}
@media (max-width: 720px) {
  .article-card {
    width: 100%;
    max-width: 85vw;
  }
}
</style>