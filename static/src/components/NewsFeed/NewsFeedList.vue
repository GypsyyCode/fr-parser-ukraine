<template>
  <el-row justify="center">
    <el-card
      class="article-card"
      v-for="article in articles"
      :key="article.index"
      @click="open(article.link)"
    >
      <el-row>
        <span class="article-title">{{
          formatTitle(article.title, article.source.title)
        }}</span>
      </el-row>
      <el-row justify="space-between" class="article-info">
        <div class="article-date">{{ getDate(article.published) }}</div>
        <div class="article-source">
          <el-link
            type="primary"
            @click="open('http://' + article.source.title)"
          >
            {{ article.source.title }}</el-link
          >
        </div>
      </el-row>
    </el-card>
  </el-row>
</template>

<script>
import MomentTZ from "moment-timezone";
import {
  ElRow,
  ElCard,
  ElLink,
} from "element-plus";

function formatDate(rawDate) {
  let date = MomentTZ.utc(rawDate);
  return date.format("DD.MM.YYYY HH:mm");
}

export default {
    components: { ElRow, ElCard, ElLink},
  props: {
    articles: {
      type: Object,
      required: true,
    },
  },
  methods: {
    formatTitle(title, source) {
      let res = title.replace("- " + source, "");
      return res.replace("| " + source, "");
    },
    getDate(rawDate) {
      return formatDate(rawDate);
    },
     open(url) {
      window.open(url);
    },
  },
};


</script>

<style>
</style>