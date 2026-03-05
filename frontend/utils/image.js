export const getUrl = (url) => {
    const config = useRuntimeConfig()
    // 如果path最后一位为/就直接拼接 否则加上/
    const path = config.public.basePROXY.slice(-1) === '/' ? config.public.basePROXY : config.public.basePROXY + '/'
    return  url && url.slice(0, 4) !== 'http' ? path + url : url
}
