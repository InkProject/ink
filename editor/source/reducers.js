import { combineReducers } from 'redux';
import Immutable from 'immutable';

const initialState = Immutable.fromJS({
    show: true,
    loading: false,
    data: []
});

function list(state = initialState, action) {
    switch (action.type) {
        case 'SHOW_LIST':
            return state.set('show', true);
        case 'HIDE_LIST':
            return state.set('show', false);
        case 'SHOW_LOADING':
            return state.set('loading', action.flag);
        case 'REFRESH_LIST':
            let newData = Object.keys(action.data).map(function(id) {
                let item = action.data[id];
                return {
                    id,
                    title: item.article.Title,
                    preview: item.article.Preview
                };
            });
            return state.set('data', Immutable.fromJS(newData));
        default:
            return state;
    }
}

function content(state = '', action) {
    switch (action.type) {
        case 'SET_CONTENT':
            return action.content;
        default:
            return state;
    }
    return state;
}

export default combineReducers({
    list: list,
    content: content
});
