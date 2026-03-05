<template>
  <div>
    <FriendsHeader :info="data?.info"></FriendsHeader>
    <FriendsLink :data="data"></FriendsLink>
  </div>
</template>
<script setup>
import {getCmsFriendInfo, getCmsFriendList} from "~/api/friends.js";

const {data} = useAsyncData("friends",async () => {
  const [infoRes,listRes] = await Promise.all([
    getCmsFriendInfo(),
    getCmsFriendList()
  ])
  return {
    info:infoRes.data.friendInfo,
    list:listRes.data.list
  }
})

</script>
