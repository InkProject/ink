import Immutable from 'immutable'

import ActionType from '../action'

export default function util(state = Immutable.fromJS({
    tip: {
        loading: false,
        show: false,
        error: false,
        content: '',
        action: null
    }
}), action) {
    switch (action.type) {
        case ActionType.UTIL_SHOW_TIP:
            return state.mergeDeep({
                tip: {
                    show: action.show,
                    error: action.error,
                    loading: action.loading,
                    content: action.content,
                    action: action.action
                }
            })
        default:
            return state
    }
}
