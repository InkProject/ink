import Immutable from 'immutable';

import { ACTION } from '../actions';

export default function menu(state = Immutable.fromJS({
    modal: {
        show: false
    }
}), action) {
    switch (action.type) {
        case ACTION.SHOW_MODAL:
            return state.mergeDeep({
                modal: {
                    show: action.flag
                }
            });
        default:
            return state;
    }
}
