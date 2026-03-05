export const getUrl = (url) => {
    if (!url) return ''
    const config = useRuntimeConfig()
    const path = config.public.basePROXY.slice(-1) === '/' ? config.public.basePROXY : config.public.basePROXY + '/'
    return url.startsWith('http') ? url : path + url
}
