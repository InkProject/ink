import Immutable from 'immutable';

import { ACTION } from '../actions';

export default function util(state = Immutable.fromJS({
    tip: {
        loading: false,
        show: false,
        content: ''
    }
}), action) {
    switch (action.type) {
        case ACTION.SHOW_TIP:
            return state.mergeDeep({
                tip: {
                    show: action.show,
                    loading: action.loading,
                    content: action.content
                }
            });
        default:
            return state;
    }
}
