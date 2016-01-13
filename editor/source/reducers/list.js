import Immutable from 'immutable';

import { ACTION } from '../actions';

export default function list(state = Immutable.fromJS({
    show: false,
    loading: false,
    data: []
}), action) {
    switch (action.type) {
        case ACTION.SHOW_LIST:
            return state.set('show', true);
        case ACTION.HIDE_LIST:
            return state.set('show', false);
        case ACTION.TOOGLE_LIST:
            return state.set('show', !state.get('show'));
        case ACTION.SHOW_LOADING:
            return state.set('loading', action.flag);
        case ACTION.REFRESH_LIST:
            let newData = Object.keys(action.data).map(function(id) {
                let item = action.data[id];
                return {
                    id,
                    name: item.path,
                    title: item.article.Title,
                    preview: item.article.Preview
                };
            });
            return state.set('data', Immutable.fromJS(newData));
        default:
            return state;
    }
}
