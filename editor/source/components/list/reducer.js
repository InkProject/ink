import Immutable from 'immutable'

import ActionType from '../action'

export default function list(state = Immutable.fromJS({
    show: false,
    loading: false,
    data: []
}), action) {
    switch (action.type) {
        case ActionType.SHOW_LIST:
            return state.set('show', true)
        case ActionType.HIDE_LIST:
            return state.set('show', false)
        case ActionType.TOOGLE_LIST:
            return state.set('show', !state.get('show'))
        case ActionType.SHOW_LOADING:
            return state.set('loading', action.flag)
        case ActionType.REFRESH_LIST:
            let newData = Object.keys(action.data).map(function(id) {
                let item = action.data[id]
                return {
                    id,
                    name: item.Path,
                    title: item.Article ? item.Article.Title : '未命名标题',
                    preview: item.Article ? item.Article.Preview : ''
                }
            })
            return state.set('data', Immutable.fromJS(newData))
        default:
            return state
    }
}
