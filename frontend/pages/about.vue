<template>
  <div>
  <AboutHeader :info="data?.info"></AboutHeader>
  <AboutTeams :data="data"></AboutTeams>
  </div>
</template>
<script setup>

import {getClientAboutInfo, getClientAboutList} from "~/api/about.js";

const {data} = useAsyncData("team",async () => {
  const [listRes,infoRes] = await Promise.all([
    getClientAboutList(),
    getClientAboutInfo()
  ])
  return {
    list:listRes.data.list,
    info:infoRes.data.info
  }
})
</script>
