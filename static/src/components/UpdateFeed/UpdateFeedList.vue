<template>
  <template v-for="timestamp in dates" :key="timestamp">
    <el-timeline-item
      :timestamp="createTime(timestamp)"
      placement="top"
      :color="color"
    >
      <el-card
        class="box-card"
        v-for="(article, index) in updates[timestamp]"
        :key="index"
        :style="getCardMargin(index)"
      >
        <p>{{ article }}</p>
      </el-card>
    </el-timeline-item>
  </template>
</template>

<script>
import { ElTimelineItem, ElCard } from "element-plus";
import MomentTZ from "moment-timezone";

async function fetchData() {
  // const news = await fetch("http://localhost:5050/updates");
  const news = await fetch("/updates");
  let newsJson = await news.json();
  const data = newsJson.message.updates;
  return data;
}
export default {
  components: {
    ElTimelineItem,
    ElCard,
  },

  async setup() {
    const color = "";
    let updates = await fetchData();
    let dates = [];

    for (let key in updates) {
      dates.push(key);
    }
    dates.sort();
    dates.reverse();
    return { updates, dates, color };
  },

  methods: {
      createTime(timestamp) {
      let num = parseInt(timestamp);
      let date = MomentTZ.tz(num * 1000, "Europe/Berlin");

      return date.format("DD.MM.YYYY HH:mm");
    },
     getCardMargin(index) {
      if (index >= 0) {
        return { "margin-bottom": "2.6rem" };
      }
    },
  }
};
</script>

<style>
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
.el-card__body {
  padding-top: 10px;
  padding-bottom: 15px;
}

.el-timeline-item__timestamp.is-top {
  color: #41b883;
  font-family: "Courier New", Courier, monospace;
}
</style>