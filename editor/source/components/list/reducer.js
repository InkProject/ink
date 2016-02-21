import Immutable from 'immutable'

import ActionType from '../action'

export default function list(state = Immutable.fromJS({
    show: false,
    data: [],
    tags: {}
}), action) {
    switch (action.type) {
        case ActionType.LIST_SHOW:
            return state.set('show', true)
        case ActionType.LIST_HIDE:
            return state.set('show', false)
        case ActionType.LIST_REFRESH:
            let tagMap = {}
            const newData = Object.keys(action.data).map(function(id) {
                const item = action.data[id]
                const tags = item.Article ? (item.Article.Tags || []) : []
                tags.forEach((tag) => {
                    tagMap[tag] ? tagMap[tag]++ : tagMap[tag] = 1
                })
                return {
                    id,
                    tags,
                    name: item.Name,
                    title: item.Article ? item.Article.Title : '未命名标题',
                    preview: item.Article ? item.Article.Preview : '',
                    draft: item.Article ? item.Article.Draft : false,
                    date: item.Date
                }
            })
            const newState = state.set('tags', tagMap)
            return newState.set('data', Immutable.fromJS(newData))
        default:
            return state
    }
}
