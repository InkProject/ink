import Immutable from 'immutable'

import ActionType from '../action'

export default function menu(state = Immutable.fromJS({
    modal: {
        show: false
    }
}), action) {
    switch (action.type) {
        case ActionType.MODAL_SHOW:
            return state.mergeDeep({
                modal: { show: action.flag }
            })
        default:
            return state
    }
}
