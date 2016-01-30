import Immutable from 'immutable'
import { ACTION } from '../actions'
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
        case ACTION.SET_HEADER:
            return state.mergeDeep({
                title: action.title,
                tags: action.tags
            })
        case ACTION.SET_CONTENT:
            let current = `${_.trim(action.config)}\n\n---\n\n${_.trim(action.content)}`
            return state.mergeDeep({
                id: action.id,
                title: action.title,
                tags: action.tags,
                config: action.config,
                content: action.content,
                current: current
            })
        case ACTION.SET_CURRENT:
            return state.mergeDeep({
                current: action.current
            })
        default:
            return state
    }
}
