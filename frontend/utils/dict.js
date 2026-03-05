export const getDictLabel = (dict,key) => {
    let label = ""
    dict.some(item => {
        if (item.value === key) {
            label = item.label
            return true
        }
    })
    return label
}