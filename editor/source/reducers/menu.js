import Immutable from 'immutable';

import { ACTION } from '../actions';

export default function menu(state = Immutable.fromJS({
    show: true,
    loading: false,
    list: {
        selected: null,
        data: []
    }
}), action) {
    switch (action.type) {
        case ACTION.SHOW_LIST:
            return state.set('show', true);
        case ACTION.HIDE_LIST:
            return state.set('show', false);
        case ACTION.SHOW_LOADING:
            return state.set('loading', action.flag);
        case ACTION.SELECT_ARTICLE:
            return state.mergeDeep({list: {selected: action.id}});
        case ACTION.REFRESH_LIST:
            let newData = Object.keys(action.data).map(function(id) {
                let item = action.data[id];
                return {
                    id,
                    title: item.article.Title,
                    preview: item.article.Preview
                };
            });
            return state.mergeDeep({list: {data: Immutable.fromJS(newData)}});
        default:
            return state;
    }
}
