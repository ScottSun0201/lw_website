import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useFavoriteStore = defineStore('favorite', () => {
    const favoriteMap = ref({})

    const setFavorites = (map) => {
        favoriteMap.value = map
    }

    const toggleFavorite = (spuId, val) => {
        favoriteMap.value = { ...favoriteMap.value, [spuId]: val }
    }

    const isFavorite = (spuId) => {
        return !!favoriteMap.value[spuId]
    }

    return { favoriteMap, setFavorites, toggleFavorite, isFavorite }
})
