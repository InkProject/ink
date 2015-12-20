import Immutable from 'immutable';

import { ACTION } from '../actions';

export default function editor(state = Immutable.fromJS({
    content: ''
}), action) {
    switch (action.type) {
        case ACTION.SET_CONTENT:
            return state.set('content', action.content);
        default:
            return state;
    }
}
