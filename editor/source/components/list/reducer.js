import Immutable from 'immutable'

import ActionType from '../action'

export default function list(state = Immutable.fromJS({
    show: false,
    data: [],
    tags: {},
    searchObj: {}
}), action) {
    switch (action.type) {
        case ActionType.LIST_SHOW:
            return state.set('show', true)
        case ActionType.LIST_HIDE:
            return state.set('show', false)
        case ActionType.LIST_REFRESH:
            let tagMap = {}
            let searchObj = {}
            let newData = Object.keys(action.data).map(function(id) {
                const item = action.data[id]
                const tags = item.Article ? (item.Article.Tags || []) : []
                tags.forEach((tag) => {
                    tagMap[tag] ? tagMap[tag]++ : tagMap[tag] = 1
                })
                const name = item.Name
                const title = item.Article ? item.Article.Title : '未命名标题'
                const preview = item.Article ? item.Article.Preview : ''
                searchObj[id] = `${name} ${title} ${tags.join(' ')} ${preview}`.toLowerCase()
                return {
                    id,
                    tags,
                    name,
                    title,
                    preview,
                    draft: item.Article ? item.Article.Draft : false,
                    date: item.Date
                }
            })
            newData = newData.sort((item1, item2) => {
                return new Date(item2.date).getTime() - new Date(item1.date).getTime()
            })
            const newState = state.set('tags', tagMap).set('searchObj', searchObj)
            return newState.set('data', Immutable.fromJS(newData))
        default:
            return state
    }
}
