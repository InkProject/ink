import Immutable from 'immutable';
import { ACTION } from '../actions';

export default function editor(state = Immutable.fromJS({
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
            });
        case ACTION.SET_CONTENT:
            return state.mergeDeep({
                title: action.title,
                tags: action.tags,
                config: action.config,
                content: action.content,
                current: action.content
            });
        default:
            return state;
    }
}
