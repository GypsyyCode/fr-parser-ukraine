<template>
  <Suspense>
    <template #default>
      <el-timeline>
        <el-scrollbar height="84vh">
          <UpdateFeedList></UpdateFeedList>
        </el-scrollbar>
      </el-timeline>
    </template>
    <template #fallback>
      <div>Loading...</div>
    </template>
  </Suspense>
</template>

<script>
// import Moment from "moment"

import { ElTimeline, ElScrollbar } from "element-plus";
import UpdateFeedList from "./UpdateFeedList.vue";

export default {
  components: {
    ElTimeline,
    ElScrollbar,
    UpdateFeedList,
  },
  methods: {},
  props: {
    isMobile: Boolean,
  },
};
</script>

<style lang="css">
@media (max-width: 720px) {
  .el-timeline {
    padding-inline-start: 10px;
    padding-inline-end: 10px;
  }
}
</style>