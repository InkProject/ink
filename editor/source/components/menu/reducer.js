import Immutable from 'immutable'

import ActionType from '../action'

export default function menu(state = Immutable.fromJS({
    focusMode: false,
    modal: {
        show: false
    }
}), action) {
    switch (action.type) {
        case ActionType.MODAL_SHOW:
            return state.mergeDeep({
                modal: { show: action.flag }
            })
        case ActionType.MENU_CHNAGE_FOCUS_MODE:
            return state.mergeDeep({
                focusMode: action.flag
            })
        default:
            return state
    }
}
