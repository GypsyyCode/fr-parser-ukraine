<template>
  <el-row justify="center">
      <span class="gravur">Made with <span style="color: #e2555567;">&#9829;</span> by Fabian Lehmann</span>
  </el-row>
  <el-row justify="center">
    <el-timeline>
      <template v-for="timestamp in keys" :key="timestamp">
        <el-timeline-item
          :timestamp="createTime(timestamp)"
          placement="top"
          :color="color"
        >
          <el-card
            class="box-card"
            v-for="(article, index) in dings[timestamp]"
            :key="index"
            :style="getCardMargin(index)"
          >
            <p>{{ article }}</p>
          </el-card>
        </el-timeline-item>
      </template>
    </el-timeline>
  </el-row>
</template>

<script>
// import Moment from "moment"
import MomentTZ from "moment-timezone";

import { ElRow, ElTimeline, ElTimelineItem, ElCard } from "element-plus";

async function fetchData() {
  // const news = await fetch("http://localhost:3000/updates");
  const news = await fetch("/updates");
  let newsJson = await news.json();
  const data = newsJson.message.updates;

  return data;
}

export default {
  components: {
    ElRow,
    ElTimeline,
    ElTimelineItem,
    ElCard,
  },
  methods: {
    createTime(timestamp) {
      let num = parseInt(timestamp);
      let date = MomentTZ.tz(num *1000, "Europe/Berlin");
      // let date = new Date(num *1000)
      console.log(date.format('DD.MM.YYYY HH:mm'));

      // let d = Sugar.Date.create(date)
      return date.format('DD.MM.YYYY HH:mm');
      // return Sugar.Date.format(d, '{dd}.{MM}.{yyyy} {HH}:{mm}')
    },
    getCardMargin(index) {
      if (index >= 0) {
        return { "margin-bottom": "2.6rem" };
      }
    },
  },

  async setup() {
    const color = "#41b883";
    let dings = await fetchData();
    let keys = [];

    for (let key in dings) {
      keys.push(key);
    }
    keys.sort();
    keys.reverse();
    for (let i = 0; i < keys.lenght; i++) {
      console.log(keys[i] + ": " + dings[keys[i]]);
    }

    return { dings, keys, color };
  },
};
</script>

<style lang="css">
.gravur{
  font-size: 0.7rem;
  color: #2c3e50a2;
  margin-bottom: 1.6rem
}
.box-card {
  font-size: 1.2rem;
  /* line-height: 1.25rem; */
  word-wrap: normal;
}
@media (min-width: 1500px) {
  .box-card {
    max-width: 40vw;
  }
}
@media (min-width: 720px) {
  .box-card {
    max-width: 50vw;
  }
}

@media (max-width: 720px) {
  .el-timeline {
    padding-inline-start: 10px;
    padding-inline-end: 10px;
  }
}
.el-card__body {
  padding-top: 10px;
  padding-bottom: 15px;
}

.el-timeline-item__timestamp.is-top {
  color: #41b883;
  font-family: "Courier New", Courier, monospace;
}
</style>