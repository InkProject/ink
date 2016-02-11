import Immutable from 'immutable'
import ActionType from '../action'
import _ from 'lodash'

const initState = Immutable.fromJS({
    id: '',
    title: '',
    tags: [],
    config: '',
    content: '',
    current: ''
})

export default function editor(state = initState, action) {
    switch (action.type) {
        case ActionType.EDITOR_SET_HEADER:
            return state.mergeDeep({
                title: action.title,
                tags: action.tags
            })
        case ActionType.EDITOR_SET_CONTENT:
            return state.mergeDeep({
                id: action.id,
                title: action.title,
                tags: action.tags,
                config: action.config,
                content: action.content,
                current: action.current
            })
        case ActionType.EDITOR_SET_CURRENT:
            return state.mergeDeep({
                current: action.current
            })
        case ActionType.EDITOR_RESET:
            return initState
        default:
            return state
    }
}
