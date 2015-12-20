import { combineReducers } from 'redux';
import Immutable from 'immutable';

import { ACTION } from './actions';

function menu(state = Immutable.fromJS({
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

function editor(state = Immutable.fromJS({
    content: ''
}), action) {
    switch (action.type) {
        case ACTION.SET_CONTENT:
            return state.set('content', action.content);
        default:
            return state;
    }
}

export default combineReducers({
    menu,
    editor
});
