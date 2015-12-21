import Immutable from 'immutable';
import { ACTION } from '../actions';

export default function editor(state = Immutable.fromJS({
    title: '',
    tags: [],
    content: ''
}), action) {
    switch (action.type) {
        case ACTION.SET_CONTENT:
            return state.set('content', action.content);
        case ACTION.SET_HEADER:
            return state
                .set('title', action.title)
                .set('tags', action.tags);
        default:
            return state;
    }
}
