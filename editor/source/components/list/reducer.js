import Immutable from 'immutable'

import ActionType from '../action'

export default function list(state = Immutable.fromJS({
    show: false,
    data: []
}), action) {
    switch (action.type) {
        case ActionType.LIST_SHOW:
            return state.set('show', true)
        case ActionType.LIST_HIDE:
            return state.set('show', false)
        case ActionType.LIST_REFRESH:
            let newData = Object.keys(action.data).map(function(id) {
                let item = action.data[id]
                return {
                    id,
                    tags: item.Article ? (item.Article.Tags || []) : [],
                    name: item.Name,
                    title: item.Article ? item.Article.Title : '未命名标题',
                    preview: item.Article ? item.Article.Preview : ''
                }
            })
            return state.set('data', Immutable.fromJS(newData))
        default:
            return state
    }
}
