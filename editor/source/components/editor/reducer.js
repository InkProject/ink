import Immutable from 'immutable'
import ActionType from '../action'
import _ from 'lodash'

export default function editor(state = Immutable.fromJS({
    id: '',
    title: '',
    tags: [],
    config: '',
    content: '',
    current: ''
}), action) {
    switch (action.type) {
        case ActionType.SET_HEADER:
            return state.mergeDeep({
                title: action.title,
                tags: action.tags
            })
        case ActionType.SET_CONTENT:
            let current = `${_.trim(action.config)}\n\n---\n\n${_.trim(action.content)}`
            return state.mergeDeep({
                id: action.id,
                title: action.title,
                tags: action.tags,
                config: action.config,
                content: action.content,
                current: current
            })
        case ActionType.SET_CURRENT:
            return state.mergeDeep({
                current: action.current
            })
        default:
            return state
    }
}
