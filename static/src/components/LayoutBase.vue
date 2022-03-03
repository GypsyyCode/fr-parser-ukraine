<template>
<el-row justify="center">
      <span class="gravur">Made with <span style="color: #e2555567;">&#9829;</span> by Fabian Lehmann</span>
  </el-row>
  <el-row justify="center">
       <!-- <el-affix position="bottom" :offset="20"> -->
    <el-tabs :tab-position="tabPosition" style="height: 200px" class="min-width" v-model="activeTab" :style="getTabStyle">
        <el-tab-pane name="feed" label="Feed">
            <UpdateFeed :isMobile="isMobile"></UpdateFeed>
            </el-tab-pane>
            
        <el-tab-pane name="article" label="Artikel">
            <Suspense>
                <template #default>
                <news-feed :isMobile="isMobile"></news-feed>
                </template>
                <template #fallback>
                <div>Loading...</div>
                </template>
            </Suspense>

        </el-tab-pane>
    </el-tabs>
    <!-- </el-affix> -->
     

  </el-row>
</template>

<script>
import {ElTabs, ElTabPane, ElRow} from "element-plus";
import UpdateFeed from './UpdateFeed/UpdateFeed.vue';
import NewsFeed from './NewsFeed/NewsFeed.vue';
import { useBreakpoints } from '@vueuse/core'

export default {
    components: {UpdateFeed, NewsFeed, ElRow, ElTabs, ElTabPane},
    setup() {
        const breakpoints = useBreakpoints({
            desktop: 720
        })

    return {breakpoints}
    },

    data() {
        return {
            activeTab: "feed",
        }
    },

    computed: {
        tabPosition() {
            console.log(this.isMobile);
            return this.isMobile ? 'bottom' : 'bottom'
        },
        isMobile() {
            return this.breakpoints.isSmaller('desktop')
        }
    }
}
</script>

<style>
.gravur{
  font-size: 0.7rem;
  color: #2c3e50a2;
  margin-bottom: 1.6rem
}
.min-width {
    display: float;
    justify-content: center;
}

.el-tabs__nav-scroll {
    margin-left: auto;
    overflow: hidden;
    margin-right: auto;
    width: 120px;
}

@media (min-width: 1500px) {
  .min-width {
    min-width: 50vw;
    max-width: 50vw;
  }
}
@media (min-width: 720px) {
  .min-width {
    min-width: 60vw;
    max-width: 60vw;
  }
}

</style>